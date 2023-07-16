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
		type = "terraform_template"
		input {
			name = "name"
			type = "array"
			value = "anything as a string"
			required = true
		}
		output {
			name = "name"
			description = "description"
			value = "anything as a string"
		}
		setting {
			name = "name"
			value = "value"
		}
		id = "id"
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
		  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
		* `profile_name` - (Optional, String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `description` - (Optional, String) The description of the project configuration.
	  * Constraints: The default value is ``. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `id` - (Required, String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `input` - (Optional, List) The outputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **input**:
		* `name` - (Required, String) The variable name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `required` - (Optional, Boolean) Whether the variable is required or not.
		* `type` - (Required, String) The variable type.
		  * Constraints: Allowable values are: `array`, `boolean`, `float`, `int`, `number`, `password`, `string`, `object`.
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
	* `type` - (Required, String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
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
* `event_notifications_crn` - (String) The CRN of the event notifications instance if one is connected to this project.
  * Constraints: The default value is ``. The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `state` - (String) The project status value.
  * Constraints: Allowable values are: `ready`, `deleting`, `deleting_failed`.


## Import

You can import the `ibm_project` resource by using `id`. The unique ID of a project.

# Syntax
```
$ terraform import ibm_project.project <id>
```
