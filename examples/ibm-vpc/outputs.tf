// This allows is_dedicated_host data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_is_dedicated_host" {
  value       = ibm_is_dedicated_host.is_dedicated_host_instance
  description = "is_dedicated_host resource instance"
}
