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
  description = "Stage environment configuration, which includes services common to all the environment regions. There must be a blueprint configuring all the services common to the stage regions. It is a terraform_template type of configuration that points to a Github repo hosting the terraform modules that can be deployed by a Schematics Workspace."
  input = {"account_id":"$configs[].name["account-stage"].input.account_id","resource_group":"stage","access_tags":["env:stage"],"logdna_name":"Name of the LogDNA stage service instance","sysdig_name":"Name of the SysDig stage service instance"}
  locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"
  name = "env-stage"
  project_id = ibm_project.project_instance.id
  setting = {"IBMCLOUD_TOOLCHAIN_ENDPOINT":"https://api.us-south.devops.dev.cloud.ibm.com"}
}
```

## Argument Reference

You can specify the following arguments for this resource.

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
* `description` - (Optional, String) The description of the project configuration.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
* `input` - (Optional, List) The input variables for the configuration definition.
Nested schema for **input**:
* `labels` - (Optional, List) A collection of configuration labels.
  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
* `locator_id` - (Optional, String) A dotted value of catalogID.versionID.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
* `name` - (Optional, String) The name of the configuration.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.
Nested schema for **setting**:

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project_config.
* `check_job` - (List) The action job performed on the project configuration.
Nested schema for **check_job**:
	* `href` - (String) A relative URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
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
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `install_job` - (List) The action job performed on the project configuration.
Nested schema for **install_job**:
	* `href` - (String) A relative URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
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
* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
* `last_approved` - (List) The last approved metadata of the configuration.
Nested schema for **last_approved**:
	* `comment` - (String) The comment left by the user who approved the configuration.
	* `is_forced` - (Boolean) The flag that indicates whether the approval was forced approved.
	* `timestamp` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `user_id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `last_save` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `needs_attention_state` - (List) The needs attention state of a configuration.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
* `output` - (List) The outputs of a Schematics template property.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **output**:
	* `description` - (String) A short explanation of the output value.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) Can be any value - a string, number, boolean, array, or object.
* `project_config_canonical_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `installed`, `installed_failed`, `installing`, `superceded`, `uninstalling`, `uninstalling_failed`, `validated`, `validating`, `validating_failed`.
* `type` - (String) The type of a project configuration manual property.
  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
* `uninstall_job` - (List) The action job performed on the project configuration.
Nested schema for **uninstall_job**:
	* `href` - (String) A relative URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
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
