---
layout: "ibm"
page_title: "IBM : ibm_logs_streams"
description: |-
  Get information about logs_streams
subcategory: "Cloud Logs"
---

# ibm_logs_streams

Provides a read-only data source to retrieve information about logs_streams. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_streams" "logs_streams" {
	instance_id = ibm_resource_instance.logs_instance.guid
  	region      = ibm_resource_instance.logs_instance.location
}
```

## Argument Reference

You can specify the following arguments for this resource.
* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_streams.
* `streams` - (List) Collection of Event Streams.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **streams**:
	* `compression_type` - (String) The compression type of the stream.
	  * Constraints: Allowable values are: `unspecified`, `gzip`.
	* `created_at` - (String) The creation time of the Event stream.
	* `dpxl_expression` - (String) The DPXL expression of the Event stream.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `ibm_event_streams` - (List) Configuration for IBM Event Streams.
	Nested schema for **ibm_event_streams**:
		* `brokers` - (String) The brokers of the IBM Event Streams.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `topic` - (String) The topic of the IBM Event Streams.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `id` - (Integer) The ID of the Event stream.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `is_active` - (Boolean) Whether the Event stream is active.
	* `name` - (String) The name of the Event stream.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `updated_at` - (String) The update time of the Event stream.

