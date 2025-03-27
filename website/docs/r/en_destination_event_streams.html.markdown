---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_event_streams'
description: |-
  Manages Event Notification Event Streams destinations.
---

# ibm_en_destination_event_streams

Create, update, or delete a IBM Event Streams destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_event_streams" "es_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Event Streams Destination"
  type                  = "event_streams"
  description           = "Event Streams Destination for event notification"
  config {
    params {
        crn = "crn:v1:bluemix:public:messagehub:us-south:a/9f007405a9fe4a5d9345fa8c131610c8:a292db6e-af78-4c0b-b3db-7d6794b40aeb::"
				endpoint = "https://n6627w6t7y62chudi.svc09.us-south.eventstreams.cloud.ibm.com"
				topic = "test_topic"
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) event_streams.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `topic` - (Required, string) Topic of Event Streams.
  - `crn` - (Required, string) CRN of the Event Streans instance.
  - `endpoint`   - (Required, string) End Point of Event Streams.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `es_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_event_streams` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_event_streams.es_en_destination <instance_guid>/<destination_id>
```
