[![PkgGoDev](https://pkg.go.dev/badge/github.com/hashicorp/terraform-plugin-go)](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go)

# terraform-plugin-go

terraform-plugin-go provides low-level Go bindings for the Terraform
plugin protocol, for integrations to be built upon. It strives to be a
minimal possible abstraction on top of the protocol, only hiding the
implementation details of the protocol while leaving its semantics
unchanged.

## Status

terraform-plugin-go is a [Go module](https://github.com/golang/go/wiki/Modules)
versioned using [semantic versioning](https://semver.org/).

The module is currently on a v0 major version, indicating our lack of
confidence in the stability of its exported API. Developers depending on it
should do so with an explicit understanding that the API may change and shift
until we hit v1.0.0, as we learn more about the needs and expectations of
developers working with the module.

We are confident in the correctness of the code and it is safe to build on so
long as the developer understands that the API may change in backwards
incompatible ways and they are expected to be tracking these changes.

## Terraform CLI Compatibility

Providers built on terraform-plugin-go will only be usable with Terraform
v0.12.0 and later. Developing providers for versions of Terraform below 0.12.0
is unsupported by the Terraform Plugin SDK team.

## Getting Started

terraform-plugin-go is targeted towards experienced Terraform developers.
Familiarity with the [Resource Instance Change
Lifecycle](https://github.com/hashicorp/terraform/blob/master/docs/resource-instance-change-lifecycle.md)
is required, and it is the provider developer's responsibility to ensure that
Terraform's requirements and invariants for responses are honored.

Provider developers are expected to create a type that implements the
[`tfprotov5.ProviderServer`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#ProviderServer)
interface. This type should be passed to
[`tfprotov5server.Serve`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5/server#Serve)
along with the name (like `"hashicorp/time"`).

Resources and data sources can be handled in resource-specific or data
source-specific functions by using
[`tfprotov5.ResourceRouter`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#ResourceRouter)
and
[`tfprotov5.DataSourceRouter`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#DataSourceRouter)
as embedded values in their `tfprotov5.ProviderServer` implementation type,
falling back on their implementations of the
[`tfprotov5.ResourceServer`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#ResourceServer)
and
[`tfprotov5.DataSourceServer`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#DataSourceServer)
interfaces to split resource and data source requests into single-purpose
functions based on which resource or data source the request is for.

When handling requests,
[`tfprotov5.DynamicValue`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#DynamicValue)
types should always have their
[`Unmarshal`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#DynamicValue.Unmarshal)
methods called; their properties should not be inspected directly. The
[`tftypes.Value`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes#Value)
returned from `Unmarshal` can be inspected to check [whether it is
known](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes#Value.IsKnown)
and subsequently converted to a plain Go type using its
[`As`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes#Value.As)
method. `As` will return an error if the `Value` is not known.

## Testing

TODO: insert information here on how to use `helper/resource` to test providers
written with terraform-plugin-go.

## Documentation

Documentation is a work in progress. The GoDoc for packages, types, functions,
and methods should have complete information, but we're working to add a
section to [terraform.io](https://terraform.io/) with more information about
the module, its common uses, and patterns developers may wish to take advantage
of.

Please bear with us as we work to get this information published, and please
[open issues](https://github.com/hashicorp/terraform-plugin-go) with requests
for the kind of documentation you would find useful.

## Scope

This module is intentionally limited in its scope. It serves as the foundation
for the provider ecosystem, so major breaking changes are incredibly expensive.
By limiting the scope of the project, we're limiting the choices it needs to
make, making it less likely that breaking changes will be required once we've
hit version 1.0.0.

To that end, terraform-plugin-go's scope is limited to providing a common gRPC
server interface and an implementation of Terraform's type system. It
specifically is not trying to be a framework, nor is it attempting to provide
any utility or helper functions that only a subset of provider developers will
rely on. Its litmus test for whether something should be included is "will
every Terraform provider need this functionality or can this functionality only
be added if it's in this module?" All other functionality should be considered
out of scope and should live in a separate module.

## Contributing

Please see [`.github/CONTRIBUTING.md`](https://github.com/hashicorp/terraform-plugin-go/blob/master/.github/CONTRIBUTING.md).

## License

This module is licensed under the [Mozilla Public License v2.0](https://github.com/hashicorp/terraform-plugin-go/blob/master/LICENSE).
