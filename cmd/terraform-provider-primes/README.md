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

$ grpcurl -plaintext $(cmd/terraform-provider-primes/reattach.sh) tfplugin5.Provider.ListResource
{
  "displayName": "2 is a prime number",
  "resourceObject": {
    "msgpack": "Ag=="
  }
}
{
  "displayName": "3 is a prime number",
  "resourceObject": {
    "msgpack": "Aw=="
  }
}
{
  "displayName": "5 is a prime number",
  "resourceObject": {
    "msgpack": "BQ=="
  }
}
{
  "displayName": "7 is a prime number",
  "resourceObject": {
    "msgpack": "Bw=="
  }
}
{
  "displayName": "11 is a prime number",
  "resourceObject": {
    "msgpack": "Cw=="
  }
}
{
  "displayName": "13 is a prime number",
  "resourceObject": {
    "msgpack": "DQ=="
  }
}
{
  "displayName": "17 is a prime number",
  "resourceObject": {
    "msgpack": "EQ=="
  }
}
{
  "displayName": "23 is a prime number",
  "resourceObject": {
    "msgpack": "Fw=="
  }
}
{
  "displayName": "29 is a prime number",
  "resourceObject": {
    "msgpack": "HQ=="
  }
}
{
  "displayName": "31 is a prime number",
  "resourceObject": {
    "msgpack": "Hw=="
  }
}
```
