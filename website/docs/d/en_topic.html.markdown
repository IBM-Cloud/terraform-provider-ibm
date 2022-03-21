---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_topic'
description: |-
  Get information about a topic
---

# ibm_en_topic

Provides a read-only data source for topic. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_en_topic" "en_topic" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  topic_id = ibm_en_topic.topic1.topic_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `topic_id` - (Required, String) Unique identifier for Topic.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the en_topic.

- `name` - (String) Name of the topic.

- `description` - (String) Description of the topic.

- `source_count` - (Integer) Number of sources.

- `sources_names` - (List) List of source name.

- `subscription_count` - (Integer) Number of subscriptions.

- `subscriptions` - (List) List of subscriptions.

  - `id`- (String) Subscription ID.

  - `name`- (String) Subscription name.

  - `description` - (String) Subscription description.

  - `destination_id` - (String) The destination ID.

  - `destination_name` - (String) The destination name.

  - `destination_type`- (String) The type of destination.

  - `from` - (String) From Email ID (it will be displayed only in case of smtp_ibm destination type).

  - `topic_id` - (String) Topic ID.

  - `topic_name` - (String) Topic name.

  - `updated_at` - (String) Last updated time.

- `updated_at` - (String) Last time the topic was updated.
