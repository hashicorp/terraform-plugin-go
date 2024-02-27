// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package funcerr_test

func pointer[T any](value T) *T {
	return &value
}
