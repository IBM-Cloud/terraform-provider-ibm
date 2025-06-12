---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_template_assignment"
description: |-
  Get information about trusted_profile_template_assignment
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_template_assignment

Provides a read-only data source to retrieve information about a trusted_profile_template_assignment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment" {
	assignment_id = "${var.assignment_id}"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `assignment_id` - (Required, String) ID of the Assignment Record.
* `include_history` - (Optional, Boolean) Defines if the entity history is included in the response.
  * Constraints: The default value is `false`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the trusted_profile_template_assignment.
* `target` - (String) Assignment target.
* `target_type` - (String) Assignment target type.
* `template_id` - (String) Template Id.
* `template_version` - (Integer) Template version.
* `account_id` - (String) Enterprise account Id.
* `status` - (String) Assignment status.
* `history` - (List) Assignment history.
Nested schema for **history**:
	* `action` - (String) Action of the history entry.
	* `iam_id` - (String) IAM ID of the identity which triggered the action.
	* `iam_id_account` - (String) Account of the identity which triggered the action.
	* `message` - (String) Message which summarizes the executed action.
	* `params` - (List) Params of the history entry.
	* `timestamp` - (String) Timestamp when the action was triggered.
* `href` - (String) Href.
* `created_at` - (String) Assignment created at.
* `created_by_id` - (String) IAMid of the identity that created the assignment.
* `entity_tag` - (String) Entity tag for this assignment record.
* `last_modified_at` - (String) Assignment modified at.
* `last_modified_by_id` - (String) IAMid of the identity that last modified the assignment.

* `resources` - (List) Status breakdown per target account of IAM resources created or errors encountered in attempting to create those IAM resources. IAM resources are only included in the response providing the assignment is not in progress. IAM resources are also only included when getting a single assignment, and excluded by list APIs.
Nested schema for **resources**:
	* `policy_template_refs` - (List) Policy resource(s) included only for trusted profile assignments with policy references.
	Nested schema for **policy_template_refs**:
		* `error_message` - (List) Body parameters for assignment error.
		Nested schema for **error_message**:
			* `error_code` - (String) Internal error code.
			* `message` - (String) Error message detailing the nature of the error.
			* `name` - (String) Name of the error.
			* `status_code` - (String) Internal status code for the error.
		* `id` - (String) Policy Template Id, only returned for a profile assignment with policy references.
		* `resource_created` - (List) Body parameters for created resource.
		Nested schema for **resource_created**:
			* `id` - (String) Id of the created resource.
		* `status` - (String) Status for the target account's assignment.
		* `version` - (String) Policy version, only returned for a profile assignment with policy references.
	* `profile` - (List)
	Nested schema for **profile**:
		* `error_message` - (List) Body parameters for assignment error.
		Nested schema for **error_message**:
			* `error_code` - (String) Internal error code.
			* `message` - (String) Error message detailing the nature of the error.
			* `name` - (String) Name of the error.
			* `status_code` - (String) Internal status code for the error.
		* `id` - (String) Policy Template Id, only returned for a profile assignment with policy references.
		* `resource_created` - (List) Body parameters for created resource.
		Nested schema for **resource_created**:
			* `id` - (String) Id of the created resource.
		* `status` - (String) Status for the target account's assignment.
		* `version` - (String) Policy version, only returned for a profile assignment with policy references.
	* `target` - (String) Target account where the IAM resource is created.
* `context` - (List) Context with key properties for problem determination.
  Nested schema for **context**:
	* `cluster_name` - (String) The cluster name.
	* `elapsed_time` - (String) The elapsed time in msec.
	* `end_time` - (String) The finish time of the request.
	* `host` - (String) The host of the server instance processing the request.
	* `instance_id` - (String) The instance ID of the server instance processing the request.
	* `operation` - (String) The operation of the inbound REST request.
	* `start_time` - (String) The start time of the request.
	* `thread_id` - (String) The thread ID of the server instance processing the request.
	* `transaction_id` - (String) The transaction ID of the inbound REST request.
	* `url` - (String) The URL of that cluster.
	* `user_agent` - (String) The user agent of the inbound REST request.