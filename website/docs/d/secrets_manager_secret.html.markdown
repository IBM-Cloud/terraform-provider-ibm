---
subcategory: "Secrets Manager"
layout: "ibm"
page_title: "IBM : secrets_manager_secret"
description: |-
  Get information about secrets_manager_secret
---

# ibm\_secrets_manager_secret

Provides a read-only data source for secrets_manager_secret. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "secrets_manager_secret" "secrets_manager_secret" {
	instance_id = "36401ffc-6280-459a-ba98-456aba10d0c7"
	secret_type = "arbitrary"
	secret_id = "7dd2022c-5f54-f96d-4c32-87309e887e5"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The Secrets Manager Instance GUID.
* `secret_type` - (Required, string) The secret type. Supported options include: arbitrary, iam_credentials, username_password.
* `secret_id` - (Required, string) The v4 UUID that uniquely identifies the secret.
* `endpoint_type` - (Optional, string) The type of the endpoint to be used for fetching secret. Supported options include: `public`, `private`. Default is `public`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the secrets_manager_secret.
* `metadata` - The metadata that describes the resource array. Nested `metadata` blocks have the following structure:
	* `collection_type` - The type of resources in the resource array.
	* `collection_total` - The number of elements in the resource array.
* `type` - The MIME type that represents the secret.
* `name` - A human-readable alias to assign to your secret.To protect your privacy, do not use personal data, such as your name or location, as an alias for your secret.
* `description` - An extended description of your secret.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret.
* `secret_group_id` - The v4 UUID that uniquely identifies the secret group to assign to this secret.If you omit this parameter, your secret is assigned to the `default` secret group.
* `labels` - Labels that you can use to filter for secrets in your instance.Up to 30 labels can be created. Labels can be between 2-30 characters, including spaces. Special characters not permitted include the angled bracket, comma, colon, ampersand, and vertical pipe character (|).To protect your privacy, do not use personal data, such as your name or location, as a label for your secret.
* `state` - The secret state based on NIST SP 800-57. States are integers and correspond to the Pre-activation = 0, Active = 1,  Suspended = 2, Deactivated = 3, and Destroyed = 5 values.
* `state_description` - A text representation of the secret state.
* `crn` - The Cloud Resource Name (CRN) that uniquely identifies your Secrets Manager resource.
* `creation_date` - The date the secret was created. The date format follows RFC 3339.
* `created_by` - The unique identifier for the entity that created the secret.
* `last_update_date` - Updates when the actual secret is modified. The date format follows RFC 3339.
* `versions` - An array that contains metadata for each secret version. Nested `versions` blocks have the following structure:
	* `id` - The ID of the secret version.
	* `creation_date` - The date that the version of the secret was created.
	* `created_by` - The unique identifier for the entity that created the secret.
	* `auto_rotated` - Indicates whether the version of the secret was created by automatic rotation.
* `expiration_date` - The date the secret material expires. The date format follows RFC 3339.You can set an expiration date on supported secret types at their creation. If you create a secret without specifying an expiration date, the secret does not expire. The `expiration_date` field is supported for the following secret types:- `arbitrary`- `username_password`.
* `secret_data` - Map of username, password if secret_type is `username_password` else map of payload if secret_type is `arbitrary`
* `payload` - The secret data assigned to an `arbitrary` secret.
* `username` - The username assigned to an  `username_password` secret.
* `password` - The password assigned to an  `username_password` secret.
* `next_rotation_date` - The date that the secret is scheduled for automatic rotation.The service automatically creates a new version of the secret on its next rotation date. This field exists only for secrets that can be auto-rotated and have an existing rotation policy.
* `ttl` - The time-to-live (TTL) or lease duration to assign to generated credentials.For `iam_credentials` secrets, the TTL defines for how long each generated API key remains valid. The value can be either an integer that specifies the number of seconds, or the string representation of a duration, such as `120m` or `24h`.
* `access_groups` - The access groups that define the capabilities of the service ID and API key that are generated for an`iam_credentials` secret.**Tip:** To find the ID of an access group, go to **Manage > Access (IAM) > Access groups** in the IBM Cloud console. Select the access group to inspect, and click **Details** to view its ID.
* `api_key` - The API key that is generated for this secret.After the secret reaches the end of its lease (see the `ttl` field), the API key is deleted automatically. If you want to continue to use the same API key for future read operations, see the `reuse_api_key` field.
* `service_id` - The service ID under which the API key (see the `api_key` field) is created. This service ID is added to the access groups that you assign for this secret.
* `reuse_api_key` - (IAM credentials) Reuse the service ID and API key for future read operations.