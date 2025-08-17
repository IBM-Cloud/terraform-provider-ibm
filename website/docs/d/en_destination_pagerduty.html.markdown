---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_pagerduty'
description: |-
  Manages Event Notification Pagerduty destinations.
---

# ibm_en_destination_pagerduty

Provides a read-only data source for Pagerduty destination. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform

data "ibm_en_destination_pagerduty" "pagerduty_en_destination" {
  instance_guid  = ibm_resource_instance.en_terraform_test_resource.guid
  destination_id = ibm_en_destination_pagerduty.destination1.destination_id
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id` - (Required, String) Unique identifier for Destination.
## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the `pagerduty_en_destination`.

- `name` - (String) Destination name.

- `description` - (String) Destination description.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscription_names` - (List) List of subscriptions.

- `type` - (String) Destination type pagerduty.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (List) Payload describing a destination configuration.
  Nested scheme for **config**:

  - `params` - (List)

  Nested scheme for **params**:

  - `api_key` - (Optional, string) The apikey required to validate user for the assigned group[The parameter has been deprecated from destination config parameter, it will be removed in future].

  - `routing_key` - (Required, string) The integration key required to route the events to pagerduty.

- `updated_at` - (String) Last updated time.