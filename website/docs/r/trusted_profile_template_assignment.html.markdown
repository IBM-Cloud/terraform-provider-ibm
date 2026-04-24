---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_template_assignment"
description: |-
  Manages trusted_profile_template_assignment.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_template_assignment

Create, update, and delete trusted_profile_template_assignments with this resource.

## Example Usage

```hcl
resource "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
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
* `template_id` - (Required, String) Template id.
* `template_version` - (Required, Integer) Template version.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the trusted_profile_template_assignment.
* `account_id` - (String) Enterprise account Id.
* `created_at` - (String) Assignment created at.
* `created_by_id` - (String) IAMid of the identity that created the assignment.
* `entity_tag` - (String) Entity tag for this assignment record.
* `href` - (String) Href.
* `last_modified_at` - (String) Assignment modified at.
* `last_modified_by_id` - (String) IAMid of the identity that last modified the assignment.
* `resources` - (List) Status breakdown per target account of IAM resources created or errors encountered in attempting to create those IAM resources. IAM resources are only included in the response providing the assignment is not in progress. IAM resources are also only included when getting a single assignment, and excluded by list APIs.
Nested schema for **resources**:
	* `policy_template_references` - (List) Policy resource(s) included only for trusted profile assignments with policy references.
	Nested schema for **policy_template_references**:
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
* `status` - (String) Assignment status.

## Import

You can import the `ibm_iam_trusted_profile_template_assignment` resource by using `id`. Assignment record Id.

### Syntax

```bash
$ terraform import ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance $id
```
