// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v52/common"
)

// UpdateBuildPipelineStageDetails The information to be updated.
type UpdateBuildPipelineStageDetails interface {

	// Stage identifier which can be renamed and is not necessarily unique
	GetDisplayName() *string

	// Optional description about the BuildStage
	GetDescription() *string

	GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatebuildpipelinestagedetails struct {
	JsonData                                []byte
	DisplayName                             *string                                  `mandatory:"false" json:"displayName"`
	Description                             *string                                  `mandatory:"false" json:"description"`
	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`
	FreeformTags                            map[string]string                        `mandatory:"false" json:"freeformTags"`
	DefinedTags                             map[string]map[string]interface{}        `mandatory:"false" json:"definedTags"`
	BuildPipelineStageType                  string                                   `json:"buildPipelineStageType"`
}

// UnmarshalJSON unmarshals json
func (m *updatebuildpipelinestagedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatebuildpipelinestagedetails updatebuildpipelinestagedetails
	s := struct {
		Model Unmarshalerupdatebuildpipelinestagedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.BuildPipelineStagePredecessorCollection = s.Model.BuildPipelineStagePredecessorCollection
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.BuildPipelineStageType = s.Model.BuildPipelineStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatebuildpipelinestagedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.BuildPipelineStageType {
	case "WAIT":
		mm := UpdateWaitStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BUILD":
		mm := UpdateBuildStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TRIGGER_DEPLOYMENT_PIPELINE":
		mm := UpdateTriggerDeploymentStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DELIVER_ARTIFACT":
		mm := UpdateDeliverArtifactStageDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updatebuildpipelinestagedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m updatebuildpipelinestagedetails) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m updatebuildpipelinestagedetails) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m updatebuildpipelinestagedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatebuildpipelinestagedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatebuildpipelinestagedetails) String() string {
	return common.PointerString(m)
}
