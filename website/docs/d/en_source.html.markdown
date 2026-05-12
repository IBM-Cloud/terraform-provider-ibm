---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_source'
description: |-
  Get information about a Source
---

# ibm_en_source

Provides a read-only data source for API sources. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_source" "en_source" {
  instance_guid  = ibm_resource_instance.en_terraform_test_resource.guid
  source_id = ibm_en_source.destination1.source_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `source_id` - (Required, String) Unique identifier for API Source.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `en_source`.

- `name` - (String) Source name.

- `description` - (String) Source description.

- `enabled` - (bool) Flag to enable/disable the api source.

- `updated_at` - (String) Last updated time.
