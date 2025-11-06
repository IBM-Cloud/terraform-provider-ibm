---
layout: "ibm"
page_title: "IBM : ibm_iam_role_template"
description: |-
  Manages iam_role_template.
subcategory: "IAM Policy Management"
---

# ibm_iam_role_template

Create, update, and delete iam_role_templates with this resource.

## Example Usage

```hcl
resource "ibm_iam_role_template" "iam_role_template_instance" {
  account_id = "account_id"
  name = "name"
  role {
		name = "name"
		display_name = "display_name"
		service_name = "service_name"
		description = "description"
		actions = [ "actions" ]
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `accept_language` - (Optional, Forces new resource, String) Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).
  * Constraints: The default value is `default`. The minimum length is `1` character.
* `account_id` - (Required, Forces new resource, String) Enterprise account ID where this template is created.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
* `committed` - (Optional, Forces new resource, Boolean) Committed status of the template. If committed is set to true, then the template version can no longer be updated.
* `description` - (Optional, Forces new resource, String) Description of the role template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the role for enterprise users managing IAM templates.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `name` - (Required, Forces new resource, String) Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `role` - (Optional, Forces new resource, List) The role properties that are created in an action resource when the template is assigned.
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

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_role_template.
* `created_at` - (String) The UTC timestamp when the role template was created.
* `created_by_id` - (String) The IAM ID of the entity that created the role template.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
* `href` - (String) The href URL that links to the role templates API by role template ID.
* `last_modified_at` - (String) The UTC timestamp when the role template was last modified.
* `last_modified_by_id` - (String) The IAM ID of the entity that last modified the role template.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
* `state` - (String) State of role template.
  * Constraints: Allowable values are: `active`, `deleted`.
* `version` - (String) The version number of the template used to identify different versions of same template.
  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.


## Import

You can import the `ibm_iam_role_template` resource by using `id`. The role template ID.

# Syntax
<pre>
$ terraform import ibm_iam_role_template.iam_role_template &lt;id&gt;
</pre>
