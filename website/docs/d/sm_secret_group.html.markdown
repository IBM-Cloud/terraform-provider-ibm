---
layout: "ibm"
page_title: "IBM : ibm_sm_secret_group" (Beta)
description: |-
  Get information about SecretGroup
subcategory: "IBM Cloud Secrets Manager API"
---

# ibm_sm_secret_group

Provides a read-only data source for SecretGroup. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_secret_group" {
  instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region        = "us-south"
  id = ibm_sm_secret_group.sm_secret_group_instance.secretGroup_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The ID of the secret group.
  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the SecretGroup.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `description` - (String) An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.

* `name` - (String) The name of your existing secret group.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

