// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateTargetAlertPolicyAssociationDetails The details used to create a new target-alert policy association.
type CreateTargetAlertPolicyAssociationDetails struct {

	// The OCID of the alert policy.
	PolicyId *string `mandatory:"true" json:"policyId"`

	// The OCID of the target.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The OCID of the compartment where the target-alert policy association is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Indicates if the target-alert policy association is enabled or disabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The display name of the target-alert policy association.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Describes the target-alert policy association.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateTargetAlertPolicyAssociationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTargetAlertPolicyAssociationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}