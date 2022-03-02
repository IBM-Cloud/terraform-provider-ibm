---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_ios'
description: |-
  Get information about a IOS subscription
---

# ibm_en_subscription_ios

Provides a read-only data source for APNS subscription. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_subscription_ios" "data_apns_subscription" {
  instance_guid = "my_instance_id"
  subscription_id = "apns_subscription_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `subscription_id` - (Required, String) Unique identifier for Subscription.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the en_subscription.

- `name` - (String) Subscription name.

- `description` - (String) Subscription description.

- `destination_id` - (String) The destination ID.

- `topic_id` - (String) Topic ID.

- `updated_at` - (String) Last updated time.
