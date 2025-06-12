---
layout: "ibm"
page_title: "IBM : ibm_en_smtp_allowed_ips"
description: |-
  Get information about en_smtp_allowed_ips
subcategory: "Event Notifications"
---

# ibm_en_smtp_allowed_ips

Provides a read-only data source to retrieve information about en_smtp_allowed_ips. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_en_smtp_allowed_ips" "en_smtp_allowed_ips" {
	en_smtp_allowed_ips_id = "en_smtp_allowed_ips_id"
	instance_id = "instance_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `en_smtp_allowed_ips_id` - (Required, Forces new resource, String) Unique identifier for SMTP.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `10` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the en_smtp_allowed_ips.
* `subnets` - (List) The SMTP allowed Ips.
  * Constraints: The list items must match regular expression `/.*/`. The maximum length is `100` items. The minimum length is `1` item.

* `updated_at` - (String) Updated at.

