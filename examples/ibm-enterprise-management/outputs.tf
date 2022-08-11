// This allows enterprise data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_enterprise" {
  value       = ibm_enterprise.enterprise_instance
  description = "enterprise resource instance"
}
// This allows enterprise_account_group data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_enterprise_account_group" {
  value       = ibm_enterprise_account_group.enterprise_account_group_instance
  description = "enterprise_account_group resource instance"
}
// This allows enterprise_account data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_enterprise_account" {
  value       = ibm_enterprise_account.enterprise_account_instance
  description = "enterprise_account resource instance"
}

output "enterprise_import_account"{
  value = ibm_enterprise_account.enterprise_import_account
  description = "enterprise_import_account resource instance"
}