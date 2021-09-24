---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scan_summary"
description: |-
  Get information about scans_summary
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scan_summary

Review information of Security and Compliance Center posture scan summary see [viewing evaluation results](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-results).

## Example usage

```terraform
data "ibm_scc_posture_scan_summary" "scans_summary" {
	profile_id = "profile_id"
	scan_id = "scan_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, String) The profile ID. This can be obtained from the Security and Compliance Center console by clicking on the profile name. The URL contains the ID.
* `scan_id` - (Required, Forces new resource, String) The Scan ID.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scans_summary.
* `controls` - (Optional, List) The list of controls on the scan summary.
Nested scheme for **controls**:
	* `control_id` - (Optional, String) The scan summary control ID.
	* `status` - (Optional, String) The control status.
	  * Constraints: Supported values are: **pass**, **unable_to_perform**
	* `external_control_id` - (Optional, String) The external control ID.
	* `control_desciption` - (Optional, String) The scan profile name.
	* `goals` - (Optional, List) The list of goals on the control.
	Nested scheme for **goals**:
		* `goal_description` - (Optional, String) The description of the goal.
		* `goal_id` - (Optional, String) The goal ID.
		* `status` - (Optional, String) The goal status.
		  * Constraints: Supported values are: **pass**, **fail**
		* `severity` - (Optional, String) The severity of the goal.
		* `completed_time` - (Optional, String) The report completed time.
		* `error` - (Optional, String) The error on goal validation.
		* `resource_result` - (Optional, List) The list of resource results.
		Nested scheme for **resource_result**:
			* `resource_name` - (Optional, String) The resource name.
			* `resource_types` - (Optional, String) The resource type.
			* `resource_status` - (Optional, String) The resource control result status.
			  * Constraints: Supported values are: **pass**, **unable_to_perform**
			* `display_expected_value` - (Optional, String) The expected results of a resource.
			* `actual_value` - (Optional, String) The actual results of a resource.
			* `results_info` - (Optional, String) The results information.
			* `not_applicable_reason` - (Optional, String) The reason for goal not applicable for a resource.
		* `goal_applicability_criteria` - (Optional, List) The criteria that defines how a profile applies.
		Nested scheme for **goal_applicability_criteria**:
			* `environment` - (Optional, List) A list of environments that a profile can be applied to.
			* `environment_category` - (Optional, List) The type of environment that a profile is able to be applied to.
			* `resource` - (Optional, List) A list of resources that a profile can be used with.
			* `resource_category` - (Optional, List) The type of resource that a profile is able to be applied to.
			* `resource_type` - (Optional, List) The resource type that the profile applies to.
			* `software_details` - (Optional, List) The software that the profile applies to.
			Nested scheme for **software_details**:
				* `name` - (Optional, String) The name of the Operating System software.
				* `version` - (Optional, String) The version of the Operating System software.
			* `os_details` - (Optional, List) The Operating System that the profile applies to.
			Nested scheme for **os_details**:
				* `name` - (Optional, String) The name of the Operating System.
				* `version` - (Optional, String) The version of the Operating System.
			* `additional_details` - (Optional, Map) Any additional details about the profile.
			* `environment_category_description` - (Optional, Map) The type of environment that your scope is targeted to.
			* `environment_description` - (Optional, Map) The environment that your scope is targeted to.
			* `resource_category_description` - (Optional, Map) The type of resource that your scope is targeted to.
			* `resource_type_description` - (Optional, Map) A further classification of the type of resource that your scope is targeted to.
			* `resource_description` - (Optional, Map) The resource that is scanned as part of your scope.
	* `resource_statistics` - (Optional, List) A scans summary controls.
	Nested scheme for **resource_statistics**:
		* `resource_pass_count` - (Optional, Integer) The resource count of pass controls.
		* `resource_fail_count` - (Optional, Integer) The resource count of fail controls.
		* `resource_unable_to_perform_count` - (Optional, Integer) The number of resources that were unable to be scanned against a control.
		* `resource_not_applicable_count` - (Optional, Integer) The resource count of not applicable(NA) controls.
* `discover_id` - (Optional, String) The scan discovery ID.
* `profile_name` - (Optional, String) The scan profile name.
* `scope_id` - (Optional, String) The scan summary scope ID.

