---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_ios'
description: |-
  Get information about a IOS destination
---

# _ibm_en_destination_ios

Provides a read-only data source for destination. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_destination_ios" "ios_en_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  destination_id = ibm_en_destination_ios.destinationiosp8.destination_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id` - (Required, String) Unique identifier for Destination.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `ios_en_destination`.

- `name` - (String) Destination name.

- `description` - (String) Destination description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Destination type push_ios.

- `certificate_content_type` - (String) The type of certificate, Values are p8/p12.

- `certificate` - (binary) Certificate file. The file type allowed is .p8 and .p12

- `config` - (List) Payload describing a destination configuration.
  Nested scheme for **config**:

  - `params` - (List)

  Nested scheme for **params**:

  - `cert_type` - (String) The Certificate type. Values are p8/p12.

  - `is_sandbox` - (boolean) The flag for sandbox/production environment.

  - `password` - (String) The password string for p12 certificate.

  - `team_id` - (String) The team_id value in case P8 certificate.

  - `key_id` - (String) The team_id value in case P8 certificate.

  - `bundle_id` - (String) The team_id value in case P8 certificate.

- `updated_at` - (String) Last updated time.