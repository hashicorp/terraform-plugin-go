package tftypes_test

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func ExampleValue_As_string() {
	// Values come over the wire, usually from a DynamicValue for this
	// example, we're just building one inline
	val := tftypes.NewValue(tftypes.String, "hello, world")

	var salutation string

	// we need to use a pointer so we can modify the value, just like
	// json.Unmarshal
	err := val.As(&salutation)
	if err != nil {
		panic(err)
	}

	fmt.Println(salutation)
	// Output:
	// hello, world
}

func ExampleValue_As_stringNull() {
	type exampleResource struct {
		salutation         string
		nullableSalutation *string
	}

	// let's see what happens when we have a null value
	val := tftypes.NewValue(tftypes.String, nil)

	var res exampleResource

	// we can use a pointer to a variable, but the variable can't hold nil,
	// so we'll get the empty value. You can use this if you don't care
	// about null, and consider it equivalent to the empty value.
	err := val.As(&res.salutation)
	if err != nil {
		panic(err)
	}

	// we can use a pointer to a pointer to a variable, which can hold nil,
	// so we'll be able to distinguish between a null and an empty string
	err = val.As(&res.nullableSalutation)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.salutation)
	fmt.Println(res.nullableSalutation)
	// Output:
	//
	// <nil>
}

type exampleResource struct {
	name         string
	suppliedName *bool // true for yes, false for no, nil for we haven't asked
}

// fill the tftypes.ValueConverter interface to control how As works
// we want a pointer to exampleResource so we can change the properties
func (e *exampleResource) FromTerraform5Value(val tftypes.Value) error {
	// this is an object type, so we're always going to get a
	// `tftypes.Value` that coerces to a map[string]tftypes.Value
	// as input
	v := map[string]tftypes.Value{}
	err := val.As(&v)
	if err != nil {
		return err
	}

	// now that we can get to the tftypes.Value for each field,
	// call its As method and assign the result to the appropriate
	// variable.

	err = v["name"].As(&e.name)
	if err != nil {
		return err
	}

	err = v["supplied_name"].As(&e.suppliedName)
	if err != nil {
		return err
	}

	return nil
}

func ExampleValue_As_interface() {
	// our tftypes.Value would usually come over the wire as a
	// DynamicValue, but for simplicity, let's just declare one inline here
	val := tftypes.NewValue(tftypes.Object{
		AttributeTypes: map[string]tftypes.Type{
			"name":          tftypes.String,
			"supplied_name": tftypes.Bool,
		},
	}, map[string]tftypes.Value{
		"name":          tftypes.NewValue(tftypes.String, "ozymandias"),
		"supplied_name": tftypes.NewValue(tftypes.Bool, nil),
	})

	// exampleResource has FromTerraform5Value method defined on it, see
	// value_example_test.go for implementation details. We'd put the
	// function and type inline here, but apparently Go can't have methods
	// defined on types defined inside a function
	var res exampleResource

	// call As as usual
	err := val.As(&res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.name)
	fmt.Println(res.suppliedName)
	// Output:
	// ozymandias
	// <nil>
}
