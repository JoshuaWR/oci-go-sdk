// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v44/common"
)

// CreateConnectionValidationDetails The properties used in create connection validation operations.
type CreateConnectionValidationDetails struct {
	DataAsset CreateDataAssetDetails `mandatory:"false" json:"dataAsset"`

	Connection CreateConnectionDetails `mandatory:"false" json:"connection"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m CreateConnectionValidationDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateConnectionValidationDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DataAsset        createdataassetdetails  `json:"dataAsset"`
		Connection       createconnectiondetails `json:"connection"`
		RegistryMetadata *RegistryMetadata       `json:"registryMetadata"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.DataAsset.UnmarshalPolymorphicJSON(model.DataAsset.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DataAsset = nn.(CreateDataAssetDetails)
	} else {
		m.DataAsset = nil
	}

	nn, e = model.Connection.UnmarshalPolymorphicJSON(model.Connection.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Connection = nn.(CreateConnectionDetails)
	} else {
		m.Connection = nil
	}

	m.RegistryMetadata = model.RegistryMetadata

	return
}
