---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscriptions'
description: |-
  List all the subscription
---

# ibm_en_subscriptions

Provides a read-only data source for en_subscriptions. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_subscriptions" "en_subscriptions" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `search_key` - (Optional, String) Filter the subscription by name.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the en_subscriptions.

- `subscriptions` - (Required, List) List of subscriptions.

  - `id` - (Required, String) ID of the subscription.

  - `name` - (Required, String) Name of the subscription.

  - `description` - (Required, String) Description of the subscription.

  - `destination_id` - (Required, String) ID of the destination.

  - `destination_name` - (Required, String) Name of the destination.

  - `destination_type` - (Required, String) The type of destination.

  - `topic_id` - (Required, String) ID of the topic.

  - `topic_name` - (Required, String) Name of the topic.

  - `updated_at` - (Required, String) Last updated time of the subscription.

- `total_count` - (Required, Integer) Number of subscriptions.
