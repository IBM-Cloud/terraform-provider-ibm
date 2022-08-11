// This allows managed_key data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_hpcs_managed_key" {
  value       = ibm_managed_key.managed_key_instance
  description = "managed_key resource instance"
}
// This allows key_template data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_hpcs_key_template" {
  value       = ibm_key_template.key_template_instance
  description = "key_template resource instance"
}
// This allows keystore data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_hpcs_keystore" {
  value       = ibm_keystore.keystore_instance
  description = "keystore resource instance"
  sensitive   = true
}
// This allows vault data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_hpcs_vault" {
  value       = ibm_vault.vault_instance
  description = "vault resource instance"
}
