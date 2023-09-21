---
layout: "ibm"
page_title: "IBM : ibm_project_config"
description: |-
  Manages project_config.
subcategory: "Projects API"
---

# ibm_project_config

Create, update, and delete project_configs with this resource.

## Example Usage

```hcl
resource "ibm_project_config" "project_config_instance" {
  definition {
		name = "name"
		description = "description"
		labels = [ "labels" ]
		authorizations {
			trusted_profile {
				id = "id"
				target_iam_id = "target_iam_id"
			}
			method = "method"
			api_key = "api_key"
		}
		compliance_profile {
			id = "id"
			instance_id = "instance_id"
			instance_location = "instance_location"
			attachment_id = "attachment_id"
			profile_name = "profile_name"
		}
		locator_id = "locator_id"
		input = {  }
		setting = {  }
		type = "terraform_template"
		output {
			name = "name"
			description = "description"
			value = "anything as a string"
		}
  }
  project_id = ibm_project.project_instance.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `definition` - (Required, List) The type and output of a project configuration.
Nested schema for **definition**:
	* `authorizations` - (Optional, List) The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (Optional, String) The IBM Cloud API Key.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
		* `method` - (Optional, String) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
		* `trusted_profile` - (Optional, List) The trusted profile for authorizations.
		Nested schema for **trusted_profile**:
			* `id` - (Optional, String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `target_iam_id` - (Optional, String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `compliance_profile` - (Optional, List) The profile required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_location` - (Optional, String) The location of the compliance instance.
		  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
		* `profile_name` - (Optional, String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `description` - (Optional, String) A project configuration description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `input` - (Optional, List) The input variables for configuration definition and environment.
	Nested schema for **input**:
	* `labels` - (Optional, List) The configuration labels.
	  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
	* `locator_id` - (Required, String) A dotted value of catalogID.versionID.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (Required, String) The configuration name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `output` - (Optional, List) The outputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **output**:
		* `description` - (Optional, String) A short explanation of the output value.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
		* `name` - (Required, String) The variable name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (Optional, String) Can be any value - a string, number, boolean, array, or object.
	* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.
	Nested schema for **setting**:
	* `type` - (Computed, String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project_config.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
* `last_save` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `needs_attention_state` - (List) The needs attention state of a configuration.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
* `project_config_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superceded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`.
* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.
* `updated_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `version` - (Integer) The version of the configuration.


## Import

You can import the `ibm_project_config` resource by using `id`.
The `id` property can be formed from `project_id`, and `id` in the following format:

```
<project_id>/<id>
```
* `project_id`: A string. The unique project ID.
* `id`: A string. The unique config ID.

# Syntax
```
$ terraform import ibm_project_config.project_config <project_id>/<id>
```
