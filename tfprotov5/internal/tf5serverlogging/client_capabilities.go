// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5serverlogging

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// ConfigureProviderClientCapabilities generates a TRACE "Announced client capabilities" log.
func ConfigureProviderClientCapabilities(ctx context.Context, capabilities *tfprotov5.ConfigureProviderClientCapabilities) {
	responseFields := map[string]interface{}{
		logging.KeyClientCapabilityDeferralAllowed: false,
	}

	if capabilities != nil {
		responseFields[logging.KeyClientCapabilityDeferralAllowed] = capabilities.DeferralAllowed
	}

	logging.ProtocolTrace(ctx, "Announced client capabilities", responseFields)
}

// ReadDataSourceClientCapabilities generates a TRACE "Announced client capabilities" log.
func ReadDataSourceClientCapabilities(ctx context.Context, capabilities *tfprotov5.ReadDataSourceClientCapabilities) {
	responseFields := map[string]interface{}{
		logging.KeyClientCapabilityDeferralAllowed: false,
	}

	if capabilities != nil {
		responseFields[logging.KeyClientCapabilityDeferralAllowed] = capabilities.DeferralAllowed
	}

	logging.ProtocolTrace(ctx, "Announced client capabilities", responseFields)
}

// ReadResourceClientCapabilities generates a TRACE "Announced client capabilities" log.
func ReadResourceClientCapabilities(ctx context.Context, capabilities *tfprotov5.ReadResourceClientCapabilities) {
	responseFields := map[string]interface{}{
		logging.KeyClientCapabilityDeferralAllowed: false,
	}

	if capabilities != nil {
		responseFields[logging.KeyClientCapabilityDeferralAllowed] = capabilities.DeferralAllowed
	}

	logging.ProtocolTrace(ctx, "Announced client capabilities", responseFields)
}

// PlanResourceChangeClientCapabilities generates a TRACE "Announced client capabilities" log.
func PlanResourceChangeClientCapabilities(ctx context.Context, capabilities *tfprotov5.PlanResourceChangeClientCapabilities) {
	responseFields := map[string]interface{}{
		logging.KeyClientCapabilityDeferralAllowed: false,
	}

	if capabilities != nil {
		responseFields[logging.KeyClientCapabilityDeferralAllowed] = capabilities.DeferralAllowed
	}

	logging.ProtocolTrace(ctx, "Announced client capabilities", responseFields)
}

// ImportResourceStateClientCapabilities generates a TRACE "Announced client capabilities" log.
func ImportResourceStateClientCapabilities(ctx context.Context, capabilities *tfprotov5.ImportResourceStateClientCapabilities) {
	responseFields := map[string]interface{}{
		logging.KeyClientCapabilityDeferralAllowed: false,
	}

	if capabilities != nil {
		responseFields[logging.KeyClientCapabilityDeferralAllowed] = capabilities.DeferralAllowed
	}

	logging.ProtocolTrace(ctx, "Announced client capabilities", responseFields)
}
