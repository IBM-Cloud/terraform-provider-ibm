---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_ibmsource'
description: |-
  Manages Event Notifications IBM Sources.
---

# ibm_en_ibmsource

 update a IBM Cloud source registered with Event Notifications by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_ibmsource" "en_ibmsource" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  source_id = [for s in toset(data.ibm_en_sources.listsources.sources): s.id if s.type == "resource-lifecycle-events"].0
  enabled       = true
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `source_id` - (Optional, String) The Source id of the IBM Cloud source integrated with Event Notifications..

- `enabled` - (Optional, bool) The enabled flag to enbale the IBM Cloud source.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_ibmsource`.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_ibmsource` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `source_id` in the following format:

```
<instance_guid>/<source_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `source_id`: A string. Unique identifier for Source.

**Example**

```
$ terraform import ibm_en_ibmsource.en_ibmsource <instance_guid>/<source_id>
```
