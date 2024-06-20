---
layout: "ibm"
page_title: "IBM : ibm_scc_report_summary"
description: |-
  Get information about scc_report_summary
subcategory: "Security and Compliance Center"
---

# ibm_scc_report_summary

Retrieve information about a report summary from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_report_summary" "scc_report_summary" {
    instance_id = "00000000-1111-2222-3333-444444444444"
	report_id = "report_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_summary.
* `account` - (List) The account that is associated with a report.
Nested schema for **account**:
	* `id` - (String) The account ID.
	* `name` - (String) The account name.
	* `type` - (String) The account type.

* `controls` - (List) The compliance stats.
Nested schema for **controls**:
	* `compliant_count` - (Integer) The number of compliant checks.
	* `not_compliant_count` - (Integer) The number of checks that are not compliant.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `total_count` - (Integer) The total number of checks.
	* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.
	* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.

* `evaluations` - (List) The evaluation stats.
Nested schema for **evaluations**:
	* `completed_count` - (Integer) The total number of completed evaluations.
	* `error_count` - (Integer) The number of evaluations that started, but did not finish, and ended with errors.
	* `failure_count` - (Integer) The number of failed evaluations.
	* `pass_count` - (Integer) The number of passed evaluations.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `total_count` - (Integer) The total number of evaluations.

* `isntance_id` - (String) Instance ID.

* `resources` - (List) The resource summary.
Nested schema for **resources**:
	* `compliant_count` - (Integer) The number of compliant checks.
	* `not_compliant_count` - (Integer) The number of checks that are not compliant.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `top_failed` - (List) The top 10 resources that have the most failures.
	  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
	Nested schema for **top_failed**:
		* `account` - (String) The account that owns the resource.
		* `completed_count` - (Integer) The total number of completed evaluations.
		* `error_count` - (Integer) The number of evaluations that started, but did not finish, and ended with errors.
		* `failure_count` - (Integer) The number of failed evaluations.
		* `id` - (String) The resource ID.
		* `name` - (String) The resource name.
		* `pass_count` - (Integer) The number of passed evaluations.
		* `service` - (String) The service that is managing the resource.
		* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
		  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
		* `tags` - (List) The collection of different types of tags.
		Nested schema for **tags**:
			* `access` - (List) The collection of access tags.
			  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
			* `service` - (List) The collection of service tags.
			  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
			* `user` - (List) The collection of user tags.
			  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
		* `total_count` - (Integer) The total number of evaluations.
	* `total_count` - (Integer) The total number of checks.
	* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.
	* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.

* `score` - (List) The compliance score.
Nested schema for **score**:
	* `passed` - (Integer) The number of successful evaluations.
	* `percent` - (Integer) The percentage of successful evaluations.
	* `total_count` - (Integer) The total number of evaluations.

