---
layout: "ibm"
page_title: "IBM : ibm_logs_data_usage_metrics"
description: |-
  Manages logs_data_usage_metrics.
subcategory: "Cloud Logs"
---

# ibm_logs_data_usage_metrics

Create, update, and delete logs_data_usage_metricss with this resource.

## Example Usage

```hcl
resource "ibm_logs_data_usage_metrics" "logs_data_usage_metrics_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  enabled     = true
}
```

## Argument Reference

You can specify the following arguments for this resource.
* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `enabled` - (Required, Boolean) The "enabled" parameter for metrics export.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_data_usage_metrics resource.


## Import

You can import the `ibm_logs_data_usage_metrics` resource by using `id`. Data Usage ID. `id` combination of `region`and `instance_id`

# Syntax
<pre>
$ terraform import ibm_logs_data_usage_metrics.logs_data_usage_metrics < region >/< instance_id >;
</pre>

# Example
```
$ terraform import ibm_logs_data_usage_metrics.logs_data_usage_metrics eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f
```
