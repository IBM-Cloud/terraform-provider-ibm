---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scan_summary"
description: |-
  Get information about scans_summary
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scan_summary

Provides a read-only data source for scans_summary. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_scan_summary" "scans_summary" {
	profile_id = "profile_id"
	scan_id = "scan_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, String) The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `scan_id` - (Required, String) Your Scan ID.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scans_summary.
* `controls` - (List) The list of controls on the scan summary.
Nested scheme for **controls**:
	* `desciption` - (String) The scan profile name.
	* `external_control_id` - (String) The external control ID.
	* `goals` - (List) The list of goals on the control.
	Nested scheme for **goals**:
		* `completed_time` - (String) The report completed time.
		* `description` - (String) The description of the goal.
		* `error` - (String) The error on goal validation.
		* `id` - (String) The goal ID.
		* `resource_result` - (List) The list of resource results.
		Nested scheme for **resource_result**:
			* `actual_value` - (String) The actual results of a resource.
			* `display_expected_value` - (String) The expected results of a resource.
			* `name` - (String) The resource name.
			* `not_applicable_reason` - (String) The reason for goal not applicable for a resource.
			* `results_info` - (String) The results information.
			* `status` - (String) The resource control result status.
			  * Constraints: Allowable values are: `pass`, `unable_to_perform`.
			* `types` - (String) The resource type.
		* `severity` - (String) The severity of the goal.
		* `status` - (String) The goal status.
		  * Constraints: Allowable values are: `pass`, `fail`.
	* `id` - (String) The scan summary control ID.
	* `resource_statistics` - (List) A scans summary controls.
	Nested scheme for **resource_statistics**:
		* `fail_count` - (Integer) The resource count of fail controls.
		* `not_applicable_count` - (Integer) The resource count of not applicable(na) controls.
		* `pass_count` - (Integer) The resource count of pass controls.
		* `unable_to_perform_count` - (Integer) The number of resources that were unable to be scanned against a control.
	* `status` - (String) The control status.
	  * Constraints: Allowable values are: `pass`, `unable_to_perform`.

* `discover_id` - (String) The scan discovery ID.

* `profile_name` - (String) The scan profile name.

* `scope_id` - (String) The scan summary scope ID.

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scan_summary is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
