provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.region
}

// Create secrets_manager_secrets data source
data "ibm_secrets_manager_secrets" "secrets_manager_secrets_instance" {
  instance_id = var.secrets_manager_instance_id
  secret_type = var.secrets_manager_secrets_secret_type
}

// Create secrets_manager_secret data source
data "ibm_secrets_manager_secret" "secrets_manager_secret_instance" {
  instance_id = var.secrets_manager_instance_id
  secret_type = var.secrets_manager_secret_secret_type
  secret_id   = var.secrets_manager_secret_id
}