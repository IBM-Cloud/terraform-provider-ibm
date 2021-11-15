---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scan_summaries"
description: |-
  Get information about scan_summaries
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scan_summaries

Review information of Security and Compliance Center scan summaries see [viewing evaluation results](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-results).

## Example usage

```terraform
data "ibm_scc_posture_scan_summaries" "scan_summaries" {
	profile_id = "profile_id"
	scan_id = "262"
	scope_id = "scope_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, String) The profile ID. This can be obtained from the Security and Compliance Center console by clicking on the profile name. The URL contains the ID.
* `scan_id` - (Optional, String) The scan ID of the scan.
* `scope_id` - (Required, String) The scope ID. This can be obtained from the Security and Compliance Center console by clicking on the scope name. The URL contains the ID.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scan_summaries.
* `first` - (Optional, List) The URL of the first scan summary.
Nested scheme for **first**:
	* `href` - (Optional, String) The URL of the first scan summary.
* `last` - (Optional, List) The URL of the last scan summary.
Nested scheme for **last**:
	* `href` - (Optional, String) The URL of the last scan summary.
* `previous` - (Optional, List) The URL of the previous scan summary.
Nested scheme for **previous**:
	* `href` - (Optional, String) The URL of the previous scan summary.
* `summaries` - (Optional, List) Summaries.
Nested scheme for **summaries**:
	* `scan_id` - (Optional, String) The ID of the scan.
	* `scan_name` - (Optional, String) A system generated name that is the combination of 12 characters in the scope name and 12 characters of a profile name.
	* `scope_id` - (Optional, String) The ID of the scan.
	* `scope_name` - (Optional, String) The name of the scope.
	* `report_run_by` - (Optional, String) The entity that ran the report.
	* `start_time` - (Optional, String) The date and time the scan was run.
	* `end_time` - (Optional, String) The date and time the scan completed.
	* `status` - (Optional, String) The status of the collector as it completes a scan. 
	  * Constraints:
		* Supported values are: **pending**, **discovery_started**, **discovery_completed**, **error_in_discovery**, **gateway_aborted**, **controller_aborted**, **not_accepted**, **waiting_for_refine**, **validation_started**, **validation_completed**, **sent_to_collector**, **discovery_in_progress**, **validation_in_progress**, **error_in_validation**, **discovery_result_posted_with_error**, **discovery_result_posted_no_error**, **validation_result_posted_with_error**, **validation_result_posted_no_error**, **fact_collection_started**, **fact_collection_in_progress**, **fact_collection_completed**, **error_in_fact_collection**, **fact_validation_started**, **fact_validation_in_progress**, **fact_validation_completed**, **error_in_fact_validation**, **abort_task_request_received**, **error_in_abort_task_request**, **abort_task_request_completed**, **user_aborted**, **abort_task_request_failed**, **remediation_started**, **remediation_in_progress**, **error_in_remediation**, **remediation_completed**, **inventory_started**, **inventory_in_progress**, **inventory_completed**, **error_in_inventory**, **inventory_completed_with_error**
	* `profile` - (Optional, List) The result of a profile.
	Nested scheme for **profile**:
		* `profile_id` - (Optional, String) The ID of the profile.
		* `profile_name` - (Optional, String) The name of the profile.
		* `profile_type` - (Optional, String) The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
		  * Constraints: 
			* Supported values are: **standard**, **authored**, **custom**, **standard_cv**, **templategroup**, **standard_certificate**
		* `validation_result` - (Optional, List) The result of a scan.
		Nested scheme for **validation_result**:
			* `goals_pass_count` - (Optional, Integer) The number of goals that passed the scan.
			* `goals_unable_to_perform_count` - (Optional, Integer) The number of goals that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
			* `goals_not_applicable_count` - (Optional, Integer) The number of goals that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `goals_fail_count` - (Optional, Integer) The number of goals that failed the scan.
			* `goals_total_count` - (Optional, Integer) The total number of goals that were included in the scan.
			* `controls_pass_count` - (Optional, Integer) The number of controls that passed the scan.
			* `controls_fail_count` - (Optional, Integer) The number of controls that failed the scan.
			* `controls_not_applicable_count` - (Optional, Integer) The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `controls_unable_to_perform_count` - (Optional, Integer) The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
			* `controls_total_count` - (Optional, Integer) The total number of controls that were included in the scan.
	* `group_profiles` - (Optional, List) The result of a group profile.
	Nested scheme for **group_profiles**:
		* `group_profile_id` - (Optional, String) The group ID of profile.
		* `group_profile_name` - (Optional, String) The group name of the profile.
		* `profile_type` - (Optional, String) The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
		  * Constraints:
			* Supported values are **standard**, **authored**, **custom**, **standard_cv**, **templategroup**, **standard_certificate**
		* `validation_result` - (Optional, List) The result of a scan.
		Nested scheme for **validation_result**:
			* `goals_pass_count` - (Optional, Integer) The number of goals that passed the scan.
			* `goals_unable_to_perform_count` - (Optional, Integer) The number of goals that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
			* `goals_not_applicable_count` - (Optional, Integer) The number of goals that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `goals_fail_count` - (Optional, Integer) The number of goals that failed the scan.
			* `goals_total_count` - (Optional, Integer) The total number of goals that were included in the scan.
			* `controls_pass_count` - (Optional, Integer) The number of controls that passed the scan.
			* `controls_fail_count` - (Optional, Integer) The number of controls that failed the scan.
			* `controls_not_applicable_count` - (Optional, Integer) The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `controls_unable_to_perform_count` - (Optional, Integer) The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
			* `controls_total_count` - (Optional, Integer) The total number of controls that were included in the scan.
		* `profiles` - (Optional, List) The result of a each profile in group profile.
		Nested scheme for **profiles**:
			* `profile_id` - (Optional, String) The ID of the profile.
			* `profile_name` - (Optional, String) The name of the profile.
			* `profile_type` - (Optional, String) The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
			  * Constraints: 
				* Supported values are: **standard**, **authored**, **custom**, **standard_cv**, **templategroup**, **standard_certificate**
			* `validation_result` - (Optional, List) The result of a scan.
			Nested scheme for **validation_result**:
				* `controls_pass_count` - (Optional, Integer) The number of controls that passed the scan.
				* `controls_fail_count` - (Optional, Integer) The number of controls that failed the scan.
				* `controls_not_applicable_count` - (Optional, Integer) The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
				* `controls_unable_to_perform_count` - (Optional, Integer) The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
				* `controls_total_count` - (Optional, Integer) The total number of controls that were included in the scan.

