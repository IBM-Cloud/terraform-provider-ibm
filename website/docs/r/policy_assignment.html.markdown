---
layout: "ibm"
page_title: "IBM : ibm_policy_assignment"
description: |-
  Manages policy_assignment.
subcategory: "IAM Policy Management"
---

# ibm_policy_assignment

Create, update, and delete policy_assignments with this resource.

## Example Usage

```hcl
resource "ibm_iam_policy_template" "policy_s2s_template" {
	name = "TerraformS2STest"
	policy {
		type = "authorization"
		description = "description"
		resource {
			attributes {
				key = "serviceName"
				operator = "stringEquals"
				value = "kms"
			}
		}
		subject {
			attributes {
				key = "serviceName"
				operator = "stringEquals"
				value = "compliance"
			}
		}
		roles = ["Reader"]
	}
	committed=true
}
resource "ibm_iam_policy_template_version" "template_version" {
	template_id = ibm_iam_policy_template.policy_s2s_template.template_id
	policy {
		type = "authorization"
		description = "description"
		resource {
			attributes {
				key = "serviceName"
				operator = "stringEquals"
				value = "appid"
			}
		}
		subject {
			attributes {
				key = "serviceName"
				operator = "stringEquals"
				value = "compliance"
			}
		}
		roles = ["Reader"]
}
committed=true
}

resource "ibm_iam_policy_assignment" "policy_assignment" {
	version ="1.0"
	target  ={
		type = "Account"
		id = "<target-accountId>"
	}
	
	templates{
		id = ibm_iam_policy_template.policy_s2s_template.template_id 
		version = ibm_iam_policy_template.policy_s2s_template.version
	}
	template_version=ibm_iam_policy_template_version.template_version.version
}

resource "ibm_iam_policy_assignment" "policy_assignment" {
	version ="1.0"
	target  ={
		type = "Account Group"
		id = "<target-accountgroupId>"
	}
	
	templates{
		id = ibm_iam_policy_template.policy_s2s_template.template_id 
		version = ibm_iam_policy_template.policy_s2s_template.version
	}
	template_version=ibm_iam_policy_template_version.template_version.version
}

resource "ibm_iam_policy_assignment" "policy_assignment" {
	version ="1.0"
	target  ={
		type = "Enterprise"
		id = "<target-enterpriseId>"
	}
	
	templates{
		id = ibm_iam_policy_template.policy_s2s_template.template_id 
		version = ibm_iam_policy_template.policy_s2s_template.version
	}
	template_version=ibm_iam_policy_template_version.template_version.version
}
```
**Note**: Above configuration is to create policy template versions and assign to a target
enterprise account. Update this parameter(***template_version***) and terraform apply again to update the assignment

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, String) Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).
  * Constraints: The default value is `default`. The minimum length is `1` character.
* `templates` - (Required, List) The set of properties required for a policy assignment.
Nested schema for **templates**:
	* `id` - (Required, String) ID of the template.
		* Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^policyTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `version` - (Required, String) template version .
		* Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `target` - (Required, List) assignment target account and type.
Nested schema for **target**:
	* `id` - (Required, String) ID of the target account.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
	* `type` - (Required, String) Assignment target type.
	  * Constraints: Allowable values are: `Account`, `Account Group` and `Enterprise`. The maximum length is `30` characters. The minimum length is `1` character.
* `version` - (Required, String) specify version of response body format.
  * Constraints: Allowable values are: `1.0`. The minimum length is `1` character.

## Timeouts section

The resource includes default timeout settings for the following operations:

* `create` - (Timeout) Defaults to 30 minutes.
* `update` - (Timeout) Defaults to 30 minutes.
* `delete` - (Timeout) Defaults to 30 minutes.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the policy_assignment.
* `account_id` - (String) The account GUID that the policies assignments belong to..
* `created_at` - (String) The UTC timestamp when the policy assignment was created.
* `created_by_id` - (String) The iam ID of the entity that created the policy assignment.
* `href` - (String) The href URL that links to the policies assignments API by policy assignment ID.
* `last_modified_at` - (String) The UTC timestamp when the policy assignment was last modified.
* `last_modified_by_id` - (String) The iam ID of the entity that last modified the policy assignment.
* `resources` - (List) Object for each account assigned.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
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
		* `resource_created` - (List) On success, includes the  policy assigned.
		Nested schema for **resource_created**:
			* `id` - (String) policy id.
		* `status` - (String) policy status.
	* `target` - (List) policy template details.
	Nested schema for **target**:
		* `id` - (String) policy template id.
		  * Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^policyTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
		* `version` - (String) policy template version.
		  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `status` - (String) The policy assignment status.
  * Constraints: Allowable values are: `in_progress`, `succeeded`, `succeed_with_errors`, `failed`.
* `subject` - (List) subject details of access type assignment.
Nested schema for **subject**:
	* `id` - (String)
	  * Constraints: The minimum length is `1` character. The value must match regular expression `/^((IBMid)|(iam-ServiceId)|(AccessGroupId)|(iam-Profile)|(SL)|([a-zA-Z0-9]{3,10}))-/`.
	* `type` - (String)
	  * Constraints: Allowable values are: `iam_id`, `access_group_id`. The minimum length is `1` character.
* `template` - (List) policy template details.
Nested schema for **template**:
	* `id` - (String) policy template id.
	  * Constraints: The maximum length is `51` characters. The minimum length is `1` character. The value must match regular expression `/^policyTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `version` - (String) policy template version.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.


## Import

You can import the `ibm_iam_policy_assignment` resource by using `id`. Policy assignment ID.

# Syntax
<pre>
$ terraform import ibm_iam_policy_assignment.policy_assignment &lt;id&gt;
</pre>
