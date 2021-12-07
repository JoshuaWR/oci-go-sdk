// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"github.com/oracle/oci-go-sdk/v53/common"
)

// SubscriptionSku SKU information.
type SubscriptionSku struct {

	// Stock keeping unit ID.
	Sku *string `mandatory:"true" json:"sku"`

	// Quantity of the stock units.
	Quantity *int `mandatory:"false" json:"quantity"`

	// Description of the stock units.
	Description *string `mandatory:"false" json:"description"`

	// Sales order line identifier.
	GsiOrderLineId *string `mandatory:"false" json:"gsiOrderLineId"`

	// Description of the covered product belonging to this SKU.
	LicensePartDescription *string `mandatory:"false" json:"licensePartDescription"`

	// Base metric for billing the service.
	MetricName *string `mandatory:"false" json:"metricName"`

	// Denotes if the SKU is considered as a parent or child.
	IsBaseServiceComponent *bool `mandatory:"false" json:"isBaseServiceComponent"`

	// Denotes if an additional test instance can be provisioned by the SAAS application.
	IsAdditionalInstance *bool `mandatory:"false" json:"isAdditionalInstance"`

	// Date-time when the SKU was created.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Date-time when the SKU ended.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`
}

func (m SubscriptionSku) String() string {
	return common.PointerString(m)
}