---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: event_streams_quota"
description: |-
  Manages a quota of an IBM Event Streams service instance.
---

# ibm_event_streams_quota

Create, update or delete a quota of an Event Streams service instance. Both the default quota and user quotas may be managed. Quotas are only available on Event Streams Enterprise plan service instances. For more information about Event Streams quotas, see [Setting Kafka quotas](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-enabling_kafka_quotas).

## Example usage

### Sample 1: Create an Event Streams service instance and apply a default quota

Using `entity = default` in the quota resource sets the default quota, which applies to all users for which no user quota has been set.

```terraform
resource "ibm_resource_instance" "es_instance_1" {
  name              = "terraform-integration-1"
  service           = "messagehub"
  plan              = "enterprise-3nodes-2tb" 
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id

  timeouts {
     create = "3h"
     update = "1h"
     delete = "15m"
  }
}

resource "ibm_event_streams_quota" "es_quota_1" {
  resource_instance_id  = ibm_resource_instance.es_instance_1.id
  entity                = "default"
  producer_byte_rate    = 16384
  consumer_byte_rate    = 32768
}

```

### Sample 2: Create a user quota on an existing Event Streams instance

The quota is set for the user with the given IAM ID. The producer rate is limited, the consumer rate is unlimited (-1).

```terraform
data "ibm_resource_instance" "es_instance_2" {
  name              = "terraform-integration-2"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_quota" "es_quota_2" {
  resource_instance_id  = ibm_resource_instance.es_instance_2.id
  entity                = "iam-ServiceId-00001111-2222-3333-4444-555566667777"
  producer_byte_rate    = 16384
  consumer_byte_rate    = -1
}

```

## Argument reference

You must specify the following arguments for this resource.

- `resource_instance_id` - (Required, String) The ID or the CRN of the Event Streams service instance.
- `entity` - (Required, String) Either `default` to set the default quota, or an IAM ID for a user quota.
- `producer_byte_rate` - (Required, Integer) The producer quota in bytes/second. Use -1 for no quota.
- `consumer_byte_rate` - (Required, Integer) The consumer quota in bytes/second. Use -1 for no quota.

## Attribute reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - (String) The ID of the quota in CRN format. The last field of the CRN is either `default`, or the IAM ID of the user. See the examples in the import section.

## Import

The `ibm_event_streams_quota` resource can be imported by using the ID in CRN format. The three colon-separated parameters of the `CRN` are:
  - instance CRN  = CRN of the Event Streams instance
  - resource type = quota
  - quota entity = `default` or the IAM ID of the user
  
**Syntax**

```
$ terraform import ibm_event_streams_quota.es_quota <crn>

```

**Examples**

```
$ terraform import ibm_event_streams_quota.es_quota_default crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:ffffffff-eeee-dddd-cccc-bbbbaaaa9999:quota:default
$ terraform import ibm_event_streams_quota.es_quota_user crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:ffffffff-eeee-dddd-cccc-bbbbaaaa9999:quota:iam-ServiceId-00001111-2222-3333-4444-555566667777
```
