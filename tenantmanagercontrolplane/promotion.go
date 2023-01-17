// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Promotion Promotion information for a subscription.
type Promotion struct {

	// How long the promotion related to the subscription, if any, is valid in duration units.
	Duration *int `mandatory:"false" json:"duration"`

	// Unit for the duration.
	DurationUnit *string `mandatory:"false" json:"durationUnit"`

	// Total amount of credit for the promotion related to the subscription if there is one.
	Amount *float32 `mandatory:"false" json:"amount"`

	// Current status of the promotion related to the subscription if there is one.
	Status PromotionStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Whether or not customer intends to pay once the promotion is done.
	IsIntentToPay *bool `mandatory:"false" json:"isIntentToPay"`

	// Currency unit associated with the promotion.
	CurrencyUnit *string `mandatory:"false" json:"currencyUnit"`

	// Date-time for when the promotion starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Date-time for when the promotion ends.
	TimeExpired *common.SDKTime `mandatory:"false" json:"timeExpired"`
}

func (m Promotion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Promotion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPromotionStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPromotionStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PromotionStatusEnum Enum with underlying type: string
type PromotionStatusEnum string

// Set of constants representing the allowable values for PromotionStatusEnum
const (
	PromotionStatusInitialized PromotionStatusEnum = "INITIALIZED"
	PromotionStatusActive      PromotionStatusEnum = "ACTIVE"
	PromotionStatusExpired     PromotionStatusEnum = "EXPIRED"
)

var mappingPromotionStatusEnum = map[string]PromotionStatusEnum{
	"INITIALIZED": PromotionStatusInitialized,
	"ACTIVE":      PromotionStatusActive,
	"EXPIRED":     PromotionStatusExpired,
}

var mappingPromotionStatusEnumLowerCase = map[string]PromotionStatusEnum{
	"initialized": PromotionStatusInitialized,
	"active":      PromotionStatusActive,
	"expired":     PromotionStatusExpired,
}

// GetPromotionStatusEnumValues Enumerates the set of values for PromotionStatusEnum
func GetPromotionStatusEnumValues() []PromotionStatusEnum {
	values := make([]PromotionStatusEnum, 0)
	for _, v := range mappingPromotionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPromotionStatusEnumStringValues Enumerates the set of values in String for PromotionStatusEnum
func GetPromotionStatusEnumStringValues() []string {
	return []string{
		"INITIALIZED",
		"ACTIVE",
		"EXPIRED",
	}
}

// GetMappingPromotionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPromotionStatusEnum(val string) (PromotionStatusEnum, bool) {
	enum, ok := mappingPromotionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
