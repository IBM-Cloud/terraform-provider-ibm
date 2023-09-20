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
  definition {
		name = "name"
		description = "description"
		destroy_on_delete = true
  }
  location = "us-south"
  resource_group = "Default"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `definition` - (Optional, List) The definition of the project.
Nested schema for **definition**:
	* `description` - (Optional, String) A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `destroy_on_delete` - (Required, Boolean) The policy that indicates whether the resources are destroyed or not when a project is deleted.
	  * Constraints: The default value is `true`.
	* `name` - (Required, String) The name of the project.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
* `location` - (Required, String) The location where the project's data and tools are created.
  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
* `resource_group` - (Required, String) The resource group where the project's data and tools are created.
  * Constraints: The maximum length is `40` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project.
* `configs` - (List) The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **configs**:
	* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `definition` - (List) The name and description of a project configuration.
	Nested schema for **definition**:
		* `description` - (String) A project configuration description.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
		* `name` - (String) The configuration name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `href` - (String) A relative URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
	* `last_approved` - (List) The last approved metadata of the configuration.
	Nested schema for **last_approved**:
		* `comment` - (String) The comment left by the user who approved the configuration.
		  * Constraints: The default value is ``.
		* `is_forced` - (Boolean) The flag that indicates whether the approval was forced approved.
		* `timestamp` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
		* `user_id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `last_deployed` - (List) The action job performed on the project configuration.
	Nested schema for **last_deployed**:
		* `href` - (String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `job` - (List) A brief summary of an action.
		Nested schema for **job**:
			* `id` - (String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `summary` - (List) The summaries of jobs that were performed on the configuration.
			Nested schema for **summary**:
				* `apply_messages` - (Map) The messages of apply jobs on the configuration.
				* `apply_summary` - (Map) The summary of the apply jobs on the configuration.
				* `destroy_messages` - (Map) The messages of destroy jobs on the configuration.
				* `destroy_summary` - (Map) The summary of the destroy jobs on the configuration.
				* `message_summary` - (Map) The message summaries of jobs on the configuration.
				* `plan_messages` - (Map) The messages of plan jobs on the configuration.
				* `plan_summary` - (Map) The summary of the plan jobs on the configuration.
		* `result` - (String) The result of the last action.
		  * Constraints: Allowable values are: `failed`, `passed`.
	* `last_save` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `last_undeployed` - (List) The action job performed on the project configuration.
	Nested schema for **last_undeployed**:
		* `href` - (String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `job` - (List) A brief summary of an action.
		Nested schema for **job**:
			* `id` - (String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `summary` - (List) The summaries of jobs that were performed on the configuration.
			Nested schema for **summary**:
				* `apply_messages` - (Map) The messages of apply jobs on the configuration.
				* `apply_summary` - (Map) The summary of the apply jobs on the configuration.
				* `destroy_messages` - (Map) The messages of destroy jobs on the configuration.
				* `destroy_summary` - (Map) The summary of the destroy jobs on the configuration.
				* `message_summary` - (Map) The message summaries of jobs on the configuration.
				* `plan_messages` - (Map) The messages of plan jobs on the configuration.
				* `plan_summary` - (Map) The summary of the plan jobs on the configuration.
		* `result` - (String) The result of the last action.
		  * Constraints: Allowable values are: `failed`, `passed`.
	* `last_validated` - (List) The action job performed on the project configuration.
	Nested schema for **last_validated**:
		* `cost_estimate` - (List) The cost estimate of the configuration.It only exists after the first configuration validation.
		Nested schema for **cost_estimate**:
			* `currency` - (String) The currency of the cost estimate of the configuration.
			* `diff_total_hourly_cost` - (String) The difference between current and past total hourly cost estimates of the configuration.
			* `diff_total_monthly_cost` - (String) The difference between current and past total monthly cost estimates of the configuration.
			* `past_total_hourly_cost` - (String) The past total hourly cost estimate of the configuration.
			* `past_total_monthly_cost` - (String) The past total monthly cost estimate of the configuration.
			* `time_generated` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
			* `total_hourly_cost` - (String) The total hourly cost estimate of the configuration.
			* `total_monthly_cost` - (String) The total monthly cost estimate of the configuration.
			* `user_id` - (String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `version` - (String) The version of the cost estimate of the configuration.
		* `cra_logs` - (List) The Code Risk Analyzer logs of the configuration.
		Nested schema for **cra_logs**:
			* `cra_version` - (String) The version of the Code Risk Analyzer logs of the configuration.
			* `schema_version` - (String) The schema version of Code Risk Analyzer logs of the configuration.
			* `status` - (String) The status of the Code Risk Analyzer logs of the configuration.
			* `summary` - (Map) The summary of the Code Risk Analyzer logs of the configuration.
			* `timestamp` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
		* `href` - (String) A relative URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `job` - (List) A brief summary of an action.
		Nested schema for **job**:
			* `id` - (String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `summary` - (List) The summaries of jobs that were performed on the configuration.
			Nested schema for **summary**:
				* `apply_messages` - (Map) The messages of apply jobs on the configuration.
				* `apply_summary` - (Map) The summary of the apply jobs on the configuration.
				* `destroy_messages` - (Map) The messages of destroy jobs on the configuration.
				* `destroy_summary` - (Map) The summary of the destroy jobs on the configuration.
				* `message_summary` - (Map) The message summaries of jobs on the configuration.
				* `plan_messages` - (Map) The messages of plan jobs on the configuration.
				* `plan_summary` - (Map) The summary of the plan jobs on the configuration.
		* `result` - (String) The result of the last action.
		  * Constraints: Allowable values are: `failed`, `passed`.
	* `needs_attention_state` - (List) The needs attention state of a configuration.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	* `project_id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superceded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`.
	* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.
	* `updated_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `version` - (Integer) The version of the configuration.
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
