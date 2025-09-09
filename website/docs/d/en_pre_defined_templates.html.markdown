---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_pre_defined_templates'
description: |-
  List all the Pre Defined Templates
---

# ibm_en_pre_defined_templates

Provides a read-only data source for Pre Defined templates. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_pre_defined_templates" "pre_defined_templates" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  source = "logs"
  type = "slack.notification"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `search_key` - (Optional, String) Filter the template by name or type.

- `source` - (Required, String) Source Type.

- `type` - (Required, String) Template type.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `pre_defined_templates`.

- `templates` - (List) List of templates.

  - `id` - The unique identifier of the `pre_defined_templates`.

  - `name` - (String) The Template name.

  - `description` - (String) The template description.

  - `type` - (String) Template type.

  - `source` - (String) Source Type.

  - `updated_at` - (String) Last updated time.

- `total_count` - (Integer) Total number of destinations.
