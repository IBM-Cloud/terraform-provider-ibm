---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_android'
description: |-
  Get information about a FCM destination
---

# ibm_en_destination_android

Provides a read-only data source for FCM destination. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_destination_android" "android_en_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  destination_id = ibm_en_destination_android.destinationandroidnew.destination_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id` - (Required, String) Unique identifier for Destination.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `android_en_destination`.

- `name` - (String) Destination name.

- `description` - (String) Destination description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Destination type push_android.

- `config` - (List) Payload describing a destination configuration.
  Nested scheme for **config**:

  - `params` - (List)

  Nested scheme for **params**:

  - `sender_id` - (String) Sender ID for your FCM Destination Configured.

  - `server_key` - (String) Server Key for FCM Destination configured.

- `updated_at` - (String) Last updated time.
