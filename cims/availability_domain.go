// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Support Management API
//
// Use the Support Management API to manage support requests. For more information, see Getting Help and Contacting Support (https://docs.cloud.oracle.com/iaas/Content/GSG/Tasks/contactingsupport.htm). **Note**: Before you can create service requests with this API, you need to have an Oracle Single Sign On (SSO) account, and you need to register your Customer Support Identifier (CSI) with My Oracle Support.
//

package cims

import (
	"strings"
)

// AvailabilityDomainEnum Enum with underlying type: string
type AvailabilityDomainEnum string

// Set of constants representing the allowable values for AvailabilityDomainEnum
const (
	AvailabilityDomainDev1            AvailabilityDomainEnum = "DEV_1"
	AvailabilityDomainDev2            AvailabilityDomainEnum = "DEV_2"
	AvailabilityDomainDev3            AvailabilityDomainEnum = "DEV_3"
	AvailabilityDomainIntegNext1      AvailabilityDomainEnum = "INTEG_NEXT_1"
	AvailabilityDomainIntegStable1    AvailabilityDomainEnum = "INTEG_STABLE_1"
	AvailabilityDomainSeaAd1          AvailabilityDomainEnum = "SEA_AD_1"
	AvailabilityDomainSeaAd2          AvailabilityDomainEnum = "SEA_AD_2"
	AvailabilityDomainSeaAd3          AvailabilityDomainEnum = "SEA_AD_3"
	AvailabilityDomainPhxAd1          AvailabilityDomainEnum = "PHX_AD_1"
	AvailabilityDomainPhxAd2          AvailabilityDomainEnum = "PHX_AD_2"
	AvailabilityDomainPhxAd3          AvailabilityDomainEnum = "PHX_AD_3"
	AvailabilityDomainUsAshburnAd1    AvailabilityDomainEnum = "US_ASHBURN_AD_1"
	AvailabilityDomainUsAshburnAd2    AvailabilityDomainEnum = "US_ASHBURN_AD_2"
	AvailabilityDomainUsAshburnAd3    AvailabilityDomainEnum = "US_ASHBURN_AD_3"
	AvailabilityDomainUsAshburnAd4    AvailabilityDomainEnum = "US_ASHBURN_AD_4"
	AvailabilityDomainEuFrankfurt1Ad1 AvailabilityDomainEnum = "EU_FRANKFURT_1_AD_1"
	AvailabilityDomainEuFrankfurt1Ad2 AvailabilityDomainEnum = "EU_FRANKFURT_1_AD_2"
	AvailabilityDomainEuFrankfurt1Ad3 AvailabilityDomainEnum = "EU_FRANKFURT_1_AD_3"
	AvailabilityDomainUkLondon1Ad1    AvailabilityDomainEnum = "UK_LONDON_1_AD_1"
	AvailabilityDomainUkLondon1Ad2    AvailabilityDomainEnum = "UK_LONDON_1_AD_2"
	AvailabilityDomainUkLondon1Ad3    AvailabilityDomainEnum = "UK_LONDON_1_AD_3"
	AvailabilityDomainCaToronto1Ad1   AvailabilityDomainEnum = "CA_TORONTO_1_AD_1"
	AvailabilityDomainApTokyo1Ad1     AvailabilityDomainEnum = "AP_TOKYO_1_AD_1"
	AvailabilityDomainApSeoul1Ad1     AvailabilityDomainEnum = "AP_SEOUL_1_AD_1"
	AvailabilityDomainApMumbai1Ad1    AvailabilityDomainEnum = "AP_MUMBAI_1_AD_1"
	AvailabilityDomainSaSaopaulo1Ad1  AvailabilityDomainEnum = "SA_SAOPAULO_1_AD_1"
	AvailabilityDomainMeJeddah1Ad1    AvailabilityDomainEnum = "ME_JEDDAH_1_AD_1"
	AvailabilityDomainApOsaka1Ad1     AvailabilityDomainEnum = "AP_OSAKA_1_AD_1"
	AvailabilityDomainApSydney1Ad1    AvailabilityDomainEnum = "AP_SYDNEY_1_AD_1"
	AvailabilityDomainEuZurich1Ad1    AvailabilityDomainEnum = "EU_ZURICH_1_AD_1"
	AvailabilityDomainEuAmsterdam1Ad1 AvailabilityDomainEnum = "EU_AMSTERDAM_1_AD_1"
	AvailabilityDomainApMelbourne1Ad1 AvailabilityDomainEnum = "AP_MELBOURNE_1_AD_1"
	AvailabilityDomainCaMontreal1Ad1  AvailabilityDomainEnum = "CA_MONTREAL_1_AD_1"
	AvailabilityDomainApHyderabad1Ad1 AvailabilityDomainEnum = "AP_HYDERABAD_1_AD_1"
	AvailabilityDomainApChuncheon1Ad1 AvailabilityDomainEnum = "AP_CHUNCHEON_1_AD_1"
	AvailabilityDomainNoAd            AvailabilityDomainEnum = "NO_AD"
)

var mappingAvailabilityDomainEnum = map[string]AvailabilityDomainEnum{
	"DEV_1":               AvailabilityDomainDev1,
	"DEV_2":               AvailabilityDomainDev2,
	"DEV_3":               AvailabilityDomainDev3,
	"INTEG_NEXT_1":        AvailabilityDomainIntegNext1,
	"INTEG_STABLE_1":      AvailabilityDomainIntegStable1,
	"SEA_AD_1":            AvailabilityDomainSeaAd1,
	"SEA_AD_2":            AvailabilityDomainSeaAd2,
	"SEA_AD_3":            AvailabilityDomainSeaAd3,
	"PHX_AD_1":            AvailabilityDomainPhxAd1,
	"PHX_AD_2":            AvailabilityDomainPhxAd2,
	"PHX_AD_3":            AvailabilityDomainPhxAd3,
	"US_ASHBURN_AD_1":     AvailabilityDomainUsAshburnAd1,
	"US_ASHBURN_AD_2":     AvailabilityDomainUsAshburnAd2,
	"US_ASHBURN_AD_3":     AvailabilityDomainUsAshburnAd3,
	"US_ASHBURN_AD_4":     AvailabilityDomainUsAshburnAd4,
	"EU_FRANKFURT_1_AD_1": AvailabilityDomainEuFrankfurt1Ad1,
	"EU_FRANKFURT_1_AD_2": AvailabilityDomainEuFrankfurt1Ad2,
	"EU_FRANKFURT_1_AD_3": AvailabilityDomainEuFrankfurt1Ad3,
	"UK_LONDON_1_AD_1":    AvailabilityDomainUkLondon1Ad1,
	"UK_LONDON_1_AD_2":    AvailabilityDomainUkLondon1Ad2,
	"UK_LONDON_1_AD_3":    AvailabilityDomainUkLondon1Ad3,
	"CA_TORONTO_1_AD_1":   AvailabilityDomainCaToronto1Ad1,
	"AP_TOKYO_1_AD_1":     AvailabilityDomainApTokyo1Ad1,
	"AP_SEOUL_1_AD_1":     AvailabilityDomainApSeoul1Ad1,
	"AP_MUMBAI_1_AD_1":    AvailabilityDomainApMumbai1Ad1,
	"SA_SAOPAULO_1_AD_1":  AvailabilityDomainSaSaopaulo1Ad1,
	"ME_JEDDAH_1_AD_1":    AvailabilityDomainMeJeddah1Ad1,
	"AP_OSAKA_1_AD_1":     AvailabilityDomainApOsaka1Ad1,
	"AP_SYDNEY_1_AD_1":    AvailabilityDomainApSydney1Ad1,
	"EU_ZURICH_1_AD_1":    AvailabilityDomainEuZurich1Ad1,
	"EU_AMSTERDAM_1_AD_1": AvailabilityDomainEuAmsterdam1Ad1,
	"AP_MELBOURNE_1_AD_1": AvailabilityDomainApMelbourne1Ad1,
	"CA_MONTREAL_1_AD_1":  AvailabilityDomainCaMontreal1Ad1,
	"AP_HYDERABAD_1_AD_1": AvailabilityDomainApHyderabad1Ad1,
	"AP_CHUNCHEON_1_AD_1": AvailabilityDomainApChuncheon1Ad1,
	"NO_AD":               AvailabilityDomainNoAd,
}

var mappingAvailabilityDomainEnumLowerCase = map[string]AvailabilityDomainEnum{
	"dev_1":               AvailabilityDomainDev1,
	"dev_2":               AvailabilityDomainDev2,
	"dev_3":               AvailabilityDomainDev3,
	"integ_next_1":        AvailabilityDomainIntegNext1,
	"integ_stable_1":      AvailabilityDomainIntegStable1,
	"sea_ad_1":            AvailabilityDomainSeaAd1,
	"sea_ad_2":            AvailabilityDomainSeaAd2,
	"sea_ad_3":            AvailabilityDomainSeaAd3,
	"phx_ad_1":            AvailabilityDomainPhxAd1,
	"phx_ad_2":            AvailabilityDomainPhxAd2,
	"phx_ad_3":            AvailabilityDomainPhxAd3,
	"us_ashburn_ad_1":     AvailabilityDomainUsAshburnAd1,
	"us_ashburn_ad_2":     AvailabilityDomainUsAshburnAd2,
	"us_ashburn_ad_3":     AvailabilityDomainUsAshburnAd3,
	"us_ashburn_ad_4":     AvailabilityDomainUsAshburnAd4,
	"eu_frankfurt_1_ad_1": AvailabilityDomainEuFrankfurt1Ad1,
	"eu_frankfurt_1_ad_2": AvailabilityDomainEuFrankfurt1Ad2,
	"eu_frankfurt_1_ad_3": AvailabilityDomainEuFrankfurt1Ad3,
	"uk_london_1_ad_1":    AvailabilityDomainUkLondon1Ad1,
	"uk_london_1_ad_2":    AvailabilityDomainUkLondon1Ad2,
	"uk_london_1_ad_3":    AvailabilityDomainUkLondon1Ad3,
	"ca_toronto_1_ad_1":   AvailabilityDomainCaToronto1Ad1,
	"ap_tokyo_1_ad_1":     AvailabilityDomainApTokyo1Ad1,
	"ap_seoul_1_ad_1":     AvailabilityDomainApSeoul1Ad1,
	"ap_mumbai_1_ad_1":    AvailabilityDomainApMumbai1Ad1,
	"sa_saopaulo_1_ad_1":  AvailabilityDomainSaSaopaulo1Ad1,
	"me_jeddah_1_ad_1":    AvailabilityDomainMeJeddah1Ad1,
	"ap_osaka_1_ad_1":     AvailabilityDomainApOsaka1Ad1,
	"ap_sydney_1_ad_1":    AvailabilityDomainApSydney1Ad1,
	"eu_zurich_1_ad_1":    AvailabilityDomainEuZurich1Ad1,
	"eu_amsterdam_1_ad_1": AvailabilityDomainEuAmsterdam1Ad1,
	"ap_melbourne_1_ad_1": AvailabilityDomainApMelbourne1Ad1,
	"ca_montreal_1_ad_1":  AvailabilityDomainCaMontreal1Ad1,
	"ap_hyderabad_1_ad_1": AvailabilityDomainApHyderabad1Ad1,
	"ap_chuncheon_1_ad_1": AvailabilityDomainApChuncheon1Ad1,
	"no_ad":               AvailabilityDomainNoAd,
}

// GetAvailabilityDomainEnumValues Enumerates the set of values for AvailabilityDomainEnum
func GetAvailabilityDomainEnumValues() []AvailabilityDomainEnum {
	values := make([]AvailabilityDomainEnum, 0)
	for _, v := range mappingAvailabilityDomainEnum {
		values = append(values, v)
	}
	return values
}

// GetAvailabilityDomainEnumStringValues Enumerates the set of values in String for AvailabilityDomainEnum
func GetAvailabilityDomainEnumStringValues() []string {
	return []string{
		"DEV_1",
		"DEV_2",
		"DEV_3",
		"INTEG_NEXT_1",
		"INTEG_STABLE_1",
		"SEA_AD_1",
		"SEA_AD_2",
		"SEA_AD_3",
		"PHX_AD_1",
		"PHX_AD_2",
		"PHX_AD_3",
		"US_ASHBURN_AD_1",
		"US_ASHBURN_AD_2",
		"US_ASHBURN_AD_3",
		"US_ASHBURN_AD_4",
		"EU_FRANKFURT_1_AD_1",
		"EU_FRANKFURT_1_AD_2",
		"EU_FRANKFURT_1_AD_3",
		"UK_LONDON_1_AD_1",
		"UK_LONDON_1_AD_2",
		"UK_LONDON_1_AD_3",
		"CA_TORONTO_1_AD_1",
		"AP_TOKYO_1_AD_1",
		"AP_SEOUL_1_AD_1",
		"AP_MUMBAI_1_AD_1",
		"SA_SAOPAULO_1_AD_1",
		"ME_JEDDAH_1_AD_1",
		"AP_OSAKA_1_AD_1",
		"AP_SYDNEY_1_AD_1",
		"EU_ZURICH_1_AD_1",
		"EU_AMSTERDAM_1_AD_1",
		"AP_MELBOURNE_1_AD_1",
		"CA_MONTREAL_1_AD_1",
		"AP_HYDERABAD_1_AD_1",
		"AP_CHUNCHEON_1_AD_1",
		"NO_AD",
	}
}

// GetMappingAvailabilityDomainEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAvailabilityDomainEnum(val string) (AvailabilityDomainEnum, bool) {
	enum, ok := mappingAvailabilityDomainEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
