---
subcategory: "Secrets Manager"
layout: "ibm"
page_title: "IBM : ibm_secrets_manager_secrets"
description: |-
  Get information about secrets manager secrets.
---

# ibm_secrets_manager_secrets (Deprecated)
Retrieve information about the secrets manager secret data sources. For more information, about getting started with secrets manager, see [about secrets manager](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-getting-started).

## Example usage

```terraform
data "ibm_secrets_manager_secrets" "secrets_manager_secrets" {
  instance_id = "36401ffc-6280-459a-ba98-456aba10d0c7"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `endpoint_type` - (Optional, String) The type of the endpoint used to fetch secret. Supported options are `public`, and `private`. Default is `public`.
- `instance_id` - (Required, String) The secrets manager instance GUID.
- `secret_type` - (Optional, String) The secret type. Supported options are `arbitrary`, `iam_credentials`, `username_password`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 


- `id` - (String) The unique identifier of the secrets manager secrets.
- `metadata` - (String) The metadata that describes the resource array. Nested `metadata` blocks have the following structure.

  Nested scheme for `metadata`:
	- `collection_type` - (String) The type of resources in the resource array.
	- `collection_total` - (String) The number of elements in the resource array.
- `secrets`-  (String) A collection of secrets. Nested `secrets` blocks have the following structure.

  Nested scheme for `secrets`:
	- `access_groups` - (String) The access groups that define the capabilities of the service ID and API key that are generated for an `iam_credentials` secret. **Tip** To find the ID of an access group, go to **Manage > Access (IAM) > Access groups** in the IBM Cloud console. Select the access group to inspect, and click **Details** to view its ID.
    - `api_key` - (String) The API key that is generated for this secret.After the secret reaches the end of its lease (see the `ttl` field), the API key is deleted automatically. If you want to continue to use the same API key for future read operations, see the `reuse_api_key` field.
	- `crn` - (String) The Cloud Resource Name (CRN) that uniquely identifies your secrets manager resource.
    - `creation_date` - (String) The date the secret was created. The date format follows `RFC 3339`.
    - `created_by` - (String) The unique identifier for the entity that created the secret.
	- `description` - (String) An extended description of your secret. To protect your privacy, do not use personal data, such as your name or location, as a description for your secret.
	- `expiration_date` - (String) The date the secret material expires. The date format follows `RFC 3339` format. You can set an expiration date on supported secret types at their creation. If you create a secret without specifying an expiration date, the secret does not expire. The `expiration_date` field is supported for the following secret types `arbitrary`, and `username_password`.
	- `labels` - (String) Labels that you can use to filter for secrets in your instance. Only 30 labels can be created. Labels can be between `2-30` characters, including spaces. Special characters are not permitted include the angled bracket, comma, colon, ampersand, and vertical pipe character (`- `). To protect your privacy, do not use personal data, such as your name or location, as a label for your secret.
	- `last_update_date` - (String) Updates when the actual secret is modified. The date format follows `RFC 3339`.
	- `name` - (String) A human readable alias to assign to your secret. To protect your privacy, do not use personal data, such as your name or location, as an alias for your secret.
	- `secret_group_id` - (String) The `v4` UUID that uniquely identifies the secret group to assign to this secret. If you omit this parameter, your secret is assigned to the default secret group.
	- `secret_id ` - (String) The `v4` UUID that uniquely identifies the secret.
	- `state` - (String) The secret state based on `NIST SP 800-57`. States are integers and correspond to the `Pre-activation = 0`, `Active = 1`, `Suspended = 2`, `Deactivated = 3`, and `Destroyed = 5` values.
	- `state_description` - (String) A text representation of the secret state.
    - `secret_type`-  (String) The secret type.
	- `type` - (String) The `MIME` type that represents the secret.
    - `versions` - (String) An array that contains metadata for each secret version. Nested `versions` blocks have the following structure.

	  Nested scheme for `versions`:
	  - `auto_rotated` - (String) Indicates whether the version of the secret  created by automatic rotation.
	  - `creation_date` - (String) The date that the version of the secret was created.
	  - `created_by` - (String) The unique identifier for the entity that created the secret.
	  - `id` - (String) The ID of the secret version.
  - `next_rotation_date` - (String) The date that the secret is scheduled for automatic rotation. The service automatically creates a new version of the secret on its next rotation date. This field exists only for secrets that can be auto rotated and an existing rotation policy.
  - `payload` - (String) The secret data assigned to an `arbitrary` secret.
  - `password` - (String) The password assigned to an `username_password` secret.
  - `reuse_api_key` - (String) (IAM credentials) Reuse the service ID and API key for future read operations.
  - `secret_data` - (String) Map of username, password if secret_type is `username_password` else map of payload if secret_type is `arbitrary`.
  - `service_id` - (String) The service ID in which the API key (see the `api_key` field) is created. This service ID is added to the access groups that you assign for this secret.
   - `ttl` - (String) The time-to-live (`TTL`) or lease duration to assign to generated credentials. For `iam_credentials` secrets, the `TTL` defines for how long each generated API key remains valid. The value can be either an integer that specifies the number of seconds, or the string representation of a duration, such as 120 minutes or 24 hours.
  - `username` - (String) The username assigned to an `username_password` secret.

