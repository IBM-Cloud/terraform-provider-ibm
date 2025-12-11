---
layout: "ibm"
page_title: "IBM : ibm_metrics_router_targets"
description: |-
  Get information about metrics_router_targets
subcategory: "Metrics Routing API Version 3"
---

# ibm_metrics_router_targets

Provides a read-only data source to retrieve information about metrics_router_targets. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_metrics_router_targets" "metrics_router_targets" {
	name = ibm_metrics_router_target.metrics_router_target_instance.name
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The name of the target resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the metrics_router_targets.
* `targets` - (List) A list of target resources.
  * Constraints: The maximum length is `32` items. The minimum length is `0` items.
Nested schema for **targets**:
	* `created_at` - (String) The timestamp of the target creation time.
	* `crn` - (String) The crn of the target resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters.
	* `destination_crn` - (String) Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Metrics Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/metrics-router?topic=metrics-router-target-monitoring&interface=ui#target-monitoring-ui) for details.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:\/]+$/`.
	* `id` - (String) The UUID of the target resource.
	  * Constraints: The maximum length is `1028` characters. The minimum length is `24` characters.
	* `managed_by` - (String) Present when the target is enterprise-managed (`managed_by: enterprise`). For account-managed targets this field is omitted.
	  * Constraints: Allowable values are: `enterprise`, `account`.
	* `name` - (String) The name of the target resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
	* `region` - (String) Include this optional field if you used it to create a target in a different region other than the one you are connected.
	  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters.
	* `target_type` - (String) The type of the target.
	  * Constraints: Allowable values are: `sysdig_monitor`.
	* `updated_at` - (String) The timestamp of the target last updated time.

