// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package funcerr_test

func pointer[T any](value T) *T {
	return &value
}
