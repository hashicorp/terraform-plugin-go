#!/usr/bin/env bash

set -euo pipefail

rm -rf tfprotov6/internal/tfplugin5
rm -rf tfprotov6/internal/tfplugin6
(cd tfprotov5/internal/fromproto && gofmt -r '"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5" -> "github.com/pulumi/terraform/pkg/tfplugin5"' -w *.go)
(cd tfprotov5/internal/toproto && gofmt -r '"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5" -> "github.com/pulumi/terraform/pkg/tfplugin5"' -w *.go)
(cd tfprotov6/internal/fromproto && gofmt -r '"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6" -> "github.com/pulumi/terraform/pkg/tfplugin6"' -w *.go)
(cd tfprotov6/internal/toproto && gofmt -r '"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6" -> "github.com/pulumi/terraform/pkg/tfplugin6"' -w *.go)

(cd tfprotov5/tf5server && gofmt -r '"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5" -> "github.com/pulumi/terraform/pkg/tfplugin5"' -w *.go)
(cd tfprotov6/tf6server && gofmt -r '"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6" -> "github.com/pulumi/terraform/pkg/tfplugin6"' -w *.go)
