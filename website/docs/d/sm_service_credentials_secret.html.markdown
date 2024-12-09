---
layout: "ibm"
page_title: "IBM : ibm_sm_service_credentials_secret"
description: |-
  Get information about ServiceCredentialsSecret
subcategory: "Secrets Manager"
---

# ibm_sm_service_credentials_secret

Provides a read-only data source for a service credentials secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.
The data source can be defined by providing the secret ID or the secret and secret group names.

## Example Usage

By secret id
```hcl
data "ibm_sm_service_credentials_secret" "service_credentials_secret" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```

By secret name and group name
```hcl
data "ibm_sm_service_credentials_secret" "service_credentials_secret" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name          = "secret-name"
  secret_group_name = "group-name"
}
```

### Example to access resource credentials using credentials attribute:

```terraform
data "ibm_sm_service_credentials_secret" "service_credentials_secret" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
output "access_key_id" {
  value = data.ibm_sm_service_credentials_secret.service_credentials_secret.credentials["cos_hmac_keys.access_key_id"]
}
output "secret_access_key" {
  value = data.ibm_sm_service_credentials_secret.service_credentials_secret.credentials["cos_hmac_keys.secret_access_key"]
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
* `name` - (Optional, String) The human-readable name of your secret. To be used in combination with `secret_group_name`.
    * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `^[A-Za-z0-9][A-Za-z0-9]*(?:_*-*\\.*[A-Za-z0-9]+)*$`.
* `secret_group_name` - (Optional, String) The name of your existing secret group. To be used in combination with `name`.
    * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.
    * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.

* `credentials` - (List) The properties of the service credentials secret payload.
  Nested scheme for **credentials**:
      * `apikey` - (String) The API key that is generated for this secret.
      * `cos_hmac_keys` - (String) The Cloud Object Storage HMAC keys that are returned after you create a service credentials secret.
        Nested scheme for **cos_hmac_keys**:
            * `access_key_id` - (String) The access key ID for Cloud Object Storage HMAC credentials.
            * `secret_access_key` - (String) The secret access key ID for Cloud Object Storage HMAC credentials.
      * `endpoints` - (String) The endpoints that are returned after you create a service credentials secret.
      * `iam_apikey_description` - (String) The description of the generated IAM API key.
      * `iam_apikey_name` - (String) The name of the generated IAM API key.
      * `iam_role_crn` - (String) The IAM role CRN that is returned after you create a service credentials secret.
      * `iam_serviceid_crn` - (String) The IAM serviceId CRN that is returned after you create a service credentials secret.
      * `resource_instance_id` - (String) The resource instance CRN that is returned after you create a service credentials secret.

* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
    * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `custom_metadata` - (Map) The secret metadata that a user can customize.

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
