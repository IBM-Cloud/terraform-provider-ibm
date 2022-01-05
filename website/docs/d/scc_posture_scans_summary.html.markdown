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
* `controls` - (Optional, List) The list of controls on the scan summary.
Nested scheme for **controls**:
	* `desciption` - (Optional, String) The scan profile name.
	* `external_control_id` - (Optional, String) The external control ID.
	* `goals` - (Optional, List) The list of goals on the control.
	Nested scheme for **goals**:
		* `completed_time` - (Optional, String) The report completed time.
		* `description` - (Optional, String) The description of the goal.
		* `error` - (Optional, String) The error on goal validation.
		* `id` - (Optional, String) The goal ID.
		* `resource_result` - (Optional, List) The list of resource results.
		Nested scheme for **resource_result**:
			* `actual_value` - (Optional, String) The actual results of a resource.
			* `display_expected_value` - (Optional, String) The expected results of a resource.
			* `name` - (Optional, String) The resource name.
			* `not_applicable_reason` - (Optional, String) The reason for goal not applicable for a resource.
			* `results_info` - (Optional, String) The results information.
			* `status` - (Optional, String) The resource control result status.
			  * Constraints: Allowable values are: `pass`, `unable_to_perform`.
			* `types` - (Optional, String) The resource type.
		* `severity` - (Optional, String) The severity of the goal.
		* `status` - (Optional, String) The goal status.
		  * Constraints: Allowable values are: `pass`, `fail`.
	* `id` - (Optional, String) The scan summary control ID.
	* `resource_statistics` - (Optional, List) A scans summary controls.
	Nested scheme for **resource_statistics**:
		* `fail_count` - (Optional, Integer) The resource count of fail controls.
		* `not_applicable_count` - (Optional, Integer) The resource count of not applicable(na) controls.
		* `pass_count` - (Optional, Integer) The resource count of pass controls.
		* `unable_to_perform_count` - (Optional, Integer) The number of resources that were unable to be scanned against a control.
	* `status` - (Optional, String) The control status.
	  * Constraints: Allowable values are: `pass`, `unable_to_perform`.

* `discover_id` - (Optional, String) The scan discovery ID.

* `id` - (Optional, String) The scan ID.

* `profile_name` - (Optional, String) The scan profile name.

* `scope_id` - (Optional, String) The scan summary scope ID.

