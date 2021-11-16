// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"github.com/oracle/oci-go-sdk/v52/common"
)

// NotebookSessionConfigurationDetails Details for the notebook session configuration.
type NotebookSessionConfigurationDetails struct {

	// The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint.
	Shape *string `mandatory:"true" json:"shape"`

	// A notebook session instance is provided with a VNIC for network access.  This specifies the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs.
	BlockStorageSizeInGBs *int `mandatory:"false" json:"blockStorageSizeInGBs"`

	NotebookSessionShapeConfigDetails *NotebookSessionShapeConfigDetails `mandatory:"false" json:"notebookSessionShapeConfigDetails"`
}

func (m NotebookSessionConfigurationDetails) String() string {
	return common.PointerString(m)
}
