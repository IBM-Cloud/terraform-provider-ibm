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

// Provision sm_secret_group resource instance
resource "ibm_sm_secret_group" {
  description = var.sm_secret_group_description
}

// Provision sm_imported_certificate resource instance
resource "ibm_sm_imported_certificate" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_imported_certificate_custom_metadata) }
  description = var.sm_imported_certificate_description
  expiration_date = var.sm_imported_certificate_expiration_date
  labels = var.sm_imported_certificate_labels
  secret_group_id = var.sm_imported_certificate_secret_group_id
  secret_type = var.sm_imported_certificate_secret_type
  certificate = var.sm_imported_certificate_certificate
  intermediate = var.sm_imported_certificate_intermediate
  private_key = var.sm_imported_certificate_private_key
}

// Provision sm_public_certificate resource instance
resource "ibm_sm_public_certificate" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_public_certificate_custom_metadata) }
  description = var.sm_public_certificate_description
  expiration_date = var.sm_public_certificate_expiration_date
  labels = var.sm_public_certificate_labels
  secret_group_id = var.sm_public_certificate_secret_group_id
  secret_type = var.sm_public_certificate_secret_type
  rotation {
    auto_rotate = true
    interval = 1
    unit = "day"
  }
}

// Provision sm_kv_secret resource instance
resource "ibm_sm_kv_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_kv_secret_custom_metadata) }
  description = var.sm_kv_secret_description
  labels = var.sm_kv_secret_labels
  secret_group_id = var.sm_kv_secret_secret_group_id
  secret_type = var.sm_kv_secret_secret_type
  data = { my_key = jsonencode(var.sm_kv_secret_data) }
}

// Provision sm_iam_credentials_secret resource instance
resource "ibm_sm_iam_credentials_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_iam_credentials_secret_custom_metadata) }
  description = var.sm_iam_credentials_secret_description
  labels = var.sm_iam_credentials_secret_labels
  secret_group_id = var.sm_iam_credentials_secret_secret_group_id
  secret_type = var.sm_iam_credentials_secret_secret_type
  ttl = var.sm_iam_credentials_secret_ttl
  access_groups = var.sm_iam_credentials_secret_access_groups
  service_id = var.sm_iam_credentials_secret_service_id
  reuse_api_key = var.sm_iam_credentials_secret_reuse_api_key
  rotation {
    auto_rotate = true
    interval = 1
    unit = "day"
  }
}

// Provision sm_arbitrary_secret resource instance
resource "ibm_sm_arbitrary_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_arbitrary_secret_custom_metadata) }
  description = var.sm_arbitrary_secret_description
  expiration_date = var.sm_arbitrary_secret_expiration_date
  labels = var.sm_arbitrary_secret_labels
  secret_group_id = var.sm_arbitrary_secret_secret_group_id
  secret_type = var.sm_arbitrary_secret_secret_type
  payload = var.sm_arbitrary_secret_payload
}

// Provision sm_username_password_secret resource instance
resource "ibm_sm_username_password_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_username_password_secret_custom_metadata) }
  description = var.sm_username_password_secret_description
  expiration_date = var.sm_username_password_secret_expiration_date
  labels = var.sm_username_password_secret_labels
  secret_group_id = var.sm_username_password_secret_secret_group_id
  secret_type = var.sm_username_password_secret_secret_type
  rotation {
    auto_rotate = true
    interval = 1
    unit = "day"
  }
  username = var.sm_username_password_secret_username
  password = var.sm_username_password_secret_password
}

// Provision sm_arbitrary_secret_version resource instance
resource "ibm_sm_arbitrary_secret_version" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  payload = var.sm_arbitrary_secret_version_payload
  version_custom_metadata = { my_key = jsonencode(var.sm_arbitrary_secret_version_version_custom_metadata) }
  secret_id = var.sm_arbitrary_secret_version_secret_id
}

// Provision sm_private_certificate resource instance
resource "ibm_sm_private_certificate" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  custom_metadata = { my_key = jsonencode(var.sm_private_certificate_custom_metadata) }
  description = var.sm_private_certificate_description
  expiration_date = var.sm_private_certificate_expiration_date
  labels = var.sm_private_certificate_labels
  secret_group_id = var.sm_private_certificate_secret_group_id
  secret_type = var.sm_private_certificate_secret_type
  rotation {
    auto_rotate = true
    interval = 1
    unit = "day"
  }
  certificate_template = var.sm_private_certificate_certificate_template
}

// Provision sm_configuration_private_certificate_root_CA resource instance
resource "ibm_sm_configuration_private_certificate_root_CA" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  config_type = var.sm_configuration_private_certificate_root_CA_config_type
  crl_disable = var.sm_configuration_private_certificate_root_CA_crl_disable
  crl_distribution_points_encoded = var.sm_configuration_private_certificate_root_CA_crl_distribution_points_encoded
  issuing_certificates_urls_encoded = var.sm_configuration_private_certificate_root_CA_issuing_certificates_urls_encoded
  ttl = var.sm_configuration_private_certificate_root_CA_ttl
}

// Provision sm_configuration_private_certificate_intermediate_CA resource instance
resource "ibm_sm_configuration_private_certificate_intermediate_CA" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  config_type = var.sm_configuration_private_certificate_intermediate_CA_config_type
  crl_disable = var.sm_configuration_private_certificate_intermediate_CA_crl_disable
  crl_distribution_points_encoded = var.sm_configuration_private_certificate_intermediate_CA_crl_distribution_points_encoded
  issuing_certificates_urls_encoded = var.sm_configuration_private_certificate_intermediate_CA_issuing_certificates_urls_encoded
  signing_method = var.sm_configuration_private_certificate_intermediate_CA_signing_method
}

// Provision sm_configuration_private_certificate_template resource instance
resource "ibm_sm_configuration_private_certificate_template" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  config_type = var.sm_configuration_private_certificate_template_config_type
  certificate_authority = var.sm_configuration_private_certificate_template_certificate_authority
  allowed_secret_groups = var.sm_configuration_private_certificate_template_allowed_secret_groups
  allow_localhost = var.sm_configuration_private_certificate_template_allow_localhost
  allowed_domains = var.sm_configuration_private_certificate_template_allowed_domains
  allowed_domains_template = var.sm_configuration_private_certificate_template_allowed_domains_template
  allow_bare_domains = var.sm_configuration_private_certificate_template_allow_bare_domains
  allow_subdomains = var.sm_configuration_private_certificate_template_allow_subdomains
  allow_glob_domains = var.sm_configuration_private_certificate_template_allow_glob_domains
  allow_any_name = var.sm_configuration_private_certificate_template_allow_any_name
  enforce_hostnames = var.sm_configuration_private_certificate_template_enforce_hostnames
  allow_ip_sans = var.sm_configuration_private_certificate_template_allow_ip_sans
  allowed_uri_sans = var.sm_configuration_private_certificate_template_allowed_uri_sans
  allowed_other_sans = var.sm_configuration_private_certificate_template_allowed_other_sans
  server_flag = var.sm_configuration_private_certificate_template_server_flag
  client_flag = var.sm_configuration_private_certificate_template_client_flag
  code_signing_flag = var.sm_configuration_private_certificate_template_code_signing_flag
  email_protection_flag = var.sm_configuration_private_certificate_template_email_protection_flag
  key_usage = var.sm_configuration_private_certificate_template_key_usage
  ext_key_usage = var.sm_configuration_private_certificate_template_ext_key_usage
  ext_key_usage_oids = var.sm_configuration_private_certificate_template_ext_key_usage_oids
  use_csr_common_name = var.sm_configuration_private_certificate_template_use_csr_common_name
  use_csr_sans = var.sm_configuration_private_certificate_template_use_csr_sans
  require_cn = var.sm_configuration_private_certificate_template_require_cn
  policy_identifiers = var.sm_configuration_private_certificate_template_policy_identifiers
  basic_constraints_valid_for_non_ca = var.sm_configuration_private_certificate_template_basic_constraints_valid_for_non_ca
}

// Provision sm_configuration_public_certificate_CA_Lets_Encrypt resource instance
resource "ibm_sm_configuration_public_certificate_CA_Lets_Encrypt" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  config_type = var.sm_configuration_public_certificate_CA_Lets_Encrypt_config_type
  lets_encrypt_environment = var.sm_configuration_public_certificate_CA_Lets_Encrypt_lets_encrypt_environment
  lets_encrypt_private_key = var.sm_configuration_public_certificate_CA_Lets_Encrypt_lets_encrypt_private_key
  lets_encrypt_preferred_chain = var.sm_configuration_public_certificate_CA_Lets_Encrypt_lets_encrypt_preferred_chain
}

// Provision sm_en_registration resource instance
resource "ibm_sm_en_registration" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  event_notifications_instance_crn = var.sm_en_registration_event_notifications_instance_crn
  event_notifications_source_name = var.sm_en_registration_event_notifications_source_name
  event_notifications_source_description = var.sm_en_registration_event_notifications_source_description
}

// Create sm_secret_group data source
data "ibm_sm_secret_group" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = ibm_sm_secret_group.sm_secret_group_instance.id
}


// Create sm_secret_groups data source
data "ibm_sm_secret_groups" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_secret_version_action data source
data "ibm_sm_secret_version_action" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_secret_version_action_secret_id
  id = var.sm_secret_version_action_id
  secret_version_action_prototype = var.sm_secret_version_action_secret_version_action_prototype
}

// Create sm_public_certificate_action_validate_manual_dns data source
data "ibm_sm_public_certificate_action_validate_manual_dns" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_public_certificate_action_validate_manual_dns_id
  secret_action_prototype = var.sm_public_certificate_action_validate_manual_dns_secret_action_prototype
}

// Create sm_private_certificate_action_revoke data source
data "ibm_sm_private_certificate_action_revoke" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_private_certificate_action_revoke_id
  secret_action_prototype = var.sm_private_certificate_action_revoke_secret_action_prototype
}

// Create sm_private_certificate_configuration_action_sign_csr data source
data "ibm_sm_private_certificate_configuration_action_sign_csr" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_private_certificate_configuration_action_sign_csr_name
  config_action_prototype = var.sm_private_certificate_configuration_action_sign_csr_config_action_prototype
}

// Create sm_secrets data source
data "ibm_sm_secrets" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_secrets_locks data source
data "ibm_sm_secrets_locks" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_secret_versions data source
data "ibm_sm_secret_versions" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_secret_versions_secret_id
}

// Create sm_secret_version_metadata data source
data "ibm_sm_secret_version_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_secret_version_metadata_secret_id
  id = var.sm_secret_version_metadata_id
}

// Create sm_imported_certificate_metadata data source
data "ibm_sm_imported_certificate_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_imported_certificate_metadata_id
}

// Create sm_public_certificate_metadata data source
data "ibm_sm_public_certificate_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_public_certificate_metadata_id
}

// Create sm_kv_secret_metadata data source
data "ibm_sm_kv_secret_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_kv_secret_metadata_id
}

// Create sm_iam_credentials_secret_metadata data source
data "ibm_sm_iam_credentials_secret_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_iam_credentials_secret_metadata_id
}

// Create sm_arbitrary_secret_metadata data source
data "ibm_sm_arbitrary_secret_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_arbitrary_secret_metadata_id
}

// Create sm_username_password_secret_metadata data source
data "ibm_sm_username_password_secret_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_username_password_secret_metadata_id
}

// Create sm_imported_certificate data source
data "ibm_sm_imported_certificate" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_imported_certificate_id
}

// Create sm_public_certificate data source
data "ibm_sm_public_certificate" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_public_certificate_id
}

// Create sm_kv_secret data source
data "ibm_sm_kv_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_kv_secret_id
}

// Create sm_iam_credentials_secret data source
data "ibm_sm_iam_credentials_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_iam_credentials_secret_id
}

// Create sm_arbitrary_secret data source
data "ibm_sm_arbitrary_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_arbitrary_secret_id
}

// Create sm_username_password_secret data source
data "ibm_sm_username_password_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_username_password_secret_id
}

// Create arbitrary_sm_secret_version data source
data "ibm_arbitrary_sm_secret_version" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.arbitrary_sm_secret_version_secret_id
  id = var.arbitrary_sm_secret_version_id
}

// Create sm_private_certificate data source
data "ibm_sm_private_certificate" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_private_certificate_id
}

// Create sm_private_certificate_metadata data source
data "ibm_sm_private_certificate_metadata" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  id = var.sm_private_certificate_metadata_id
}

// Create sm_configuration_private_certificate_root_CA data source
data "ibm_sm_configuration_private_certificate_root_CA" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_configuration_private_certificate_root_CA_name
}

// Create sm_configuration_private_certificate_intermediate_CA data source
data "ibm_sm_configuration_private_certificate_intermediate_CA" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_configuration_private_certificate_intermediate_CA_name
}

// Create sm_configuration_private_certificate_template data source
data "ibm_sm_configuration_private_certificate_template" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_configuration_private_certificate_template_name
}

// Create sm_configurations data source
data "ibm_sm_configurations" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_configuration_public_certificate_CA_Lets_Encrypt data source
data "ibm_sm_configuration_public_certificate_CA_Lets_Encrypt" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_configuration_public_certificate_CA_Lets_Encrypt_name
}

// Create sm_en_registration data source
data "ibm_sm_en_registration" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}