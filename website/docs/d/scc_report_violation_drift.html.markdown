---
layout: "ibm"
page_title: "IBM : ibm_scc_report_violation_drift"
description: |-
  Get information about scc_report_violation_drift
subcategory: "Security and Compliance Center"
---

# ibm_scc_report_violation_drift

Retrieve information about a report violation drift from a read-only data source. Then, yo can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_report_violation_drift" "scc_report_violation_drift" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    report_id = "report_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `scan_time_duration` - (Optional, Integer) The duration of the `scan_time` timestamp in number of days.
  * Constraints: The default value is `0`. The maximum value is `366`. The minimum value is `0`.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_violation_drift.
* `data_points` - (List) The list of report violations data points.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **data_points**:
	* `controls` - (List) The compliance stats.
	Nested schema for **controls**:
		* `compliant_count` - (Integer) The number of compliant checks.
		* `not_compliant_count` - (Integer) The number of checks that are not compliant.
		* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
		  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
		* `total_count` - (Integer) The total number of checks.
		* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.
		* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.
	* `report_group_id` - (String) The group ID that is associated with the report. The group ID combines profile, scope, and attachment IDs.
	* `report_id` - (String) The ID of the report.
	* `scan_time` - (String) The date when the scan was run.

* `home_account_id` - (String) The ID of the home account.

