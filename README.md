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

## Go Compatibility

This project follows the [support policy](https://golang.org/doc/devel/release.html#policy) of Go as its support policy. The two latest major releases of Go are supported by the project.

Currently, that means Go **1.18** or later must be used when including this project as a dependency.

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
source-specific functions by using implementations of the
[`tfprotov5.ResourceServer`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#ResourceServer)
and
[`tfprotov5.DataSourceServer`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov5#DataSourceServer)
interfaces, using the provider-level implementations of the interfaces to route
to the correct resource or data source level implementations using
`req.TypeName`.

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

The Terraform Plugin SDK's [`helper/resource`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource) package can be used to test providers written using terraform-plugin-go. While we are working on a testing framework for terraform-plugin-go providers that is independent of the Plugin SDK, this may take some time, so we recommend writing tests in the meantime using the plugin SDK, which will not be a runtime dependency of your provider.

You must supply a factory for your provider server by setting `ProtoV5ProviderFactories` on each [`TestCase`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource#TestCase). For example:

```go
package myprovider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceFoo(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"myprovider": func() (tfprotov5.ProviderServer, error) {
				return Server(), nil
			},
		},
		Steps: []resource.TestStep{
			{
				Config: `"data" "myprovider_foo" "bar" {}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.myprovider_foo.bar", "current", regexp.MustCompile(`[0-9]+`)),
				),
			},
		},
	})
}
```

## Debugging

Provider servers can be instrumented with debugging tooling, such as [`delve`](https://github.com/go-delve/delve/), by using the [`WithManagedDebug()`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server#WithManagedDebug) and [`WithDebug()`](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-go/tfprotov6/tf6server#WithDebug) `ServeOpt`. In this mode, Terraform CLI no longer manages the server lifecycle and instead connects to the running provider server via a reattach configuration supplied by the `TF_REATTACH_PROVIDERS` environment variable. The `WithDebug()` implementation is meant for advanced use cases which require manually handling the reattach configuration, such as managing providers with [terraform-exec](https://pkg.go.dev/github.com/hashicorp/terraform-exec), while the `WithManagedDebug()` implementation is suitable for provider `main()` functions. For example:

```go
func main() {
	debugFlag := flag.Bool("debug", false, "Start provider in debug mode.")
	flag.Parse()

	opts := []tf6server.ServeOpt{}

	if *debugFlag {
		opts = append(opts, tf6server.WithManagedDebug())
	}

	tf6server.Serve("registry.terraform.io/namespace/example", /* Provider function */, opts...)
}
```

### Protocol Data

To write raw protocol MessagePack or JSON data to disk, set the `TF_LOG_SDK_PROTO_DATA_DIR` environment variable. During Terraform execution, this directory will get populated with `{TIME}_{RPC}_{MESSAGE}_{FIELD}.{EXTENSION}` named files. Tooling such as [`jq`](https://stedolan.github.io/jq/) can be used to inspect the JSON data. Tooling such as [`fq`](https://github.com/wader/fq) or [`msgpack2json`](https://pkg.go.dev/github.com/nokute78/msgpack-microscope/cmd/msgpack2json) can be used to inspect the MessagePack data.

## Documentation

Documentation is a work in progress. The GoDoc for packages, types, functions,
and methods should have complete information, but we're working to add a
section to [terraform.io](https://terraform.io/) with more information about
the module, its common uses, and patterns developers may wish to take advantage
of.

Please bear with us as we work to get this information published, and please
[open
issues](https://github.com/hashicorp/terraform-plugin-go/issues/new/choose)
with requests for the kind of documentation you would find useful.

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

Please see [`.github/CONTRIBUTING.md`](https://github.com/hashicorp/terraform-plugin-go/blob/main/.github/CONTRIBUTING.md).

### Unit Testing

Run `go test ./...` or `make test` after any changes.

### Linting

Ensure the following tooling is installed:

- [`golangci-lint](https://golangci-lint.run/): Aggregate Go linting tool.

Run `golangci-lint run ./...` or `make lint` after any changes.

### Protocol Updates

Ensure the following tooling is installed:

- [`protoc`](https://github.com/protocolbuffers/protobuf): Protocol Buffers compiler.
- [`protoc-gen-go`](https://pkg.go.dev/google.golang.org/protobuf/cmd/protoc-gen-go): Go plugin for Protocol Buffers compiler. e.g. `go install google.golang.org/protobuf/cmd/protoc-gen-go`
- [`protoc-gen-go-grpc`](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc): Go gRPC plugin for Protocol Buffers compiler. e.g. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`

The Protocol Buffers definitions can be found in `tfprotov5/internal/tfplugin5` and `tfprotov6/internal/tfplugin6`.

Run `make protoc` to recompile the Protocol Buffers files after any changes.

## License

This module is licensed under the [Mozilla Public License v2.0](https://github.com/hashicorp/terraform-plugin-go/blob/main/LICENSE).
