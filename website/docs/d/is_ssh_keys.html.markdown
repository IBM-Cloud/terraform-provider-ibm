---
layout: "ibm"
page_title: "IBM : ibm_is_ssh_keys"
description: |-
  Get information about KeyCollection
subcategory: "Virtual Private Cloud API"
---

# ibm_is_ssh_keys

Provides a read-only data source for KeyCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_ssh_keys" "is_ssh_keys" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the KeyCollection.
* `first` - (Required, List) A link to the first page of resources.
Nested scheme for **first**:
	* `href` - (Required, String) The URL for a page of resources.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `keys` - (Required, List) Collection of keys.
Nested scheme for **keys**:
	* `created_at` - (Required, String) The date and time that the key was created.
	* `crn` - (Required, String) The CRN for this key.
	* `fingerprint` - (Required, String) The fingerprint for this key.  The value is returned base64-encoded and prefixed with the hash algorithm (always `SHA256`).
	* `href` - (Required, String) The URL for this key.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The unique identifier for this key.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `length` - (Required, Integer) The length of this key (in bits).
	  * Constraints: Allowable values are: `2048`, `4096`.
	* `name` - (Required, String) The unique user-defined name for this key. If unspecified, the name will be a hyphenated list of randomly-selected words.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$/`.
	* `public_key` - (Required, String) The public SSH key, consisting of two space-separated fields: the algorithm name, and the base64-encoded key.
	* `resource_group` - (Required, List) The resource group for this key.
	Nested scheme for **resource_group**:
		* `href` - (Required, String) The URL for this resource group.
		  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (Required, String) The unique identifier for this resource group.
		  * Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
		* `name` - (Required, String) The user-defined name for this resource group.
		  * Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`.
	* `type` - (Required, String) The crypto-system used by this key.
	  * Constraints: The default value is `rsa`. Allowable values are: `rsa`.

* `limit` - (Required, Integer) The maximum number of resources that can be returned by the request.
  * Constraints: The maximum value is `100`. The minimum value is `1`.

* `next` - (Optional, List) A link to the next page of resources. This property is present for all pagesexcept the last page.
Nested scheme for **next**:
	* `href` - (Required, String) The URL for a page of resources.
	  * Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `total_count` - (Required, Integer) The total number of resources across all pages.
  * Constraints: The minimum value is `0`.

