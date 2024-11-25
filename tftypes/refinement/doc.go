// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// The refinement package contains the interfaces and structs that represent unknown value refinement data. Refinements contain
// additional constraints about unknown values and what their eventual known values can be. In certain scenarios, Terraform can
// use these constraints to produce known results from unknown values. (like evaluating a count expression comparing an unknown
// value to "null")
//
// Unknown value refinements can be added to a `tftypes.Value` via the `(tftypes.Value).Refine` method. Refinements on an unknown
// `tftypes.Value` can be retrieved via the `(tftypes.Value).Refinements()` method.
package refinement
