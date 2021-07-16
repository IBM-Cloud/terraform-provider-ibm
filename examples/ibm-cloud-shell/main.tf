provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create cloud_shell_account_settings data source
data "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = var.cloud_shell_account_settings_account_id
}

// Update cloud_shell_account_settings
resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = var.cloud_shell_account_settings_account_id
  rev = data.ibm_cloud_shell_account_settings.cloud_shell_account_settings
  default_enable_new_features = var.cloud_shell_account_settings_default_enable_new_features
  default_enable_new_regions = var.cloud_shell_account_settings_default_enable_new_regions
  enabled = var.cloud_shell_account_settings_enabled
  features = var.cloud_shell_account_settings_features
  regions = var.cloud_shell_account_settings_regions
  tags = var.cloud_shell_account_settings_tags
}
