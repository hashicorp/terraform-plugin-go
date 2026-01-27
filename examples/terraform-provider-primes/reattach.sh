#!/usr/bin/env bash
# Copyright IBM Corp. 2020, 2025
# SPDX-License-Identifier: MPL-2.0


source .env.reattach
echo $TF_REATTACH_PROVIDERS | jq -r '.["terraform.io/playground/primes"].Addr.Network + "://" + .["terraform.io/playground/primes"].Addr.String'
