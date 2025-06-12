// This allows sm_secret_group data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_secret_group" {
  value       = ibm_sm_secret_group.sm_secret_group_instance
  description = "sm_secret_group resource instance"
}
// This allows sm_imported_certificate data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_imported_certificate" {
  value       = ibm_sm_imported_certificate.sm_imported_certificate_instance
  description = "sm_imported_certificate resource instance"
  sensitive   = true
}
// This allows sm_public_certificate data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_public_certificate" {
  value       = ibm_sm_public_certificate.sm_public_certificate_instance
  description = "sm_public_certificate resource instance"
  sensitive   = true
}
// This allows sm_public_certificate_action_validate_manual_dns data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_public_certificate_action_validate_manual_dns" {
  value       = ibm_sm_public_certificate_action_validate_manual_dns.sm_public_certificate_action_validate_manual_dns_instance
  description = "sm_public_certificate_action_validate_manual_dns resource instance"
}
// This allows sm_kv_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_kv_secret" {
  value       = ibm_sm_kv_secret.sm_kv_secret_instance
  description = "sm_kv_secret resource instance"
  sensitive   = true
}
// This allows sm_iam_credentials_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_iam_credentials_secret" {
  value       = ibm_sm_iam_credentials_secret.sm_iam_credentials_secret_instance
  description = "sm_iam_credentials_secret resource instance"
}
// This allows sm_service_credentials_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_service_credentials_secret" {
  value       = ibm_sm_service_credentials_secret.sm_service_credentials_secret_instance
  description = "sm_service_credentials_secret resource instance"
}
// This allows sm_arbitrary_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_arbitrary_secret" {
  value       = ibm_sm_arbitrary_secret.sm_arbitrary_secret_instance
  description = "sm_arbitrary_secret resource instance"
}
// This allows sm_username_password_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_username_password_secret" {
  value       = ibm_sm_username_password_secret.sm_username_password_secret_instance
  description = "sm_username_password_secret resource instance"
}
// This allows sm_arbitrary_secret_version data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_arbitrary_secret_version" {
  value       = ibm_sm_arbitrary_secret_version.sm_arbitrary_secret_version_instance
  description = "sm_arbitrary_secret_version resource instance"
}
// This allows sm_private_certificate data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_private_certificate" {
  value       = ibm_sm_private_certificate.sm_private_certificate_instance
  description = "sm_private_certificate resource instance"
  sensitive   = true
}
// This allows sm_private_certificate_configuration_root_ca data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_private_certificate_configuration_root_ca" {
  value       = ibm_sm_private_certificate_configuration_root_ca.sm_private_certificate_configuration_root_ca_instance
  description = "sm_private_certificate_configuration_root_ca resource instance"
  sensitive   = true
}
// This allows sm_private_certificate_configuration_intermediate_ca data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_private_certificate_configuration_intermediate_ca" {
  value       = ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca_instance
  description = "sm_private_certificate_configuration_intermediate_ca resource instance"
  sensitive   = true
}
// This allows sm_private_certificate_configuration_template data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_private_certificate_configuration_template" {
  value       = ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template_instance
  description = "sm_private_certificate_configuration_template resource instance"
}
// This allows sm_private_certificate_configuration_action_sign_csr data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_private_certificate_configuration_action_sign_csr" {
  value       = ibm_sm_private_certificate_configuration_action_sign_csr.sm_private_certificate_configuration_action_sign_csr_instance
  description = "sm_private_certificate_configuration_action_sign_csr resource instance"
  sensitive   = true
}
// This allows sm_private_certificate_configuration_action_set_signed data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_private_certificate_configuration_action_set_signed" {
  value       = ibm_sm_private_certificate_configuration_action_set_signed.sm_private_certificate_configuration_action_set_signed_instance
  description = "sm_private_certificate_configuration_action_set_signed resource instance"
  sensitive   = true
}
// This allows sm_public_certificate_configuration_ca_lets_encrypt data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_public_certificate_configuration_ca_lets_encrypt" {
  value       = ibm_sm_public_certificate_configuration_ca_lets_encrypt.sm_public_certificate_configuration_ca_lets_encrypt_instance
  description = "sm_public_certificate_configuration_ca_lets_encrypt resource instance"
}
// This allows sm_public_certificate_configuration_dns_cis data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_public_certificate_configuration_dns_cis" {
  value       = ibm_sm_public_certificate_configuration_dns_cis.sm_public_certificate_configuration_dns_cis_instance
  description = "sm_public_certificate_configuration_dns_cis resource instance"
}
// This allows sm_public_certificate_configuration_dns_classic_infrastructure data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" {
  value       = ibm_sm_public_certificate_configuration_dns_classic_infrastructure.sm_public_certificate_configuration_dns_classic_infrastructure_instance
  description = "sm_public_certificate_configuration_dns_classic_infrastructure resource instance"
}
// This allows sm_en_registration data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_sm_en_registration" {
  value       = ibm_sm_en_registration.sm_en_registration_instance
  description = "sm_en_registration resource instance"
}