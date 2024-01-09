// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CancelSnoozeDeploymentUpgradeDetails The information about snooze canceling.
type CancelSnoozeDeploymentUpgradeDetails interface {
}

type cancelsnoozedeploymentupgradedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *cancelsnoozedeploymentupgradedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercancelsnoozedeploymentupgradedetails cancelsnoozedeploymentupgradedetails
	s := struct {
		Model Unmarshalercancelsnoozedeploymentupgradedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *cancelsnoozedeploymentupgradedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "DEFAULT":
		mm := DefaultCancelSnoozeDeploymentUpgradeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CancelSnoozeDeploymentUpgradeDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m cancelsnoozedeploymentupgradedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m cancelsnoozedeploymentupgradedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
