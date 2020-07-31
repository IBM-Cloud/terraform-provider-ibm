---
layout: "ibm"
page_title: "IBM: event_streams_topic"
sidebar_current: "docs-ibm-resource-event-streams-topics"
description: |-
  Manages IBM Event Streams topics.
---

# ibm_event_streams_topic

The `event_streams_topic` resource represents a topic on an Event Streams instance.

## Example Usage 1: Create Event Streams service instance and topic
```hcl
resource "ibm_resource_instance" "es_instance_1" {
  name              = "terraform-integration-1"
  service           = "messagehub"
  plan              = "standard" # "lite", "enterprise-3nodes-2tb"
  location          = "us-south" # "us-east", "eu-gb", "eu-de", "jp-tok", "au-syd"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_topic" "es_topic_1" {
  resource_instance_id = ibm_resource_instance.es_instance_1.id
  name                 = "my-es-topic"
  partitions           = 1
  config = {
    "cleanup.policy"  = "compact,delete"
    "retention.ms"    = "86400000"
    "retention.bytes" = "1073741824"
    "segment.bytes"   = "536870912"
  }
}
```

## Example Usage 2: Create topic on an existing Event Streams instance`
```hcl
data "ibm_resource_instance" "es_instance_2" {
  name              = "terraform-integration-2"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_topic" "es_topic_2" {
  resource_instance_id = data.ibm_resource_instance.es_instance_2.id
  name                 = "my-es-topic"
  partitions           = 1
  config = {
    "cleanup.policy"  = "compact,delete"
    "retention.ms"    = "86400000"
    "retention.bytes" = "1073741824"
    "segment.bytes"   = "536870912"
  }
}
```

## Example Usage 3: Create a kafka consumer application connecting to an existing Event Streams instance and its topics
```hcl
data "ibm_resource_instance" "es_instance_3" {
  name              = "terraform-integration-3"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_topic" "es_topic_3" {
  resource_instance_id = data.ibm_resource_instance.es_instance_3.id
  name                 = "my-es-topic"
}

resource "kafka_consumer_app" "es_kafka_app" {
  bootstrap_server = lookup(data.ibm_resource_instance.es_instance_3.extensions, "kafka_brokers_sasl", [])
  topics           = [data.ibm_event_streams_topic.es_topic_3.name]
  apikey           = var.es_reader_api_key
}
```

## Argument Reference

The following arguments are supported:

- `resource_instance_id` - (Required, string) The ID/CRN of the Event Streams service instance.
- `name` - (Required, string) The name of the topic.
- `partitions` - (Optional, int) The number of partitions of the topic. Default value is 1 if not specified.
- `config` - (Optional, map) The configuration parameters of a topic. Supported configurations are: `cleanup.policy`, `retention.ms`, `retention.bytes`, `segment.bytes`, `segment.ms`, `segment.index.bytes`.

## Attribute Reference

The following attributes are exported:

- `id` (string) - The ID of the topic in CRN format. eg. `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:topic:my-es-topic`
- `kafka_http_url` (string) - The API endpoint for interacting with Event Streams REST API.
- `kafka_brokers_sasl` (array of strings) - Kafka brokers addresses for interacting with Kafka native API.


## Import

The `ibm_event_streams_topic` resource can be imported using the `id`. The ID is the `CRN` (Cloud Resource Name), the `resource type` is `topic`, `resource` is the name of the topic.

```
$ terraform import ibm_event_streams_topic.es_topic <crn>

$ terraform import ibm_event_streams_topic.es_topic crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:topic:my-es-topic
```
