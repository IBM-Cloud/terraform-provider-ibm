---
layout: "ibm"
page_title: "IBM : ibm_sm_secret_groups (Beta)"
description: |-
  Get information about SecretGroupCollection
subcategory: "Secrets Manager"
---

# ibm_sm_secret_groups

Provides a read-only data source for SecretGroupCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_secret_groups" {
  instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region        = "us-south"
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the SecretGroupCollection.
* `secret_groups` - (List) A collection of secret groups.
  * Constraints: The maximum length is `201` items. The minimum length is `1` item.
Nested scheme for **secret_groups**:
	* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.
	* `description` - (String) An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.
	* `id` - (String) A v4 UUID identifier, or `default` secret group.
	  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.
	* `name` - (String) The name of your existing secret group.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.
	* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

* `total_count` - (Integer) The total number of resources in a collection.
  * Constraints: The minimum value is `0`.

