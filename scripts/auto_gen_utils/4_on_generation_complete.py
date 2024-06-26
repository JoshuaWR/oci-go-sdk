# - this script is run as the as the last step of a build step of a job triggered on commits to SDK/CLI branches
# of the form auto-preview-* or auto-public-*
# - at this point the pom.xml has been updated, and the new code has been generated
# - this script reports back to JIRA whether or not the code generation was successful and commits / pushes the generated code

import argparse
import sys
import util
import config
import re
import shared.bitbucket_utils
from datetime import datetime
from shared.buildsvc_tc_compatibility import build_log_link, build_artifacts_link

try:
    from urllib import quote
except ImportError:
    from urllib.parse import quote

PREVIEW_TC_LINK = "[Preview Auto-Generation](https://teamcity.oci.oraclecorp.com/project/Sdk_SelfService_Preview_AutoPreviewBuilds)"
PUBLIC_TC_LINK = "[Public Release Auto-Generation](https://teamcity.oci.oraclecorp.com/project/Sdk_SelfService_PublicV2_AutoPublicBuilds)"

PR_DESCRIPTION_TEMPLATE = """This {build_description} was generated by the {tc_link} jobs in TeamCiy.

For more information, see the {build_log_link}.

This {build_description} includes changes for the following issues: {issue_keys}"""

TRANSITIONED_TO_SERVICE_TEAM_WORK_REQUIRED_TEMPLATE = """Setting this ticket to '{service_team_work_required}' automatically, since the SDK status fields for all requested tools have been set to '{success}' or '{done}'
{bypassed_tools_message}
Service team, please:

1. If you had manual changes to the CLI code, cherry-pick the CLI recordings from [preview|https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/python-cli/browse?at=refs%2Fheads%2Fpreview] into this branch: [{generated_branch}|https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/python-cli/browse?at=refs%2Fheads%2F{generated_branch}].
2. Make your feature API publicly available to all customers (un-whitelisted) in Prod, then flip the 'Feature API is publicly available & un-whitelisted in Prod' JIRA ticket field value to 'Yes'."""


TRANSITIONED_TO_SERVICE_TEAM_REVIEW_REQUIRED_TEMPLATE = """Setting this ticket to '{service_team_review_required}' automatically, since the SDK status fields for all requested tools have been set to '{success}' or '{done}'
{bypassed_tools_message}
{review_text}"""


def get_preview_review_text(issue_key, generated_branch, tool_name):
    text = "Service team, please review the resulting generated source diffs and transition the issue to 'Ready for Preview' if appropriate."

    m = re.search("^.*{}-(.*)$".format(tool_name), generated_branch)
    if not m:
        print("Did not find '{}' and a timestamp in the branch: '{}'".format(tool_name, generated_branch))
        return text

    branch_suffix = m.group(1)

    print("Looking for branch with suffix: {}".format(branch_suffix))

    pr = shared.bitbucket_utils.get_newest_pullrequest_matching_branch_suffix("SDK", config.DEXREQ_REPO_NAME, branch_suffix)
    if pr:
        hrefs = util.deep_get(pr, 'links.self')
        if hrefs:
            url = util.deep_get(hrefs[0], 'href')
            text = "Service team, please review the spec diff in this pull request, and if it contains (1) your entire change, and (2) nothing else, approve the pull request and transition the issue to 'Ready for Preview':\n\n" + url

    return text


def check_all_tools_successful(issue_key, build_type, repo_name, generated_branch, running_tool_name):
    issue = util.get_dexreq_issue(issue_key)

    bypassed_tools = []
    all_successful_or_done = True
    # k -> Tool name (e.g., RubySDK), v -> Jira status field identifier
    for tool_name, jira_field_id in util.get_jira_custom_field_ids_for_tool().items():
        status = getattr(issue.fields, jira_field_id)
        if config.BYPASS_CHECK_GENERATION_PREFIX + tool_name in issue.fields.labels:
            bypassed_tools.append(tool_name)
            continue

        if not str(status) == config.CUSTOM_STATUS_SUCCESS and not str(status) == config.CUSTOM_STATUS_DONE:
            all_successful_or_done = False
            break

    bypassed_tools_message = ''
    if bypassed_tools:
        bypassed_tools_message = "\nThe following requested tools are ignored: {}\n".format(''.join(bypassed_tools))

    print('All tools successful or done?: {}'.format(str(all_successful_or_done)))
    if all_successful_or_done:
        if build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
            cli_branch = get_cli_branch_text(generated_branch)
            # All tool creation successful, in public builds, we set the overall ticket status to "Service Team Work Required"
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_SERVICE_TEAM_WORK_REQUIRED)
            util.add_jira_comment(
                issue,
                TRANSITIONED_TO_SERVICE_TEAM_WORK_REQUIRED_TEMPLATE.format(
                    service_team_work_required=config.STATUS_SERVICE_TEAM_WORK_REQUIRED,
                    success=config.CUSTOM_STATUS_SUCCESS,
                    done=config.STATUS_DONE,
                    bypassed_tools_message=bypassed_tools_message,
                    generated_branch=cli_branch,
                    repo=repo_name
                )
            )
        elif build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW:
            # All tool creation successful, in preview builds, we set the overall ticket status to "Service Team Review Required"

            review_text = get_preview_review_text(issue_key, generated_branch, running_tool_name)

            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_SERVICE_TEAM_REVIEW_REQUIRED)
            util.add_jira_comment(
                issue,
                TRANSITIONED_TO_SERVICE_TEAM_REVIEW_REQUIRED_TEMPLATE.format(
                    service_team_review_required=config.STATUS_SERVICE_TEAM_REVIEW_REQUIRED,
                    success=config.CUSTOM_STATUS_SUCCESS,
                    done=config.CUSTOM_STATUS_DONE,
                    bypassed_tools_message=bypassed_tools_message,
                    review_text=review_text
                )
            )
        elif build_type in config.BULK_BUILD_TYPES:
            # All tool creation successful, in bulk builds, we set the overall ticket status to "DEX Bulk Review"
            util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DEX_BULK_REVIEW)


def get_cli_branch_text(gen_branch):
    for tool in config.TOOL_NAMES:
        if tool in gen_branch:
            return gen_branch.replace(tool, config.CLI_NAME)
    return gen_branch


def get_message(repo, commit_message, build_pass):
    message = commit_message.strip()
    if 'nothing to commit' in repo.git.status():
        message = "{} (no change in generated code)".format(message)

    if not build_pass:
        message = "FAILED: {}\n\nNote: This branch failed to build. It cannot be merged without manual changes. If necessary, you can use this branch as a starting point to fix the build (e.g. if you made a breaking change in preview and you now have to change tests or samples).".format(message)

    return message


def get_title(name, issue_keys, build_description, build_pass, build_type):
    if build_type in config.BULK_BUILD_TYPES:
        time_stamp = datetime.now()
        title = 'Auto Generated {build_description} for {repo_name} {date_time}'.format(
            repo_name=name,
            date_time=time_stamp.strftime("%c"),
            build_description=build_description
        )
    else:
        title = 'Auto Generated {build_description} for {repo_name} {issue_keys}'.format(
            repo_name=name,
            issue_keys=', '.join(issue_keys),
            build_description=build_description
        )

    if not build_pass:
        title = "FAILED: {}".format(title)

    return title


def get_jira_message(tool_name, links, build_description):
    links_text = "\n".join(links)
    branch_text = "this branch" if len(links) == 1 else "these branches"

    template = config.BUILD_PASS_JIRA_MESSAGE_TEMPLATE if build_pass else config.BUILD_FAIL_JIRA_MESSAGE_TEMPLATE

    message = template.format(
        tool_name=tool_name,
        repos=util.join(config.REPO_NAMES_FOR_TOOL[tool_name]),
        build_log_link=build_log_link(build_id, "build log"),
        build_artifacts_link=build_artifacts_link(build_id),
        links=links_text,
        build_description=build_description,
        branch_text=branch_text
    )

    comment_type = config.COMMENT_TYPE_SUCCESS if build_pass else config.COMMENT_TYPE_ERROR

    return message, comment_type


def get_link_for_bulk_pr(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch):
    pr_url = util.create_pull_request(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch)
    print("Automatically generated pull request: {}".format(pr_url))
    return "- [Pull request for the {repo} changes|{link}]".format(repo=name, link=pr_url)


def get_link_for_individual_build_passed(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch, build_type=None):
    # If you change this, also change PR_REQUEST_LINK_TEMPLATE in autogen_issue_advisor_public.py
    pull_request_link = 'https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/{repo}/compare?targetBranch=refs%2Fheads%2F{target_branch}&sourceBranch=refs%2Fheads%2F{generated_branch}&title={title}&description={description}&targetRepoId={target_repo_id}'.format(
        repo=repo_for_link,
        generated_branch=generated_branch,
        title=quote(title),
        description=quote(description),
        target_repo_id=target_repo_id,
        target_branch=target_branch)

    if repo_for_link == 'python-cli' and build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW:
        # store PR link in preview-pr.txt file - it will be used in create_cli_design_review_ticket script.
        with open('preview-pr.txt', 'w') as filehandle:
            filehandle.write(pull_request_link)

    print('Link with diff for {} changes: {}'.format(name, pull_request_link))
    return "- [Link with diff for {repo} changes|{link}]".format(repo=name, link=pull_request_link)


def get_link_for_individual_build_failed(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch):
    branch_link = 'https://bitbucket.oci.oraclecorp.com/projects/SDK/repos/{repo}/browse?at=refs%2Fheads%2F{generated_branch}'.format(
        repo=repo_for_link,
        generated_branch=generated_branch)

    print('Link with failed branch for {}: {}'.format(name, branch_link))
    return "- [Link to failed branch for {repo} changes|{link}]".format(repo=name, link=branch_link)


if __name__ == "__main__":
    # team city build step will have already run Maven so at this point generation will be done
    # somehow determine if Python SDK / CLI build jobs succeeded or failed and report back to JIRA
    # consider invoking build_preview.sh from this script
    # previously we had two build steps, one for build.sh for the python SDK and one for build.sh for the CLI not sure if we need both of those or not
    # but we do need to wait in between build SDK and build CLI
    # right now those build steps handle publishing the artifacts so it is probably better if I leave those steps as is and just run this after
    parser = argparse.ArgumentParser(description='Post completion updates to DEXREQ tickets.')
    parser.add_argument('--build-id',
                        required=True,
                        help='The team city build id for the build that is running this script.  This is used to update the relevant JIRA tickets with links to the team city build')
    parser.add_argument('--tool',
                        default=config.CLI_NAME,
                        help='The tool for which to generate the preview. Accepted values: {}'.format(', '.join(config.TOOL_NAMES)))
    parser.add_argument('--build-type',
                        default=config.BUILD_TYPE_INDIVIDUAL_PREVIEW,
                        help='The build type to use, can be one of the following: {}'.format(', '.join(config.VALID_BUILD_TYPES)))
    parser.add_argument('--dry-run',
                        default=False,
                        action='store_true',
                        help='Perform a dry-run')

    args = parser.parse_args()
    build_id = args.build_id
    tool_name = args.tool
    build_type = args.build_type
    config.IS_DRY_RUN = args.dry_run

    shared.bitbucket_utils.setup_bitbucket(args)

    # Build type based configuration
    if build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PUBLIC:
        tc_link = PUBLIC_TC_LINK
        build_description = "Bulk Public"
        target_branch = "master"
    elif build_type == config.BUILD_TYPE_INDIVIDUAL_PUBLIC:
        tc_link = PUBLIC_TC_LINK
        build_description = "Public Release"
        target_branch = "master"
    elif build_type == config.BUILD_TYPE_BULK_PENDING_MERGE_PREVIEW:
        tc_link = PREVIEW_TC_LINK
        build_description = "Bulk Preview"
        target_branch = "preview"
    elif build_type == config.BUILD_TYPE_INDIVIDUAL_PREVIEW:
        tc_link = PREVIEW_TC_LINK
        build_description = "Preview"
        target_branch = "preview"
    else:
        raise ValueError("Unknown build type: '{}', must be one of: {}".format(
            build_type, ', '.join(config.VALID_BUILD_TYPES)))

    # Determines if a given build_tool should report status back to the associated Jira issue
    is_tool_jira_reportable = util.is_tool_jira_reportable(tool_name)

    # check if we should even run
    generation_pass, build_pass = util.were_steps_successful(tool_name)
    if not generation_pass:
        print("Generation did not pass, not proceeding.")
        sys.exit(0)  # exit, but without failing here; the generation step already failed

    if not build_pass and build_type in config.BULK_BUILD_TYPES:
        print("Build did not pass, and this was a bulk build ({}). Not proceeding.".format(build_type))
        sys.exit(0)

    # get current branches
    current_branches = {}
    for name, repo in zip(config.REPO_NAMES_FOR_TOOL[tool_name], config.REPOS_FOR_TOOL[tool_name]):
        current_branches[name] = [branch.strip()[2:] for branch in repo.git.branch().split('\n') if branch.startswith('* ')][0]

    values = list(current_branches.values())
    if not all(x == values[0] for x in values):
        sys.exit('Expected branches for {} to be in sync. Got the following: {}'.format(config.REPO_NAMES_FOR_TOOL[tool_name], current_branches))

    if not all(x.startswith("auto-") for x in values):
        sys.exit('Expected all branches for {} to start with "auto-". Got the following: {}'.format(config.REPO_NAMES_FOR_TOOL[tool_name], current_branches))

    branch = current_branches[config.REPO_NAMES_FOR_TOOL[tool_name][0]]

    # create new branches generated-{original auto-preview-* or auto-public-* branch name}
    # this job is triggered by commits to auto-preview-* or auto-public-* so we want to move to a new branch we don't retrigger this job
    prefix_to_use = config.GENERATION_BRANCH_PREFIX if build_pass else config.FAILED_BRANCH_PREFIX
    generated_branch = '{generation_branch_prefix}-{branch}'.format(generation_branch_prefix=prefix_to_use, branch=branch)
    print("Branch to push: {}".format(generated_branch))
    for repo in config.REPOS_FOR_TOOL[tool_name]:
        repo.git.checkout(b=generated_branch)

    # TODO: parse last commit message to related DEXREQ issue
    last_commit_message = util.get_last_commit_message(tool_name)

    # will be of the form 'Updating pom.xml for DEXREQ-123: Preview for RQS
    print('Found last commit: {}'.format(last_commit_message))

    issue_keys = util.parse_issue_keys_from_commit_message(last_commit_message)
    if issue_keys:
        commit_message = '{commit_prefix} [[{issue_keys}]]'.format(commit_prefix=config.GENERATION_COMMIT_MESSAGE_PREFIX, issue_keys=', '.join(issue_keys))
    else:
        commit_message = '{commit_prefix} "{last_commit_message}"'.format(commit_prefix=config.GENERATION_COMMIT_MESSAGE_PREFIX, last_commit_message=last_commit_message)

    try:
        # add all generated files and commit
        links = []
        for name, repo, repo_for_link, target_repo_id in zip(config.REPO_NAMES_FOR_TOOL[tool_name], config.REPOS_FOR_TOOL[tool_name], config.REPO_FOR_LINKS[tool_name], config.TARGET_REPO_IDS_FOR_LINKS[tool_name]):
            repo.git.add(A=True)
            message = get_message(repo, commit_message, build_pass)

            print('Commiting the following: {}'.format(message))

            repo.git.commit("-m", message, "--allow-empty")

            # Rebase/squash the last two commits together
            last_two_log_message = repo.git.log(n=2, format='Author: %cn <%aE>%nDate:  %aD%n%s%n%b')
            repo.git.reset("--soft", "HEAD~2")
            repo.git.commit("-m", last_two_log_message + "\n\n(automatically combined)", "--allow-empty")

            if config.IS_DRY_RUN:
                print('DRY-RUN: not pushing to origin HEAD')
            else:
                repo.git.push('-u','origin','HEAD')
                # now that generated-* branch is pushed, we can delete the auto-preview or auto-public branch
                util.safe_delete_branch(repo, branch)

            title = get_title(name, issue_keys, build_description, build_pass, build_type)

            description = PR_DESCRIPTION_TEMPLATE.format(
                build_log_link=build_log_link(build_id, text="build results"),
                issue_keys=', '.join(issue_keys),
                build_description=build_description,
                tc_link=tc_link
            )

            last_repo_name = config.REPO_NAMES_FOR_TOOL[tool_name][-1]
            link = None
            if build_type in config.BULK_BUILD_TYPES:
                # For the Python CLI, we build both the Python SDK (first) and the Python CLI (last).
                # We only want to send out a pull request for the last repo.
                if name == last_repo_name:
                    link = get_link_for_bulk_pr(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch)
                else:
                    print('Not sending out PR for {}, since it is not the last repo ({}) for tool {}'.format(name, last_repo_name, tool_name))
            else:
                if build_pass:
                    # For the Python CLI, we build both the Python SDK (first) and the Python CLI (last).
                    # We only want to include a link for the last repo.
                    if name == last_repo_name:
                        link = get_link_for_individual_build_passed(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch, build_type)
                    else:
                        print('Not including link for {}, since it is not the last repo ({}) for tool {}'.format(name, last_repo_name, tool_name))
                else:
                    # Build did not pass (even for the Python SDK repo of the Python CLI build)
                    link = get_link_for_individual_build_failed(repo_for_link, generated_branch, title, description, target_repo_id, name, target_branch)
            if link:
                links.append(link)

        if is_tool_jira_reportable:
            if issue_keys:
                for issue_key in issue_keys:
                    # post back to JIRA indicating that a build is ready and the change is Pending Merge by the SDK / CLI team

                    message, comment_type = get_jira_message(tool_name, links, build_description)

                    util.add_jira_comment(
                        issue_key,
                        message,
                        comment_type=comment_type
                    )

                    # update issue custom status for tool_name to 'Success'
                    issue = util.get_dexreq_issue(issue_key)
                    if build_pass:
                        util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_SUCCESS, tool_name)

                        # check if all status fields are not in 'Success'
                        check_all_tools_successful(issue_key, build_type, config.REPO_FOR_LINKS[tool_name][0], generated_branch, tool_name)
                    else:
                        util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)
                        # if an issue is already in 'DEX Support Required' based on failure for another tool, we do not want to overwrite that
                        util.transition_issue_overall_status_if_not_in_status(util.JIRA_CLIENT(), issue, desired_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION, blacklisted_status=config.STATUS_DEX_SUPPORT_REQUIRED)

                    # refresh issue from server to make sure we have most up to date state
                    issue = util.get_dexreq_issue(issue_key)

                    # this logic is a catch all to make sure if this is the last or only tool to complete
                    # that we never strand a ticket in 'Processing'
                    if issue.fields.status.name == config.STATUS_PROCESSING or issue.fields.status.name == config.STATUS_PROCESSING_BULK:
                        any_tools_processing = False
                        any_tools_failed = False
                        for jira_field_id in util.get_jira_custom_field_ids_for_tool().values():
                            status = getattr(issue.fields, jira_field_id)
                            if str(status) == config.CUSTOM_STATUS_PROCESSING:
                                any_tools_processing = True
                            if str(status) == config.CUSTOM_STATUS_FAILURE:
                                any_tools_failed = True

                        print('Any tools in Processing?: {}'.format(str(any_tools_processing)))
                        print('Any tools Failed?: {}'.format(str(any_tools_failed)))
                        if not any_tools_processing:
                            # ticket is still in 'Processing' global state but no individual tools are 'Processing'
                            # this can happen if the last job to be 'Processing' completes successfully but other tools are still in To Do / Failure / None
                            # for example, if 2 SDKs fail and we re-run a single one which succeeds, we will hit this case
                            # thus we look at the individual status fields and try to figure out which state to put the ticket in
                            if any_tools_failed:
                                util.transition_issue_overall_status_if_not_in_status(util.JIRA_CLIENT(), issue, desired_status=config.STATUS_SERVICE_TEAM_FAILURE_INVESTIGATION, blacklisted_status=config.STATUS_DEX_SUPPORT_REQUIRED)
                            else:
                                # this is a very rare case that should only happen when someone is manually tweaking SDK status fields
                                # so we put the ticket into backlog and the reporter will have to move it to a meaningful status
                                # this condition will be hit when the current job succeeded, no other jobs are in progress, and there are other status fields still in To Do or None
                                util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_BACKLOG)
            else:
                print('Did not find any issue keys in commit message (text between double brackets: "[[issue-key]]"). Not updating any JIRA issues.')
    except Exception as e:
        print(e)

        if not is_tool_jira_reportable:
            raise e

        # catch anything that went wrong and post it in the ticket so it never gets stuck in 'Processing'
        if issue_keys:
            for issue_key in issue_keys:
                issue = util.get_dexreq_issue(issue_key)
                util.transition_issue_per_tool_status(util.JIRA_CLIENT(), issue, config.CUSTOM_STATUS_FAILURE, tool_name)
                util.transition_issue_overall_status(util.JIRA_CLIENT(), issue, config.STATUS_DEX_SUPPORT_REQUIRED)
                util.add_jira_comment(
                    issue.key,
                    config.SELF_SERVICE_BUG_TEMPLATE.format(
                        exception=str(e),
                        build_log_link=build_log_link(build_id)
                    ),
                    comment_type=config.COMMENT_TYPE_ERROR
                )
