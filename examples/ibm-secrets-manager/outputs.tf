// This allows secret_group data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_secret_group" {
  value       = ibm_secret_group.secret_group_instance
  description = "secret_group resource instance"
}
