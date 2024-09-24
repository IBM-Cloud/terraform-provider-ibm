---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_mirroring_config"
description: |-
  Set the mirroring configuration of an IBM Event Streams instance.
---

# ibm_event_streams_mirroring_config


Set the mirroring topic patterns of an Event Streams service instance. This can only be performed on an Event Streams Enterprise plan service instance. For more information about the Event Streams mirroring, see [Event Streams Mirroring](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-mirroring).

**Note:** The mirroring config is a configuration property for defining topic selection pattern of the Event Streams instance, not a resource which must be created. The default value of the config is an empty list. When the terraform resource is created or updated, it changes the value of the configuration to the `mirroring_topic_patterns` argument; when the resource is deleted, it resets the value to empty list. For this reason, **only one mirroring config resource should be created** for an Event Streams service instance. Creating more than one resource for a given `resource_instance_id` will have unpredictable effects including terraform errors.

## Example usage

```terraform
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_mirroring_config" "es_mirroring_config" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
  mirroring_topic_patterns = ["topicA", "topicB"]
}
```

## Argument reference
Review the argument parameters that you can specify for your resource. 

- `resource_instance_id` - (Required, String) The ID or CRN of the Event Streams service instance.
- `mirroring_topic_patterns` - (Required, List of String) The topic selection patterns to set in instance

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created. 

- `id` - (String) The ID of the mirroring config in CRN format. For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:mirroring-config:`.

## Import

The `ibm_event_streams_mirroring_config` resource can be imported by using the rule's `CRN`, which is the `id` described above: the CRN of the service instance, with resource type "mirroring-config".

**Syntax**

```
$ terraform import ibm_event_streams_mirroring_config.es_mirroring_config <crn>
```

**Example**

```
$ terraform import ibm_event_streams_mirroring_config.es_mirroring_config crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:mirroring-config:
```
