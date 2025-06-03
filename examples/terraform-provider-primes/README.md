# terraform-provider-primes

```bash
$ TF_LOG_SDK_PROTO=trace TF_PLUGIN_MAGIC_COOKIE=[value] go run ./examples/terraform-provider-primes

{"@level":"debug","@message":"plugin address","@timestamp":"2037-06-30T10:10:00.627659+04:00","address":"...","network":"unix"}
Provider started. To attach Terraform CLI, set the TF_REATTACH_PROVIDERS environment variable with the following:

	TF_REATTACH_PROVIDERS='{"terraform.io/playground/primes":{"Protocol":"grpc","ProtocolVersion":5,"Pid":...,"Test":true,"Addr":{"Network":"unix","String":"..."}}}'

Writing reattach configuration to env file at path reattach.env

$ grpcurl -plaintext $(examples/terraform-provider-primes/reattach.sh) list
grpc.health.v1.Health
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
plugin.GRPCBroker
plugin.GRPCController
plugin.GRPCStdio
tfplugin5.Provider

$ grpcurl -plaintext $(examples/terraform-provider-primes/reattach.sh) tfplugin5.Provider.GetSchema
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

$ grpcurl -plaintext $(examples/terraform-provider-primes/reattach.sh) tfplugin5.Provider.ListResource
{
  "displayName": "primes[1]: 2",
  "resourceObject": {
    "msgpack": "gqZudW1iZXICp29yZGluYWwB"
  }
}
{
  "displayName": "primes[2]: 3",
  "resourceObject": {
    "msgpack": "gqZudW1iZXIDp29yZGluYWwC"
  }
}
{
  "displayName": "primes[3]: 5",
  "resourceObject": {
    "msgpack": "gqZudW1iZXIFp29yZGluYWwD"
  }
}
{
  "displayName": "primes[4]: 7",
  "resourceObject": {
    "msgpack": "gqZudW1iZXIHp29yZGluYWwE"
  }
}
...
```
