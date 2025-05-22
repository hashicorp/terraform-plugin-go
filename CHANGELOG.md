## 0.28.0 (May 21, 2025)

BREAKING CHANGES:

* tfprotov5+tfprotov6: Removed temporary `ProviderServerWithResourceIdentity` interface type. Use `ProviderServer` instead. ([#516](https://github.com/hashicorp/terraform-plugin-go/issues/516))
* tfprotov5+tfprotov6: `GetResourceIdentitySchemas` and `UpgradeResourceIdentity` RPC calls are now required in `ProviderServer` and `ResourceServer`. Implementations that don't support resource identity can return empty responses from the `GetResourceIdentitySchemas` method and an error message the `UpgradeResourceIdentity` method. ([#516](https://github.com/hashicorp/terraform-plugin-go/issues/516))

NOTES:

* all: To prevent compilation errors, ensure your Go module is updated to at least terraform-plugin-framework@v1.15.0, terraform-plugin-mux@v0.20.0, terraform-plugin-sdk/v2@v2.37.0, and terraform-plugin-testing@v1.13.0 before upgrading this dependency. ([#516](https://github.com/hashicorp/terraform-plugin-go/issues/516))

## 0.27.0 (May 16, 2025)

NOTES:

* tfprotov5+tfprotov6: An upcoming release will require the `GetResourceIdentitySchemas` and `UpgradeResourceIdentity` implementations as part of `ProviderServer`. ([#476](https://github.com/hashicorp/terraform-plugin-go/issues/476))
* all: This Go module has been updated to Go 1.23 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.23 release notes](https://go.dev/doc/go1.23) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#491](https://github.com/hashicorp/terraform-plugin-go/issues/491))

FEATURES:

* tfprotov5+tfprotov6: Upgraded protocols and added types to support the new resource identity feature ([#476](https://github.com/hashicorp/terraform-plugin-go/issues/476))
* Add WithManagedDebugEnvFilePath() option that allows writing TF_REATTACH_PROVIDERS to an environment file ([#484](https://github.com/hashicorp/terraform-plugin-go/issues/484))

## 0.27.0-alpha.1 (March 18, 2025)

NOTES:

* This alpha pre-release contains the protocol definitions for managed resource identity, which can used with Terraform v1.12.0-alpha20250312, to store and read identity data during plan and apply workflows. ([#476](https://github.com/hashicorp/terraform-plugin-go/issues/476))
* all: This Go module has been updated to Go 1.23 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.23 release notes](https://go.dev/doc/go1.23) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#491](https://github.com/hashicorp/terraform-plugin-go/issues/491))

## 0.26.0 (January 22, 2025)

BREAKING CHANGES:

* tfprotov5+tfprotov6: `EphemeralResourceServer` interface is now required in `ProviderServer`. Implementations not needing ephemeral resource support can return errors from the `*EphemeralResource` methods. ([#442](https://github.com/hashicorp/terraform-plugin-go/issues/442))
* tfprotov5+tfprotov6: Removed temporary `ProviderServerWithEphemeralResources` interface type. Use `EphemeralResourceServer` instead. ([#442](https://github.com/hashicorp/terraform-plugin-go/issues/442))

NOTES:

* all: To prevent compilation errors, ensure your Go module is updated to at least terraform-plugin-framework@v1.13.0, terraform-plugin-mux@v0.17.0, terraform-plugin-sdk/v2@v2.35.0, and terraform-plugin-testing@v1.11.0 before upgrading this dependency. ([#442](https://github.com/hashicorp/terraform-plugin-go/issues/442))

FEATURES:

* tfprotov5+tfprotov6: Upgraded protocols and added types to support write-only attributes ([#462](https://github.com/hashicorp/terraform-plugin-go/issues/462))

## 0.25.0 (October 30, 2024)

NOTES:

* tfprotov5+tfprotov6: An upcoming release will require the `EphemeralResourceServer` implementation as part of `ProviderServer`.  ([#441](https://github.com/hashicorp/terraform-plugin-go/issues/441))

FEATURES:

* tfprotov5+tfprotov6: Upgraded protocols and added types to support the new ephemeral resource type ([#441](https://github.com/hashicorp/terraform-plugin-go/issues/441))

## 0.24.0 (September 18, 2024)

BREAKING CHANGES:

* tfprotov5+tfprotov6: Removed temporary `ResourceServerWithMoveResourceState` interface type. Use `ResourceServer` instead. ([#408](https://github.com/hashicorp/terraform-plugin-go/issues/408))

NOTES:

* all: If using terraform-plugin-mux, it must be upgraded to v0.16.0 or later to prevent compilation errors ([#408](https://github.com/hashicorp/terraform-plugin-go/issues/408))
* all: This Go module has been updated to Go 1.22 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.22 release notes](https://go.dev/doc/go1.22) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#428](https://github.com/hashicorp/terraform-plugin-go/issues/428))

BUG FIXES:

* logging: Add existing `MoveResourceState` server capability to v5 and v6 logging ([#432](https://github.com/hashicorp/terraform-plugin-go/issues/432))

## 0.23.0 (May 06, 2024)

BREAKING CHANGES:

* tfprotov5+tfprotov6: `FunctionServer` interface is now required in `ProviderServer`. Implementations not needing function support can return errors from the `GetFunctions` and `CallFunction` methods. ([#388](https://github.com/hashicorp/terraform-plugin-go/issues/388))
* tfprotov5+tfprotov6: `MoveResourceState` method is now required in `ResourceServer`. Implementations not needing move state support can return errors from the `MoveResourceState` method. ([#388](https://github.com/hashicorp/terraform-plugin-go/issues/388))

NOTES:

* all: To prevent compilation errors, ensure your Go module is updated to at least terraform-plugin-framework@v1.6.0, terraform-plugin-mux@v0.15.0, terraform-plugin-sdk/v2@v2.33.0, and terraform-plugin-testing@v1.7.0 before upgrading this dependency. ([#388](https://github.com/hashicorp/terraform-plugin-go/issues/388))

FEATURES:

* tfprotov5+tfprotov6: Upgraded protocols and added types to support deferred actions ([#403](https://github.com/hashicorp/terraform-plugin-go/issues/403))

## 0.22.2 (April 18, 2024)

BUG FIXES:

* tftypes: Large integers are always encoded as msgpack integers rather than float values to ensure the decoded value will not be rounded to 52-bit precision ([#396](https://github.com/hashicorp/terraform-plugin-go/issues/396))

## 0.22.1 (March 11, 2024)

NOTES:

* all: Upgraded google.golang.org/protobuf dependency, which includes a fix for CVE-2024-24786. While the fixed feature is not used in this Go module, this upgrade may silence security vulnerability scanner reports that are based purely on dependency versions. ([#385](https://github.com/hashicorp/terraform-plugin-go/issues/385))

BUG FIXES:

* tftypes: Fixed an edge-case where `(Value).Equal` would panic when comparing two values with underlying `DynamicPseudoType` types and different concrete values. ([#383](https://github.com/hashicorp/terraform-plugin-go/issues/383))

## 0.22.0 (February 22, 2024)

BREAKING CHANGES:

* tfprotov5+tfprotov6: Modified the response returned from the CallFunction RPC, replacing diagnostics with function error ([#380](https://github.com/hashicorp/terraform-plugin-go/issues/380))

NOTES:

* all: This Go module has been updated to Go 1.21 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.21 release notes](https://go.dev/doc/go1.21) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#381](https://github.com/hashicorp/terraform-plugin-go/issues/381))
* all: If using terraform-plugin-framework, terraform-plugin-mux, or terraform-plugin-sdk, only upgrade this Go module when upgrading those Go modules to terraform-plugin-framework@v1.6.0, terraform-plugin-mux@v0.15.0, and terraform-plugin-sdk/v2@v2.33.0, or greater, respectively ([#380](https://github.com/hashicorp/terraform-plugin-go/issues/380))

## 0.21.0 (January 26, 2024)

NOTES:

* tfprotov5+tfprotov6: An upcoming release will require the MoveResourceState implementation as part of ResourceServer ([#364](https://github.com/hashicorp/terraform-plugin-go/issues/364))

FEATURES:

* tfprotov5+tfprotov6: Upgraded protocols and added types to support the `MoveResourceState` RPC ([#364](https://github.com/hashicorp/terraform-plugin-go/issues/364))

## 0.20.0 (December 14, 2023)

NOTES:

* tfprotov5+tfprotov6: An upcoming release will require the FunctionServer implementation as part of ProviderServer. ([#351](https://github.com/hashicorp/terraform-plugin-go/issues/351))
* Provider-defined function support is in technical preview and offered without compatibility promises until Terraform 1.8 is generally available. ([#351](https://github.com/hashicorp/terraform-plugin-go/issues/351))

FEATURES:

* tfprotov5+tfprotov6: Upgraded protocols and added types to support provider-defined functions ([#351](https://github.com/hashicorp/terraform-plugin-go/issues/351))

## 0.19.1 (November 15, 2023)

NOTES:

* This release introduces no functional changes. It does however include dependency updates, such as `google.golang.org/grpc`, which address upstream CVEs. ([#324](https://github.com/hashicorp/terraform-plugin-go/issues/324))

## 0.19.0 (September 06, 2023)

NOTES:

* all: If using terraform-plugin-framework, terraform-plugin-mux, or terraform-plugin-sdk, only upgrade this Go module when upgrading those Go modules or you may receive a `missing GetMetadata method` error when compiling ([#310](https://github.com/hashicorp/terraform-plugin-go/issues/310))
* all: This Go module has been updated to Go 1.20 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.20 release notes](https://go.dev/doc/go1.20) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#321](https://github.com/hashicorp/terraform-plugin-go/issues/321))

FEATURES:

* tfprotov5: Upgraded protocol to 5.4 and implemented `GetMetadata` RPC ([#310](https://github.com/hashicorp/terraform-plugin-go/issues/310))
* tfprotov6: Upgraded protocol to 6.4 and implemented `GetMetadata` RPC ([#310](https://github.com/hashicorp/terraform-plugin-go/issues/310))

ENHANCEMENTS:

* tfprotov5: Added `ServerCapabilities` type `GetProviderSchemaOptional` field, which when enabled can signal that the provider supports RPC operations without the `GetProviderSchema` RPC being called first ([#310](https://github.com/hashicorp/terraform-plugin-go/issues/310))
* tfprotov6: Added `ServerCapabilities` type `GetProviderSchemaOptional` field, which when enabled can signal that the provider supports RPC operations without the `GetProviderSchema` RPC being called first ([#310](https://github.com/hashicorp/terraform-plugin-go/issues/310))

## 0.18.0 (July 03, 2023)

ENHANCEMENTS:

* tftypes: Added `AttributePath` type `NextStep()` method, which returns the next step in the path without first copying via `Steps()` ([#307](https://github.com/hashicorp/terraform-plugin-go/issues/307))

BUG FIXES:

* tftypes: Significantly reduced compute and memory usage of `Value` type walking and transformation ([#307](https://github.com/hashicorp/terraform-plugin-go/issues/307))
* tftypes: Removed unnecessary memory allocations from `AttributePath` type `Equal()`, `LastStep()`, and `WithoutLastStep()` methods ([#307](https://github.com/hashicorp/terraform-plugin-go/issues/307))
* tftypes: Removed unnecessary memory allocations from `NewValue()` function ([#307](https://github.com/hashicorp/terraform-plugin-go/issues/307))

## 0.17.0 (June 28, 2023)

ENHANCEMENTS:

* tfprotov5: Added `DynamicValue` type `IsNull` method, which enables checking if the value is null without type information and fully decoding underlying data ([#305](https://github.com/hashicorp/terraform-plugin-go/issues/305))
* tfprotov6: Added `DynamicValue` type `IsNull` method, which enables checking if the value is null without type information and fully decoding underlying data ([#305](https://github.com/hashicorp/terraform-plugin-go/issues/305))

## 0.16.0 (June 16, 2023)

ENHANCEMENTS:

* all: Improved SDK logging performance when messages would be skipped due to configured logging level ([#289](https://github.com/hashicorp/terraform-plugin-go/issues/289))

BUG FIXES:

* tfprotov6/tf6server: Ensure interrupts route to StopProvider RPC ([#301](https://github.com/hashicorp/terraform-plugin-go/issues/301))

## 0.15.0 (March 27, 2023)

NOTES:

* all: This Go module has been updated to Go 1.19 per the [Go support policy](https://go.dev/doc/devel/release#policy). It is recommended to review the [Go 1.19 release notes](https://go.dev/doc/go1.19) before upgrading. Any consumers building on earlier Go versions may experience errors. ([#266](https://github.com/hashicorp/terraform-plugin-go/issues/266))

# 0.14.3 (January 5, 2023)

BUG FIXES:

* tfprotov5/tf5server: Prevented `TF_LOG_SDK_PROTO_DATA_DIR` file overlap by switching from second to millisecond file naming granularity ([#245](https://github.com/hashicorp/terraform-plugin-go/issues/245))
* tfprotov6/tf6server: Prevented `TF_LOG_SDK_PROTO_DATA_DIR` file overlap by switching from second to millisecond file naming granularity ([#245](https://github.com/hashicorp/terraform-plugin-go/issues/245))

# 0.14.2 (November 22, 2022)

BUG FIXES:

* tfprotov5: Allow diagnostic messages with incorrect UTF-8 encoding to pass through with the invalid sequences replaced with the Unicode Replacement Character. This avoids returning the unhelpful message "string field contains invalid UTF-8" in that case. ([#237](https://github.com/hashicorp/terraform-plugin-go/issues/237))
* tfprotov6: Allow diagnostic messages with incorrect UTF-8 encoding to pass through with the invalid sequences replaced with the Unicode Replacement Character. This avoids returning the unhelpful message "string field contains invalid UTF-8" in that case. ([#237](https://github.com/hashicorp/terraform-plugin-go/issues/237))

# 0.14.1 (November 8, 2022)

NOTES:

* No expected changes with this Go module's functionality. Contains updates to dependencies such as `google.golang.org/grpc` and `github.com/hashicorp/go-plugin`, which may be beneficial for consumers.

# 0.14.0 (August 15, 2022)

NOTES:

* This Go module has been updated to Go 1.18 per the [Go support policy](https://golang.org/doc/devel/release.html#policy). Any consumers building on earlier Go versions may experience errors. ([#219](https://github.com/hashicorp/terraform-plugin-go/issues/219))

ENHANCEMENTS:

* tfprotov5/tf5server: Added resource private state when protocol data output is enabled ([#221](https://github.com/hashicorp/terraform-plugin-go/issues/221))
* tfprotov6/tf6server: Added resource private state when protocol data output is enabled ([#221](https://github.com/hashicorp/terraform-plugin-go/issues/221))

BUG FIXES:

* tfprotov5/tf5server: Fixed `ApplyResourceChange` request RPC protocol data output to include `PriorState` and `ProviderMeta` fields ([#221](https://github.com/hashicorp/terraform-plugin-go/issues/221))
* tfprotov6/tf6server: Fixed `ApplyResourceChange` request RPC protocol data output to include `PriorState` and `ProviderMeta` fields ([#221](https://github.com/hashicorp/terraform-plugin-go/issues/221))

# 0.13.0 (July 28, 2022)

ENHANCEMENTS:
* tfprotov5: Added `RawState` type `UnmarshalWithOpts` method to facilitate configurable behaviour during unmarshalling ([#213](https://github.com/hashicorp/terraform-plugin-go/issues/213))
* tfprotov6: Added `RawState` type `UnmarshalWithOpts` method to facilitate configurable behaviour during unmarshalling ([#213](https://github.com/hashicorp/terraform-plugin-go/issues/213))

BUG FIXES:
* tftypes: Clarified `ValueFromJSON` error messaging with object attribute key issues ([#214](https://github.com/hashicorp/terraform-plugin-go/issues/214))

# 0.12.0 (July 15, 2022)

NOTES:

* The underlying `terraform-plugin-log` dependency has been updated to v0.6.0, which includes log filtering support and breaking changes of `With()` to `SetField()` function names. Any provider logging which calls those functions may require updates. ([#209](https://github.com/hashicorp/terraform-plugin-go/issues/209))

# 0.11.0 (July 8, 2022)

FEATURES:

* Added support for protocol version 5.3 and 6.3, which allows providers to opt into the `PlanResourceChange` RPC for resource destruction ([#205](https://github.com/hashicorp/terraform-plugin-go/issues/205))

ENHANCEMENTS:

* tfprotov5: Added `ServerCapabilities` type and `ServerCapabilities` field to `GetProviderSchemaResponse` ([#205](https://github.com/hashicorp/terraform-plugin-go/issues/205))
* tfprotov6: Added `ServerCapabilities` type and `ServerCapabilities` field to `GetProviderSchemaResponse` ([#205](https://github.com/hashicorp/terraform-plugin-go/issues/205))

# 0.10.0 (July 5, 2022)

ENHANCEMENTS:

* tfprotov5/tf5server: Added downstream RPC request duration and response diagnostics logging ([#203](https://github.com/hashicorp/terraform-plugin-go/issues/203))
* tfprotov6/tf6server: Added downstream RPC request duration and response diagnostics logging ([#203](https://github.com/hashicorp/terraform-plugin-go/issues/203))

# 0.9.1 (May 12, 2022)

BUG FIXES:
* tftypes: Prevented loss of number precision with integers between 54 and 64 bits ([#190](https://github.com/hashicorp/terraform-plugin-go/issues/190))

# 0.9.0 (April 13, 2022)

NOTES:

* This Go module has been updated to Go 1.17 per the [Go support policy](https://golang.org/doc/devel/release.html#policy). Any consumers building on earlier Go versions may experience errors. ([#175](https://github.com/hashicorp/terraform-plugin-go/issues/175))

ENHANCEMENTS:

* tftypes: Added `Type` support to `WalkAttributePath()` function ([#163](https://github.com/hashicorp/terraform-plugin-go/issues/163))

BUG FIXES:

* tfprotov5/tf5server: Ensured `@caller` in protocol logging entries accurately reflected calling code location ([#179](https://github.com/hashicorp/terraform-plugin-go/issues/179))
* tfprotov6/tf6server: Ensured `@caller` in protocol logging entries accurately reflected calling code location ([#179](https://github.com/hashicorp/terraform-plugin-go/issues/179))

# 0.8.0 (March 10, 2022)

NOTES:

* The underlying `terraform-plugin-log` dependency has been updated to v0.3.0, which includes a breaking change in the optional additional fields parameter of logging function calls to ensure correctness and catch coding errors during compilation. Any early adopter provider logging which calls those functions may require updates. ([#166](https://github.com/hashicorp/terraform-plugin-go/issues/166))

ENHANCEMENTS:

* tfprotov5: Added `ValueType()` methods to `Schema`, `SchemaAttribute`, `SchemaBlock`, and `SchemaNestedBlock` types. ([#158](https://github.com/hashicorp/terraform-plugin-go/issues/158))
* tfprotov6: Added `ValueType()` methods to `Schema`, `SchemaAttribute`, `SchemaBlock`, `SchemaNestedBlock`, and `SchemaObject` types. ([#158](https://github.com/hashicorp/terraform-plugin-go/issues/158))

# 0.7.1

BUG FIXES:

* tfprotov5/tf5server: Ensure server options are passed through on startup ([#153](https://github.com/hashicorp/terraform-plugin-go/issues/153))
* tfprotov5/tf5server: Prevent empty provider address error logs on startup ([#150](https://github.com/hashicorp/terraform-plugin-go/issues/150))

# 0.7.0 (January 20, 2022)

BREAKING CHANGES:

* tfprotov6: The `SchemaObject.MaxItems` and `SchemaObject.MinItems` fields have been removed without replacement. These were never used in the protocol and did not perform any validation. ([#143](https://github.com/hashicorp/terraform-plugin-go/issues/143))

BUG FIXES:

* tfprotov6: The `ApplyResourceChangeResponse.UnsafeToUseLegacyTypeSystem` and `PlanResourceChangeResponse.UnsafeToUseLegacyTypeSystem` fields are now properly passed into and out of the protocol ([#143](https://github.com/hashicorp/terraform-plugin-go/issues/143))

# 0.6.0 (January 18, 2022)

ENHANCEMENTS:

* tfprotov5/tf5server: Added `WithManagedDebug()` `ServeOpt`, which implements outputting reattach configuration to stdout and stopping on SIGINT ([#137](https://github.com/hashicorp/terraform-plugin-go/issues/137))
* tfprotov5/tf5server: Added support for writing protocol data to disk by setting `TF_LOG_SDK_PROTO_DATA_DIR` environment variable ([#135](https://github.com/hashicorp/terraform-plugin-go/issues/135))
* tfprotov5/tf5server: Increased maximum gRPC send and receive message size limit to 256MB ([#139](https://github.com/hashicorp/terraform-plugin-go/issues/139))
* tfprotov6/tf6server: Added `WithManagedDebug()` `ServeOpt`, which implements outputting reattach configuration to stdout and stopping on SIGINT ([#137](https://github.com/hashicorp/terraform-plugin-go/issues/137))
* tfprotov6/tf6server: Added support for writing protocol data to disk by setting `TF_LOG_SDK_PROTO_DATA_DIR` environment variable ([#135](https://github.com/hashicorp/terraform-plugin-go/issues/135))
* tfprotov6/tf6server: Increased maximum gRPC send and receive message size limit to 256MB ([#139](https://github.com/hashicorp/terraform-plugin-go/issues/139))

BUG FIXES:

* Prevent potential process leak on Windows platforms ([#123](https://github.com/hashicorp/terraform-plugin-go/issues/123))
* tftypes: Fixed regression with DynamicPseudoType handling since v0.4.0, allowing usage of known values again and preventing msgpack decoding errors in Terraform CLI ([#136](https://github.com/hashicorp/terraform-plugin-go/issues/136))

# 0.5.0 (December 07, 2021)

BREAKING CHANGES:
* `tf6server.New` and `tf5server.New` now accept a name argument (meant to be the full registry path of the provider, e.g. registry.terraform.io/hashicorp/random) and a variadic argument of `ServeOpts`, just like the `Serve` function. Providers calling `Serve` will not notice any change. Providers calling `New` directly should pass the name and any options they would like the server to be configured with. ([#93](https://github.com/hashicorp/terraform-plugin-go/issues/93))

FEATURES:
* Added support for [terraform-plugin-log](https://github.com/hashicorp/terraform-plugin-log) v0.2.0, injecting loggers for SDKs and the provider into requests and adding trace and error log lines to request handlers. ([#93](https://github.com/hashicorp/terraform-plugin-go/issues/93))

ENHANCEMENTS:
* New `WithoutLogStderrOverride`, `WithLogEnvVarName`, and `WithoutLogLocation` `ServeOpt` helper functions have been added to `tf6server` and `tf5server`. These helpers can be passed to `Serve` or `New` to control logging behaviors for SDK and provider logs. `WithoutLogStderrOverride` disables using the stderr that existed at startup instead of the one that exists when the log function is called; it is recommended to not use this unless you understand the impacts, as Terraform's default behavior is counter-intuitive. `WithLogEnvVarName` sets the name of the provider's log module and controls what follows `TF_LOG_PROVIDER_` as the environment variable to control the log's level. `WithoutLogLocation` excludes filename and line numbers from log output. ([#93](https://github.com/hashicorp/terraform-plugin-go/issues/93))
* tftypes: Added `AttributePath` `LastStep()` method and `AttributePathStep` `Equal()` method ([#112](https://github.com/hashicorp/terraform-plugin-go/issues/112))

BUG FIXES:
* Fixed a panic when using DynamicPseudoType attributes in an object. ([#116](https://github.com/hashicorp/terraform-plugin-go/issues/116))
* tftypes: Return null `Number` when `NewValue` receives `(*big.Float)(nil)` ([#114](https://github.com/hashicorp/terraform-plugin-go/issues/114))

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
