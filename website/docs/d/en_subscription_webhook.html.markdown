---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_webhook'
description: |-
  Get information about a webhook subscription
---

# ibm_en_subscription_webhook

Provides a read-only data source for subscription. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_subscription_webhook" "webhook_subscription" {
  instance_guid   = ibm_resource_instance.en_terraform_test_resource.guid
  subscription_id = ibm_en_subscription_webhook.subscriptionwebhook.subscription_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `subscription_id` - (Required, String) Unique identifier for Subscription.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the webhook_subscription.

- `name` - (String) Subscription name.

- `description` - (String) Subscription description.

- `destination_id` - (String) The destination ID.

- `destination_name` - (String) The destination name.

- `destination_type` - (String) The type of destination.

- `from` - (Optional, String) From Email ID (it will be displayed only in case of smtp_ibm destination type).

- `topic_id` - (String) Topic ID.

- `topic_name` - (String) Topic name.

- `attributes` - (Required, List)

  - `add_notification_payload` - (Boolean) Whether to add the notification payload to the email.

- `updated_at` - (String) Last updated time.
