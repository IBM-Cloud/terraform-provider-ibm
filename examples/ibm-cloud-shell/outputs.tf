// This allows cloud_shell_account_settings data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cloud_shell_account_settings" {
  value       = ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance
  description = "cloud_shell_account_settings resource instance"
}
