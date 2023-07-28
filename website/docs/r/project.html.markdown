---
layout: "ibm"
page_title: "IBM : ibm_project"
description: |-
  Manages project.
subcategory: "Projects API"
---

# ibm_project

Create, update, and delete projects with this resource.

## Example Usage

```hcl
resource "ibm_project" "project_instance" {
  configs {
		id = "id"
		project_id = "project_id"
		version = 1
		is_draft = true
		needs_attention_state = [ "anything as a string" ]
		state = "approved"
		pipeline_state = "pipeline_failed"
		update_available = true
		created_at = "2021-01-31T09:44:12Z"
		updated_at = "2021-01-31T09:44:12Z"
		last_approved {
			is_forced = true
			comment = "comment"
			timestamp = "2021-01-31T09:44:12Z"
			user_id = "user_id"
		}
		last_save = "2021-01-31T09:44:12Z"
		name = "name"
		labels = [ "labels" ]
		description = "description"
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
		input {
			name = "name"
			value = "anything as a string"
		}
		setting {
			name = "name"
			value = "value"
		}
		type = "terraform_template"
		output {
			name = "name"
			description = "description"
			value = "anything as a string"
		}
		active_draft {
			version = 1
			state = "discarded"
			pipeline_state = "pipeline_failed"
			href = "href"
		}
		definition {
			name = "name"
			labels = [ "labels" ]
			description = "description"
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
			input {
				name = "name"
				value = "anything as a string"
			}
			setting {
				name = "name"
				value = "value"
			}
			type = "terraform_template"
			output {
				name = "name"
				description = "description"
				value = "anything as a string"
			}
		}
		href = "href"
  }
  description = "A microservice to deploy on top of ACME infrastructure."
  location = "us-south"
  name = "acme-microservice"
  resource_group = "Default"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `configs` - (Optional, List) The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **configs**:
	* `active_draft` - (Optional, List) The project configuration version.
	Nested schema for **active_draft**:
		* `href` - (Optional, String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `pipeline_state` - (Optional, String) The pipeline state of the configuration. It only exists after the first configuration validation.
		  * Constraints: Allowable values are: `pipeline_failed`, `pipeline_running`, `pipeline_succeeded`.
		* `state` - (Required, String) The state of the configuration draft.
		  * Constraints: Allowable values are: `discarded`, `merged`, `active`.
		* `version` - (Required, Integer) The version number of the configuration.
	* `authorizations` - (Optional, List) The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (Optional, String) The IBM Cloud API Key.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
		* `method` - (Optional, String) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
		* `trusted_profile` - (Optional, List) The trusted profile for authorizations.
		Nested schema for **trusted_profile**:
			* `id` - (Optional, String) The unique ID of a project.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `target_iam_id` - (Optional, String) The unique ID of a project.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `compliance_profile` - (Optional, List) The profile required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (Optional, String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `id` - (Optional, String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_id` - (Optional, String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_location` - (Optional, String) The location of the compliance instance.
		  * Constraints: The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
		* `profile_name` - (Optional, String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `created_at` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `definition` - (Optional, List) The project configuration definition.
	Nested schema for **definition**:
		* `authorizations` - (Optional, List) The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.
		Nested schema for **authorizations**:
			* `api_key` - (Optional, String) The IBM Cloud API Key.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
			* `method` - (Optional, String) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
			  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
			* `trusted_profile` - (Optional, List) The trusted profile for authorizations.
			Nested schema for **trusted_profile**:
				* `id` - (Optional, String) The unique ID of a project.
				  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
				* `target_iam_id` - (Optional, String) The unique ID of a project.
				  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `compliance_profile` - (Optional, List) The profile required for compliance.
		Nested schema for **compliance_profile**:
			* `attachment_id` - (Optional, String) The unique ID of a project.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `id` - (Optional, String) The unique ID of a project.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `instance_id` - (Optional, String) The unique ID of a project.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `instance_location` - (Optional, String) The location of the compliance instance.
			  * Constraints: The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
			* `profile_name` - (Optional, String) The name of the compliance profile.
			  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
		* `description` - (Optional, String) The description of the project configuration.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
		* `input` - (Optional, List) The inputs of a Schematics template property.
		  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
		Nested schema for **input**:
			* `name` - (Required, String) The variable name.
			  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
			* `value` - (Optional, String) Can be any value - a string, number, boolean, array, or object.
		* `labels` - (Optional, List) A collection of configuration labels.
		  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
		* `locator_id` - (Required, String) A dotted value of catalogID.versionID.
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
		* `name` - (Required, String) The name of the configuration.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
		* `output` - (Optional, List) The outputs of a Schematics template property.
		  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
		Nested schema for **output**:
			* `description` - (Optional, String) A short explanation of the output value.
			  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
			* `name` - (Required, String) The variable name.
			  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
			* `value` - (Optional, String) Can be any value - a string, number, boolean, array, or object.
		* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.
		  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
		Nested schema for **setting**:
			* `name` - (Required, String) The name of the configuration setting.
			  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
			* `value` - (Required, String) The value of the configuration setting.
			  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `type` - (Optional, String) The type of a project configuration manual property.
		  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
	* `description` - (Optional, String) The description of the project configuration.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `href` - (Optional, String) A relative URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (Optional, String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `input` - (Optional, List) The inputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **input**:
		* `name` - (Required, String) The variable name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (Optional, String) Can be any value - a string, number, boolean, array, or object.
	* `is_draft` - (Optional, Boolean) The flag that indicates whether the version of the configuration is draft, or active.
	* `labels` - (Optional, List) A collection of configuration labels.
	  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
	* `last_approved` - (Optional, List) The last approved metadata of the configuration.
	Nested schema for **last_approved**:
		* `comment` - (Optional, String) The comment left by the user who approved the configuration.
		* `is_forced` - (Required, Boolean) The flag that indicates whether the approval was forced approved.
		* `timestamp` - (Required, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
		* `user_id` - (Required, String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `last_save` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `locator_id` - (Required, String) A dotted value of catalogID.versionID.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (Required, String) The name of the configuration.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `needs_attention_state` - (Optional, List) The needs attention state of a configuration.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	* `output` - (Optional, List) The outputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **output**:
		* `description` - (Optional, String) A short explanation of the output value.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
		* `name` - (Required, String) The variable name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (Optional, String) Can be any value - a string, number, boolean, array, or object.
	* `pipeline_state` - (Optional, String) The pipeline state of the configuration. It only exists after the first configuration validation.
	  * Constraints: Allowable values are: `pipeline_failed`, `pipeline_running`, `pipeline_succeeded`.
	* `project_id` - (Optional, String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **setting**:
		* `name` - (Required, String) The name of the configuration setting.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (Required, String) The value of the configuration setting.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `state` - (Optional, String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `installed`, `installed_failed`, `installing`, `superceded`, `uninstalling`, `uninstalling_failed`, `validated`, `validating`, `validating_failed`.
	* `type` - (Optional, String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
	* `update_available` - (Optional, Boolean) The flag that indicates whether a configuration update is available.
	* `updated_at` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `version` - (Optional, Integer) The version of the configuration.
* `description` - (Optional, String) A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
* `destroy_on_delete` - (Optional, Boolean) The policy that indicates whether the resources are destroyed or not when a project is deleted.
  * Constraints: The default value is `true`.
* `location` - (Required, String) The location where the project's data and tools are created.
  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
* `name` - (Required, String) The name of the project.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
* `resource_group` - (Required, String) The resource group where the project's data and tools are created.
  * Constraints: The maximum length is `40` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `cumulative_needs_attention_view` - (List) The cumulative list of needs attention items for a project. If the view is successfully retrieved, an array which could be empty is returned.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **cumulative_needs_attention_view**:
	* `config_id` - (String) A unique ID for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `config_version` - (Integer) The version number of the configuration.
	* `event` - (String) The event name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `event_id` - (String) A unique ID for that individual event.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `cumulative_needs_attention_view_error` - (Boolean) True indicates that the fetch of the needs attention items failed. It only exists if there was an error while retrieving the cumulative needs attention view.
* `definition` - (List) The definition of the project.
Nested schema for **definition**:
	* `description` - (String) A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `destroy_on_delete` - (Boolean) The policy that indicates whether the resources are destroyed or not when a project is deleted.
	  * Constraints: The default value is `true`.
	* `name` - (String) The name of the project.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
* `event_notifications_crn` - (String) The CRN of the event notifications instance if one is connected to this project.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `state` - (String) The project status value.
  * Constraints: Allowable values are: `ready`, `deleting`, `deleting_failed`.


## Import

You can import the `ibm_project` resource by using `id`. The unique ID of a project.

# Syntax
```
$ terraform import ibm_project.project <id>
```
