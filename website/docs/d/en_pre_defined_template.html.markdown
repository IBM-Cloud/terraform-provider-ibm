---
layout: "ibm"
page_title: "IBM : ibm_en_pre_defined_template"
description: |-
  Get information about en_pre_defined_template
subcategory: "Event Notifications"
---

# ibm_en_pre_defined_template

Provides a read-only data source to retrieve information about an en_pre_defined_template. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_en_pre_defined_template" "en_pre_defined_template" {
	template_id = "en_pre_defined_template_id"
	instance_guid = "instance_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `template_id` - (Required, Forces new resource, String) Unique identifier for Template.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the en_pre_defined_template.
* `description` - (String) Template description.
  * Constraints: The maximum length is `255` characters. The minimum length is `0` characters. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.

* `name` - (String) Template name.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.

* `source` - (String) The type of source.
  * Constraints: The minimum length is `1` character. The value must match regular expression `/.*/`.

- `params` - (String) base64 encoded template body

* `type` - (String) The type of template.
  * Constraints: The minimum length is `1` character. The value must match regular expression `/.*/`.

* `updated_at` - (String) Updated at.

