---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_email'
description: |-
  Manages Event Notifications Email subscription.
---

# ibm_en_subscription_email

Create, update, or delete an Email subscription by using IBM Cloud™ Event Notifications.

## Example usage for Email Subscription Creation

```terraform
resource "ibm_en_subscription_email" "email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Email Certificate Subscription"
  description      = "Subscription for Certificate expiration alert"
  destination_id   = [for s in toset(data.ibm_en_destinations.destinations.destinations): s.id if s.type == "smtp_ibm"].0
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    add_notification_payload = true
    reply_to_mail            = "compliancealert@ibm.com"
    reply_to_name            = "Compliance User"
    from_name                = "en@ibm.com"
    invited                  = ["usernew1@gmail.com", "testuser@gmail.com"]
  }
}
```

## Example usage for Email Subscription Update

```terraform
resource "ibm_en_subscription_email" "email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Email Certificate Subscription"
  description      = "Subscription for Certificate expiration alert"
  destination_id   = "email_destination_id"
  topic_id         = "topicId"
  attributes {
    add_notification_payload = true
    reply_to_mail            = "compliancealert@ibm.com"
    reply_to_name            = "Compliance User"
    from_name                = "en@ibm.com"
    invited                  = ["productionuser@ibm.com"]
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

  - `add_notification_payload` - (Optional, Boolean) Whether to add the notification payload to the email.

  - `reply_to_name` - (Optional, String) The name of the email address user to reply to.

  - `reply_to_mail` - (Optional, String) The email address to reply to.

  - `from_name` - (Optional, String) The email address from which email is sourced.

  - `invited` - (Optional, List) The email addresses to invite. Add an address by adding it to this list; remove an address by removing it from this list.

  - `subscribed` - (Computed, List) Email addresses that have accepted the invitation and are currently subscribed. Populated by the service; read-only.

  - `unsubscribed` - (Computed, List) Email addresses that have unsubscribed. Populated by the service; read-only.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `email_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `destination_type` - (String) The type of Destination.

- `destination_name` - (String) The Destination name.

- `topic_name` - (String) Name of the topic.

- `from` - (String) From Email ID (displayed only for `smtp_ibm` destination type).

- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_subscription_email` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription_email.email_en_subscription <instance_guid>/<subscription_id>
```
