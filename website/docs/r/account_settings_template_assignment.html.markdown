---
layout: "ibm"
page_title: "IBM : ibm_iam_account_settings_template_assignment"
description: |-
  Manages account_settings_template_assignment.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_account_settings_template_assignment

Create, update, and delete account_settings_template_assignments with this resource.

## Example Usage

```hcl
resource "ibm_iam_account_settings_template_assignment" "account_settings_template_assignment_instance" {
	template_id = "${var.template_id}"
	template_version = "${var.template_version}"
	target = "${var.target_account}"
	target_type = "${var.account_type}"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `target` - (Required, String) Assignment target.
* `target_type` - (Required, String) Assignment target type.
* `template_id` - (Required, String) Template Id.
* `template_version` - (Required, Integer) Template version.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the account_settings_template_assignment.
* `template_id` - (String) Template id.
* `template_version` - (Integer) Template version.
* `account_id` - (String) Enterprise account id.
* `status` - (String) Assignment status.
* `target` - (String) Assignment target.
* `target_type` - (String) Assignment target type.
* `resources` - (List) Status breakdown per target account of IAM resources created or errors encountered in attempting to create those IAM resources. IAM resources are only included in the response providing the assignment is not in progress. IAM resources are also only included when getting a single assignment, and excluded by list APIs.
  Nested schema for **resources**:
	* `account_settings` - (List)
	  Nested schema for **account_settings**:
		* `error_message` - (List) Body parameters for assignment error.
		  Nested schema for **error_message**:
			* `error_code` - (String) Internal error code.
			* `message` - (String) Error message detailing the nature of the error.
			* `name` - (String) Name of the error.
			* `status_code` - (String) Internal status code for the error.
		* `resource_created` - (List) Body parameters for created resource.
		  Nested schema for **resource_created**:
			* `id` - (String) id of the created resource.
		* `status` - (String) Status for the target account's assignment.
	* `target` - (String) Target account where the IAM resource is created.
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
* `last_modified_at` - (String) Assignment modified at.
* `last_modified_by_id` - (String) IAMid of the identity that last modified the assignment.
* `entity_tag` - (String) Entity tag for this assignment record.
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

## Import

You can import the `ibm_iam_account_settings_template_assignment` resource by using `id`. Assignment record Id.

### Syntax

```bash
$ terraform import ibm_iam_account_settings_template_assignment.account_settings_template_assignment_instance $id
```
