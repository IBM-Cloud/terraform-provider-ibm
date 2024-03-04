---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_custom_sms'
description: |-
  Get information about a custom sms subscription
---

# ibm_en_subscription_custom_sms

Provides a read-only data source for Custom SMS subscription. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_subscription_custom_sms" "custom_sms_subscription" {
  instance_guid   = ibm_resource_instance.en_terraform_test_resource.guid
  subscription_id = ibm_en_subscription_custom_sms.subscription_custom_sms.subscription_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `subscription_id` - (Required, String) Unique identifier for Subscription.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the custom_sms_subscription.

- `name` - (String) Subscription name.

- `description` - (String) Subscription description.

- `destination_id` - (String) The destination ID.

- `topic_id` - (String) Topic ID.

- `additional_properties` - (Required, List)

  - `susbscribed`- (Map) The phone number who have subscribed for topic.

  - `unsubscribed`- (List) The phone number which has opted for unsusbscribtion from that topic.

  - `invited`- (List) The phone number for invitation.

- `updated_at` - (String) Last updated time.
