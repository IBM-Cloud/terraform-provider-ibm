---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: event_streams_schema"
description: |-
  Manages IBM Event Streams schema.
---

# ibm_event_streams_schema

Create, update or delete the Event Streams schemas. The schema operations can only be performed on an Event Streams Enterprise plan service instances. For more information, about Event Streams schema, see [Event Streams Schema Registry](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_schema_registry).

## Example usage

### Sample 1: Create an Event Streams service instance and a schema


```terraform
resource "ibm_resource_instance" "es_instance_1" {
  name              = "terraform-integration-1"
  service           = "messagehub"
  plan              = "enterprise-3nodes-2tb" 
  location          = "us-south" # "us-east", "eu-gb", "eu-de", "jp-tok", "au-syd"
  resource_group_id = data.ibm_resource_group.group.id

  # parameters = {
  #   service-endpoints     = "private"                    # for enterprise instance only, Options are: "public", "public-and-private", "private". Default is "public" when not specified.
  #   private_ip_allowlist = "[10.0.0.0/32,10.0.0.1/32]" # for enterprise instance only. Specify 1 or more IP range in CIDR format.
  #   # Refer private service endpoint and IP allow list to restrict access documentation, (
/docs/EventStreams?topic=EventStreams-restrict_access) for more details.
  #   throughput   = "150"  # for enterprise instance only. Options are: "150", "300", "450". Default is "150".
  #   storage_size = "2048" # for enterprise instance only. Options are: "2048", "4096", "6144", "8192", "10240", "12288". Default is "2048".
  #   Note: When throughput is "300", storage_size starts from "4096",  when throughput is "450", storage_size starts from "6144".
  #   Refer support combinations of throughput and storage_size documentation (
/docs/EventStreams?topic=EventStreams-ES_scaling_capacity#ES_scaling_combinations) for more details.
  # }

  # timeouts {
  #   create = "3h" # use 3h when creating enterprise instance, add more 1h for each level of non-default throughput, add more 30m for each level of non-default storage_size
  #   update = "1h" # use 1h when updating enterprise instance, add more 1h for each level of non-default throughput, add more 30m for each level of non-default storage_size
  #   delete = "15m"
  # }
}

resource "ibm_event_streams_schema" "es_schema_1" {
  resource_instance_id  = ibm_resource_instance.es_instance_1.id
  schema_id           = "my-es-schema"
  schema = <<SCHEMA
   {
           "type": "record",
           "name": "record_name",
           "fields" : [
             {"name": "value_1", "type": "long"},
             {"name": "value_2", "type": "string"}
           ]
         }
  SCHEMA
}

```

### Sample 2: Create a schema on an existing Event Streams instance

Create a schema on an existing Event Streams Enterprise plan service instance.
 
```terraform
data "ibm_resource_instance" "es_instance_2" {
  name              = "terraform-integration-2"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_schema" "es_schema_2" {
  resource_instance_id  = ibm_resource_instance.es_instance_2.id
  schema_id           = "my-es-schema"
  schema = <<SCHEMA
   {
           "type": "record",
           "name": "record_name",
           "fields" : [
             {"name": "value_1", "type": "long"},
             {"name": "value_2", "type": "string"}
           ]
         }
  SCHEMA
}

```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `schema` - (Required, String) The schema in JSON format.
- `resource_instance_id` - (Required, String) The ID or the CRN of the Event Streams service instance.
- `schema_id` - (Optional, String) The unique ID to be assigned to schema. If this value is not specified, a generated `UUID` is assigned.

## Attribute reference

In addition to the above argument reference list, the following attribute reference can be accessed after the resource is created. 

- `id` - (String) The ID of the schema in CRN format. For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:schema:my-es-schema`.
- `kafka_http_url` - (String) The API endpoint for interacting with an Event Streams REST API.

## Import

The `ibm_event_streams_schema` resource can be imported by using `CRN`. The three colon-separated parameters of the `CRN` are:
  - instance CRN  = CRN of the Event Streams instance
  - resource type = schema
  - schema ID = ID of the schema
  
**Syntax**

```
$ terraform import ibm_event_streams_schema.es_schema <crn>

```

**Example**

```
$ terraform import ibm_event_streams_schema.es_schema crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:cb5a0252-8b8d-4390-b017-80b743d32839:schema:my-es-schema
```
