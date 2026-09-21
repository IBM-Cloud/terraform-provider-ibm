output "ibm_iam_account_settings" {
  value       = ibm_iam_account_settings.iam_account_settings_instance
  description = "ibm_iam_account_settings resource instance"
}

output "ibm_iam_account_settings_data" {
  value       = data.ibm_iam_account_settings.iam_account_settings_data
  description = "ibm_iam_account_settings data"
}
