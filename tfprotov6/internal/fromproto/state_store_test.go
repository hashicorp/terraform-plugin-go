// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestValidateStateStoreConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateStateStoreConfig_Request
		expected *tfprotov6.ValidateStateStoreConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateStateStoreConfig_Request{},
			expected: &tfprotov6.ValidateStateStoreConfigRequest{},
		},
		"Config": {
			in: &tfplugin6.ValidateStateStoreConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateStateStoreConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ValidateStateStoreConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ValidateStateStoreConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateStateStoreConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestConfigureStateStoreRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ConfigureStateStore_Request
		expected *tfprotov6.ConfigureStateStoreRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ConfigureStateStore_Request{},
			expected: &tfprotov6.ConfigureStateStoreRequest{},
		},
		"Config": {
			in: &tfplugin6.ConfigureStateStore_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ConfigureStateStoreRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ConfigureStateStore_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ConfigureStateStoreRequest{
				TypeName: "test",
			},
		},
		"Capabilities": {
			in: &tfplugin6.ConfigureStateStore_Request{
				Capabilities: &tfplugin6.StateStoreClientCapabilities{
					ChunkSize: 8 << 20,
				},
			},
			expected: &tfprotov6.ConfigureStateStoreRequest{
				Capabilities: &tfprotov6.ConfigureStateStoreClientCapabilities{
					ChunkSize: 8 << 20,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ConfigureStateStoreRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestReadStateBytesRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ReadStateBytes_Request
		expected *tfprotov6.ReadStateBytesRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ReadStateBytes_Request{},
			expected: &tfprotov6.ReadStateBytesRequest{},
		},
		"TypeName": {
			in: &tfplugin6.ReadStateBytes_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ReadStateBytesRequest{
				TypeName: "test",
			},
		},
		"StateID": {
			in: &tfplugin6.ReadStateBytes_Request{
				StateId: "test-id",
			},
			expected: &tfprotov6.ReadStateBytesRequest{
				StateID: "test-id",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ReadStateBytesRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetStatesRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.GetStates_Request
		expected *tfprotov6.GetStatesRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.GetStates_Request{},
			expected: &tfprotov6.GetStatesRequest{},
		},
		"TypeName": {
			in: &tfplugin6.GetStates_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.GetStatesRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.GetStatesRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestDeleteStateRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.DeleteState_Request
		expected *tfprotov6.DeleteStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.DeleteState_Request{},
			expected: &tfprotov6.DeleteStateRequest{},
		},
		"TypeName": {
			in: &tfplugin6.DeleteState_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.DeleteStateRequest{
				TypeName: "test",
			},
		},
		"StateID": {
			in: &tfplugin6.DeleteState_Request{
				StateId: "test-id",
			},
			expected: &tfprotov6.DeleteStateRequest{
				StateID: "test-id",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.DeleteStateRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestLockStateRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.LockState_Request
		expected *tfprotov6.LockStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.LockState_Request{},
			expected: &tfprotov6.LockStateRequest{},
		},
		"TypeName": {
			in: &tfplugin6.LockState_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.LockStateRequest{
				TypeName: "test",
			},
		},
		"StateID": {
			in: &tfplugin6.LockState_Request{
				StateId: "test-id",
			},
			expected: &tfprotov6.LockStateRequest{
				StateID: "test-id",
			},
		},
		"Operation": {
			in: &tfplugin6.LockState_Request{
				Operation: "apply",
			},
			expected: &tfprotov6.LockStateRequest{
				Operation: "apply",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.LockStateRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestUnlockStateRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.UnlockState_Request
		expected *tfprotov6.UnlockStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.UnlockState_Request{},
			expected: &tfprotov6.UnlockStateRequest{},
		},
		"TypeName": {
			in: &tfplugin6.UnlockState_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.UnlockStateRequest{
				TypeName: "test",
			},
		},
		"StateID": {
			in: &tfplugin6.UnlockState_Request{
				StateId: "test-id",
			},
			expected: &tfprotov6.UnlockStateRequest{
				StateID: "test-id",
			},
		},
		"LockID": {
			in: &tfplugin6.UnlockState_Request{
				LockId: "test-lock",
			},
			expected: &tfprotov6.UnlockStateRequest{
				LockID: "test-lock",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.UnlockStateRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestWriteStateBytesChunk(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in           *tfplugin6.WriteStateBytes_RequestChunk
		expected     *tfprotov6.WriteStateBytesChunk
		expectedDiag *tfprotov6.Diagnostic
	}{
		"nil": {
			in: nil,
			expectedDiag: &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unexpected empty state chunk in WriteStateBytes",
				Detail:   "An empty state byte chunk was received. This is a bug in Terraform that should be reported to the maintainers.",
			},
		},
		"zero": {
			in: &tfplugin6.WriteStateBytes_RequestChunk{},
			expectedDiag: &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unexpected state chunk data received in WriteStateBytes",
				Detail: "An invalid state byte chunk was received with no range start/end information. " +
					"This is a bug in Terraform that should be reported to the maintainers.",
			},
		},
		"Chunk-no-Range": {
			in: &tfplugin6.WriteStateBytes_RequestChunk{
				Bytes:       []byte(`{"test": "data"}`),
				TotalLength: 16,
			},
			expectedDiag: &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "Unexpected state chunk data received in WriteStateBytes",
				Detail:   "An invalid state byte chunk was received with no range start/end information. This is a bug in Terraform that should be reported to the maintainers.",
			},
		},
		"Chunk-with-Meta": {
			in: &tfplugin6.WriteStateBytes_RequestChunk{
				Meta: &tfplugin6.RequestChunkMeta{
					TypeName: "test",
					StateId:  "test-state-id",
				},
				Bytes:       []byte(`{"test": "data"}`),
				TotalLength: 16,
				Range: &tfplugin6.StateByteRange{
					Start: 0,
					End:   16,
				},
			},
			expected: &tfprotov6.WriteStateBytesChunk{
				Meta: &tfprotov6.WriteStateChunkMeta{
					TypeName: "test",
					StateID:  "test-state-id",
				},
				StateByteChunk: tfprotov6.StateByteChunk{
					Bytes:       []byte(`{"test": "data"}`),
					TotalLength: 16,
					Range: tfprotov6.StateByteRange{
						Start: 0,
						End:   16,
					},
				},
			},
		},
		"Chunk-without-Meta": {
			in: &tfplugin6.WriteStateBytes_RequestChunk{
				Bytes:       []byte(`{"test": "data"}`),
				TotalLength: 16,
				Range: &tfplugin6.StateByteRange{
					Start: 0,
					End:   16,
				},
			},
			expected: &tfprotov6.WriteStateBytesChunk{
				StateByteChunk: tfprotov6.StateByteChunk{
					Bytes:       []byte(`{"test": "data"}`),
					TotalLength: 16,
					Range: tfprotov6.StateByteRange{
						Start: 0,
						End:   16,
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, diag := fromproto.WriteStateBytesChunk(testCase.in)
			if diag != nil {
				if testCase.expectedDiag == nil {
					t.Fatalf("wanted no diagnostic, got diagnostic: %v", diag)
				}

				if diff := cmp.Diff(diag, testCase.expectedDiag); diff != "" {
					t.Fatalf("unexpected difference: %s", diff)
				}
			}

			if diag == nil && testCase.expectedDiag != nil {
				t.Fatalf("got no diagnostic, wanted diagnostic: %v", testCase.expectedDiag)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
