---
layout: "ibm"
page_title: "IBM : ibm_logs_stream"
description: |-
  Get information about logs_stream
subcategory: "Cloud Logs"
---

# ibm_logs_stream

Provides a read-only data source to retrieve information about a logs_stream. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_stream" "logs_stream" {
	instance_id 	= ibm_logs_stream.logs_stream_instance.instance_id
	region     		= ibm_logs_stream.logs_stream_instance.region
	logs_streams_id = ibm_logs_stream.logs_stream_instance.streams_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_streams_id` - (Required, String)  Streams ID.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_stream.
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
* `is_active` - (Boolean) Whether the Event stream is active.
* `name` - (String) The name of the Event stream.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `updated_at` - (String) The update time of the Event stream.

