// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

// TODO: pkg docs for this whole file
type ActionSchema struct {
	Schema *Schema
	Type   ActionSchemaType
}

type ActionSchemaType interface {
	isActionSchemaType()
}

var (
	_ ActionSchemaType = UnlinkedActionSchemaType{}
	_ ActionSchemaType = LifecycleActionSchemaType{}
	_ ActionSchemaType = LinkedActionSchemaType{}
)

type UnlinkedActionSchemaType struct{}

func (a UnlinkedActionSchemaType) isActionSchemaType() {}

type LifecycleActionSchemaType struct {
	Executes       LifecycleExecutionOrder
	LinkedResource *LinkedResourceSchema
}

func (a LifecycleActionSchemaType) isActionSchemaType() {}

const (
	LifecycleExecutionOrderInvalid LifecycleExecutionOrder = 0
	LifecycleExecutionOrderBefore  LifecycleExecutionOrder = 1
	LifecycleExecutionOrderAfter   LifecycleExecutionOrder = 2
)

type LifecycleExecutionOrder int32

func (l LifecycleExecutionOrder) String() string {
	switch l {
	case 0:
		return "INVALID"
	case 1:
		return "BEFORE"
	case 2:
		return "AFTER"
	case 3:
	}
	return "UNKNOWN"
}

type LinkedResourceSchema struct {
	TypeName    string
	Description string
}

type LinkedActionSchemaType struct {
	LinkedResources []*LinkedResourceSchema
}

func (a LinkedActionSchemaType) isActionSchemaType() {}
