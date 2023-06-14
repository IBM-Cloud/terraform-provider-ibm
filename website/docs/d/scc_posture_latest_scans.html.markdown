---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_latest_scans"
description: |-
  Get information about list_latest_scans
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_latest_scans

Provides a read-only data source for list_latest_scans. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_latest_scans" "list_latest_scans" {
	scan_id = "262"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `scan_id` - (Optional, String) The ID of the scan.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the list_latest_scans.
* `first` - (List) The URL of a page.
Nested scheme for **first**:
	* `href` - (String) The URL of a page.

* `last` - (List) The URL of a page.
Nested scheme for **last**:
	* `href` - (String) The URL of a page.

* `latest_scans` - (List) The details of a scan.
Nested scheme for **latest_scans**:
	* `end_time` - (String) The date and time the scan completed.
	* `group_profile_id` - (String) The group ID of profile.
	* `group_profile_name` - (String) The group name of the profile.
	* `profiles` - (List) Profiles array.
	Nested scheme for **profiles**:
		* `id` - (String) An auto-generated unique identifier for the scope.
		* `name` - (String) The name of the profile.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
		* `type` - (String) The type of profile.
		  * Constraints: Allowable values are: `predefined`, `custom`, `template_group`.
	* `report_run_by` - (String) The entity that ran the report.
	* `report_setting_id` - (String) The unique ID for Scan that is created.
	* `result` - (List) The result of a scan.The above values will not be avaialble if no scopes are available.
	Nested scheme for **result**:
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
	* `scan_id` - (String) The ID of the scan.
	* `scan_name` - (String) A system generated name that is the combination of 12 characters in the scope name and 12 characters of a profile name.
	* `scope_id` - (String) The scope ID of the scan.
	* `scope_name` - (String) The name of the scope.
	* `start_time` - (String) The date and time the scan was run.

* `previous` - (List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (String) The URL of a page.
!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_latest_scans is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
