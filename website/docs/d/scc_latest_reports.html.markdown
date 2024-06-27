---
layout: "ibm"
page_title: "IBM : ibm_scc_latest_reports"
description: |-
  Get information about scc_latest_reports
subcategory: "Security and Compliance Center"
---

# ibm_scc_latest_reports

Retrieve information about the latest reports from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_latest_reports" "scc_latest_reports" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    sort = "profile_name"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `sort` - (Optional, String) This field sorts results by using a valid sort field. To learn more, see [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[\\-]?[a-z0-9_]+$/`.
* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_latest_reports.
* `controls_summary` - (List) The compliance stats.
Nested schema for **controls_summary**:
	* `compliant_count` - (Integer) The number of compliant checks.
	* `not_compliant_count` - (Integer) The number of checks that are not compliant.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `total_count` - (Integer) The total number of checks.
	* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.
	* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.

* `evaluations_summary` - (List) The evaluation stats.
Nested schema for **evaluations_summary**:
	* `completed_count` - (Integer) The total number of completed evaluations.
	* `error_count` - (Integer) The number of evaluations that started, but did not finish, and ended with errors.
	* `failure_count` - (Integer) The number of failed evaluations.
	* `pass_count` - (Integer) The number of passed evaluations.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `total_count` - (Integer) The total number of evaluations.

* `home_account_id` - (String) The ID of the home account.

* `reports` - (List) The list of reports.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **reports**:
	* `account` - (List) The account that is associated with a report.
	Nested schema for **account**:
		* `id` - (String) The account ID.
		* `name` - (String) The account name.
		* `type` - (String) The account type.
	* `attachment` - (List) The attachment that is associated with a report.
	Nested schema for **attachment**:
		* `description` - (String) The description of the attachment.
		* `id` - (String) The attachment ID.
		* `name` - (String) The name of the attachment.
		* `schedule` - (String) The attachment schedule.
		* `scope` - (List) The scope of the attachment.
		  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
		Nested schema for **scope**:
			* `environment` - (String) The environment that relates to this scope.
			* `id` - (String) The unique identifier for this scope.
			* `properties` - (List) The properties that are supported for scoping by this environment.
			  * Constraints: The maximum length is `99999` items. The minimum length is `0` items.
			Nested schema for **properties**:
				* `name` - (String) The property name.
				* `value` - (String) The property value.
	* `cos_object` - (String) The Cloud Object Storage object that is associated with the report.
	* `created_on` - (String) The date when the report was created.
	* `group_id` - (String) The group ID that is associated with the report. The group ID combines profile, scope, and attachment IDs.
	* `id` - (String) The ID of the report.
	* `instance_id` - (String) Instance ID.
	* `profile` - (List) The profile information.
	Nested schema for **profile**:
		* `id` - (String) The profile ID.
		* `name` - (String) The profile name.
		* `version` - (String) The profile version.
	* `scan_time` - (String) The date when the scan was run.
	* `type` - (String) The type of the scan.

* `score` - (List) The compliance score.
Nested schema for **score**:
	* `passed` - (Integer) The number of successful evaluations.
	* `percent` - (Integer) The percentage of successful evaluations.
	* `total_count` - (Integer) The total number of evaluations.

