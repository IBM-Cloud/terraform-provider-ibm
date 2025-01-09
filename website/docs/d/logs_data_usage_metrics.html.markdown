---
layout: "ibm"
page_title: "IBM : ibm_logs_data_usage_metrics"
description: |-
  Get information about logs_data_usage_metrics
subcategory: "Cloud Logs"
---

# ibm_logs_data_usage_metrics

Provides a read-only data source to retrieve information about logs_data_usage_metrics. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_data_usage_metrics" "logs_data_usage_metrics" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_data_usage_metrics.
* `enabled` - (Boolean) The "enabled" parameter for metrics export.

