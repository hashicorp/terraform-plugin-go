resource "time_static" "example" {}

output "current_time" {
  value = time_static.example.rfc3339
}