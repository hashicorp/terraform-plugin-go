// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var timestamp = "2024-08-16T16:56:57Z"

func testGoTime() time.Time {
	rfc3339, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return time.Time{}
	}
	return rfc3339
}

func testPbTimestamp() *timestamppb.Timestamp {
	return timestamppb.New(testGoTime())
}
