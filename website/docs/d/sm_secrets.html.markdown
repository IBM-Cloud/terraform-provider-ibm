---
layout: "ibm"
page_title: "IBM : ibm_sm_secrets"
description: |-
  Get information about sm_secrets
subcategory: "Secrets Manager"
---

# ibm_sm_secrets

Provides a read-only data source for sm_secrets. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_secrets" "secrets" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `sort` - (Optional, String) - Sort a collection of secrets by the specified field in ascending order. To sort in descending order use the `-` character. 
	* Constraints: Allowable values are: `id`, `created_at`, `updated_at`, `expiration_date`, `secret_type`, `name`.
* `search` - (Optional, String) - Obtain a collection of secrets that contain the specified string in one or more of the fields: `id`, `name`, `description`, `labels`, `secret_type`.
* `groups` - (Optional, String) - Filter secrets by groups. You can apply multiple filters by using a comma-separated list of secret group IDs. If you need to filter secrets that are in the default secret group, use the `default` keyword.
* `secret_types` - (Optional, List) - Filter secrets by secret types. You can apply multiple filters by using a comma-separated list of secret types.
* `match_all_labels` - (Optional, String) - Filter secrets by a label or a combination of labels (comma-separated list).

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the sm_secrets.
* `secrets` - (List) A collection of secret metadata. Note that the list of metadata attributes conatains attributes that are common to all types of secrets, as well as attributes that are specific to cetrain secret types. A type specific attribute is included in every secret but the value is empty for secrets of other types. The common attributes are: `name, id, description, secret_type, crn, created_by, created_at, updated_at, downloaded, secret_group_id, state, state_description, versions_total`.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items. 
Nested scheme for **secrets**:
    * `access_groups` - (List) Access Groups that you can use for an `iam_credentials` secret.Up to 10 Access Groups can be used for each secret.
      * Constraints: The list items must match regular expression `/^AccessGroupId-[a-z0-9-]+[a-z0-9]$/`. The maximum length is `10` items. The minimum length is `1` item.
    * `alt_names` - (List) With the Subject Alternative Name field, you can specify additional host names to be protected by a single SSL certificate.
      * Constraints: The list items must match regular expression `/^(.*?)$/`. The maximum length is `99` items. The minimum length is `0` items.
    * `api_key_id` - (String) The ID of the API key that is generated for this secret.
    * `bundle_certs` - (Boolean) Indicates whether the issued certificate is bundled with intermediate certificates.
    * `ca` - (String) The name that is assigned to the certificate authority configuration.
    * `certificate_authority` - (String) The intermediate certificate authority that signed this certificate.
    * `certificate_template` - (String) The name of the certificate template.
      * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:_?-?\\.?[A-Za-z0-9]+)*$/`.
    * `common_name` - (String) The Common Name (AKA CN) represents the server name protected by the SSL certificate.
      * Constraints: The maximum length is `64` characters. The minimum length is `4` characters. The value must match regular expression `/^(\\*\\.)?(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\\-]*[A-Za-z0-9])\\.?$/`.
    * `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
    * `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
      * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
    * `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
      * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
    * `custom_metadata` - (Map) The secret metadata that a user can customize.
		* Constraints: Nested JSONs are supported in Terraform only as string-encoded JSONs.
    * `description` - (String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
      * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
    * `dns` - (String) The name that is assigned to the DNS provider configuration.
    * `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.
    * `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.
    * `id` - (String) A UUID identifier.
      * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.
    * `intermediate_included` - (Boolean) Indicates whether the certificate was imported with an associated intermediate certificate.
    * `issuance_info` - (List) Issuance information that is associated with your certificate.
	Nested scheme for **issuance_info**:
		* `auto_rotated` - (Boolean) Indicates whether the issued certificate is configured with an automatic rotation policy.
		* `challenges` - (List) The set of challenges. It is returned only when ordering public certificates by using manual DNS configuration.
		  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
		Nested scheme for **challenges**:
			* `domain` - (String) The challenge domain.
			* `expiration` - (String) The challenge expiration date. The date format follows RFC 3339.
			* `status` - (String) The challenge status.
			* `txt_record_name` - (String) The TXT record name.
			* `txt_record_value` - (String) The TXT record value.
		* `dns_challenge_validation_time` - (String) The date that a user requests to validate DNS challenges for certificates that are ordered with a manual DNS provider. The date format follows RFC 3339.
		* `error_code` - (String) A code that identifies an issuance error.This field, along with `error_message`, is returned when Secrets Manager successfully processes your request, but the certificate authority is unable to issue a certificate.
		* `error_message` - (String) A human-readable message that provides details about the issuance error.
		* `ordered_on` - (String) The date when the certificate is ordered. The date format follows RFC 3339.
		* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
		  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.
		* `state_description` - (String) A text representation of the secret state.
		  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.
	* `issuer` - (String) The distinguished name that identifies the entity that signed and issued the certificate.
	  * Constraints: The maximum length is `128` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
	* `key_algorithm` - (String) The identifier for the cryptographic algorithm used to generate the public key that is associated with the certificate.
	  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
	* `labels` - (List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
	  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.
	* `locks_total` - (Integer) The number of locks of the secret.
	  * Constraints: The maximum value is `1000`. The minimum value is `0`.
	* `name` - (String) The human-readable name of your secret.
	  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters.
	* `next_rotation_date` - (String) The date that the secret is scheduled for automatic rotation.The service automatically creates a new version of the secret on its next rotation date. This field exists only for secrets that have an existing rotation policy.
	* `private_key_included` - (Boolean) Indicates whether the certificate was imported with an associated private key.
	* `reuse_api_key` - (Boolean) Determines whether to use the same service ID and API key for future read operations on an`iam_credentials` secret. The value is always `true` for IAM credentials secrets managed by Terraform.
	* `revocation_time_rfc3339` - (String) The date and time that the certificate was revoked. The date format follows RFC 3339.
	* `revocation_time_seconds` - (Integer) The timestamp of the certificate revocation.
	* `rotation` - (List) Determines whether Secrets Manager rotates your secrets automatically.
	Nested scheme for **rotation**:
		* `auto_rotate` - (Boolean) Determines whether Secrets Manager rotates your secret automatically.Default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval.
		* `interval` - (Integer) The length of the secret rotation time interval.
		  * Constraints: The minimum value is `1`.
		* `rotate_keys` - (Boolean) Determines whether Secrets Manager rotates the private key for your public certificate automatically.Default is `false`. If it is set to `true`, the service generates and stores a new private key for your rotated certificate.
		* `unit` - (String) The units for the secret rotation time interval.
		  * Constraints: Allowable values are: `day`, `month`.
	* `secret_group_id` - (String) A UUID identifier, or `default` secret group.
	  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.
	* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
	  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.
	* `serial_number` - (String) The unique serial number that was assigned to a certificate by the issuing certificate authority.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[^a-fA-F0-9]/`.
	* `service_id` - (String) The service ID under which the API key (see the `api_key` field) is created.If you omit this parameter, Secrets Manager generates a new service ID for your secret at its creation and adds it to the access groups that you assign.Optionally, you can use this field to provide your own service ID if you prefer to manage its access directly or retain the service ID after your secret expires, is rotated, or deleted. If you provide a service ID, do not include the `access_groups` parameter.
	  * Constraints: The maximum length is `50` characters. The minimum length is `40` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:-?[A-Za-z0-9]+)*$/`.
	* `service_id_is_static` - (Boolean) Indicates whether an `iam_credentials` secret was created with a static service ID.If it is set to `true`, the service ID for the secret was provided by the user at secret creation. If it is set to `false`, the service ID was generated by Secrets Manager.
	* `signing_algorithm` - (String) The identifier for the cryptographic algorithm that was used by the issuing certificate authority to sign a certificate.
	  * Constraints: The maximum length is `64` characters. The minimum length is `4` characters.
	* `source_service` - (List) The properties required for creating the service credentials for the specified source service instance.
	Nested scheme for **source_service**:
		* `instance` - (List) The source service instance identifier.
		Nested scheme for **instance**:
			* `crn` - (String) A CRN that uniquely identifies a service credentials source.
		* `iam` - (List) The source service IAM data is returned in case IAM credentials where created for this secret.
		Nested scheme for **iam**:
			* `apikey` - (String) The IAM apikey metadata for the IAM credentials that were generated.
            Nested scheme for **apikey**:
				* `name` - (String) The IAM API key name for the generated service credentials.
				* `description` - (String) The IAM API key description for the generated service credentials.
			* `role` - (String) The IAM role for the generate service credentials.
			Nested scheme for **role**:
				* `crn` - (String) The IAM role CRN assigned to the generated service credentials.
			* `serviceid` - (String) The IAM serviceid for the generated service credentials.
			Nested scheme for **serviceid**:
				* `crn` - (String) The IAM Service ID CRN.
		* `parameters` - (Map) The collection of parameters for the service credentials target.
		* `resource_key` - (List) The source service resource key data of the generated service credentials.
		Nested scheme for **resource_key**:
			* `crn` - (String) The resource key CRN of the generated service credentials.
			* `name` - (String) The resource key name of the generated service credentials.
		* `role` - (List) The service-specific custom role object.
		Nested scheme for **role**:
			* `crn` - (String) The CRN role identifier for creating a service-id.
	* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
	  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.
	* `state_description` - (String) A text representation of the secret state.
	  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.
	* `ttl` - (String) The time-to-live (TTL) or lease duration to assign to generated credentials.For `iam_credentials` secrets, the TTL defines for how long each generated API key remains valid. The value can be either an integer that specifies the number of seconds, or the string representation of a duration, such as `120m` or `24h`.Minimum duration is 1 minute. Maximum is 90 days.
	  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.
	* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.
	* `validity` - (List) The date and time that the certificate validity period begins and ends.
	Nested scheme for **validity**:
		* `not_after` - (String) The date-time format follows RFC 3339.
		* `not_before` - (String) The date-time format follows RFC 3339.
	* `versions_total` - (Integer) The number of versions of the secret.
	  * Constraints: The maximum value is `50`. The minimum value is `0`.

