---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2_autoscale"
description: |-
  Get Information about Autoscale configurations of IBM Db2 instance.
---

# ibm_db2_autoscale

Retrieve information about Autoscale configurations of an existing [IBM Db2 Instance](https://cloud.ibm.com/docs/Db2onCloud).

## Example Usage

```hcl
data "ibm_db2_autoscale" "db2_autoscale" {
    deployment_id = "<encoded_crn>"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, String) Encoded CRN of the instance this autoscale relates to.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.
* `auto_scaling_allow_plan_limit` - (Boolean) Indicates the maximum number of scaling actions that are allowed within a specified time period.
* `auto_scaling_enabled` - (Boolean) Indicates if automatic scaling is enabled or not.
* `auto_scaling_max_storage` - (Integer) The maximum limit for automatically increasing storage capacity to handle growing data needs.
* `auto_scaling_over_time_period` - (Integer) Defines the time period over which auto-scaling adjustments are monitored and applied.
* `auto_scaling_pause_limit` - (Integer) Specifies the duration to pause auto-scaling actions after a scaling event has occurred.
* `auto_scaling_threshold` - (Integer) Specifies the resource utilization level that triggers an auto-scaling.
* `storage_unit` - (String) Specifies the unit of measurement for storage capacity.
* `storage_utilization_percentage` - (Integer) Represents the percentage of total storage capacity currently in use.
* `support_auto_scaling` - (Boolean) Indicates whether a system or service can automatically adjust resources based on demand.