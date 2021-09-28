---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_latest_scans"
description: |-
  Get information about list_latest_scans
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_latest_scans

Review information about the security and compliance center posture latest scans. For more information, about latest scans, see [viewing evaluation results](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-results).

## Example usage

```terraform
data "ibm_scc_posture_latest_scans" "list_latest_scans" {
	scan_id = "262"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `scan_id` - (Optional, String) The ID of the scan.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the `list_latest_scans`.
* `first` - (Optional, List) The URL of the first page of scans.
Nested scheme for **first**:
	* `href` - (Optional, String) The URL of the first page of scans.

* `last` - (Optional, List) The URL of the last page of scans.
Nested scheme for **last**:
	* `href` - (Optional, String) The URL of the last page of scans.

* `latest_scans` - (Optional, List) The details of a scan.
Nested scheme for **latest_scans**:
	* `scan_id` - (Optional, String) The ID of the scan.
	* `scan_name` - (Optional, String) A system generated name that is the combination of 12 characters in the scope name and 12 characters of a profile name.
	* `scope_id` - (Optional, String) The ID of the scan.
	* `scope_name` - (Optional, String) The name of the scope.
	* `profile_id` - (Optional, String) The ID of the profile.
	* `profile_name` - (Optional, String) The name of the profile.
	* `profile_type` - (Optional, String) The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).
	  * Constraints: Allowable values are: **standard**, **authored**, **custom**, **standard_cv**, **temmplategroup**, **standard_certificate**
	* `group_profile_id` - (Optional, String) The group ID of profile.
	* `group_profile_name` - (Optional, String) The group name of the profile.
	* `report_run_by` - (Optional, String) The entity that ran the report.
	* `start_time` - (Optional, String) The date and time the scan was run.
	* `end_time` - (Optional, String) The date and time the scan completed.
	* `result` - (Optional, List) The result of a scan.
	Nested scheme for **result**:
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

* `previous` - (Optional, List) The URL of the previous page of scans.
Nested scheme for **previous**:
	* `href` - (Optional, String) The URL of the previous page of scans.

