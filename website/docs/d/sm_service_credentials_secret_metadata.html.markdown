---
layout: "ibm"
page_title: "IBM : ibm_sm_service_credentials_secret_metadata"
description: |-
  Get information about ServiceCredentialsSecretMetadata
subcategory: "Secrets Manager"
---

# ibm_sm_service_credentials_secret_metadata

Provides a read-only data source for the metadata of an service credentials secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_sm_service_credentials_secret_metadata" "service_credentials_secret_metadata" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `secret_id` - (Optional, String) The ID of the secret.
    * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
    * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
    
* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
    * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `custom_metadata` - (Map) The secret metadata that a user can customize.
  * Constraints: Nested JSONs are supported in Terraform only as string-encoded JSONs.

* `description` - (String) An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
    * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.

* `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.

* `labels` - (List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
    * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.

* `locks_total` - (Integer) The number of locks of the secret.
    * Constraints: The maximum value is `1000`. The minimum value is `0`.

* `name` - (String) The human-readable name of your secret.
    * Constraints: The maximum length is `256` characters. The minimum length is `2` characters.

* `next_rotation_date` - (String) The date that the secret is scheduled for automatic rotation.The service automatically creates a new version of the secret on its next rotation date. This field exists only for secrets that have an existing rotation policy.

* `retrieved_at` - (String) The date when the data of the secret was last retrieved. The date format follows RFC 3339. Epoch date if there is no record of secret data retrieval.

* `rotation` - (List) Determines whether Secrets Manager rotates your secrets automatically.
  Nested scheme for **rotation**:
  * `auto_rotate` - (Boolean) Determines whether Secrets Manager rotates your secret automatically.Default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval.
  * `interval` - (Integer) The length of the secret rotation time interval.
  * Constraints: The minimum value is `1`.
  * `unit` - (String) The units for the secret rotation time interval.
  * Constraints: Allowable values are: `day`, `month`.

* `secret_group_id` - (String) A UUID identifier, or `default` secret group.
    * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
    * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `source_service` - (List) The properties required for creating the service credentials for the specified source service instance.
  Nested scheme for **source_service**:
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
  * `resource_key` - (List) The source service resource key data of the generated service credentials.
  Nested scheme for **resource_key**:
  * `crn` - (String) The resource key CRN of the generated service credentials.
  * `name` - (String) The resource key name of the generated service credentials.

* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
    * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.

* `state_description` - (String) A text representation of the secret state.
    * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.

* `ttl` - (String) The time-to-live (TTL) or lease duration to assign to generated credentials. The TTL defines for how long generated credentials remain valid. The value should be a string that specifies the number of seconds. Minimum duration is 86400 (1 day). Maximum is 7776000 seconds (90 days).
    * Constraints: The maximum length is `7` characters. The minimum length is `2` characters.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `versions_total` - (Integer) The number of versions of the secret.
    * Constraints: The maximum value is `50`. The minimum value is `0`.

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.
