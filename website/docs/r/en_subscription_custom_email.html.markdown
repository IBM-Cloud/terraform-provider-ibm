---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_custom_email'
description: |-
  Manages Event Notifications Custom Email subscription.
---

# ibm_en_subscription_custom_email

Create, update, or delete a Custom Email subscription by using IBM Cloud™ Event Notifications.

## Example usage for Custom Email Subscription Creation

```terraform
resource "ibm_en_subscription_custom_email" "custom_domain_email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Custom Domain Email Subscription"
  description      = "Subscription for Certificate expiration alert"
  destination_id   = ibm_resource_instance.ibm_en_subscription_custom_email.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
      add_notification_payload = true
      reply_to_mail = "en@ibm.com"
      reply_to_name = "EYS ORG"
      from_name="ABC ORG"
			from_mail="Testuser@mailx.com"
      invited = ["test@gmail.com"]
  }
}
```

## Example usage for Email Subscription Updation

```terraform
resource "ibm_en_subscription_custom_email" "custom_domain_email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Custom Domain Email Subscription"
  description      = "Subscription for Certificate expiration alert"
  destination_id   = ibm_resource_instance.ibm_en_subscription_custom_email.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
      add_notification_payload = true
      reply_to_mail            = "en@ibm.com"
      reply_to_name            = "EYS ORG"
      from_name                = "ABC ORG"
			from_mail                = "Testuser@mailx.com"
      add                      = ["productionuser@ibm.com"]
      remove                   = ["testuser@gamil.com"]
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

  - `from_name` - (Optional, String) The user name from which email is addressed.

  - `from_name` - (Optional, String) The email address user from which email is addressed(Should belong to the custom domain).

  - `invited`- (List) The Email address to send the email to.

  - `add`- (List) The Email address to add in case of updating the list of email addressses

  - `reomve`- (List) The Email address list to be provided in case of removing the email addresses from subscription

  - `template_id_notification` - (Optional, String) The templete id for notification.

  - `template_id_invitation` - (Optional, String) The templete id for invitation.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `custom_domain_email_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_subscription_custom_email` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription_custom_email.custom_domain_email_en_subscription <instance_guid>/<subscription_id>
```
