---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_mirroring_config"
description: |-
  Get information about an IBM Event Streams mirroring configuration resource.
---

# ibm_event_streams_mirroring_config


Retrieve information about the mirroring config of an Event Streams service instance. This can only be performed on an Event Streams Enterprise plan service instance. For more information about the Event Streams mirroring, see [Event Streams Mirroring](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-mirroring).

## Example usage

```terraform
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_mirroring_config" "es_mirroring_config" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
}
```

## Argument reference
Review the argument parameters that you can specify for your data source. 

- `resource_instance_id` - (Required, string) The ID or CRN of the Event Streams service instance.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your data source is created. 

- `id` - (String) The ID of the mirroring config in CRN format. For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:mirroring-config:`.
- `mirroring_topic_patterns` - (List of String) The current topic selection patterns in the Event Streams instance.
