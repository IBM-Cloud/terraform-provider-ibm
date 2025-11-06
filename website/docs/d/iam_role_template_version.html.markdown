---
layout: "ibm"
page_title: "IBM : ibm_iam_role_template_version"
description: |-
  Get information about iam_role_template_version
subcategory: "IAM Policy Management"
---

# ibm_iam_role_template_version

Provides a read-only data source to retrieve information about an iam_role_template_version. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_role_template_version" "iam_role_template_version" {
	role_template_id = "role_template_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `role_template_id` - (Required, Forces new resource, String) The role template ID.
  * Constraints: The maximum length is `49` characters. The minimum length is `1` character. The value must match regular expression `/^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
* `state` - (Optional, String) Role template state.
  * Constraints: Allowable values are: `active`, `deleted`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_role_template_version.
* `versions` - (List) List of role templates versions.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **versions**:
	* `account_id` - (String) Enterprise account ID where this template is created.
	  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9-]*$/`.
	* `committed` - (Boolean) Committed status of the template. If committed is set to true, then the template version can no longer be updated.
	* `created_at` - (String) The UTC timestamp when the role template was created.
	* `created_by_id` - (String) The IAM ID of the entity that created the role template.
	  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
	* `description` - (String) Description of the role template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the role for enterprise users managing IAM templates.
	  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
	* `href` - (String) The href URL that links to the role templates API by role template ID.
	* `id` - (String) The role template ID.
	  * Constraints: The maximum length is `49` characters. The minimum length is `1` character. The value must match regular expression `/^roleTemplate-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `last_modified_at` - (String) The UTC timestamp when the role template was last modified.
	* `last_modified_by_id` - (String) The IAM ID of the entity that last modified the role template.
	  * Constraints: The maximum length is `250` characters. The minimum length is `1` character.
	* `name` - (String) Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `role` - (List) The role properties that are created in an action resource when the template is assigned.
	Nested schema for **role**:
		* `actions` - (List) The actions of the role.
		  * Constraints: The minimum length is `1` item.
		* `description` - (String) Description of the role.
		  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
		* `display_name` - (String) The display the name of the role that is shown in the console.
		  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^((?!<|>).)*$/`.
		* `name` - (String) The name of the role that is used in the CRN. This must be alphanumeric and capitalized.
		  * Constraints: The maximum length is `30` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Z]{1}[A-Za-z0-9]{0,29}$/`.
		* `service_name` - (String) The service name that the role refers.
		  * Constraints: The maximum length is `300` characters. The minimum length is `1` character.
	* `state` - (String) State of role template.
	  * Constraints: Allowable values are: `active`, `deleted`.
	* `version` - (String) The version number of the template used to identify different versions of same template.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

