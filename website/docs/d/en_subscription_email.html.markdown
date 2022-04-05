---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_email'
description: |-
  Get information about a Email subscription
---

# ibm_en_subscription_email

Provides a read-only data source for Email subscription. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_subscription_email" "email_subscription" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  subscription_id = ibm_en_subscription_email.subscriptionemail.subscription_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `subscription_id` - (Required, String) Unique identifier for Subscription.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the email_subscription.

- `name` - (String) Subscription name.

- `description` - (String) Subscription description.

- `destination_id` - (String) The destination ID.

- `topic_id` - (String) Topic ID.

- `add_notification_payload` - (Boolean) Whether to add the notification payload to the email.

- `additional_properties` - (Required, List)

  - `reply_to_name` - (String) The Email User Name to reply to.

  - `reply_to_mail` - (String) The email address to reply to.

  - `signing_enabled`- (Boolean) Signing webhook attributes.

  - `to`- (Map) The Email address to send the email to.

  - `unsubscribed`- (List) The Email address which are unsusbscribed.

  - `invited`- (List) The Email address for invitation.

- `updated_at` - (String) Last updated time.
