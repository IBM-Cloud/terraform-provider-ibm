---
layout: "ibm"
page_title: "IBM : ibm_sm_iam_credentials_secret"
description: |-
  Get information about a custom credentials secret
subcategory: "Secrets Manager"
---

# ibm_sm_iam_credentials_secret

Provides a read-only data source for a custom credentials secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.
The data source can be defined by providing the secret ID or the secret and secret group names.

## Example Usage

By secret id
```hcl
data "ibm_sm_custom_credentials_secret" "my_secret" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_id = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
}
```


By secret name and group name
```hcl
data "ibm_sm_custom_credentials_secret" "my_secret" {
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

* `configuration` - (String) The name of the Custom Credentials configuration.

* `created_at` - (String) The date when the secret was created. The date format follows RFC 3339.

* `created_by` - (String) The unique identifier that is associated with the entity that created the secret.

* `credentials_content` - (List) The credentials that were generated for this secret.
  Nested scheme for **credentials_content**:
  * `integer_values` - (Map) Values of integer credentials.
  * `string_values` - (Map) Values of string credentials.
  * `boolean_values` - (Map) Values of boolean credentials.

* `crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.

* `custom_metadata` - (Map) The secret metadata that a user can customize.
  * Constraints: Nested JSONs are supported in Terraform only as string-encoded JSONs.

* `description` - (String) An extended description of your secret.

* `downloaded` - (Boolean) Indicates whether the secret data that is associated with a secret version was retrieved in a call to the service API.

* `labels` - (List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.

* `locks_total` - (Integer) The number of locks of the secret.

* `name` - (String) The human-readable name of your secret.

* `next_rotation_date` - (String) The date that the secret is scheduled for automatic rotation.The service automatically creates a new version of the secret on its next rotation date. This field exists only for secrets that have an existing rotation policy.

* `parameters` - (List) The parameters that were passed to the Code Engine job.
  Nested scheme for **parameters**:
  * `integer_values` - (Map) Values of integer parameters.
  * `string_values` - (Map) Values of string parameters.
  * `boolean_values` - (Map) Values of boolean parameters.

* `retrieved_at` - (String) The date when the data of the secret was last retrieved. The date format follows RFC 3339. Epoch date if there is no record of secret data retrieval.

* `rotation` - (List) Determines whether Secrets Manager rotates your secrets automatically.
Nested scheme for **rotation**:
    * `auto_rotate` - (Boolean) Determines whether Secrets Manager rotates your secret automatically.Default is `false`. If `auto_rotate` is set to `true` the service rotates your secret based on the defined interval.
    * `interval` - (Integer) The length of the secret rotation time interval.
      * Constraints: The minimum value is `1`.
    * `unit` - (String) The units for the secret rotation time interval.
      * Constraints: Allowable values are: `day`, `month`.

* `secret_group_id` - (String) A UUID identifier, or `default` secret group.

* `secret_type` - (String) The secret type. 

* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.

* `state_description` - (String) A text representation of the secret state.
  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.

* `ttl` - (String) The time-to-live or lease duration (in seconds) to assign to generated credentials. Minimum duration is 86400 seconds (one day).

* `updated_at` - (String) The date when the secret was recently modified. The date format follows RFC 3339.

* `versions_total` - (Integer) The number of versions of the secret.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.
