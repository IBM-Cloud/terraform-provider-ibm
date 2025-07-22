---
layout: "ibm"
page_title: "IBM : ibm_iam_action_control_assignment"
description: |-
  Manages action_control_assignment.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_action_control_assignment

Create, update, and delete action_control_assignments with this resource.

## Example Usage

```hcl
resource "ibm_iam_action_control_template" "action_template" {
	name = "TerraformActionControlTest"
	action_control {
		actions = ["am-test-service.test.delete" ]
		service_name="am-test-service"
		}
	committed=true
}
resource "ibm_iam_action_control_template_version" "template_version" {
	action_control_template_id = ibm_iam_action_control_template.action_template.action_control_template_id
	action_control {
		actions = ["am-test-service.test.delete", "am-test-service.test.create" ]
		service_name="am-test-service"
	}
	committed=true
}

resource "ibm_iam_action_control_assignment" "action_control_assignment" {
	target  ={
		type = "Account" or "Account Group" or "Enterprise"
		id = "<target-accountId>"
	}
	
	templates{
		id = ibm_iam_action_control_template.action_template.action_control_template_id
		version = ibm_iam_action_control_template.action_template.version
	}
}

resource "ibm_iam_action_control_assignment" "action_control_assignment" {
	target  ={
		type = "Account"  # or "Account Group" or "Enterprise"
		id = "<target-accountId>"
	}
	
	templates{
		id = ibm_iam_action_control_template.action_template.action_control_template_id
		version = ibm_iam_action_control_template.action_template.version
	}

	 # Optional: Use this during update to assign a specific template version
	template_version=ibm_iam_action_control_template_version.template_version.version
}

```
**Note**: Above configuration is to create action control template versions and assign to a target
enterprise account. Add this parameter(***template_version***) and terraform apply again to update the assignment

## Argument Reference

You can specify the following arguments for this resource.
* `templates` - (Required, List) The set of properties required for a ActionControl assignment.
Nested schema for **templates**:
	* `id` - (Required, String) ID of the template.
		* Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^actionControlTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `version` - (Required, String) template version .
		* Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `target` - (Required, List) assignment target account and type.
Nested schema for **target**:
	* `id` - (Required, String) ID of the target account.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
	* `type` - (Required, String) Assignment target type.
	  * Constraints: Allowable values are: `Account`, `Account Group` and `Enterprise`. The maximum length is `30` characters. The minimum length is `1` character.

## Timeouts section

The resource includes default timeout settings for the following operations:

* `create` - (Timeout) Defaults to 30 minutes.
* `update` - (Timeout) Defaults to 30 minutes.
* `delete` - (Timeout) Defaults to 30 minutes.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the action_control_assignment.
* `account_id` - (String) The account GUID that the action control assignments belong to..
* `created_at` - (String) The UTC timestamp when the action control assignment was created.
* `created_by_id` - (String) The iam ID of the entity that created the action control assignment.
* `href` - (String) The href URL that links to the actionControl assignments API by action control assignment ID.
* `last_modified_at` - (String) The UTC timestamp when the action control assignment was last modified.
* `last_modified_by_id` - (String) The iam ID of the entity that last modified the action control assignment.
* `resources` - (List) Object for each account assigned.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested schema for **resources**:
	* `action control` - (List) Set of properties for the assigned resource.
	Nested schema for **action_control**:
		* `error_message` - (List) The error response from API.
		Nested schema for **error_message**:
			* `errors` - (List) The errors encountered during the response.
			  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
			Nested schema for **errors**:
				* `code` - (String) The API error code for the error.
				  * Constraints: Allowable values are: `insufficent_permissions`, `invalid_body`, `invalid_token`, `missing_required_query_parameter`, `not_found`, `action_control_conflict_error`, `action_control_not_found`, `request_not_processed`, `role_conflict_error`, `role_not_found`, `too_many_requests`, `unable_to_process`, `unsupported_content_type`, `action_control_template_conflict_error`, `action_control_template_not_found`, `action_control_assignment_not_found`, `action_control_assignment_conflict_error`.
				* `details` - (List) Additional error details.
				Nested schema for **details**:
					* `conflicts_with` - (List) Details of conflicting resource.
					Nested schema for **conflicts_with**:
						* `etag` - (String) The revision number of the resource.
						* `action_control` - (String) The conflicting action_control id.
				* `message` - (String) The error message returned by the API.
				* `more_info` - (String) Additional info for error.
			* `status_code` - (Integer) The http error code of the response.
			* `name` - (String) Name of the error.
			* `errorCode` - (String) error code.
			* `message` - (String) Error message detailing the nature of the error.
			* `code` - (String) error code.
		* `resource_created` - (List) On success, includes the  action_control assigned.
		Nested schema for **resource_created**:
			* `id` - (String) action_control id.
		* `status` - (String) action_control status.
	* `target` - (List) action_control template details.
	Nested schema for **target**:
		* `id` - (String) action_control template id.
		  * Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^ActionControlTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
		* `version` - (String) action_control template version.
		  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `status` - (String) The action_control assignment status.
  * Constraints: Allowable values are: `in_progress`, `succeeded`, `succeed_with_errors`, `failed`.
* `template` - (List) action_control template details.
Nested schema for **template**:
	* `id` - (String) action_control template id.
	  * Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^actionControlTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `version` - (String) action_control template version.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.


## Import

You can import the `ibm_iam_action_control_assignment` resource by using `id`. ActionControl assignment ID.

# Syntax
<pre>
$ terraform import ibm_iam_action_control_assignment.action_control_assignment &lt;id&gt;
</pre>
