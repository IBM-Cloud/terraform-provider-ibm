---
layout: "ibm"
page_title: "IBM : ibm_iam_role_template_version"
description: |-
  Manages iam_role_template_version.
subcategory: "IAM Policy Management"
---

# ibm_iam_role_template_version

Create, update, and delete iam_role_template_versions with this resource.

## Example Usage

```hcl
resource "ibm_iam_role_template_version" "iam_role_template_version_instance" {
  role {
		name = "name"
		display_name = "display_name"
		service_name = "service_name"
		description = "description"
		actions = [ "actions" ]
  }
  role_template_id = "role_template_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `committed` - (Optional, Boolean) Committed status of the template. If committed is set to true, then the template version can no longer be updated.
* `description` - (Optional, String) Description of the role template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the role for enterprise users managing IAM templates.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `name` - (Optional, String) Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `role` - (Required, List) The role properties that are created in an action resource when the template is assigned.
Nested schema for **role**:
	* `actions` - (Required, List) The actions of the role.
	  * Constraints: The minimum length is `1` item.
	* `description` - (Optional, String) Description of the role.
	  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `display_name` - (Required, String) The display the name of the role that is shown in the console.
	  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^((?!<|>).)*$/`.
	* `name` - (Required, String) The name of the role that is used in the CRN. This must be alphanumeric and capitalized.
	  * Constraints: The maximum length is `30` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Z]{1}[A-Za-z0-9]{0,29}$/`.
	* `service_name` - (Required, String) The service name that the role refers.
	  * Constraints: The maximum length is `300` characters. The minimum length is `1` character.
* `role_template_id` - (Required, Forces new resource, String) The role template ID.
  * Constraints: The maximum length is `49` characters. The minimum length is `1` character. The value must match regular expression `/^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_role_template_version.
* `account_id` - (String) Enterprise account ID where this template is created.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
* `created_at` - (String) The UTC timestamp when the role template was created.
* `created_by_id` - (String) The IAM ID of the entity that created the role template.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
* `href` - (String) The href URL that links to the role templates API by role template ID.
* `id` - (String) The role template ID.
  * Constraints: The maximum length is `49` characters. The minimum length is `1` character. The value must match regular expression `/^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
* `last_modified_at` - (String) The UTC timestamp when the role template was last modified.
* `last_modified_by_id` - (String) The IAM ID of the entity that last modified the role template.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
* `state` - (String) State of role template.
  * Constraints: Allowable values are: `active`, `deleted`.

* `etag` - ETag identifier for iam_role_template_version.

## Import

You can import the `ibm_iam_role_template_version` resource by using `version`.
The `version` property can be formed from and `role_template_id` in the following format:

<pre>
&lt;role_template_id&gt;
</pre>
* `role_template_id`: A string. The role template ID.

# Syntax
<pre>
$ terraform import ibm_iam_role_template_version.iam_role_template_version &lt;role_template_id&gt;
</pre>
