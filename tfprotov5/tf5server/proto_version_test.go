// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5server

import (
	"bufio"
	"os"
	"regexp"
	"testing"
)

// MAINTAINER NOTE: This test is a best effort for ensuring that the protocol version variables in the tf5server package
// stay in sync with the actual protocol file.
func Test_EnsureVersionConstantMatchesProtoFile(t *testing.T) {
	t.Parallel()

	file, err := os.Open("../internal/tfplugin5/tfplugin5.proto")
	if err != nil {
		t.Fatalf("error opening proto file: %s", err)
	}
	defer file.Close()

	protoFileComment := regexp.MustCompile(`(?:Terraform Plugin RPC protocol version )(\d+.\d+)+`)

	var expectedProtocolVersion string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := protoFileComment.FindStringSubmatch(line)

		if len(matches) > 1 {
			expectedProtocolVersion = matches[1]
			break
		}
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("error scanning proto file: %s", err)
	}

	if expectedProtocolVersion == "" {
		t.Fatalf("couldn't find version comment in proto file: %s", err)
	}

	if protocolVersion != expectedProtocolVersion {
		t.Errorf("protocol version Go variable is different from proto file - expected: %s, got: %s\n", expectedProtocolVersion, protocolVersion)
		t.Log("MAINTAINER NOTE: Update tf5server.protocolVersionMajor and tf5server.protocolVersionMinor to match the proto file.")
	}
}
