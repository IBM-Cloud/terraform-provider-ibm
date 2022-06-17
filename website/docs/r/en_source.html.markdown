---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_source'
description: |-
  Manages Event Notifications API Sources.
---

# ibm_en_source

Create, update, or delete a source by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_source" "en_source" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name          = "EN Source"
  description   = "API source for Event Notifications destinations"
  enabled       = true
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Source name.

- `description` - (Optional, String) The Source description.

- `enabled` - (Optional, bool) The enabled flag to enbale the created API source.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_source`.
- `source_id` - (String) The unique identifier of the created destination.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_source` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `source_id` in the following format:

```
<instance_guid>/<source_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `source_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_source.en_source <instance_guid>/<destination_id>
```
