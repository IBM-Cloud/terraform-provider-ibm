provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.region
}

// Provision sm_secret_group resource instance
resource "ibm_sm_secret_group" "sm_secret_group_instance" {
  description = var.sm_secret_group_description
}

// Provision sm_imported_certificate resource instance
resource "ibm_sm_imported_certificate" "sm_imported_certificate_instance" {
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
resource "ibm_sm_public_certificate" "sm_public_certificate_instance" {
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
resource "ibm_sm_kv_secret" "sm_kv_secret_instance" {
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
resource "ibm_sm_iam_credentials_secret"  "sm_iam_credentials_secret_instance" {
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
resource "ibm_sm_arbitrary_secret" "sm_arbitrary_secret_instance" {
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
resource "ibm_sm_username_password_secret" "sm_username_password_secret_instance" {
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

// Provision sm_private_certificate resource instance
resource "ibm_sm_private_certificate" "sm_private_certificate_instance" {
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

// Provision sm_private_certificate_configuration_root_ca resource instance
resource "ibm_sm_private_certificate_configuration_root_ca" "sm_private_certificate_configuration_root_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  crl_disable = var.sm_private_certificate_configuration_root_ca_crl_disable
  crl_distribution_points_encoded = var.sm_private_certificate_configuration_root_ca_crl_distribution_points_encoded
  issuing_certificates_urls_encoded = var.sm_private_certificate_configuration_root_ca_issuing_certificates_urls_encoded
  ttl = var.sm_private_certificate_configuration_root_ca_ttl
}

// Provision sm_private_certificate_configuration_intermediate_ca resource instance
resource "ibm_sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  crl_disable = var.sm_private_certificate_configuration_intermediate_ca_crl_disable
  crl_distribution_points_encoded = var.sm_private_certificate_configuration_intermediate_ca_crl_distribution_points_encoded
  issuing_certificates_urls_encoded = var.sm_private_certificate_configuration_intermediate_ca_issuing_certificates_urls_encoded
  signing_method = var.sm_private_certificate_configuration_intermediate_ca_signing_method
}

// Provision sm_private_certificate_configuration_template resource instance
resource "ibm_sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  certificate_authority = var.sm_private_certificate_configuration_template_certificate_authority
  allowed_secret_groups = var.sm_private_certificate_configuration_template_allowed_secret_groups
  allow_localhost = var.sm_private_certificate_configuration_template_allow_localhost
  allowed_domains = var.sm_private_certificate_configuration_template_allowed_domains
  allowed_domains_template = var.sm_private_certificate_configuration_template_allowed_domains_template
  allow_bare_domains = var.sm_private_certificate_configuration_template_allow_bare_domains
  allow_subdomains = var.sm_private_certificate_configuration_template_allow_subdomains
  allow_glob_domains = var.sm_private_certificate_configuration_template_allow_glob_domains
  allow_any_name = var.sm_private_certificate_configuration_template_allow_any_name
  enforce_hostnames = var.sm_private_certificate_configuration_template_enforce_hostnames
  allow_ip_sans = var.sm_private_certificate_configuration_template_allow_ip_sans
  allowed_uri_sans = var.sm_private_certificate_configuration_template_allowed_uri_sans
  allowed_other_sans = var.sm_private_certificate_configuration_template_allowed_other_sans
  server_flag = var.sm_private_certificate_configuration_template_server_flag
  client_flag = var.sm_private_certificate_configuration_template_client_flag
  code_signing_flag = var.sm_private_certificate_configuration_template_code_signing_flag
  email_protection_flag = var.sm_private_certificate_configuration_template_email_protection_flag
  key_usage = var.sm_private_certificate_configuration_template_key_usage
  ext_key_usage = var.sm_private_certificate_configuration_template_ext_key_usage
  ext_key_usage_oids = var.sm_private_certificate_configuration_template_ext_key_usage_oids
  use_csr_common_name = var.sm_private_certificate_configuration_template_use_csr_common_name
  use_csr_sans = var.sm_private_certificate_configuration_template_use_csr_sans
  require_cn = var.sm_private_certificate_configuration_template_require_cn
  policy_identifiers = var.sm_private_certificate_configuration_template_policy_identifiers
  basic_constraints_valid_for_non_ca = var.sm_private_certificate_configuration_template_basic_constraints_valid_for_non_ca
}

// Provision sm_public_certificate_configuration_ca_lets_encrypt resource instance
resource "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  lets_encrypt_environment = var.sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_environment
  lets_encrypt_private_key = var.sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_private_key
  lets_encrypt_preferred_chain = var.sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_preferred_chain
}

// Provision sm_public_certificate_configuration_dns_cis resource instance
resource "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  cloud_internet_services_apikey = var.sm_public_certificate_configuration_dns_cis_cloud_internet_services_apikey
  cloud_internet_services_crn = var.sm_public_certificate_configuration_dns_cis_cloud_internet_services_crn
}

// Provision sm_public_certificate_configuration_dns_classic_infrastructure resource instance
resource "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  classic_infrastructure_username = var.sm_public_certificate_configuration_dns_classic_infrastructure_classic_infrastructure_username
  classic_infrastructure_password = var.sm_public_certificate_configuration_dns_classic_infrastructure_classic_infrastructure_password
}

// Provision sm_en_registration resource instance
resource "ibm_sm_en_registration" "sm_en_registration_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  event_notifications_instance_crn = var.sm_en_registration_event_notifications_instance_crn
  event_notifications_source_name = var.sm_en_registration_event_notifications_source_name
  event_notifications_source_description = var.sm_en_registration_event_notifications_source_description
}

// Create sm_secret_group data source
data "ibm_sm_secret_group" "sm_secret_group_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_group_id = ibm_sm_secret_group.sm_secret_group_instance.secret_group_id
}

// Create sm_secret_groups data source
data "ibm_sm_secret_groups" "sm_secret_groups_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_secrets data source
data "ibm_sm_secrets" "sm_secrets_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_imported_certificate_metadata data source
data "ibm_sm_imported_certificate_metadata" "sm_imported_certificate_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_imported_certificate_metadata_id
}

// Create sm_public_certificate_metadata data source
data "ibm_sm_public_certificate_metadata" "sm_public_certificate_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_public_certificate_metadata_id
}

// Create sm_kv_secret_metadata data source
data "ibm_sm_kv_secret_metadata" "sm_kv_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_kv_secret_metadata_id
}

// Create sm_iam_credentials_secret_metadata data source
data "ibm_sm_iam_credentials_secret_metadata" "sm_iam_credentials_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_iam_credentials_secret_metadata_id
}

// Create sm_arbitrary_secret_metadata data source
data "ibm_sm_arbitrary_secret_metadata" "sm_arbitrary_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_arbitrary_secret_metadata_id
}

// Create sm_username_password_secret_metadata data source
data "ibm_sm_username_password_secret_metadata" "sm_username_password_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_username_password_secret_metadata_id
}

// Create sm_imported_certificate data source
data "ibm_sm_imported_certificate" "sm_imported_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_imported_certificate_id
}

// Create sm_public_certificate data source
data "ibm_sm_public_certificate" "sm_public_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_public_certificate_id
}

// Create sm_kv_secret data source
data "ibm_sm_kv_secret" "sm_kv_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_kv_secret_id
}

// Create sm_iam_credentials_secret data source
data "ibm_sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_iam_credentials_secret_id
}

// Create sm_arbitrary_secret data source
data "ibm_sm_arbitrary_secret" "sm_arbitrary_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_arbitrary_secret_id
}

// Create sm_username_password_secret data source
data "ibm_sm_username_password_secret" "sm_username_password_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_username_password_secret_id
}

// Create sm_private_certificate data source
data "ibm_sm_private_certificate" "sm_private_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_private_certificate_id
}

// Create sm_private_certificate_metadata data source
data "ibm_sm_private_certificate_metadata" "sm_private_certificate_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  secret_id = var.sm_private_certificate_metadata_id
}

// Create sm_private_certificate_configuration_root_ca data source
data "ibm_sm_private_certificate_configuration_root_ca" "sm_private_certificate_configuration_root_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_private_certificate_configuration_root_ca_name
}

// Create sm_private_certificate_configuration_intermediate_ca data source
data "ibm_sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_private_certificate_configuration_intermediate_ca_name
}

// Create sm_private_certificate_configuration_template data source
data "ibm_sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_private_certificate_configuration_template_name
}

// Create sm_configurations data source
data "ibm_sm_configurations" "sm_configurations_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
}

// Create sm_public_certificate_configuration_ca_lets_encrypt data source
data "ibm_sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_public_certificate_configuration_ca_lets_encrypt_name
}

// Create sm_public_certificate_configuration_dns_cis data source
data "ibm_sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_public_certificate_configuration_dns_cis_name
}

// Create sm_public_certificate_configuration_dns_classic_infrastructure data source
data "ibm_sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name = var.sm_public_certificate_configuration_dns_classic_infrastructure_name
}

// Create sm_en_registration data source
data "ibm_sm_en_registration" "sm_en_registration_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
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
