#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


source .env.reattach
echo $TF_REATTACH_PROVIDERS | jq -r '.["terraform.io/playground/primes"].Addr.Network + "://" + .["terraform.io/playground/primes"].Addr.String'
