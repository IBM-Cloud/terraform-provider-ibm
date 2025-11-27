---
layout: "ibm"
page_title: "IBM : ibm_iam_role_assignment"
description: |-
  Manages iam_role_assignment.
subcategory: "IAM Policy Management"
---

# ibm_iam_role_assignment

Create, update, and delete iam_role_assignments with this resource.

## Example Usage

```hcl
resource "ibm_iam_role_assignment" "iam_role_assignment_instance" {
  target {
		type = "Account"
		id = "id"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).
  * Constraints: The default value is `default`. The minimum length is `1` character.
* `templates` - (Required, List) The set of properties required for a Role Template assignment.
Nested schema for **templates**:
	* `id` - (Required, String) ID of the template.
		* Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `version` - (Required, String) template version .
		* Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `target` - (Required, List) assignment target account and type.
Nested schema for **target**:
	* `id` - (Required, String) ID of the target account.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
	* `type` - (Required, String) Assignment target type.
	  * Constraints: Allowable values are: `Account`. The maximum length is `30` characters. The minimum length is `1` character.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_role_assignment.
* `account_id` - (String) The account GUID that the role assignments belong to.
* `created_at` - (String) The UTC timestamp when the role assignment was created.
* `created_by_id` - (String) The IAM ID of the entity that created the role assignment.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
* `href` - (String) The href URL that links to the role assignments API by role assignment ID.
* `last_modified_at` - (String) The UTC timestamp when the role assignment was last modified.
* `last_modified_by_id` - (String) The IAM ID of the entity that last modified the role assignment.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
* `operation` - (String) The current operation of the role assignment.
  * Constraints: Allowable values are: `create`, `apply`, `update`, `remove`.
* `resources` - (List) Resources created when role template is assigned.
  * Constraints: The minimum length is `1` item.
Nested schema for **resources**:
	* `role` - (List) Set of properties of the assigned resource or error message if assignment failed.
	Nested schema for **role**:
		* `error_message` - (List) Body parameters for assignment error.
		Nested schema for **error_message**:
			* `code` - (String) Internal status code for the error.
			* `error_code` - (String) error code.
			* `errors` - (List) The errors encountered during the response.
			  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
			Nested schema for **errors**:
				* `code` - (String) The API error code for the error.
				  * Constraints: Allowable values are: `insufficent_permissions`, `invalid_body`, `invalid_token`, `missing_required_query_parameter`, `not_found`, `policy_conflict_error`, `policy_not_found`, `request_not_processed`, `role_conflict_error`, `role_not_found`, `too_many_requests`, `unable_to_process`, `unsupported_content_type`, `policy_template_conflict_error`, `policy_template_not_found`, `policy_assignment_not_found`, `policy_assignment_conflict_error`, `resource_not_found`, `action_control_template_not_found`, `action_control_assignment_not_found`, `role_template_conflict_error`, `role_template_not_found`, `role_assignment_not_found`.
				* `details` - (List) Additional error details.
				Nested schema for **details**:
					* `conflicts_with` - (List) Details of conflicting resource.
					Nested schema for **conflicts_with**:
						* `etag` - (String) The revision number of the resource.
						* `policy` - (String) The conflicting policy ID.
						* `role` - (String) The conflicting role of ID.
				* `message` - (String) The error message returned by the API.
				* `more_info` - (String) Additional info for error.
			* `message` - (String) Error message detailing the nature of the error.
			* `name` - (String) Name of the error.
		* `resource_created` - (List) On success, it includes the role assigned.
		Nested schema for **resource_created**:
			* `id` - (String) role id.
	* `target` - (List) assignment target account and type.
	Nested schema for **target**:
		* `id` - (String) ID of the target account.
		  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
		* `type` - (String) Assignment target type.
		  * Constraints: Allowable values are: `Account`. The maximum length is `30` characters. The minimum length is `1` character.
* `status` - (String) The role assignment status.
  * Constraints: Allowable values are: `accepted`, `failure`, `in_progress`, `superseded`.
* `template` - (List) The role template id and version that will be assigned.
Nested schema for **template**:
	* `id` - (String) Action control template ID.
	  * Constraints: The maximum length is `58` characters. The minimum length is `1` character. The value must match regular expression `/^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `version` - (String) Action control template version.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

* `etag` - ETag identifier for iam_role_assignment.

## Import

You can import the `ibm_iam_role_assignment` resource by using `id`. Role template assignment ID.

# Syntax
<pre>
$ terraform import ibm_iam_role_assignment.iam_role_assignment &lt;id&gt;
</pre>
