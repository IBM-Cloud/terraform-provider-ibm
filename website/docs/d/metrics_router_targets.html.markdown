---
layout: "ibm"
page_title: "IBM : ibm_metrics_router_targets"
description: |-
  Get information about metrics_router_targets
subcategory: "IBM Cloud Metrics Routing"
---

# ibm_metrics_router_targets

Provides a read-only data source for metrics_router_targets. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_metrics_router_targets" "metrics_router_targets" {
	name = ibm_metrics_router_target.metrics_router_target.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Optional, String) The name of the target resource.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the metrics_router_targets.
* `targets` - (List) A list of target resources.
  * Constraints: The maximum length is `32` items. The minimum length is `0` items.
Nested scheme for **targets**:
	* `created_at` - (String) The timestamp of the target creation time.
	* `crn` - (String) The crn of the target resource.
	* `destination_crn` - (String) The CRN of the destination service instance or resource. Ensure you have a service authorization between IBM Cloud Metrics Routing and your Cloud resource. Read [S2S authorization](https://cloud.ibm.com/docs/metrics-router?topic=metrics-router-target-monitoring&interface=ui#target-monitoring-ui) for details.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:\/]+$/`.
	* `id` - (String) The UUID of the target resource.
	* `name` - (String) The name of the target resource.
	* `region` - (String) Include this optional field if you used it to create a target in a different region other than the one you are connected.
	* `target_type` - (String) The type of the target.
	  * Constraints: Allowable values are: `sysdig_monitor`.
	* `updated_at` - (String) The timestamp of the target last updated time.

