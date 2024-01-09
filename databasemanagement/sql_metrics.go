// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlMetrics Metrics of the Sql in the Sql tuning set.
type SqlMetrics struct {

	// Total CPU time consumed by the Sql.
	CpuTime *int64 `mandatory:"false" json:"cpuTime"`

	// Elapsed time of the Sql.
	ElapsedTime *int64 `mandatory:"false" json:"elapsedTime"`

	// Sum total number of buffer gets.
	BufferGets *int64 `mandatory:"false" json:"bufferGets"`

	// Sum total number of disk reads.
	DiskReads *int64 `mandatory:"false" json:"diskReads"`

	// Sum total number of direct path writes.
	DirectWrites *int64 `mandatory:"false" json:"directWrites"`

	// Total executions of this SQL statement.
	Executions *int64 `mandatory:"false" json:"executions"`
}

func (m SqlMetrics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlMetrics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
