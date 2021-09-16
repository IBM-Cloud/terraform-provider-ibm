---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_schema"
description: |-
  Get information about an IBM Event Streams schema resource.
---

# ibm_event_streams_schema

Retrieve information about the Event Streams schema data sources. For more information, see documentation on Event Streams [schema registry](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_schema_registry)

## Example Usage

```terraform
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_schema" "es_schema" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
  schema_id = "my-es-schema"
}
```

## Argument Reference
Following are the argument parameters that you can specify for your data source:

- `schema_id` - (Required, string) The ID of a schema.
- `resource_instance_id` - (Required, string) The ID or CRN of the Event Streams service instance.

## Attribute Reference

In addition to the argument reference list, the following attribute reference can be accessed after data source is created:

- `id` - (String) The ID of the schema in CRN format. For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:schema:my-es-schema`.
- `kafka_http_url` - (String) The API endpoint for interacting with Event Streams REST API.

## Import

The `ibm_event_streams_schema` resource can be imported by using `CRN`. The three colon-separated parameters of the `CRN` are:
  - ID = CRN of the Event Streams instance
  - resource type = schema
  - resource = ID of the schema
  
**Syntax**

```
$ terraform import ibm_event_streams_schema.es_schema <crn>

```

**Example**

```
$ terraform import ibm_event_streams_schema.es_schema crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:schema:my-es-schema
```
