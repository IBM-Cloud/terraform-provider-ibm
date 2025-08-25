---
layout: "ibm"
page_title: "IBM : ibm_sm_arbitrary_secret_metadata"
description: |-
  Get information about ArbitrarySecretMetadata
subcategory: "Secrets Manager"
---

# ibm_sm_arbitrary_secret_metadata

Provides a read-only data source for the metadata of an arbitrary secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_arbitrary_secret_metadata" "arbitrary_secret_metadata" {
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
* `secret_id` - (Required, String) The ID of the secret.
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

* `expiration_date` - (String) The date a secret is expired. The date format follows RFC 3339.

* `labels` - (List) Labels that you can use to search for secrets in your instance.Up to 30 labels can be created.
  * Constraints: The list items must match regular expression `/(.*?)/`. The maximum length is `30` items. The minimum length is `0` items.

* `locks_total` - (Integer) The number of locks of the secret.
  * Constraints: The maximum value is `1000`. The minimum value is `0`.

* `name` - (String) The human-readable name of your secret.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. 

* `secret_group_id` - (String) A UUID identifier, or `default` secret group.
  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

* `secret_type` - (String) The secret type. Supported types are arbitrary, certificates (imported, public, and private), IAM credentials, key-value, and user credentials.
  * Constraints: Allowable values are: `arbitrary`, `imported_cert`, `public_cert`, `iam_credentials`, `kv`, `username_password`, `private_cert`.

* `state` - (Integer) The secret state that is based on NIST SP 800-57. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`,  `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
  * Constraints: Allowable values are: `0`, `1`, `2`, `3`, `5`.

* `state_description` - (String) A text representation of the secret state.
  * Constraints: Allowable values are: `pre_activation`, `active`, `suspended`, `deactivated`, `destroyed`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `versions_total` - (Integer) The number of versions of the secret.
  * Constraints: The maximum value is `50`. The minimum value is `0`.

