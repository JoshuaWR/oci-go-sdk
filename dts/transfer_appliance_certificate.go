// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// DTS API
//
// A description of the DTS API
//

package dts

import (
	"github.com/oracle/oci-go-sdk/common"
)

// TransferApplianceCertificate The representation of TransferApplianceCertificate
type TransferApplianceCertificate struct {
	Certificate *string `mandatory:"false" json:"certificate"`
}

func (m TransferApplianceCertificate) String() string {
	return common.PointerString(m)
}