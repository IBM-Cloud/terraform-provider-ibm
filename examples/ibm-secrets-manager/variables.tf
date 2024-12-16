variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "secrets_manager_instance_id" {
  description = "Secrets Manager Instance GUID"
  type        = string
}

variable "region" {
  description = "Secrets Manager Instance region"
  default     = null
}

variable "endpoint_type" {
  description = "Secrets Manager endpoint type"
  type        = string
  default     = "private"
}

// Resource arguments for sm_secret_group
variable "sm_secret_group_description" {
  description = "An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this group."
}

variable "sm_secret_group_name" {
  description = "The name of your existing secret group."
  type        = string
  default     = "my-group-name"
}

// Resource arguments for sm_imported_certificate
variable "sm_imported_certificate_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-imported-cert-secret"
}
variable "sm_imported_certificate_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_imported_certificate_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_imported_certificate_expiration_date" {
  description = "The date a secret is expired. The date format follows RFC 3339."
  type        = string
  default     = "2022-04-12T23:20:50.520Z"
}
variable "sm_imported_certificate_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_imported_certificate_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_imported_certificate_certificate" {
  description = "The PEM-encoded contents of your certificate."
  type        = string
  default     = "certificate"
}
variable "sm_imported_certificate_intermediate" {
  description = "(Optional) The PEM-encoded intermediate certificate to associate with the root certificate."
  type        = string
  default     = "intermediate"
}
variable "sm_imported_certificate_private_key" {
  description = "(Optional) The PEM-encoded private key to associate with the certificate."
  type        = string
  default     = "private_key"
}

// Resource arguments for sm_public_certificate
variable "sm_public_certificate_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-public-cert-secret"
}
variable "sm_public_certificate_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_public_certificate_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_public_certificate_expiration_date" {
  description = "The date a secret is expired. The date format follows RFC 3339."
  type        = string
  default     = "2022-04-12T23:20:50.520Z"
}
variable "sm_public_certificate_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_public_certificate_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}

// Resource arguments for sm_public_certificate_action_validate_manual_dns
variable "sm_public_certificate_action_validate_manual_dns_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Resource arguments for sm_kv_secret
variable "sm_kv_secret_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-kv-secret"
}
variable "sm_kv_secret_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_kv_secret_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_kv_secret_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_kv_secret_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_kv_secret_data" {
  description = "The payload data of a key-value secret."
  type        = any
  default     = "anything as a string"
}

// Resource arguments for sm_iam_credentials_secret
variable "sm_iam_credentials_secret_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-iam-credentials-secret"
}
variable "sm_iam_credentials_secret_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_iam_credentials_secret_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_iam_credentials_secret_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_iam_credentials_secret_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_iam_credentials_secret_ttl" {
  description = "The time-to-live (TTL) or lease duration to assign to generated credentials.For `iam_credentials` secrets, the TTL defines for how long each generated API key remains valid. The value can be either an integer that specifies the number of seconds, or the string representation of a duration, such as `120m` or `24h`.Minimum duration is 1 minute. Maximum is 90 days."
  type        = string
  default     = "30m"
}
variable "sm_iam_credentials_secret_access_groups" {
  description = "Access Groups that you can use for an `iam_credentials` secret.Up to 10 Access Groups can be used for each secret."
  type        = list(string)
  default     = [ "AccessGroupId-45884031-54be-4dd7-86ff-112511e92699" ]
}
variable "sm_iam_credentials_secret_reuse_api_key" {
  description = "Determines whether to use the same service ID and API key for future read operations on an`iam_credentials` secret.If it is set to `true`, the service reuses the current credentials. If it is set to `false`, a new service ID and API key are generated each time that the secret is read or accessed."
  type        = bool
  default     = true
}

// Resource arguments for sm_service_credentials_secret
variable "sm_service_credentials_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-service-credentials-secret"
}
variable "sm_service_credentials_secret_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_service_credentials_secret_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_service_credentials_secret_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_service_credentials_secret_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_service_credentials_secret_source_service_instance_crn" {
  description = "A CRN that uniquely identifies a service credentials source"
  type        = string
  default     = "crn:v1:staging:public:cloud-object-storage:global:a/111f5fb10986423e9saa8512f1db7e65:111133c8-49ea-41xe-8c40-122038246f5b::"
}
variable "sm_service_credentials_secret_source_service_role_crn" {
  description = "The service-specific custom role object, CRN role is accepted. Refer to the service’s documentation for supported roles."
  type        = string
  default     = "crn:v1:bluemix:public:iam::::serviceRole:Writer"
}
variable "sm_service_credentials_secret_source_service_parameters" {
  description = "Configuration options represented as key-value pairs. Service-defined options are used in the generation of credentials for some services."
  type        = string
  default     = {}
}
variable "sm_service_credentials_secret_ttl" {
  description = "The time-to-live (TTL) or lease duration to assign to generated credentials. The TTL defines for how long generated credentials remain valid. The value should be a string that specifies the number of seconds. Minimum duration is 86400 (1 day). Maximum is 7776000 seconds (90 days)."
  type        = string
  default     = "86401"
}


// Resource arguments for sm_arbitrary_secret
variable "sm_arbitrary_secret_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-arbitrary-secret"
}
variable "sm_arbitrary_secret_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_arbitrary_secret_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_arbitrary_secret_expiration_date" {
  description = "The date a secret is expired. The date format follows RFC 3339."
  type        = string
  default     = "2022-04-12T23:20:50.520Z"
}
variable "sm_arbitrary_secret_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_arbitrary_secret_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_arbitrary_secret_payload" {
  description = "The arbitrary secret's data payload."
  type        = string
  default     = "secret-credentials"
}

// Resource arguments for sm_username_password_secret
variable "sm_username_password_secret_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-username-password-secret"
}
variable "sm_username_password_secret_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_username_password_secret_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_username_password_secret_expiration_date" {
  description = "The date a secret is expired. The date format follows RFC 3339."
  type        = string
  default     = "2022-04-12T23:20:50.520Z"
}
variable "sm_username_password_secret_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_username_password_secret_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_username_password_secret_username" {
  description = "The username that is assigned to the secret."
  type        = string
  default     = "username"
}
variable "sm_username_password_secret_password" {
  description = "The password that is assigned to the secret."
  type        = string
  default     = "password"
}

// Resource arguments for sm_private_certificate
variable "sm_private_certificate_name" {
  description = "The human-readable name of your secret."
  type        = string
  default     = "my-private-certificate-secret"
}
variable "sm_private_certificate_custom_metadata" {
  description = "The secret metadata that a user can customize."
  type        = any
  default     = "anything as a string"
}
variable "sm_private_certificate_description" {
  description = "An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group."
  type        = string
  default     = "Extended description for this secret."
}
variable "sm_private_certificate_expiration_date" {
  description = "The date a secret is expired. The date format follows RFC 3339."
  type        = string
  default     = "2022-04-12T23:20:50.520Z"
}
variable "sm_private_certificate_labels" {
  description = "Labels that you can use to search for secrets in your instance.Up to 30 labels can be created."
  type        = list(string)
  default     = [ "my-label" ]
}
variable "sm_private_certificate_secret_group_id" {
  description = "A UUID identifier, or `default` secret group."
  type        = string
  default     = "default"
}
variable "sm_private_certificate_certificate_template" {
  description = "The name of the certificate template."
  type        = string
  default     = "cert-template-1"
}

// Resource arguments for sm_private_certificate_configuration_root_ca
variable "sm_private_certificate_configuration_root_ca_name" {
  description = "A human-readable unique name to assign to your configuration."
  type        = string
  default     = "my_root_ca"
}
variable "sm_private_certificate_configuration_root_ca_crl_disable" {
  description = "Disables or enables certificate revocation list (CRL) building.If CRL building is disabled, a signed but zero-length CRL is returned when downloading the CRL. If CRL building is enabled, it will rebuild the CRL."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_root_ca_crl_distribution_points_encoded" {
  description = "Determines whether to encode the certificate revocation list (CRL) distribution points in the certificates that are issued by this certificate authority."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_root_ca_issuing_certificates_urls_encoded" {
  description = "Determines whether to encode the URL of the issuing certificate in the certificates that are issued by this certificate authority."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_root_ca_ttl" {
  description = "The requested time-to-live (TTL) for certificates that are created by this CA. This field's value cannot be longer than the `max_ttl` limit.The value can be supplied as a string representation of a duration in hours, for example '8760h'. In the API response, this value is returned in seconds (integer)."
  type        = string
  default     = "8760h"
}

// Resource arguments for sm_private_certificate_configuration_intermediate_ca
variable "sm_private_certificate_configuration_intermediate_ca_name" {
  description = "A human-readable unique name to assign to your configuration."
  type        = string
  default     = "my_intermediate_ca"
}
variable "sm_private_certificate_configuration_intermediate_ca_crl_disable" {
  description = "Disables or enables certificate revocation list (CRL) building.If CRL building is disabled, a signed but zero-length CRL is returned when downloading the CRL. If CRL building is enabled, it will rebuild the CRL."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_intermediate_ca_crl_distribution_points_encoded" {
  description = "Determines whether to encode the certificate revocation list (CRL) distribution points in the certificates that are issued by this certificate authority."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_intermediate_ca_issuing_certificates_urls_encoded" {
  description = "Determines whether to encode the URL of the issuing certificate in the certificates that are issued by this certificate authority."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_intermediate_ca_signing_method" {
  description = "The signing method to use with this certificate authority to generate private certificates.You can choose between internal or externally signed options. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-intermediate-certificate-authorities)."
  type        = string
  default     = "internal"
}

// Resource arguments for sm_private_certificate_configuration_template
variable "sm_private_certificate_configuration_template_name" {
  description = "A human-readable unique name to assign to your configuration."
  type        = string
  default     = "my_template"
}
variable "sm_private_certificate_configuration_template_certificate_authority" {
  description = "The name of the intermediate certificate authority."
  type        = string
  default     = "certificate_authority"
}
variable "sm_private_certificate_configuration_template_allowed_secret_groups" {
  description = "Scopes the creation of private certificates to only the secret groups that you specify.This field can be supplied as a comma-delimited list of secret group IDs."
  type        = string
  default     = "allowed_secret_groups"
}
variable "sm_private_certificate_configuration_template_allow_localhost" {
  description = "Determines whether to allow `localhost` to be included as one of the requested common names."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allowed_domains" {
  description = "The domains to define for the certificate template. This property is used along with the `allow_bare_domains` and `allow_subdomains` options."
  type        = list(string)
  default     = [ "allowed_domains" ]
}
variable "sm_private_certificate_configuration_template_allowed_domains_template" {
  description = "Determines whether to allow the domains that are supplied in the `allowed_domains` field to contain access control list (ACL) templates."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allow_bare_domains" {
  description = "Determines whether to allow clients to request private certificates that match the value of the actual domains on the final certificate.For example, if you specify `example.com` in the `allowed_domains` field, you grant clients the ability to request a certificate that contains the name `example.com` as one of the DNS values on the final certificate.**Important:** In some scenarios, allowing bare domains can be considered a security risk."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allow_subdomains" {
  description = "Determines whether to allow clients to request private certificates with common names (CN) that are subdomains of the CNs that are allowed by the other certificate template options. This includes wildcard subdomains.For example, if `allowed_domains` has a value of `example.com` and `allow_subdomains`is set to `true`, then the following subdomains are allowed: `foo.example.com`, `bar.example.com`, `*.example.com`.**Note:** This field is redundant if you use the `allow_any_name` option."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allow_glob_domains" {
  description = "Determines whether to allow glob patterns, for example, `ftp*.example.com`, in the names that are specified in the `allowed_domains` field.If set to `true`, clients are allowed to request private certificates with names that match the glob patterns."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allow_any_name" {
  description = "Determines whether to allow clients to request a private certificate that matches any common name."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_enforce_hostnames" {
  description = "Determines whether to enforce only valid host names for common names, DNS Subject Alternative Names, and the host section of email addresses."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allow_ip_sans" {
  description = "Determines whether to allow clients to request a private certificate with IP Subject Alternative Names."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_allowed_uri_sans" {
  description = "The URI Subject Alternative Names to allow for private certificates.Values can contain glob patterns, for example `spiffe://hostname/_*`."
  type        = list(string)
  default     = [ "allowed_uri_sans" ]
}
variable "sm_private_certificate_configuration_template_allowed_other_sans" {
  description = "The custom Object Identifier (OID) or UTF8-string Subject Alternative Names (SANs) to allow for private certificates.The format for each element in the list is the same as OpenSSL: `<oid>:<type>:<value>` where the current valid type is `UTF8`. To allow any value for an OID, use `*` as its value. Alternatively, specify a single `*` to allow any `other_sans` input."
  type        = list(string)
  default     = ["2.5.4.5;UTF8:*"]
}
variable "sm_private_certificate_configuration_template_server_flag" {
  description = "Determines whether private certificates are flagged for server use."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_client_flag" {
  description = "Determines whether private certificates are flagged for client use."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_code_signing_flag" {
  description = "Determines whether private certificates are flagged for code signing use."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_email_protection_flag" {
  description = "Determines whether private certificates are flagged for email protection use."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_key_usage" {
  description = "The allowed key usage constraint to define for private certificates.You can find valid values in the [Go x509 package documentation](https://pkg.go.dev/crypto/x509#KeyUsage). Omit the `KeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list."
  type        = list(string)
  default     = ["DigitalSignature","KeyAgreement","KeyEncipherment"]
}
variable "sm_private_certificate_configuration_template_ext_key_usage" {
  description = "The allowed extended key usage constraint on private certificates.You can find valid values in the [Go x509 package documentation](https://golang.org/pkg/crypto/x509/#ExtKeyUsage). Omit the `ExtKeyUsage` part of the value. Values are not case-sensitive. To specify no key usage constraints, set this field to an empty list."
  type        = list(string)
  default     = [ "ext_key_usage" ]
}
variable "sm_private_certificate_configuration_template_ext_key_usage_oids" {
  description = "A list of extended key usage Object Identifiers (OIDs)."
  type        = list(string)
  default     = [ "ext_key_usage_oids" ]
}
variable "sm_private_certificate_configuration_template_use_csr_common_name" {
  description = "When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the common name (CN) from a certificate signing request (CSR) instead of the CN that's included in the data of the certificate.Does not include any requested Subject Alternative Names (SANs) in the CSR. To use the alternative names, include the `use_csr_sans` property."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_use_csr_sans" {
  description = "When used with the `private_cert_configuration_action_sign_csr` action, this field determines whether to use the Subject Alternative Names(SANs) from a certificate signing request (CSR) instead of the SANs that are included in the data of the certificate.Does not include the common name in the CSR. To use the common name, include the `use_csr_common_name` property."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_require_cn" {
  description = "Determines whether to require a common name to create a private certificate.By default, a common name is required to generate a certificate. To make the `common_name` field optional, set the `require_cn` option to `false`."
  type        = bool
  default     = true
}
variable "sm_private_certificate_configuration_template_policy_identifiers" {
  description = "A list of policy Object Identifiers (OIDs)."
  type        = list(string)
  default     = [ "policy_identifiers" ]
}
variable "sm_private_certificate_configuration_template_basic_constraints_valid_for_non_ca" {
  description = "Determines whether to mark the Basic Constraints extension of an issued private certificate as valid for non-CA certificates."
  type        = bool
  default     = true
}

// Resource arguments for sm_private_certificate_configuration_action_sign_csr
variable "sm_private_certificate_configuration_action_sign_csr_name" {
  description = "The name that uniquely identifies a configuration."
  type        = string
  default     = "my_root_ca"
}
variable "sm_private_certificate_configuration_action_sign_csr_csr" {
  description = "The certificate signing request."
  type        = string
  default     = "csr"
}

// Resource arguments for sm_private_certificate_configuration_action_set_signed
variable "sm_private_certificate_configuration_action_set_signed_name" {
  description = "The name that uniquely identifies a configuration."
  type        = string
  default     = "my_intermediate_ca"
}
variable "sm_private_certificate_configuration_action_set_signed_certificate" {
  description = "The PEM-encoded certificate."
  type        = string
  default     = "certificate"
}

// Resource arguments for sm_public_certificate_configuration_ca_lets_encrypt
variable "sm_public_certificate_configuration_ca_lets_encrypt_name" {
  description = "A human-readable unique name to assign to your configuration."
  type        = string
  default     = "my-ca-lets-encrypt-config"
}
variable "sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_environment" {
  description = "The configuration of the Let's Encrypt CA environment."
  type        = string
  default     = "production"
}
variable "sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_private_key" {
  description = "The PEM encoded private key of your Lets Encrypt account."
  type        = string
  default     = "lets_encrypt_private_key"
}
variable "sm_public_certificate_configuration_ca_lets_encrypt_lets_encrypt_preferred_chain" {
  description = "Prefer the chain with an issuer matching this Subject Common Name."
  type        = string
  default     = "lets_encrypt_preferred_chain"
}

// Resource arguments for sm_public_certificate_configuration_dns_cis
variable "sm_public_certificate_configuration_dns_cis_cloud_internet_services_name" {
  description = "A human-readable unique name to assign to your configuration."
  type        = string
  default     = "my-dns-cis-cloud-internet-services-config"
}
variable "sm_public_certificate_configuration_dns_cis_cloud_internet_services_apikey" {
  description = "An IBM Cloud API key that can to list domains in your Cloud Internet Services instance.To grant Secrets Manager the ability to view the Cloud Internet Services instance and all of its domains, the API key must be assigned the Reader service role on Internet Services (`internet-svcs`).If you need to manage specific domains, you can assign the Manager role. For production environments, it is recommended that you assign the Reader access role, and then use the[IAM Policy Management API](https://cloud.ibm.com/apidocs/iam-policy-management#create-policy) to control specific domains. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-prepare-order-certificates#authorize-specific-domains)."
  type        = string
  default     = "cloud_internet_services_apikey"
}
variable "sm_public_certificate_configuration_dns_cis_cloud_internet_services_crn" {
  description = "A CRN that uniquely identifies an IBM Cloud resource."
  type        = string
  default     = "cloud_internet_services_crn"
}

// Resource arguments for sm_public_certificate_configuration_dns_classic_infrastructure
variable "sm_public_certificate_configuration_dns_classic_infrastructure_name" {
  description = "A human-readable unique name to assign to your configuration."
  type        = string
  default     = "my-dns-classic-infrastructure-config"
}
variable "sm_public_certificate_configuration_dns_classic_infrastructure_classic_infrastructure_username" {
  description = "The username that is associated with your classic infrastructure account.In most cases, your classic infrastructure username is your `<account_id>_<email_address>`. For more information, see the [docs](https://cloud.ibm.com/docs/account?topic=account-classic_keys)."
  type        = string
  default     = "classic_infrastructure_username"
}
variable "sm_public_certificate_configuration_dns_classic_infrastructure_classic_infrastructure_password" {
  description = "Your classic infrastructure API key.For information about viewing and accessing your classic infrastructure API key, see the [docs](https://cloud.ibm.com/docs/account?topic=account-classic_keys)."
  type        = string
  default     = "classic_infrastructure_password"
}

// Resource arguments for sm_en_registration
variable "sm_en_registration_event_notifications_instance_crn" {
  description = "A CRN that uniquely identifies an IBM Cloud resource."
  type        = string
  default     = "crn:v1:bluemix:public:event-notifications:us-south:a/22018f3c34ff4ff193698d15ca316946:578ad1a4-2fd8-4e66-95d5-79a842ba91f8::"
}
variable "sm_en_registration_event_notifications_source_name" {
  description = "The name that is displayed as a source that is in your Event Notifications instance."
  type        = string
  default     = "My Secrets Manager"
}
variable "sm_en_registration_event_notifications_source_description" {
  description = "An optional description for the source  that is in your Event Notifications instance."
  type        = string
  default     = "Optional description of this source in an Event Notifications instance."
}

// Data source arguments for sm_secret_group
variable "sm_secret_group_id" {
  description = "The ID of the secret group."
  type        = string
  default     = "d898bb90-82f6-4d61-b5cc-b079b66cfa76"
}


// Data source arguments for sm_secret_version_action
variable "sm_secret_version_action_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
variable "sm_secret_version_action_id" {
  description = "The ID of the secret version."
  type        = string
  default     = "eb4cf24d-9cae-424b-945e-159788a5f535"
}
variable "sm_secret_version_action_secret_version_action_prototype" {
  description = "The request body to specify the properties of the action to create a secret version."
  type        = list(object({ example=string }))
  default     = {"action_type":"private_cert_action_revoke_certificate"}
}

// Data source arguments for sm_public_certificate_action_validate_manual_dns
variable "sm_public_certificate_action_validate_manual_dns_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
variable "sm_public_certificate_action_validate_manual_dns_secret_action_prototype" {
  description = "Specify the properties for your secret action."
  type        = list(object({ example=string }))
  default     = {"action_type":"private_cert_action_revoke_certificate"}
}

// Data source arguments for sm_private_certificate_action_revoke
variable "sm_private_certificate_action_revoke_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
variable "sm_private_certificate_action_revoke_secret_action_prototype" {
  description = "Specify the properties for your secret action."
  type        = list(object({ example=string }))
  default     = {"action_type":"private_cert_action_revoke_certificate"}
}

// Data source arguments for sm_secret_versions
variable "sm_secret_versions_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_secret_version_metadata
variable "sm_secret_version_metadata_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
variable "sm_secret_version_metadata_id" {
  description = "The ID of the secret version."
  type        = string
  default     = "eb4cf24d-9cae-424b-945e-159788a5f535"
}

// Data source arguments for sm_imported_certificate_metadata
variable "sm_imported_certificate_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_public_certificate_metadata
variable "sm_public_certificate_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_kv_secret_metadata
variable "sm_kv_secret_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_iam_credentials_secret_metadata
variable "sm_iam_credentials_secret_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_service_credentials_secret_metadata
variable "sm_service_credentials_secret_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}


// Data source arguments for sm_arbitrary_secret_metadata
variable "sm_arbitrary_secret_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_username_password_secret_metadata
variable "sm_username_password_secret_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_imported_certificate
variable "sm_imported_certificate_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_public_certificate
variable "sm_public_certificate_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_kv_secret
variable "sm_kv_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_iam_credentials_secret
variable "sm_iam_credentials_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_service_credentials_secret
variable "sm_service_credentials_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_arbitrary_secret
variable "sm_arbitrary_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_username_password_secret
variable "sm_username_password_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for arbitrary_sm_secret_version
variable "arbitrary_sm_secret_version_secret_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
variable "arbitrary_sm_secret_version_id" {
  description = "The ID of the secret version."
  type        = string
  default     = "eb4cf24d-9cae-424b-945e-159788a5f535"
}

// Data source arguments for sm_private_certificate
variable "sm_private_certificate_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_private_certificate_metadata
variable "sm_private_certificate_metadata_id" {
  description = "The ID of the secret."
  type        = string
  default     = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}

// Data source arguments for sm_private_certificate_configuration_root_ca
variable "sm_private_certificate_configuration_root_ca_name" {
  description = "The name of the configuration."
  type        = string
  default     = "configuration-name"
}

// Data source arguments for sm_private_certificate_configuration_intermediate_ca
variable "sm_private_certificate_configuration_intermediate_ca_name" {
  description = "The name of the configuration."
  type        = string
  default     = "configuration-name"
}

// Data source arguments for sm_private_certificate_configuration_template
variable "sm_private_certificate_configuration_template_name" {
  description = "The name of the configuration."
  type        = string
  default     = "configuration-name"
}


// Data source arguments for sm_public_certificate_configuration_ca_lets_encrypt
variable "sm_public_certificate_configuration_ca_lets_encrypt_name" {
  description = "The name of the configuration."
  type        = string
  default     = "configuration-name"
}

// Data source arguments for sm_public_certificate_configuration_dns_cis
variable "sm_public_certificate_configuration_dns_cis_name" {
  description = "The name of the configuration."
  type        = string
  default     = "configuration-name"
}

// Data source arguments for sm_public_certificate_configuration_dns_classic_infrastructure
variable "sm_public_certificate_configuration_dns_classic_infrastructure_name" {
  description = "The name of the configuration."
  type        = string
  default     = "configuration-name"
}