// This output allows account_settings_template_assignment_instance data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "account_settings_template_instance" {
  value       = ibm_iam_account_settings_template.account_settings_template_instance
  description = "ibm_iam_account_settings_template_assignment resource instance"
}

output "account_settings_template_assignment_instance" {
  value       = ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance
  description = "ibm_iam_account_settings_template_assignment resource instance"
}