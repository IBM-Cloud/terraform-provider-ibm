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
		id = "id"
		project_id = "project_id"
		version = 1
		is_draft = true
		needs_attention_state = [ "anything as a string" ]
		state = "approved"
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
		cra_logs {
			cra_version = "cra_version"
			schema_version = "schema_version"
			status = "status"
			summary = { "key" = "anything as a string" }
			timestamp = "2021-01-31T09:44:12Z"
		}
		cost_estimate {
			version = "version"
			currency = "currency"
			total_hourly_cost = "total_hourly_cost"
			total_monthly_cost = "total_monthly_cost"
			past_total_hourly_cost = "past_total_hourly_cost"
			past_total_monthly_cost = "past_total_monthly_cost"
			diff_total_hourly_cost = "diff_total_hourly_cost"
			diff_total_monthly_cost = "diff_total_monthly_cost"
			time_generated = "2021-01-31T09:44:12Z"
			user_id = "user_id"
		}
		check_job {
			id = "id"
			href = "href"
			summary {
				plan_summary = { "key" = "anything as a string" }
				apply_summary = { "key" = "anything as a string" }
				destroy_summary = { "key" = "anything as a string" }
				message_summary = { "key" = "anything as a string" }
				plan_messages = { "key" = "anything as a string" }
				apply_messages = { "key" = "anything as a string" }
				destroy_messages = { "key" = "anything as a string" }
			}
		}
		install_job {
			id = "id"
			href = "href"
			summary {
				plan_summary = { "key" = "anything as a string" }
				apply_summary = { "key" = "anything as a string" }
				destroy_summary = { "key" = "anything as a string" }
				message_summary = { "key" = "anything as a string" }
				plan_messages = { "key" = "anything as a string" }
				apply_messages = { "key" = "anything as a string" }
				destroy_messages = { "key" = "anything as a string" }
			}
		}
		uninstall_job {
			id = "id"
			href = "href"
			summary {
				plan_summary = { "key" = "anything as a string" }
				apply_summary = { "key" = "anything as a string" }
				destroy_summary = { "key" = "anything as a string" }
				message_summary = { "key" = "anything as a string" }
				plan_messages = { "key" = "anything as a string" }
				apply_messages = { "key" = "anything as a string" }
				destroy_messages = { "key" = "anything as a string" }
			}
		}
		output {
			name = "name"
			description = "description"
			value = "anything as a string"
		}
		type = "terraform_template"
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
			* `id` - (Optional, String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `target_iam_id` - (Optional, String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `check_job` - (Optional, List) The action job performed on the project configuration.
	Nested schema for **check_job**:
		* `href` - (Optional, String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `summary` - (Optional, List) The summaries of jobs that were performed on the configuration.
		Nested schema for **summary**:
			* `apply_messages` - (Optional, Map) The messages of apply jobs on the configuration.
			* `apply_summary` - (Optional, Map) The summary of the apply jobs on the configuration.
			* `destroy_messages` - (Optional, Map) The messages of destroy jobs on the configuration.
			* `destroy_summary` - (Optional, Map) The summary of the destroy jobs on the configuration.
			* `message_summary` - (Optional, Map) The message summaries of jobs on the configuration.
			* `plan_messages` - (Optional, Map) The messages of plan jobs on the configuration.
			* `plan_summary` - (Optional, Map) The summary of the plan jobs on the configuration.
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
	* `cost_estimate` - (Optional, List) The cost estimate of the configuration.It only exists after the first configuration validation.
	Nested schema for **cost_estimate**:
		* `currency` - (Optional, String) The currency of the cost estimate of the configuration.
		* `diff_total_hourly_cost` - (Optional, String) The difference between current and past total hourly cost estimates of the configuration.
		* `diff_total_monthly_cost` - (Optional, String) The difference between current and past total monthly cost estimates of the configuration.
		* `past_total_hourly_cost` - (Optional, String) The past total hourly cost estimate of the configuration.
		* `past_total_monthly_cost` - (Optional, String) The past total monthly cost estimate of the configuration.
		* `time_generated` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
		* `total_hourly_cost` - (Optional, String) The total hourly cost estimate of the configuration.
		* `total_monthly_cost` - (Optional, String) The total monthly cost estimate of the configuration.
		* `user_id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `version` - (Optional, String) The version of the cost estimate of the configuration.
	* `cra_logs` - (Optional, List) The Code Risk Analyzer logs of the configuration.
	Nested schema for **cra_logs**:
		* `cra_version` - (Optional, String) The version of the Code Risk Analyzer logs of the configuration.
		* `schema_version` - (Optional, String) The schema version of Code Risk Analyzer logs of the configuration.
		* `status` - (Optional, String) The status of the Code Risk Analyzer logs of the configuration.
		* `summary` - (Optional, Map) The summary of the Code Risk Analyzer logs of the configuration.
		* `timestamp` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `created_at` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `description` - (Optional, String) The description of the project configuration.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `id` - (Optional, String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `input` - (Optional, List) The input variables for the configuration definition.
	Nested schema for **input**:
	* `install_job` - (Optional, List) The action job performed on the project configuration.
	Nested schema for **install_job**:
		* `href` - (Optional, String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `summary` - (Optional, List) The summaries of jobs that were performed on the configuration.
		Nested schema for **summary**:
			* `apply_messages` - (Optional, Map) The messages of apply jobs on the configuration.
			* `apply_summary` - (Optional, Map) The summary of the apply jobs on the configuration.
			* `destroy_messages` - (Optional, Map) The messages of destroy jobs on the configuration.
			* `destroy_summary` - (Optional, Map) The summary of the destroy jobs on the configuration.
			* `message_summary` - (Optional, Map) The message summaries of jobs on the configuration.
			* `plan_messages` - (Optional, Map) The messages of plan jobs on the configuration.
			* `plan_summary` - (Optional, Map) The summary of the plan jobs on the configuration.
	* `is_draft` - (Optional, Boolean) The flag that indicates whether the version of the configuration is draft, or active.
	* `labels` - (Optional, List) A collection of configuration labels.
	  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
	* `last_approved` - (Optional, List) The last approved metadata of the configuration.
	Nested schema for **last_approved**:
		* `comment` - (Optional, String) The comment left by the user who approved the configuration.
		* `is_forced` - (Required, Boolean) The flag that indicates whether the approval was forced approved.
		* `timestamp` - (Required, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
		* `user_id` - (Required, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `last_save` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `locator_id` - (Optional, String) A dotted value of catalogID.versionID.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (Optional, String) The name of the configuration.
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
	* `project_id` - (Optional, String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.
	Nested schema for **setting**:
	* `state` - (Optional, String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `installed`, `installed_failed`, `installing`, `superceded`, `uninstalling`, `uninstalling_failed`, `validated`, `validating`, `validating_failed`.
	* `type` - (Optional, String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
	* `uninstall_job` - (Optional, List) The action job performed on the project configuration.
	Nested schema for **uninstall_job**:
		* `href` - (Optional, String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `id` - (Optional, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `summary` - (Optional, List) The summaries of jobs that were performed on the configuration.
		Nested schema for **summary**:
			* `apply_messages` - (Optional, Map) The messages of apply jobs on the configuration.
			* `apply_summary` - (Optional, Map) The summary of the apply jobs on the configuration.
			* `destroy_messages` - (Optional, Map) The messages of destroy jobs on the configuration.
			* `destroy_summary` - (Optional, Map) The summary of the destroy jobs on the configuration.
			* `message_summary` - (Optional, Map) The message summaries of jobs on the configuration.
			* `plan_messages` - (Optional, Map) The messages of plan jobs on the configuration.
			* `plan_summary` - (Optional, Map) The summary of the plan jobs on the configuration.
	* `update_available` - (Optional, Boolean) The flag that indicates whether a configuration update is available.
	* `updated_at` - (Optional, String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `version` - (Optional, Integer) The version of the configuration.
* `description` - (Optional, String) A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
* `destroy_on_delete` - (Optional, Boolean) The policy that indicates whether the resources are destroyed or not when a project is deleted.
  * Constraints: The default value is `true`.
* `location` - (Required, String) The location where the project's data and tools are created.
  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
* `name` - (Optional, String) The name of the project.
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
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `state` - (String) The project status value.
  * Constraints: Allowable values are: `ready`, `deleting`, `deleting_failed`.


## Import

You can import the `ibm_project` resource by using `id`. The unique ID.

# Syntax
```
$ terraform import ibm_project.project <id>
```
