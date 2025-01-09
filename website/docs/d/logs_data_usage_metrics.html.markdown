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
  query = "daily"
	range = "last_week"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `query` - (Optional, String) Query to filter daily or detailed the data usage, by default it will use daily one.
  * Constraints: The default value is `daily`. Allowable values are: `daily`, `detailed`.
* `range` - (Optional, String) Range of days to get the data usage for, by default it will use current month.
  * Constraints: The default value is `current_month`. Allowable values are: `current_month`, `last_30_days`, `last_90_days`, `last_week`.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_data_usage_metrics.
* `enabled` - (Boolean) The "enabled" parameter for metrics export.

