---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_topic"
description: |-
  Get information about an IBM Event Streams topic resource.
---

# ibm_event_streams_topic


Review the [Event Streams](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-about) resource that you can connect, administer, developed with Event Streams and integrate with the other services. 

## Example Usage

```terraform
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
Review the argument parameters that you can specify for your data source. 

- `name` - (Required, string) The name of the topic.
- `resource_instance_id` - (Required, string) The ID or CRN of the Event Streams service instance.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your data source is created. 

- `id` - (String) The ID of the topic in CRN format. For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:topic:my-es-topic`.
- `kafka_http_url` - (String) The API endpoint for interacting with Event Streams REST API.
- `kafka_brokers_sasl` - (Array of strings) Kafka brokers uses for interacting with Kafka native API.
