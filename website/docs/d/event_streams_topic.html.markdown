---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_topic"
description: |-
  Get information about an IBM Event Streams topic resource.
---

# ibm_event_streams_topic

Import the name of an existing Event Streams topic as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_topic" "es_topic" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
  name                 = "my-es-topic"
}
```

## Argument Reference

The following arguments are supported:
- `resource_instance_id` - (Required, string) The ID/CRN of the Event Streams service instance.
- `name` - (Required, string) The name of the topic.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` (string) - The ID of the topic in CRN format. eg. `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:topic:my-es-topic`.
- `kafka_http_url` (string) - The API endpoint for interacting with Event Streams REST API.
- `kafka_brokers_sasl` (array of strings) - Kafka brokers addresses for interacting with Kafka native API.

