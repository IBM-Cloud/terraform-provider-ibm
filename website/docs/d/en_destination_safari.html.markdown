---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_safari'
description: |-
  Get information about a Safari destination
---

# ibm_en_destination_safari

Provides a read-only data source for Webhook destination. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_destination_safari" "safari_en_destination" {
  instance_guid  = ibm_resource_instance.en_terraform_test_resource.guid
  destination_id = ibm_en_destination_safari.destination1.destination_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id` - (Required, String) Unique identifier for Destination.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `safari_en_destination`.

- `name` - (String) Destination name.

- `description` - (String) Destination description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Destination type push_safari.

- `config` - (List) Payload describing a destination configuration.
  Nested scheme for **config**:

  - `params` - (List)

  Nested scheme for **params**:

  - `cert_type` - (String) The Certificate type. Value is p12.

  - `password` - (String) The password string for p12 certificate.

  - `url_format_string` - (String) Website formatted url .

  - `website_name` - (String) The name of website.

  - `website_push_id` - (String) Website push ID  .

  - `website_url` - (String) Website url.


- `updated_at` - (String) Last updated time.
