---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_event_streams'
description: |-
  Manages Event Notifications subscription.
---

# ibm_en_subscription_event_streams

Create, update, or delete a Event Streams subscription by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_subscription_event_streams" "es_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "EN Event Streams subscription"
  description      = "Subscription for event streams destination in Event Notifications"
  destination_id   = ibm_en_destination_event_streams.destination1.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    template_id_notification = "e40843c8-hgft-4717-8ee4-f923f2786a34"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Requires, String) Subscription name.

- `description` - (Optional, String) Subscription description.

- `destination_id` - (Requires, String) Destination ID.

- `topic_id` - (Required, String) Topic ID.

- `attributes` - (Optional, List) Subscription attributes.
  Nested scheme for **attributes**:

  - `template_id_notification` - (Optional, String) The templete id for notification.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `es_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_subscription_event_streams` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription_event_streams.es_subscription <instance_guid>/<subscription_id>
```
