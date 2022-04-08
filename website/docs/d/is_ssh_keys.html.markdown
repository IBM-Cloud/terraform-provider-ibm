---
layout: "ibm"
page_title: "IBM : ibm_is_ssh_keys"
description: |-
  Get information about KeyCollection
subcategory: "VPC infrastructure"
---

# ibm_is_ssh_keys

Provides a read-only data source for KeyCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_ssh_keys" "example" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the KeyCollection.
- `keys` - (List) Collection of keys.
Nested scheme for **keys**:
	- `created_at` - (String) The date and time that the key was created.
	- `crn` - (String) The CRN for this key.
	- `fingerprint` - (String) The fingerprint for this key.  The value is returned base64-encoded and prefixed with the hash algorithm (always `SHA256`).
	- `href` - (String) The URL for this key.
	  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `id` - (String) The unique identifier for this key.
	  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	- `length` - (Integer) The length of this key (in bits).
	  - Constraints: Allowable values are: `2048`, `4096`.
	- `name` - (String) The unique user-defined name for this key. If unspecified, the name will be a hyphenated list of randomly-selected words.
	  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$/`.
	- `public_key` - (String) The public SSH key, consisting of two space-separated fields: the algorithm name, and the base64-encoded key.
	- `resource_group` - (List) The resource group for this key.
	Nested scheme for **resource_group**:
		- `href` - (String) The URL for this resource group.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this resource group.
		  - Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
		- `name` - (String) The user-defined name for this resource group.
		  - Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`.
	- `type` - (String) The crypto-system used by this key.
	  - Constraints: The default value is `rsa`. Allowable values are: `rsa`.
