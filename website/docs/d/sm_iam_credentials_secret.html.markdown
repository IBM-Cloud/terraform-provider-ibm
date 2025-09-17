---
layout: "ibm"
page_title: "IBM : ibm_sm_iam_credentials_secret"
description: |-
  Get information about IAMCredentialsSecret
subcategory: "Secrets Manager"
---

# ibm_sm_iam_credentials_secret

Provides a read-only data source for an IAM credentials secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.
The data source can be defined by providing the secret ID or the secret and secret group names.

## Example Usage

By secret id
```hcl
data "ibm_sm_iam_credentials_secret" "iam_secret" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```


By secret name and group name
```hcl
data "ibm_sm_iam_credentials_secret" "iam_secret" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  name          = "secret-name"
  secret_group_name = "group-name"
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

* `access_groups` - (List) Access Groups that you can use for an `iam_credentials` secret.Up to 10 Access Groups can be used for each secret.
  * Constraints: The list items must match regular expression `/^AccessGroupId-[a-z0-9-]+[a-z0-9]$/`. The maximum length is `10` items. The minimum length is `1` item.

* `account_id` - (String) The ID of the account in which the IAM credentials are created. This field is omitted if the target account is the same as the account of the Secrets Manager instance.

* `api_key` - (String) The API key that is generated for this secret. After the secret reaches the end of its lease (see the `ttl` field), the API key is deleted automatically.
  * Constraints: The maximum length is `60` characters. The minimum length is `5` characters. The value must match regular expression `/^(?:[A-Za-z0-9_\\-]{4})*(?:[A-Za-z0-9_\\-]{2}==|[A-Za-z0-9_\\-]{3}=)?$/`.

* `api_key_id` - (String) The ID of the API key that is generated for this secret.

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

* `reuse_api_key` - (Boolean) Determines whether to use the same service ID and API key for future read operations on an`iam_credentials` secret. The value is always `true` for IAM credentials secrets managed by Terraform.

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

* `service_id` - (String) The service ID under which the API key (see the `api_key` field) is created.If you omit this parameter, Secrets Manager generates a new service ID for your secret at its creation and adds it to the access groups that you assign.Optionally, you can use this field to provide your own service ID if you prefer to manage its access directly or retain the service ID after your secret expires, is rotated, or deleted. If you provide a service ID, do not include the `access_groups` parameter.
  * Constraints: The maximum length is `50` characters. The minimum length is `40` characters. The value must match regular expression `/^[A-Za-z0-9][A-Za-z0-9]*(?:-?[A-Za-z0-9]+)*$/`.

* `service_id_is_static` - (Boolean) Indicates whether an `iam_credentials` secret was created with a static service ID.If it is set to `true`, the service ID for the secret was provided by the user at secret creation. If it is set to `false`, the service ID was generated by Secrets Manager.

* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.

* `state_description` - (String) A text representation of the secret state.
  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.

* `ttl` - (String) The time-to-live (TTL) or lease duration to assign to generated credentials.For `iam_credentials` secrets, the TTL defines for how long each generated API key remains valid. The value can be either an integer that specifies the number of seconds, or the string representation of a duration, such as `120m` or `24h`.Minimum duration is 1 minute. Maximum is 90 days.
  * Constraints: The maximum length is `10` characters. The minimum length is `2` characters. The value must match regular expression `/^[0-9]+[s,m,h,d]{0,1}$/`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `versions_total` - (Integer) The number of versions of the secret.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.
