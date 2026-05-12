---
layout: "ibm"
page_title: "IBM : ibm_logs_router_target"
description: |-
  Manages logs_router_target.
subcategory: "Logs Routing API Version 3"
---

# ibm_logs_router_target

Create, update, and delete logs_router_targets with this resource.

## Example Usage

```hcl
resource "ibm_logs_router_target" "logs_router_target_instance" {
  destination_crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
  managed_by = "enterprise"
  name = "my-lr-target"
  region = "us-south"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `destination_crn` - (Required, String) Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Logs Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/logs-router?topic=logs-router-target-monitoring&interface=ui#target-monitoring-ui) for details.
  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 \\-._:\/]+$/`.
* `managed_by` - (Optional, String) Present when the target is enterprise-managed (`managed_by: enterprise`). For account-managed targets this field is omitted.
  * Constraints: Allowable values are: `enterprise`, `account`.
* `name` - (Required, String) The name of the target resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character.
* `region` - (Optional, String) Include this optional field if you used it to create a target in a different region other than the one you are connected.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_router_target.
* `created_at` - (String) The timestamp of the target creation time.
* `crn` - (String) The crn of the target resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `3` characters.
* `target_type` - (String) The type of the target.
  * Constraints: Allowable values are: `cloud_logs`.
* `updated_at` - (String) The timestamp of the target last updated time.
* `write_status` - (List) The status of the write attempt to the target with the provided endpoint parameters.
Nested schema for **write_status**:
	* `last_failure` - (String) The timestamp of the failure.
	* `reason_for_last_failure` - (String) Detailed description of the cause of the failure.
	* `status` - (String) The status such as failed or success.
	  * Constraints: Allowable values are: `success`, `failed`.


## Import

You can import the `ibm_logs_router_target` resource by using `id`. The UUID of the target resource.

# Syntax
<pre>
$ terraform import ibm_logs_router_target.logs_router_target &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_logs_router_target.logs_router_target f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6
```
