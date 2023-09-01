---
layout: "ibm"
page_title: "IBM : ibm_scc_report_controls"
description: |-
  Get information about scc_report_controls
subcategory: "Results"
---

# ibm_scc_report_controls

Provides a read-only data source to retrieve information about scc_report_controls. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_report_controls" "scc_report_controls" {
	report_id = "report_id"
	status = "compliant"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `control_category` - (Optional, String) A control category value.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `control_description` - (Optional, String) The description of the control.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\s]+$/`.
* `control_id` - (Optional, String) The ID of the control.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `control_name` - (Optional, String) The name of the control.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `sort` - (Optional, String) This field sorts controls by using a valid sort field. To learn more, see [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
  * Constraints: Allowable values are: `control_name`, `control_category`, `status`.
* `status` - (Optional, String) The compliance status value.
  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_controls.
* `compliant_count` - (Integer) The number of compliant checks.

* `controls` - (List) The list of controls that are in the report.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **controls**:
	* `compliant_count` - (Integer) The number of compliant checks.
	* `control_category` - (String) The control category.
	* `control_description` - (String) The control description.
	* `control_library_id` - (String) The control library ID.
	* `control_library_version` - (String) The control library version.
	* `control_name` - (String) The control name.
	* `control_path` - (String) The control path.
	* `control_specifications` - (List) The list of specifications that are on the page.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **control_specifications**:
		* `assessments` - (List) The list of assessments.
		  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
		Nested schema for **assessments**:
			* `assessment_description` - (String) The assessment description.
			* `assessment_id` - (String) The assessment ID.
			* `assessment_method` - (String) The assessment method.
			* `assessment_type` - (String) The assessment type.
			* `parameter_count` - (Integer) The number of parameters of this assessment.
			* `parameters` - (List) The list of parameters of this assessment.
			  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
			Nested schema for **parameters**:
				* `parameter_display_name` - (String) The parameter display name.
				* `parameter_name` - (String) The parameter name.
				* `parameter_type` - (String) The parameter type.
				* `parameter_value` - (String) The property value.
		* `compliant_count` - (Integer) The number of compliant checks.
		* `component_id` - (String) The component ID.
		* `control_specification_description` - (String) The component description.
		* `control_specification_id` - (String) The control specification ID.
		* `environment` - (String) The environment.
		* `not_compliant_count` - (Integer) The number of checks that are not compliant.
		* `responsibility` - (String) The responsibility for managing control specifications.
		* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
		  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
		* `total_count` - (Integer) The total number of checks.
		* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.
		* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.
	* `id` - (String) The control ID.
	* `not_compliant_count` - (Integer) The number of checks that are not compliant.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `total_count` - (Integer) The total number of checks.
	* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.
	* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.

* `home_account_id` - (String) The ID of the home account.

* `not_compliant_count` - (Integer) The number of checks that are not compliant.

* `total_count` - (Integer) The total number of checks.

* `unable_to_perform_count` - (Integer) The number of checks that are unable to perform.

* `user_evaluation_required_count` - (Integer) The number of checks that require a user evaluation.

