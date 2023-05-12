---
layout: "ibm"
page_title: "IBM : ibm_sm_secret_group"
description: |-
  Get information about SecretGroup
subcategory: "Secrets Manager"
---

# ibm_sm_secret_group

Provides a read-only data source for a secret group. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_secret_group" "secret_group" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
  secret_group_id = ibm_sm_secret_group.sm_secret_group_instance.secret_group_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
* `secret_group_id` - (Required, String) The ID of the secret group.
  * Constraints: The maximum length is `36` characters. The minimum length is `7` characters. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|default)$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.
* `created_at` - (String) The date when a resource was created. The date format follows RFC 3339.

* `description` - (String) An extended description of your secret group.To protect your privacy, do not use personal data, such as your name or location, as a description for your secret group.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/(.*?)/`.

* `name` - (String) The name of your existing secret group.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/(.*?)/`.

* `updated_at` - (String) The date when a resource was recently modified. The date format follows RFC 3339.

