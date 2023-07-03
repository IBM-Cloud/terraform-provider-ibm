---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scan_summaries"
description: |-
  Get information about scan_summaries
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scan_summaries

Provides a read-only data source for scan_summaries. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_scan_summaries" "scan_summaries" {
	report_setting_id = "report_setting_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `report_setting_id` - (Required, String) The report setting ID. This can be obtained from the /validations/latest_scans API call.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scan_summaries.
* `first` - (List) The URL of a page.
Nested scheme for **first**:
	* `href` - (String) The URL of a page.

* `last` - (List) The URL of a page.
Nested scheme for **last**:
	* `href` - (String) The URL of a page.

* `previous` - (List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (String) The URL of a page.

* `summaries` - (List) Summaries.
Nested scheme for **summaries**:
	* `end_time` - (String) The date and time the scan completed.
	* `group_profiles` - (List) The list of group profiles.
	Nested scheme for **group_profiles**:
		* `id` - (String) The ID of the profile.
		* `name` - (String) The name of the profile.
		* `type` - (String) The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
		  * Constraints: Allowable values are: `standard`, `authored`, `custom`, `standard_cv`, `temmplategroup`, `standard_certificate`, `predefined`.
		* `validation_result` - (List) The result of a scan.The above values will not be avaialble if no scopes are available.
		Nested scheme for **validation_result**:
			* `controls_fail_count` - (Integer) The number of controls that failed the scan.
			* `controls_not_applicable_count` - (Integer) The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `controls_pass_count` - (Integer) The number of controls that passed the scan.
			* `controls_total_count` - (Integer) The total number of controls that were included in the scan.
			* `controls_unable_to_perform_count` - (Integer) The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
			* `goals_fail_count` - (Integer) The number of goals that failed the scan.
			* `goals_not_applicable_count` - (Integer) The number of goals that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `goals_pass_count` - (Integer) The number of goals that passed the scan.
			* `goals_total_count` - (Integer) The total number of goals that were included in the scan.
			* `goals_unable_to_perform_count` - (Integer) The number of goals that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
	* `id` - (String) The ID of the scan.
	* `name` - (String) A system generated name that is the combination of 12 characters in the scope name and 12 characters of a profile name.
	* `profiles` - (List) The list of profiles.
	Nested scheme for **profiles**:
		* `id` - (String) The ID of the profile.
		* `name` - (String) The name of the profile.
		* `type` - (String) The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
		  * Constraints: Allowable values are: `standard`, `authored`, `custom`, `standard_cv`, `temmplategroup`, `standard_certificate`, `predefined`.
		* `validation_result` - (List) The result of a scan.The above values will not be avaialble if no scopes are available.
		Nested scheme for **validation_result**:
			* `controls_fail_count` - (Integer) The number of controls that failed the scan.
			* `controls_not_applicable_count` - (Integer) The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `controls_pass_count` - (Integer) The number of controls that passed the scan.
			* `controls_total_count` - (Integer) The total number of controls that were included in the scan.
			* `controls_unable_to_perform_count` - (Integer) The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
			* `goals_fail_count` - (Integer) The number of goals that failed the scan.
			* `goals_not_applicable_count` - (Integer) The number of goals that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.
			* `goals_pass_count` - (Integer) The number of goals that passed the scan.
			* `goals_total_count` - (Integer) The total number of goals that were included in the scan.
			* `goals_unable_to_perform_count` - (Integer) The number of goals that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.
	* `report_run_by` - (String) The entity that ran the report.
	* `scope_id` - (String) The ID of the scope.
	* `scope_name` - (String) The name of the scope.
	* `start_time` - (String) The date and time the scan was run.
	* `status` - (String) The status of the collector as it completes a scan.
	  * Constraints: Allowable values are: `pending`, `discovery_started`, `discovery_completed`, `error_in_discovery`, `gateway_aborted`, `controller_aborted`, `not_accepted`, `waiting_for_refine`, `validation_started`, `validation_completed`, `sent_to_collector`, `discovery_in_progress`, `validation_in_progress`, `error_in_validation`, `discovery_result_posted_with_error`, `discovery_result_posted_no_error`, `validation_result_posted_with_error`, `validation_result_posted_no_error`, `fact_collection_started`, `fact_collection_in_progress`, `fact_collection_completed`, `error_in_fact_collection`, `fact_validation_started`, `fact_validation_in_progress`, `fact_validation_completed`, `error_in_fact_validation`, `abort_task_request_received`, `error_in_abort_task_request`, `abort_task_request_completed`, `user_aborted`, `abort_task_request_failed`, `remediation_started`, `remediation_in_progress`, `error_in_remediation`, `remediation_completed`, `inventory_started`, `inventory_in_progress`, `inventory_completed`, `error_in_inventory`, `inventory_completed_with_error`.

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scan_summaries is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
