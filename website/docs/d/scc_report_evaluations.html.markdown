---
layout: "ibm"
page_title: "IBM : ibm_scc_report_evaluations"
description: |-
  Get information about scc_report_evaluations
subcategory: "Results"
---

# ibm_scc_report_evaluations

Provides a read-only data source to retrieve information about scc_report_evaluations. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_report_evaluations" "scc_report_evaluations" {
	report_id = "report_id"
	status = "failure"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `assessment_id` - (Optional, String) The ID of the assessment.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `component_id` - (Optional, String) The ID of component.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\-]+$/`.
* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `status` - (Optional, String) The evaluation status value.
  * Constraints: Allowable values are: `pass`, `failure`, `error`, `skipped`.
* `target_id` - (Optional, String) The ID of the evaluation target.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `target_name` - (Optional, String) The name of the evaluation target.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_evaluations.
* `evaluations` - (List) The list of evaluations that are on the page.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **evaluations**:
	* `assessment` - (List) The control specification assessment.
	Nested schema for **assessment**:
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
	* `component_id` - (String) The component ID.
	* `control_id` - (String) The control ID.
	* `details` - (List) The evaluation details.
	Nested schema for **details**:
		* `properties` - (List) The evaluation properties.
		  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
		Nested schema for **properties**:
			* `expected_value` - (String) The property value.
			* `found_value` - (String) The property value.
			* `operator` - (String) The property operator.
			* `property` - (String) The property name.
			* `property_description` - (String) The property description.
	* `evaluate_time` - (String) The time when the evaluation was made.
	* `home_account_id` - (String) The ID of the home account.
	* `reason` - (String) The reason for the evaluation failure.
	* `report_id` - (String) The ID of the report that is associated to the evaluation.
	* `status` - (String) The allowed values of an evaluation status.
	  * Constraints: Allowable values are: `pass`, `failure`, `error`, `skipped`.
	* `target` - (List) The evaluation target.
	Nested schema for **target**:
		* `account_id` - (String) The target account ID.
		* `id` - (String) The target ID.
		* `resource_crn` - (String) The target resource CRN.
		* `resource_name` - (String) The target resource name.
		* `service_name` - (String) The target service name.

* `first` - (List) The page reference.
Nested schema for **first**:
	* `href` - (String) The URL for the first and next page.

* `home_account_id` - (String) The ID of the home account.

