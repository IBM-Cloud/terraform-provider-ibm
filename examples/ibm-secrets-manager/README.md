# Example for SecretsManagerV2

This example illustrates how to use the SecretsManagerV2

These types of resources are supported:
* SecretGroup
* ImportedCertificate
* PublicCertificate
* KVSecret
* IAMCredentialsSecret
* ArbitrarySecret
* UsernamePasswordSecret
* PrivateCertificate
* PrivateCertificateConfigurationRootCA
* PrivateCertificateConfigurationIntermediateCA
* PrivateCertificateConfigurationTemplate
* PublicCertificateConfigurationCALetsEncrypt
* PublicCertificateConfigurationDNSCloudInternetServices
* PublicCertificateConfigurationDNSClassicInfrastructure
* NotificationsRegistration
* IAMCredentialsConfiguration

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## SecretsManagerV2 resources

sm_secret_group resource:

```hcl
resource "sm_secret_group" "sm_secret_group_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name          = var.sm_secret_group_name
  description   = var.sm_secret_group_description
}
```
sm_imported_certificate resource:

```hcl
resource "sm_imported_certificate" "sm_imported_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_imported_certificate_name
  custom_metadata = var.sm_imported_certificate_custom_metadata
  description = var.sm_imported_certificate_description
  expiration_date = var.sm_imported_certificate_expiration_date
  labels = var.sm_imported_certificate_labels
  secret_group_id = var.sm_imported_certificate_secret_group_id
  certificate = var.sm_imported_certificate_certificate
  intermediate = var.sm_imported_certificate_intermediate
  private_key = var.sm_imported_certificate_private_key
}
```
sm_public_certificate resource:

```hcl
resource "sm_public_certificate" "sm_public_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_public_certificate_name
  custom_metadata = var.sm_public_certificate_custom_metadata
  description = var.sm_public_certificate_description
  expiration_date = var.sm_public_certificate_expiration_date
  labels = var.sm_public_certificate_labels
  secret_group_id = var.sm_public_certificate_secret_group_id
  rotation = var.sm_public_certificate_rotation
}
```
sm_kv_secret resource:

```hcl
resource "sm_kv_secret" "sm_kv_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_kv_secret_name
  custom_metadata = var.sm_kv_secret_custom_metadata
  description = var.sm_kv_secret_description
  labels = var.sm_kv_secret_labels
  secret_group_id = var.sm_kv_secret_secret_group_id
  data = var.sm_kv_secret_data
}
```
sm_iam_credentials_secret resource:

```hcl
resource "sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_iam_credentials_secret_name
  custom_metadata = var.sm_iam_credentials_secret_custom_metadata
  description = var.sm_iam_credentials_secret_description
  labels = var.sm_iam_credentials_secret_labels
  secret_group_id = var.sm_iam_credentials_secret_secret_group_id
  ttl = var.sm_iam_credentials_secret_ttl
  access_groups = var.sm_iam_credentials_secret_access_groups
  reuse_api_key = var.sm_iam_credentials_secret_reuse_api_key
  rotation = var.sm_iam_credentials_secret_rotation
}
```
sm_service_credentials_secret resource:

```hcl
resource "ibm_sm_service_credentials_secret" "sm_service_credentials_secret" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  endpoint_type    = var.endpoint_type
  name 			= var.sm_service_credentials_secret_name
  custom_metadata = { my_key = jsonencode(var.sm_service_credentials_secret_custom_metadata) }
  description = var.sm_service_credentials_secret_description
  labels = var.sm_service_credentials_secret_labels
  rotation = var.sm_service_credentials_secret_rotation
  secret_group_id = var.sm_service_credentials_secret_secret_group_id
  source_service {
	instance {
		crn = var.sm_service_credentials_secret_source_service_instance_crn
	}
	role {
		crn = var.sm_service_credentials_secret_source_service_role_crn
	}
	parameters = var.sm_service_credentials_secret_source_service_parameters
  }
  ttl = var.sm_service_credentials_secret_ttl
}
```
sm_arbitrary_secret resource:

```hcl
resource "sm_arbitrary_secret" "sm_arbitrary_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_arbitrary_secret_name
  custom_metadata = var.sm_arbitrary_secret_custom_metadata
  description = var.sm_arbitrary_secret_description
  expiration_date = var.sm_arbitrary_secret_expiration_date
  labels = var.sm_arbitrary_secret_labels
  secret_group_id = var.sm_arbitrary_secret_secret_group_id
  payload = var.sm_arbitrary_secret_payload
}
```
sm_username_password_secret resource:

```hcl
resource "sm_username_password_secret" "sm_username_password_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_username_password_secret_name
  custom_metadata = var.sm_username_password_secret_custom_metadata
  description = var.sm_username_password_secret_description
  expiration_date = var.sm_username_password_secret_expiration_date
  labels = var.sm_username_password_secret_labels
  secret_group_id = var.sm_username_password_secret_secret_group_id
  rotation = var.sm_username_password_secret_rotation
  username = var.sm_username_password_secret_username
  password = var.sm_username_password_secret_password
}
```
sm_private_certificate resource:

```hcl
resource "sm_private_certificate" "sm_private_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name 			= var.sm_private_certificate_name
  custom_metadata = var.sm_private_certificate_custom_metadata
  description = var.sm_private_certificate_description
  expiration_date = var.sm_private_certificate_expiration_date
  labels = var.sm_private_certificate_labels
  secret_group_id = var.sm_private_certificate_secret_group_id
  rotation = var.sm_private_certificate_rotation
  certificate_template = var.sm_private_certificate_certificate_template
}
```
sm_private_certificate_configuration_root_ca resource:

```hcl
resource "sm_private_certificate_configuration_root_ca" "sm_private_certificate_configuration_root_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name          = var.sm_private_certificate_configuration_root_ca_name
  crl_disable = var.sm_private_certificate_configuration_root_ca_crl_disable
  crl_distribution_points_encoded = var.sm_private_certificate_configuration_root_ca_crl_distribution_points_encoded
  issuing_certificates_urls_encoded = var.sm_private_certificate_configuration_root_ca_issuing_certificates_urls_encoded
  ttl = var.sm_private_certificate_configuration_root_ca_ttl
}
```
sm_private_certificate_configuration_intermediate_ca resource:

```hcl
resource "sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name           = var.sm_private_certificate_configuration_intermediate_ca_name
  crl_disable = var.sm_private_certificate_configuration_intermediate_ca_crl_disable
  crl_distribution_points_encoded = var.sm_private_certificate_configuration_intermediate_ca_crl_distribution_points_encoded
  issuing_certificates_urls_encoded = var.sm_private_certificate_configuration_intermediate_ca_issuing_certificates_urls_encoded
  signing_method = var.sm_private_certificate_configuration_intermediate_ca_signing_method
}
```
sm_private_certificate_configuration_template resource:

```hcl
resource "sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name          = var.sm_private_certificate_configuration_template_name
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
```
sm_public_certificate_configuration_ca_lets_encrypt resource:

```hcl
resource "sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name          = var.sm_public_certificate_configuration_ca_lets_encrypt_name
  lets_encrypt_environment = var.sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_environment
  lets_encrypt_private_key = var.sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_private_key
  lets_encrypt_preferred_chain = var.sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_preferred_chain
}
```
sm_public_certificate_configuration_dns_cis resource:

```hcl
resource "sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name          = var.sm_public_certificate_configuration_dns_cis_name
  cloud_internet_services_apikey = var.sm_public_certificate_configuration_dns_cis_cloud_internet_services_apikey
  cloud_internet_services_crn = var.sm_public_certificate_configuration_dns_cis_cloud_internet_services_crn
}
```
sm_public_certificate_configuration_dns_classic_infrastructure resource:

```hcl
resource "sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name          = var.sm_public_certificate_configuration_dns_classic_infrastructure_name
  classic_infrastructure_username = var.sm_public_certificate_configuration_dns_classic_infrastructure_classic_infrastructure_username
  classic_infrastructure_password = var.sm_public_certificate_configuration_dns_classic_infrastructure_classic_infrastructure_password
}
```
sm_en_registration resource:

```hcl
resource "sm_en_registration" "sm_en_registration_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  event_notifications_instance_crn = var.sm_en_registration_event_notifications_instance_crn
  event_notifications_source_name = var.sm_en_registration_event_notifications_source_name
  event_notifications_source_description = var.sm_en_registration_event_notifications_source_description
}
```

## SecretsManagerV2 Data sources

sm_secret_group data source:

```hcl
data "sm_secret_group" "sm_secret_group_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_group_id = ibm_sm_secret_group.sm_secret_group_instance.secret_group_id
}
```
sm_secret_groups data source:

```hcl
data "sm_secret_groups" "sm_secret_groups_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
}
```
sm_secrets data source:

```hcl
data "sm_secrets" "sm_secrets_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
}
```

sm_imported_certificate_metadata data source:

```hcl
data "sm_imported_certificate_metadata" "sm_imported_certificate_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_imported_certificate_metadata_id
}
```
sm_public_certificate_metadata data source:

```hcl
data "sm_public_certificate_metadata" "sm_public_certificate_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_public_certificate_metadata_id
}
```
sm_kv_secret_metadata data source:

```hcl
data "sm_kv_secret_metadata" "sm_kv_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_kv_secret_metadata_id
}
```
sm_iam_credentials_secret_metadata data source:

```hcl
data "sm_iam_credentials_secret_metadata" "sm_iam_credentials_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_iam_credentials_secret_metadata_id
}
```
sm_service_credentials_secret_metadata data source:

```hcl
data "sm_service_credentials_secret_metadata" "sm_service_credentials_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_service_credentials_secret_metadata_id
}
```
sm_arbitrary_secret_metadata data source:

```hcl
data "sm_arbitrary_secret_metadata" "sm_arbitrary_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_arbitrary_secret_metadata_id
}
```
sm_username_password_secret_metadata data source:

```hcl
data "sm_username_password_secret_metadata" "sm_username_password_secret_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_username_password_secret_metadata_id
}
```
sm_imported_certificate data source:

```hcl
data "sm_imported_certificate" "sm_imported_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_imported_certificate_id
}
```
sm_public_certificate data source:

```hcl
data "sm_public_certificate" "sm_public_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_public_certificate_id
}
```
sm_kv_secret data source:

```hcl
data "sm_kv_secret" "sm_kv_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_kv_secret_id
}
```
sm_iam_credentials_secret data source:

```hcl
data "sm_iam_credentials_secret" "sm_iam_credentials_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_iam_credentials_secret_id
}
```
sm_service_credentials_secret data source:

```hcl
data "sm_service_credentials_secret" "sm_service_credentials_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_service_credentials_secret_id
}
```
sm_arbitrary_secret data source:

```hcl
data "sm_arbitrary_secret" "sm_arbitrary_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_arbitrary_secret_id
}
```
sm_username_password_secret data source:

```hcl
data "sm_username_password_secret" "sm_username_password_secret_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_username_password_secret_id
}
```

sm_private_certificate data source:

```hcl
data "sm_private_certificate" "sm_private_certificate_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_private_certificate_id
}
```
sm_private_certificate_metadata data source:

```hcl
data "sm_private_certificate_metadata" "sm_private_certificate_metadata_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  secret_id = var.sm_private_certificate_metadata_id
}
```
sm_private_certificate_configuration_root_ca data source:

```hcl
data "sm_private_certificate_configuration_root_ca" "sm_private_certificate_configuration_root_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name = var.sm_private_certificate_configuration_root_ca_name
}
```
sm_private_certificate_configuration_intermediate_ca data source:

```hcl
data "sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name = var.sm_private_certificate_configuration_intermediate_ca_name
}
```
sm_private_certificate_configuration_template data source:

```hcl
data "sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name = var.sm_private_certificate_configuration_template_name
}
```

sm_public_certificate_configuration_ca_lets_encrypt data source:

```hcl
data "sm_public_certificate_configuration_ca_lets_encrypt" "sm_public_certificate_configuration_ca_lets_encrypt_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name = var.sm_public_certificate_configuration_ca_lets_encrypt_name
}
```
sm_public_certificate_configuration_dns_cis data source:

```hcl
data "sm_public_certificate_configuration_dns_cis" "sm_public_certificate_configuration_dns_cis_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name = var.sm_public_certificate_configuration_dns_cis_name
}
```
sm_public_certificate_configuration_dns_classic_infrastructure data source:

```hcl
data "sm_public_certificate_configuration_dns_classic_infrastructure" "sm_public_certificate_configuration_dns_classic_infrastructure_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
  name = var.sm_public_certificate_configuration_dns_classic_infrastructure_name
}
```
sm_en_registration data source:

```hcl
data "sm_en_registration" "sm_en_registration_instance" {
  instance_id   = var.secrets_manager_instance_id
  region        = var.region
}
```
## Assumptions

1. TODO

## Notes


## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.22.0 |

## Inputs

| Name                                   | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Type           | Default   | Required |
|----------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------|-----------|----------|
| ibmcloud\_api\_key                     | IBM Cloud API key                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | `string`       |           | true     |
| region                                 | Secrets Manager Instance region                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | `string`       | us-south  | false    |
| secrets\_manager\_instance\_id         | Secrets Manager Instance GUID                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | `string`       |           | true     |
| instance\_id                           | Secrets Manager Instance GUID                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | `string`       |           | true     |
| endpoint\_type                         | Secrets manager endpoint type                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | `string`       | `private` | false    |
| description                            | An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.                                                                                                                                                                                                                                                                                                                                    | `string`       | false     |
| custom_metadata                        | The secret metadata that a user can customize.                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `map()`        | false     |
| description                            | An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.                                                                                                                                                                                                                                                                                                                                          | `string`       | false     |
| expiration_date                        | The date a secret is expired. The date format follows RFC 3339.                                                                                                                                                                                                                                                                                                                                                                                                                                           | ``             | false     |
| labels                                 | Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.                                                                                                                                                                                                                                                                                                                                                                                                            | `list(string)` | false     |
| secret_group_id                        | A UUID identifier, or `default` secret group.                                                                                                                                                                                                                                                                                                                                                                                                                                                             | `string`       | false     |
| secret_type                            | The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.                                                                                                                                                                                                                                                                                                                                                           | `string`       | false     |
| certificate                            | The PEM-encoded contents of your certificate.                                                                                                                                                                                                                                                                                                                                                                                                                                                             | `string`       | false     |
| intermediate                           | (Optional) The PEM-encoded intermediate certificate to associate with the root certificate.                                                                                                                                                                                                                                                                                                                                                                                                               | `string`       | false     |
| private_key                            | (Optional) The PEM-encoded private key to associate with the certificate.                                                                                                                                                                                                                                                                                                                                                                                                                                 | `string`       | false     |
| custom_metadata                        | The secret metadata that a user can customize.                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `map()`        | false     |
| rotation                               | Determines whether Secrets Manager rotates your secrets automatically.                                                                                                                                                                                                                                                                                                                                                                                                                                    | ``             | false     |
| source_service                         | The properties required for creating the service credentials for the specified source service instance.                                                                                                                                                                                                                                                                                                                                                                                                   | ``             | false     |
| data                                   | The payload data of a key-value secret.                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `map()`        | false     |
| ttl                                    | The time-to-live (TTL) or lease duration to assign to generated credentials.The TTL defines for how long generated credentials remain valid. For iam_credentials secret TTL is mandatory. The minimum duration is 1 minute. The maximum is 90 days. For service_credentials secret TTL is optional, if set the minimum duration is 1 day. The maximum is 90 days. The TTL defaults to 0 which means no TTL.                                                                                               | `string`       | false     |
| access_groups                          | Access Groups that you can use for an `iam_credentials` secret.Up to 10 Access Groups can be used for each secret.                                                                                                                                                                                                                                                                                                                                                                                        | `list(string)` | false     |
| service_id                             | The service ID under which the API key (see the `api_key` field) is created.If you omit this parameter, Secrets Manager generates a new service ID for your secret at its creation and adds it to the access groups that you assign.Optionally, you can use this field to provide your own service ID if you prefer to manage its access directly or retain the service ID after your secret expires, is rotated, or deleted. If you provide a service ID, do not include the `access_groups` parameter.  | `string`       | false     |
| reuse_api_key                          | Determines whether to use the same service ID and API key for future read operations on an`iam_credentials` secret.If it is set to `true`, the service reuses the current credentials. If it is set to `false`, a new service ID and API key are generated each time that the secret is read or accessed.                                                                                                                                                                                                 | `bool`         | false     |
| payload                                | The arbitrary secret's data payload.                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | `string`       | false     |
| username                               | The username that is assigned to the secret.                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `string`       | false     |
| password                               | The password that is assigned to the secret.                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `string`       | false     |
| secret_id                              | The ID of the secret.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `string`       | true      |
| certificate_template                   | The name of the certificate template.                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `string`       | false     |
| config_type                            | Th configuration type.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `string`       | false     |
| crl_disable                            | Disables or enables certificate revocation list (CRL) building.If CRL building is disabled, a signed but zero-length CRL is returned when downloading the CRL. If CRL building is enabled, it will rebuild the CRL.                                                                                                                                                                                                                                                                                       | `bool`         | false     |
| crl_distribution_points_encoded        | Determines whether to encode the certificate revocation list (CRL) distribution points in the certificates that are issued by this certificate authority.                                                                                                                                                                                                                                                                                                                                                 | `bool`         | false     |
| issuing_certificates_urls_encoded      | Determines whether to encode the URL of the issuing certificate in the certificates that are issued by this certificate authority.                                                                                                                                                                                                                                                                                                                                                                        | `bool`         | false     |
| ttl                                    | The requested time-to-live (TTL) for certificates that are created by this CA. This field's value cannot be longer than the `max_ttl` limit.The value can be supplied as a string representation of a duration in hours, for example '8760h'. In the API response, this value is returned in seconds (integer).                                                                                                                                                                                           | `string`       | false     |
| signing_method                         | The signing method to use with this certificate authority to generate private certificates.You can choose between internal or externally signed options. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-intermediate-certificate-authorities).                                                                                                                                                                                                     | `string`       | false     |
| certificate_authority                  | The name of the intermediate certificate authority.                                                                                                                                                                                                                                                                                                                                                                                                                                                       | `string`       | false     |
| allowed_secret_groups                  | Scopes the creation of private certificates to only the secret groups that you specify.This field can be supplied as a comma-delimited list of secret group IDs.                                                                                                                                                                                                                                                                                                                                          | `string`       | false     |
| allow_localhost                        | Determines whether to allow `localhost` to be included as one of the requested common names.                                                                                                                                                                                                                                                                                                                                                                                                              | `bool`         | false     |
| allowed_domains                        | The domains to define for the certificate template. This property is used along with the `allow_bare_domains` and `allow_subdomains` options.                                                                                                                                                                                                                                                                                                                                                             | `list(string)` | false     |
| allowed_domains_template               | Determines whether to allow the domains that are supplied in the `allowed_domains` field to contain access control list (ACL) templates.                                                                                                                                                                                                                                                                                                                                                                  | `bool`         | false     |
| allow_bare_domains                     | Determines whether to allow clients to request private certificates that match the value of the actual domains on the final certificate.For example, if you specify `example.com` in the `allowed_domains` field, you grant clients the ability to request a certificate that contains the name `example.com` as one of the DNS values on the final certificate.**Important:** In some scenarios, allowing bare domains can be considered a security risk.                                                | `bool`         | false     |
| allow_subdomains                       | Determines whether to allow clients to request private certificates with common names (CN) that are subdomains of the CNs that are allowed by the other certificate template options. This includes wildcard subdomains.For example, if `allowed_domains` has a value of `example.com` and `allow_subdomains`is set to `true`, then the following subdomains are allowed: `foo.example.com`, `bar.example.com`, `*.example.com`.**Note:** This field is redundant if you use the `allow_any_name` option. | `bool`         | false     |
| allow_glob_domains                     | Determines whether to allow glob patterns, for example, `ftp*.example.com`, in the names that are specified in the `allowed_domains` field.If set to `true`, clients are allowed to request private certificates with names that match the glob patterns.                                                                                                                                                                                                                                                 | `bool`         | false     |
| allow_any_name                         | Determines whether to allow clients to request a private certificate that matches any common name.                                                                                                                                                                                                                                                                                                                                                                                                        | `bool`         | false     |
| enforce_hostnames                      | Determines whether to enforce only valid host names for common names, DNS Subject Alternative Names, and the host section of email addresses.                                                                                                                                                                                                                                                                                                                                                             | `bool`         | false     |
| allow_ip_sans                          | Determines whether to allow clients to request a private certificate with IP Subject Alternative Names.                                                                                                                                                                                                                                                                                                                                                                                                   | `bool`         | false     |
| allowed_uri_sans                       | The URI Subject Alternative Names to allow for private certificates.Values can contain glob patterns, for example `spiffe://hostname/_*`.                                                                                                                                                                                                                                                                                                                                                                 | `list(string)` | false     |
| allowed_other_sans                     | The custom Object Identifier (OID) or UTF8-string Subject Alternative Names (SANs) to allow for private certificates.The format for each element in the list is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`. To allow any value for an OID, use `*` as its value. Alternatively, specify a single `*` to allow any `other_sans` input.                                                                                                                             | `list(string)` | false     |
| server_flag                            | Determines whether private certificates are flagged for server use.                                                                                                                                                                                                                                                                                                                                                                                                                                       | `bool`         | false     |
| client_flag                            | Determines whether private certificates are flagged for client use.                                                                                                                                                                                                                                                                                                                                                                                                                                       | `bool`         | false     |
| code_signing_flag                      | Determines whether private certificates are flagged for code signing use.                                                                                                                                                                                                                                                                                                                                                                                                                                 | `bool`         | false     |
| email_protection_flag                  | Determines whether private certificates are flagged for email protection use.                                                                                                                                                                                                                                                                                                                                                                                                                             | `bool`         | false     |
| key_usage                              | The allowed key usage constraint to define for private certificates.You can find valid values in the [Go x509 package documentation](https://pkg.go.dev/crypto/x509#KeyUsage). Omit the `KeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list.                                                                                                                                                                                | `list(string)` | false     |
| ext_key_usage                          | The allowed extended key usage constraint on private certificates.You can find valid values in the [Go x509 package documentation](https://golang.org/pkg/crypto/x509/#ExtKeyUsage). Omit the `ExtKeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list.                                                                                                                                                                       | `list(string)` | false     |
| ext_key_usage_oids                     | A list of extended key usage Object Identifiers (OIDs).                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `list(string)` | false     |
| use_csr_common_name                    | When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the common name (CN) from a certificate signing request (CSR) instead of the CN that's included in the data of the certificate.Does not include any requested Subject Alternative Names (SANs) in the CSR. To use the alternative names, include the `use_csr_sans` property.                                                                                                                | `bool`         | false     |
| use_csr_sans                           | When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the Subject Alternative Names(SANs) from a certificate signing request (CSR) instead of the SANs that are included in the data of the certificate.Does not include the common name in the CSR. To use the common name, include the `use_csr_common_name` property.                                                                                                                           | `bool`         | false     |
| require_cn                             | Determines whether to require a common name to create a private certificate.By default, a common name is required to generate a certificate. To make the `common_name` field optional, set the `require_cn` option to `false`.                                                                                                                                                                                                                                                                            | `bool`         | false     |
| policy_identifiers                     | A list of policy Object Identifiers (OIDs).                                                                                                                                                                                                                                                                                                                                                                                                                                                               | `list(string)` | false     |
| basic_constraints_valid_for_non_ca     | Determines whether to mark the Basic Constraints extension of an issued private certificate as valid for non-CA certificates.                                                                                                                                                                                                                                                                                                                                                                             | `bool`         | false     |
| lets_encrypt_environment               | The configuration of the Let's Encrypt CA environment.                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `string`       | false     |
| lets_encrypt_private_key               | The PEM encoded private key of your Lets Encrypt account.                                                                                                                                                                                                                                                                                                                                                                                                                                                 | `string`       | false     |
| lets_encrypt_preferred_chain           | Prefer the chain with an issuer matching this Subject Common Name.                                                                                                                                                                                                                                                                                                                                                                                                                                        | `string`       | false     |
| event_notifications_instance_crn       | A CRN that uniquely identifies an IBM Cloud resource.                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `string`       | true      |
| event_notifications_source_name        | The name that is displayed as a source that is in your Event Notifications instance.                                                                                                                                                                                                                                                                                                                                                                                                                      | `string`       | true      |
| event_notifications_source_description | An optional description for the source  that is in your Event Notifications instance.                                                                                                                                                                                                                                                                                                                                                                                                                     | `string`       | false     |
| secret_group_id                        | The ID of the secret group.                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | `string`       | true      |
| secret_id                              | The ID of the secret.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `string`       | true      |
| name                                   | The name of the configuration.                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `string`       | true      |

## Outputs

| Name                                                 | Description                                                 |
|------------------------------------------------------|-------------------------------------------------------------|
| secrets\_manager\_secrets                            | secrets\_manager\_secrets object                            |
| secrets\_manager\_secret                             | secrets\_manager\_secret object                             |
| sm_secret_group                                      | sm_secret_group object                                      |
| sm_imported_certificate                              | sm_imported_certificate object                              |
| sm_public_certificate                                | sm_public_certificate object                                |
| sm_kv_secret                                         | sm_kv_secret object                                         |
| sm_iam_credentials_secret                            | sm_iam_credentials_secret object                            |
| sm_service_credentials_secret                        | sm_service_credentials_secret object                        |
| sm_arbitrary_secret                                  | sm_arbitrary_secret object                                  |
| sm_username_password_secret                          | sm_username_password_secret object                          |
| sm_private_certificate                               | sm_private_certificate object                               |
| sm_private_certificate_configuration_root_ca         | sm_private_certificate_configuration_root_ca object         |
| sm_private_certificate_configuration_intermediate_ca | sm_private_certificate_configuration_intermediate_ca object |
| sm_private_certificate_configuration_template        | sm_private_certificate_configuration_template object        |
| sm_public_certificate_configuration_ca_lets_encrypt  | sm_public_certificate_configuration_ca_lets_encrypt object  |
| sm_en_registration                                   | sm_en_registration object                                   |
| sm_secret_group                                      | sm_secret_group object                                      |
| sm_secret_groups                                     | sm_secret_groups object                                     |
| sm_secrets                                           | sm_secrets object                                           |
| sm_imported_certificate_metadata                     | sm_imported_certificate_metadata object                     |
| sm_public_certificate_metadata                       | sm_public_certificate_metadata object                       |
| sm_kv_secret_metadata                                | sm_kv_secret_metadata object                                |
| sm_iam_credentials_secret_metadata                   | sm_iam_credentials_secret_metadata object                   |
| sm_service_credentials_secret_metadata               | sm_service_credentials_secret_metadata object               |
| sm_arbitrary_secret_metadata                         | sm_arbitrary_secret_metadata object                         |
| sm_username_password_secret_metadata                 | sm_username_password_secret_metadata object                 |
| sm_private_certificate_metadata                      | sm_private_certificate_metadata object                      |
| sm_configurations                                    | sm_configurations object                                    |
