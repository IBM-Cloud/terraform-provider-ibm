---
subcategory: "Event Streams"
layout: "ibm"
page_title: "IBM: ibm_event_streams_quota"
description: |-
  Get information about a quota of an IBM Event Streams service instance.
---

# ibm_event_streams_quota

Retrieve information about a quota of an Event Streams instance. Both the default quota and user quotas may be managed. Quotas are only available on Event Streams Enterprise plan service instances. For more information about Event Streams quotas, see [Setting Kafka quotas](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-enabling_kafka_quotas).

## Example usage

To retrieve the default quota:

```terraform
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_quota" "es_quota_default" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
  entity               = "default"
}
```

To retrieve a user quota, for a user with the given IAM ID:

```terraform
data "ibm_resource_instance" "es_instance" {
  name              = "terraform-integration"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_quota" "es_quota_user" {
  resource_instance_id = data.ibm_resource_instance.es_instance.id
  entity               = "iam-ServiceId-00001111-2222-3333-4444-555566667777"
}

## Argument reference

You must specify the following arguments for this data source.

- `resource_instance_id` - (Required, String) The ID or CRN of the Event Streams service instance.
- `entity` - (Required, String) Either `default` to set the default quota, or an IAM ID for a user quota.

## Attribute reference

After your data source is created, you can read values from the listed arguments and the following attributes.

- `id` - (String) The ID of the quota in CRN format. The last field of the CRN is either `default`, or the IAM ID of the user. For example, `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:ffffffff-eeee-dddd-cccc-bbbbaaaa9999:quota:default`, or `crn:v1:bluemix:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:ffffffff-eeee-dddd-cccc-bbbbaaaa9999:quota:iam-ServiceId-00001111-2222-3333-4444-555566667777`.
- `producer_byte_rate` - (Integer) The producer quota in bytes/second. If no producer quota is set, this is -1.
- `consumer_byte_rate` - (Integer) The consumer quota in bytes/second. If no consumer quota is set, this is -1.
