---
layout: "ibm"
page_title: "IBM : ibm_logs_router_targets"
description: |-
  Get information about logs_router_targets
subcategory: "Logs Routing API Version 3"
---

# ibm_logs_router_targets

Provides a read-only data source to retrieve information about logs_router_targets. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_router_targets" "logs_router_targets" {
	name = ibm_logs_router_target.logs_router_target_instance.name
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The name of the target resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_router_targets.
* `targets` - (List) A list of target resources.
  * Constraints: The maximum length is `32` items. The minimum length is `0` items.
Nested schema for **targets**:
	* `created_at` - (String) The timestamp of the target creation time.
	* `crn` - (String) The crn of the target resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters.
	* `destination_crn` - (String) Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Logs Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/logs-router?topic=logs-router-target-monitoring&interface=ui#target-monitoring-ui) for details.
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
	  * Constraints: Allowable values are: `cloud_logs`.
	* `updated_at` - (String) The timestamp of the target last updated time.
	* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
	Nested schema for **write_status**:
		* `last_failure` - (String) The timestamp of the failure.
		* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
		* `status` - (String) The status such as failed or success.
		  * Constraints: Allowable values are: `success`, `failed`.

