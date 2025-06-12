// This output allows account_settings_template data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_iam_account_settings_template" {
  value       = ibm_iam_account_settings_template.account_settings_template_instance
  description = "account_settings_template resource instance"
}

output "account_settings_template_new_version" {
  value       = ibm_iam_account_settings_template.account_settings_template_new_version
  description = "account_settings_template resource instance"
}

