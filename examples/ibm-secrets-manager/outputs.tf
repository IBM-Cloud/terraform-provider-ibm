output "secrets_manager_secrets" {
  value = data.ibm_secrets_manager_secrets.secrets_manager_secrets_instance
}

output "secrets_manager_secret" {
  value = data.ibm_secrets_manager_secret.secrets_manager_secret_instance
}