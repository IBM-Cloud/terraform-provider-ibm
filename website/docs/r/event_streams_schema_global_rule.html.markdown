---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_schema_global_rule"
description: |-
  Set the global compatibility rule of an IBM Event Streams service instance.
---

# ibm_event_streams_schema_global_rule

Set the value of the global compatibility rule of an Event Streams service instance. This can only be performed on an Event Streams Enterprise plan service instance. For more information about the Event Streams schema registry, see [Event Streams Schema Registry](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_schema_registry).

**Note:** The global compatibility rule is a configuration property of the Event Streams schema registry, not a resource which must be created. The default value of the rule is "NONE". When the terraform resource is created or updated, it changes the value of the configuration to the `config` argument; when the resource is deleted, it resets the value to "NONE". For this reason, **only one global compatibility rule resource should be created** for an Event Streams service instance. Creating more than one resource for a given `resource_instance_id` will have unpredictable effects including terraform errors.

## Example usage

```terraform
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_schema_global_rule" "es_schema_global_rule" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
  config = "FORWARD"
}
```

## Argument reference
Following are the argument parameters that you can specify for your data source:

- `resource_instance_id` - (Required, String) The ID or CRN of the Event Streams service instance.
- `config` - (Required, String) The value of the global compatibility rule in the instance; one of "NONE", "FULL", "FULL_TRANSITIVE", "FORWARD", "FORWARD_TRANSITIVE", "BACKWARD", or "BACKWARD_TRANSITIVE".

## Attribute reference

In addition to the argument reference list, the following attribute reference can be accessed after the resource is created:

- `id` - (String) The ID of the schema global compatibility rule in CRN format. This will be the CRN of the service instance, with resource type "schema-global-compatibility-rule". For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:schema-global-compatibility-rule:`.

## Import

The `ibm_event_streams_schema_global_rule` resource can be imported by using the rule's `CRN`, which is the `id` described above: the CRN of the service instance, with resource type "schema-global-compatibility-rule".

**Syntax**

```
$ terraform import ibm_event_streams_schema_global_rule.es_global_rule <crn>

```

**Example**

```
$ terraform import ibm_event_streams_schema_global_rule.es_global_rule crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:schema-global-compatibility-rule:
```

