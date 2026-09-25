---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_sms'
description: |-
  Manages Event Notifications SMS subscription.
---

# ibm_en_subscription_sms

Create, update, or delete an SMS subscription by using IBM Cloud™ Event Notifications.

## Example usage for SMS Subscription Creation

```terraform
resource "ibm_en_subscription_sms" "sms_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "News Subscription"
  description      = "SMS subscription for news alert"
  destination_id   = [for s in toset(data.ibm_en_destinations.destinations.destinations): s.id if s.type == "sms_ibm"].0
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    invited = ["+15678923404", "+19643567389"]
  }
}
```

## Example usage for SMS Subscription Update

```terraform
resource "ibm_en_subscription_sms" "sms_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "IBM SMS Certificate Subscription"
  description      = "Subscription for Certificate expiration alert"
  destination_id   = [for s in toset(data.ibm_en_destinations.destinations.destinations): s.id if s.type == "sms_ibm"].0
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    invited = ["+19643744902"]
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) Subscription name.

- `description` - (Optional, String) Subscription description.

- `destination_id` - (Required, Forces new resource, String) Destination ID.

- `topic_id` - (Required, Forces new resource, String) Topic ID.

- `attributes` - (Optional, List) Subscription attributes.

  Nested scheme for **attributes**:

  - `invited` - (Optional, List) The phone numbers to invite. Add a number by adding it to this list; remove a number by removing it from this list.

  - `subscribed` - (Computed, List) Phone numbers that have accepted the invitation and are currently subscribed. Populated by the service; read-only.

  - `unsubscribed` - (Computed, List) Phone numbers that have unsubscribed. Populated by the service; read-only.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `sms_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `destination_type` - (String) The type of Destination.

- `destination_name` - (String) The Destination name.

- `topic_name` - (String) Name of the topic.

- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_subscription_sms` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription_sms.sms_subscription <instance_guid>/<subscription_id>
```
