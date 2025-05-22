# terraform-provider-primes

```bash
$ while :; do TF_PLUGIN_MAGIC_COOKIE=d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2 go run ./cmd/terraform-provider-primes; done

{"@level":"debug","@message":"plugin address","@timestamp":"2025-05-21T18:15:01.627659-04:00","address":"/var/folders/17/pz_1q8cn6pq53qjmvh72dzpm0000gn/T/plugin2908212597","network":"unix"}
Provider started. To attach Terraform CLI, set the TF_REATTACH_PROVIDERS environment variable with the following:

	TF_REATTACH_PROVIDERS='{"terraform.io/playground/primes":{"Protocol":"grpc","ProtocolVersion":5,"Pid":47254,"Test":true,"Addr":{"Network":"unix","String":"/var/folders/17/pz_1q8cn6pq53qjmvh72dzpm0000gn/T/plugin2908212597"}}}'

Writing reattach configuration to env file at path reattach.env

$ grpcurl -plaintext $(cmd/terraform-provider-primes/reattach.sh) list
grpc.health.v1.Health
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
plugin.GRPCBroker
plugin.GRPCController
plugin.GRPCStdio
tfplugin5.Provider

$ grpcurl -plaintext $(cmd/terraform-provider-primes/reattach.sh) tfplugin5.Provider.GetSchema
{
  "provider": {},
  "listResourceSchemas": {
    "prime": {
      "block": {
        "attributes": [
          {
            "name": "number",
            "type": "Im51bWJlciI=",
            "description": "The nth prime"
          },
          {
            "name": "ordinal",
            "type": "Im51bWJlciI=",
            "description": "n"
          }
        ]
      }
    }
  }
}

$ grpcurl -plaintext $(cmd/terraform-provider-primes/reattach.sh) tfplugin5.Provider.ListResource
{
  "displayName": "The 1st prime is 2",
  "resourceObject": {
    "msgpack": "gqZudW1iZXICp29yZGluYWwB"
  }
}
{
  "displayName": "The 2nd prime is 3",
  "resourceObject": {
    "msgpack": "gqZudW1iZXIDp29yZGluYWwC"
  }
}
{
  "displayName": "The 3rd prime is 5",
  "resourceObject": {
    "msgpack": "gqZudW1iZXIFp29yZGluYWwD"
  }
}
{
  "displayName": "The 4th prime is 7",
  "resourceObject": {
    "msgpack": "gqZudW1iZXIHp29yZGluYWwE"
  }
}
...
```
