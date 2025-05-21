#!/usr/bin/env bash

source reattach.env
echo $TF_REATTACH_PROVIDERS | jq -r '.["terraform.io/playground/primes"].Addr.Network + "://" + .["terraform.io/playground/primes"].Addr.String'
