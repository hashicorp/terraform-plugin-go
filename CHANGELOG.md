# 0.4.0 (September 24, 2021)

BREAKING CHANGES:
* The `AttributeType` property of `tftypes.Map` has been renamed to `ElementType`. ([#105](https://github.com/hashicorp/terraform-plugin-go/issues/105))
* The `tf5server` package's import path is now `github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server` instead of `github.com/hashicorp/terraform-plugin-go/tfprotov5/server`. ([#103](https://github.com/hashicorp/terraform-plugin-go/issues/103))
* The `tf6server` package's import path is now `github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server` instead of `github.com/hashicorp/terraform-plugin-go/tfprotov6/server`. ([#103](https://github.com/hashicorp/terraform-plugin-go/issues/103))
* With the release of Go 1.17, Go 1.16 is now the lowest supported version of Go to use with terraform-plugin-go. ([#102](https://github.com/hashicorp/terraform-plugin-go/issues/102))
* tftypes: The `Is()` method on types implementing the `Type` interface now only perform a shallow type check, no longer comparing any underlying attribute or element type(s). Use the `Equal()` method for deep type checking. ([#94](https://github.com/hashicorp/terraform-plugin-go/issues/94))
* tftypes: `(AttributePath).WithElementKeyInt()` now has an `int` parameter instead of `int64`. Using `(AttributePath).WithElementKeyInt()` within `for ... range` loops no longer requires `int64()` conversion. ([#101](https://github.com/hashicorp/terraform-plugin-go/issues/101))

ENHANCEMENTS:
* tftypes: All types implementing the `Type` interface now provide an `Equal()` method for deep type checking and `UsableAs()` method for type conformance checking. ([#94](https://github.com/hashicorp/terraform-plugin-go/issues/94))

BUG FIXES:
* tftypes: Ensure `NewValue()` panics on `DynamicPseudoType` and known values. ([#94](https://github.com/hashicorp/terraform-plugin-go/issues/94))
* tftypes: Fixed elements of `Tuple` and `Map` and attributes of `Object` having `DynamicPseudoType` as their type when unmarshaling JSON values from Terraform. ([#94](https://github.com/hashicorp/terraform-plugin-go/issues/94))
* tftypes: Fixed elements of `Tuple` and `Map` and attributes of `Object` having `DynamicPseudoType` as their type when unmarshaling msgpack values from Terraform. ([#100](https://github.com/hashicorp/terraform-plugin-go/issues/100))
* tftypes: Prevent potential panic unmarshaling null DynamicPseudoType in msgpack ([#99](https://github.com/hashicorp/terraform-plugin-go/issues/99))
* tftypes: Return error instead of panic when calling `(Value).Diff()` with either `Value` missing type ([#104](https://github.com/hashicorp/terraform-plugin-go/issues/104))
* tftypes: Return error instead of panic when calling `Transform()` with `Value` missing type ([#104](https://github.com/hashicorp/terraform-plugin-go/issues/104))

# 0.3.1 (June 24, 2021)

BUG FIXES:
* Fixed AttributePaths pointing to the root of the value to be omitted instead of prefixing the error with `: `. ([#87](https://github.com/hashicorp/terraform-plugin-go/issues/87))
* Fixed a panic when `.String()` is called on an empty `tftypes.Value`. ([#86](https://github.com/hashicorp/terraform-plugin-go/issues/86))
* Fixed a panic when calling `tftypes.Value.As` and passing a pointer to an uninstantiated *big.Float. ([#85](https://github.com/hashicorp/terraform-plugin-go/issues/85))
* Fixed a panic when calling `tftypes.Value.As` and passing a pointer to an uninstantiated *bool, *string, *map[string]Value or *[]Value. ([#88](https://github.com/hashicorp/terraform-plugin-go/issues/88))
* Fixed a panic when comparing the empty value of a tftypes.Value to a non-empty value. ([#90](https://github.com/hashicorp/terraform-plugin-go/issues/90))

# 0.3.0 (April 21, 2021)

BREAKING CHANGES:
* Previously, `tftypes.NewValue` would panic if the Go type supplied wasn't a valid Go type for _any_ `tftypes.Type`. Now `tftypes.NewValue` will panic if the Go type supplied isn't a valid Go type for the _specific_ `tftypes.Type` supplied. ([#67](https://github.com/hashicorp/terraform-plugin-go/issues/67))
* Removed support for `*float32` (and `float32`, which was only documented and never implemented) when creating a `tftypes.Number` using `tftypes.NewValue`. We can't find a lossless way to convert a `float32` to a `*big.Float` and so require provider developers to choose the lossy conversion they find acceptable. ([#67](https://github.com/hashicorp/terraform-plugin-go/issues/67))
* Removed the now-unnecessary `tftypes.ValueComparer` helper, which helped `github.com/google/go-cmp` compare `tftypes.Value`s. `tftypes.Value`s now have an `Equal` method that `go-cmp` can use, and don't need any special options passed anymore. ([#67](https://github.com/hashicorp/terraform-plugin-go/issues/67))
* The `tftypes` package has been moved to the root of the module and is no longer under the `tfprotov5` package. Providers can automatically rewrite their import paths using a command like `sed -i 's/"github.com\/hashicorp\/terraform-plugin-go\/tfprotov5\/tftypes"/"github.com\/hashicorp\/terraform-plugin-go\/tftypes"/g' **/*.go` on Unix-like systems. ([#70](https://github.com/hashicorp/terraform-plugin-go/issues/70))
* With the release of Go 1.16, Go 1.15 is now the lowest supported version of Go to use with terraform-plugin-go. ([#62](https://github.com/hashicorp/terraform-plugin-go/issues/62))
* `tftypes.AttributePath` is now referenced as a pointer instead of a value pretty much everywhere it is used. This enables much more ergonomic use with `tfprotov5.Diagnostic` values. ([#68](https://github.com/hashicorp/terraform-plugin-go/issues/68))
* `tftypes.AttributePath`'s `Steps` property is now internal-only. Use `tftypes.AttributePath.Steps()` to access the list of `tftypes.AttributePathSteps`, and `tftypes.NewAttributePath` or `tftypes.NewAttributePathWithSteps` to create a new `tftypes.AttributePath`. ([#68](https://github.com/hashicorp/terraform-plugin-go/issues/68))
* `tftypes.String`, `tftypes.Number`, `tftypes.Bool`, and `tftypes.DynamicPseudoType` are now represented by a different Go type. Uses of `==` and `switch` on them will no longer work. The recommended way to compare any type is using `Is`. ([#58](https://github.com/hashicorp/terraform-plugin-go/issues/58))
* `tftypes.Value`s no longer have an `Is` method. Use `tftypes.Value.Type().Is` instead. ([#58](https://github.com/hashicorp/terraform-plugin-go/issues/58))
* tftypes.AttributePath.WithAttributeName, WithElementKeyString, WithElementKeyInt, and WithElementKeyValue no longer accept pointers and mutate the AttributePath. They now copy the AttributePath, and return a version of it with the new AttributePathStep appended. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))

FEATURES:
* Added tftypes.Diff function to return the elements and attributes that are different between two tftypes.Values. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* Added tftypes.Walk and tftypes.Transform functions for the tftypes.Value type, allowing providers to traverse and mutate a tftypes.Value, respectively. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* `tftypes.Value`s now have a `Type` method, exposing their `tftypes.Type`. ([#58](https://github.com/hashicorp/terraform-plugin-go/issues/58))

ENHANCEMENTS:
* A number of methods in `tftypes` are benefitting from a better error message for `tftypes.AttributePathError`s, which are returned in various places, and will now surface the path associated with the error as part of the error message. ([#68](https://github.com/hashicorp/terraform-plugin-go/issues/68))
* Added Equal method to tftypes.Type implementations, allowing them to be compared using github.com/google/go-cmp. ([#74](https://github.com/hashicorp/terraform-plugin-go/issues/74))
* Added a Copy method to tftypes.Value, returning a clone of the tftypes.Value such that modifying the clone is guaranteed to not modify the original. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* Added a String method to tftypes.AttributePath to return a string representation of the tftypes.AttributePath. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* Added a String method to tftypes.Value, returning a string representation of the tftypes.Value. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* Added a `tftypes.ValidateValue` function that returns an error if the combination of the `tftypes.Type` and Go type passed when panic when passed to `tftypes.NewValue`. ([#67](https://github.com/hashicorp/terraform-plugin-go/issues/67))
* Added an Equal method to tftypes.AttributePath to compares two tftypes.AttributePaths. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* Added an Equal method to tftypes.Value to compare two tftypes.Values. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* Added support for OptionalAttributes to tftypes.Objects, allowing for objects with attributes that can be set or can be omitted. See https://www.terraform.io/docs/language/expressions/type-constraints.html#experimental-optional-object-type-attributes for more information on optional attributes in objects. ([#74](https://github.com/hashicorp/terraform-plugin-go/issues/74))
* Added support for `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `int`, `int8`, `int16`, `int32`, `int64`, and `float64` conversions when creating a `tftypes.Number` with `tftypes.NewValue`. These were mistakenly omitted previously. ([#67](https://github.com/hashicorp/terraform-plugin-go/issues/67))
* Added support for version 6 of the Terraform protocol, in a new tfprotov6 package. ([#71](https://github.com/hashicorp/terraform-plugin-go/issues/71))
* Updated the String method of all tftypes.Type implementations to include any element or attribute types in the string as well. ([#60](https://github.com/hashicorp/terraform-plugin-go/issues/60))
* `tftypes.AttributePathError` is now exported. Provider developers can use `errors.Is` and `errors.As` to check for `tftypes.AttributePathError`s, `errors.Unwrap` to get to the underlying error, and the `Path` property on a `tftypes.AttributePathError` to access the `tftypes.AttributePath` the error is associated with. `tftypes.AttributePath.NewError` and `tftypes.AttributePath.NewErrorf` are still the supported ways to create a `tftypes.AttributePathError`. ([#68](https://github.com/hashicorp/terraform-plugin-go/issues/68))

BUG FIXES:
* Fixed a bug in `tftypes.Value.IsFullyKnown` that would cause a panic when calling `IsFullyKnown` on `tftypes.Value` with a `tftypes.Type` of Map, Object, List, Set, or Tuple if the `tftypes.Value` was null. ([#69](https://github.com/hashicorp/terraform-plugin-go/issues/69))
* Fixed a bug where `*uint8`, `*uint16`, and `*uint32` would be coerced to `int64`s as part of their conversion in `tftypes.NewValue`. This may have had no impact, as all those types can be represented in an `int64`, but to be sure our conversion is accurate, the conversion was fixed to convert them to a `uint64` instead. ([#67](https://github.com/hashicorp/terraform-plugin-go/issues/67))

# 0.2.1 (January 07, 2021)

BUG FIXES:

* Fixed a bug that could cause a crash when a provider was prematurely stopped. ([#49](https://github.com/hashicorp/terraform-plugin-go/issues/49))

# 0.2.0 (November 20, 2020)

ENHANCEMENTS:

* `tftypes.NewValue` can now accept a wider array of standard library types which will be automatically converted to their standard representation ([#46](https://github.com/hashicorp/terraform-plugin-go/issues/46)] [[#47](https://github.com/hashicorp/terraform-plugin-go/issues/47))
* `tfprotov5.RawState` now has an `Unmarshal` method, just like `tfprotov5.DynamicValue`, yielding a `tftypes.Value`. ([#42](https://github.com/hashicorp/terraform-plugin-go/issues/42))

# 0.1.0 (November 02, 2020)

Initial release.
