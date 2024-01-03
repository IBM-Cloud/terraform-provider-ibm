---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_custom_sms'
description: |-
  Get information about a Custom SMS destination
---

# ibm_en_destination_custom_sms

Provides a read-only data source for Custom SMS destination. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_destination_custom_sms" "custom_sms_en_destination" {
  instance_guid  = ibm_resource_instance.en_terraform_test_resource.guid
  destination_id = ibm_en_destination_custom_sms.destination1.destination_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id` - (Required, String) Unique identifier for Destination.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `custom_sms_en_destination`.

- `name` - (String) Destination name.

- `description` - (String) Destination description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Destination type smtp_custom.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (List) Payload describing a destination configuration.
  Nested scheme for **config**:

  - `params` - (List)

  Nested scheme for **params**:

  - `domain` - (String) The Custom Domain.

- `updated_at` - (String) Last updated time.
