---
layout: "ibm"
page_title: "IBM : ibm_iam_policy_assignment"
description: |-
  Get information about policy_assignments
subcategory: "IAM Policy Management"
---

# ibm_iam_policy_assignment

Provides a read-only data source to retrieve information about policy_assignments. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_policy_assignments" "policy_assignment" {
	account_id = "account_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).
  * Constraints: The default value is `default`. The minimum length is `1` character.
* `account_id` - (Required, String) The account GUID in which the policies belong to.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the policy_assignment.
* `policy_assignments` - (List) List of policy assignments.
  * Constraints: The minimum length is `0` items.
Nested schema for **policy_assignments**:
	* `assignment_id` - (String) Passed in value to correlate with other assignments.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character.
	* `created_at` - (String) The UTC timestamp when the policy assignment was created.
	* `created_by_id` - (String) The iam ID of the entity that created the policy assignment.
	* `href` - (String) The href URL that links to the policies assignments API by policy assignment ID.
	* `id` - (String) Policy assignment ID.
	* `last_modified_at` - (String) The UTC timestamp when the policy assignment was last modified.
	* `last_modified_by_id` - (String) The iam ID of the entity that last modified the policy assignment.
	* `resources` - (List) Object for each account assigned.
	  * Constraints: The minimum length is `1` item.
	Nested schema for **resources**:
		* `policy` - (List) Set of properties for the assigned resource.
		Nested schema for **policy**:
			* `error_message` - (List) The error response from API.
			Nested schema for **error_message**:
				* `errors` - (List) The errors encountered during the response.
				  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
				Nested schema for **errors**:
					* `code` - (String) The API error code for the error.
					  * Constraints: Allowable values are: `insufficent_permissions`, `invalid_body`, `invalid_token`, `missing_required_query_parameter`, `not_found`, `policy_conflict_error`, `policy_not_found`, `request_not_processed`, `role_conflict_error`, `role_not_found`, `too_many_requests`, `unable_to_process`, `unsupported_content_type`, `policy_template_conflict_error`, `policy_template_not_found`, `policy_assignment_not_found`, `policy_assignment_conflict_error`.
					* `details` - (List) Additional error details.
					Nested schema for **details**:
						* `conflicts_with` - (List) Details of conflicting resource.
						Nested schema for **conflicts_with**:
							* `etag` - (String) The revision number of the resource.
							* `policy` - (String) The conflicting policy id.
							* `role` - (String) The conflicting role id.
					* `message` - (String) The error message returned by the API.
					* `more_info` - (String) Additional info for error.
				* `status_code` - (Integer) The http error code of the response.
				* `trace` - (String) The unique transaction id for the request.
			* `resource_created` - (List) On success, includes the  policy assigned.
			Nested schema for **resource_created**:
				* `id` - (String) policy id.
			* `status` - (String) The policy assignment status.
			  * Constraints: Allowable values are: `in_progress`, `succeeded`, `succeed_with_errors`, `failed`.
		* `target` - (String) Account ID where resources are assigned.
		  * Constraints: The minimum length is `1` character.
	* `target` - (String) assignment target id.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character.
	* `target_type` - (String) Assignment target type.
	  * Constraints: Allowable values are: `Account`. The maximum length is `30` characters. The minimum length is `1` character.
	* `template_id` - (String) policy template id.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character.
	* `template_version` - (String) policy template version.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character.