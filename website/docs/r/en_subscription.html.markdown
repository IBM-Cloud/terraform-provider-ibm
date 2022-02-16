---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription'
description: |-
  Manages Event Notifications subscription.
---

# ibm_en_subscription

Create, update, or delete a subscription by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_subscription" "en_subscription" {
  instance_guid    = "instance_guid"
  name           = "name"
  description    = "description"
  destination_id = "destinationId"
  topic_id       = "topicId"
  attributes {
    add_notification_payload = true
    signing_enabled          = true
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

  - `add_notification_payload` - (Optional, Boolean) Add notification payload.

  - `signing_enabled` - (Optional, Boolean) Signing enabled.

  - `add_notification_payload` - (Optional, Boolean) Whether to add the notification payload to the email.

  - `reply_to` - (Optional, String) The email address to reply to.

  - `to` - (Optional, List) The phone number to send the SMS to.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `en_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `destination_name` - (Required, String) The destination name.

- `destination_type`- (Required, String) The type of destination.

- `topic_name` - (Required, String) Topic name.

- `from` - (Optional, String) From Email ID (it will be displayed only in case of smtp_ibm destination type).

- `updated_at` - (Required, String) Last updated time.

## Import

You can import the `ibm_en_subscription` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription.en_subscription <instance_guid>/<subscription_id>
```
