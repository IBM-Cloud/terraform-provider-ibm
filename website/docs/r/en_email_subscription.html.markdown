---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_email'
description: |-
  Manages Event Notifications SMS subscription.
---

# ibm_en_subscription_email

Create, update, or delete a Email subscription by using IBM Cloud™ Event Notifications.

## Example usage for Email Subscription Creation

```terraform
resource "ibm_en_subscription_email" "email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name           = "Email Certificate Subscription"
  description    = "Subscription for Certificate expiration alert"
  destination_id = "email_destination_id"
  topic_id       = ibm_en_topic.topic1.topic_id
  attributes {
      add_notification_payload = true
      reply_to_mail = "compliancealert@ibm.com"
      reply_to_name = "Compliance User"
      from_name="en@ibm.com"
      to = ["usernew1@gmail.com","testuser@gamil.com"]
  }
}
```

## Example usage for Email Subscription Updation

```terraform
resource "ibm_en_subscription_email" "email_subscription" {
  instance_guid    = "my_instance_guid"
  name           = "Email Certificate Subscription"
  description    = "Subscription for Certificate expiration alert"
  destination_id = "email_destination_id"
  topic_id       = "topicId"
  attributes {
      add_notification_payload = true
      reply_to_mail = "compliancealert@ibm.com"
      reply_to_name = "Compliance User"
      from_name="en@ibm.com"
      add = ["productionuser@ibm.com"]
      unsubscribed = ["testuser@gamil.com"]
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

  - `reply_to_name` - (String) The Email User Name to reply to.

  - `reply_to_mail` - (String) The email address to reply to.

  - `from_name` - (Optional, String) The email address user from which email is addressed.

  - `to`- (List) The Email address to send the email to.

  - `add`- (List) The Email address to add in case of updating the list of email addressses

  - `unsubscribed`- (List) The Email address list to be provided in case of removing the email addresses from subscription

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `email_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `updated_at` - (Required, String) Last updated time.

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
