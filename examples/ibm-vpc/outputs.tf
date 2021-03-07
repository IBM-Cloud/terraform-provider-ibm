// This allows is_dedicated_host_group data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_is_dedicated_host_group" {
  value       = ibm_is_dedicated_host_group.is_dedicated_host_group_instance
  description = "is_dedicated_host_group resource instance"
}
