---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_custom_sms'
description: |-
  Manages Event Notifications Custom SMS subscription.
---

# ibm_en_subscription_custom_sms

Create, update, or delete a Custom SMS Destination subscription by using IBM Cloud™ Event Notifications.

## Example usage for SMS Subscription Creation

```terraform
resource "ibm_en_subscription_sms" "sms_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Non IBM SMS Destination Subscription"
  description      = "Custom SMS Destination subscription"
  destination_id   = ibm_en_destination_custom_sms.custom_sms_en_destination.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    invited = ["+15678923404", "+19643567389"]
  }
}
```

## Example usage for SMS Subscription Updation

```terraform
resource "ibm_en_subscription_custom_sms" "custom_sms_subscription" {
  instance_guid    = "my_instance_guid"
  name             = "Non IBM SMS Destination Subscription"
  description      = "Subscription for Certificate expiration alert"
  destination_id   = ibm_en_destination_custom_sms.custom_sms_en_destination.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
     add = ["+19643744902"]
     remove = ["+19807485102"]
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

  - `invited` - (Optional, List) The phone number to send the SMS to.

  - `add`- (List) The phone number to add in case of updating the list of contact

  - `reomve`- (List) The phone number list to be provided in case of removing the contact number from subscription

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `custon_sms_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_subscription_custom_sms` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription_custom_sms.custom_sms_subscription <instance_guid>/<subscription_id>
```
