---
layout: "ibm"
page_title: "IBM : ibm_iam_role_assignment"
description: |-
  Get information about role_assignments
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_role_assignment

Provides a read-only data source to retrieve information about role_assignments. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_role_assignment" "role_assignment" {
	assignment_id = "role_assignment_id"
}
```

## Argument Reference
You can specify the following arguments for this data source.

* `assignment_id` - (Required, String) The role assignment ID.

## Timeouts section

The resource includes default timeout settings for the following operations:

* `create` - (Timeout) Defaults to 30 minutes.
* `update` - (Timeout) Defaults to 30 minutes.
* `delete` - (Timeout) Defaults to 30 minutes.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the role_assignment.
* `assignment_id` - (String) Passed in value to correlate with other assignments.
	* Constraints: The maximum length is `50` characters. The minimum length is `1` character.
* `created_at` - (String) The UTC timestamp when the role assignment was created.
* `created_by_id` - (String) The iam ID of the entity that created the role assignment.
* `href` - (String) The href URL that links to the policies assignments API by role assignment ID.
* `id` - (String) role assignment ID.
* `last_modified_at` - (String) The UTC timestamp when the role assignment was last modified.
* `last_modified_by_id` - (String) The iam ID of the entity that last modified the role assignment.
* `resources` - (List) Object for each account assigned.
	* Constraints: The minimum length is `1` item.
Nested schema for **resources**:
	* `role` - (List) Set of properties for the assigned resource.
	Nested schema for **role**:
		* `error_message` - (List) The error response from API.
		Nested schema for **error_message**:
			* `errors` - (List) The errors encountered during the response.
			Nested schema for **errors**:
				* `code` - (String) The API error code for the error.
				* `details` - (List) Additional error details.
				Nested schema for **details**:
					Nested schema for **conflicts_with**:
						* `etag` - (String) The revision number of the resource.
						* `role` - (String) The conflicting role id.
				* `message` - (String) The error message returned by the API.
				* `more_info` - (String) Additional info for error.
			* `status_code` - (Integer) The http error code of the response.
			* `trace` - (String) The unique transaction id for the request.
			* `name` - (String) Name of the error.
			* `errorCode` - (String) error code.
			* `message` - (String) Error message detailing the nature of the error.
			* `code` - (String) error code.
		* `resource_created` - (List) On success, includes the  role assigned.
		Nested schema for **resource_created**:
			* `id` - (String) role id.
		* `status` - (String) The role assignment status.
	* `target` - (String) Account ID where resources are assigned.
* `target` - (Map) assignment target details.
Nested schema for **target**:
	* `id` - (String) The role assignment target account id.
	* `type` - (String) The target type.
* `template` - (Map) template details
Nested schema for **template**:
	* `id` - (String) The role assignment template id.
	* `version` - (String) The orchestrator template version.
* `account_id` - (String) Enterprise account ID where template will be created.
