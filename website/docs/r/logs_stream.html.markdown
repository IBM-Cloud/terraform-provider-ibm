---
layout: "ibm"
page_title: "IBM : ibm_logs_stream"
description: |-
  Manages logs_stream.
subcategory: "Cloud Logs"
---

# ibm_logs_stream

Create, update, and delete logs_streams with this resource.

## Example Usage

```hcl
resource "ibm_logs_stream" "logs_stream_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  compression_type = "gzip"
  is_active = true
  dpxl_expression = "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"
  ibm_event_streams {
		brokers = "kafka01.example.com:9093"
		topic = "live.screen.v2"
  }
  name = "Live Screen"
}
```

## Argument Reference

You can specify the following arguments for this resource.
* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `compression_type` - (Optional, String) The compression type of the stream.
  * Constraints: Allowable values are: `unspecified`, `gzip`.
* `dpxl_expression` - (Required, String) The DPXL expression of the Event stream.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `ibm_event_streams` - (Optional, List) Configuration for IBM Event Streams.
Nested schema for **ibm_event_streams**:
	* `brokers` - (Required, String) The brokers of the IBM Event Streams.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `topic` - (Required, String) The topic of the IBM Event Streams.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `is_active` - (Optional, Boolean) Whether the Event stream is active.
* `name` - (Required, String) The name of the Event stream.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_stream resource.
* `streams_id` - The unique identifier of the logs_stream.
* `created_at` - (String) The creation time of the Event stream.
* `updated_at` - (String) The update time of the Event stream.


## Import

You can import the `ibm_logs_stream` resource by using `id`. `id`. `id` combination of `region`, `instance_id` and `streams_id`.

# Syntax
<pre>
$ terraform import ibm_logs_stream.logs_stream eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/1;
</pre>
