---
layout: "ibm"
page_title: "IBM : ibm_scc_report_resources"
description: |-
  Get information about scc_report_resources
subcategory: "Security and Compliance Center"
---

# ibm_scc_report_resources

Retrieve information about report resources from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_report_resources" "scc_report_resources" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    report_id = "report_id"
    status = "compliant"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `account_id` - (Optional, String) The ID of the account owning a resource.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `component_id` - (Optional, String) The ID of component.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\-]+$/`.
* `id` - (Optional, String) The ID of the resource.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `report_id` - (Required, Forces new resource, String) The ID of the scan that is associated with a report.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `resource_name` - (Optional, String) The name of the resource.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-]+$/`.
* `sort` - (Optional, String) This field sorts resources by using a valid sort field. To learn more, see [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
  * Constraints: Allowable values are: `account_id`, `component_id`, `resource_name`, `status`.
* `status` - (Optional, String) The compliance status value.
  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_report_resources.
* `first` - (List) The page reference.
Nested schema for **first**:
	* `href` - (String) The URL for the first and next page.

* `home_account_id` - (String) The ID of the home account.

* `resources` - (List) The list of resource evaluation summaries that are on the page.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **resources**:
	* `account` - (List) The account that is associated with a report.
	Nested schema for **account**:
		* `id` - (String) The account ID.
		* `name` - (String) The account name.
		* `type` - (String) The account type.
	* `completed_count` - (Integer) The total number of completed evaluations.
	* `component_id` - (String) The ID of the component.
	* `environment` - (String) The environment.
	* `error_count` - (Integer) The number of evaluations that started, but did not finish, and ended with errors.
	* `failure_count` - (Integer) The number of failed evaluations.
	* `id` - (String) The resource CRN.
	* `pass_count` - (Integer) The number of passed evaluations.
	* `report_id` - (String) The ID of the report.
	* `resource_name` - (String) The resource name.
	* `status` - (String) The allowed values of an aggregated status for controls, specifications, assessments, and resources.
	  * Constraints: Allowable values are: `compliant`, `not_compliant`, `unable_to_perform`, `user_evaluation_required`.
	* `total_count` - (Integer) The total number of evaluations.

