// This allows iam_account_settings data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_iam_account_settings" {
  value       = ibm_iam_account_settings.iam_account_settings_instance
  description = "iam_account_settings resource instance"
}
