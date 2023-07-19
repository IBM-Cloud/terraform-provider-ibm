---
layout: "ibm"
page_title: "IBM : ibm_project_config"
description: |-
  Get information about project_config
subcategory: "Projects API"
---

# ibm_project_config

Provides a read-only data source to retrieve information about a project_config. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_project_config" "project_config_instance" {
	id = ibm_project_config.project_config_instance.project_config_id
	project_id = ibm_project_config.project_config_instance.project_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `id` - (Required, Forces new resource, String) The unique config ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the project_config.
* `authorizations` - (List) The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.
Nested schema for **authorizations**:
	* `api_key` - (String) The IBM Cloud API Key.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `method` - (String) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `trusted_profile` - (List) The trusted profile for authorizations.
	Nested schema for **trusted_profile**:
		* `id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `target_iam_id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `compliance_profile` - (List) The profile required for compliance.
Nested schema for **compliance_profile**:
	* `attachment_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `instance_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `instance_location` - (String) The location of the compliance instance.
	  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
	* `profile_name` - (String) The name of the compliance profile.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.

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
	* `user_id` - (String) The unique ID of a project.
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

* `description` - (String) The description of the project configuration.
  * Constraints: The default value is ``. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.

* `input` - (List) The outputs of a Schematics template property.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **input**:
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `required` - (Boolean) Whether the variable is required or not.
	* `type` - (String) The variable type.
	  * Constraints: Allowable values are: `array`, `boolean`, `float`, `int`, `number`, `password`, `string`, `object`.
	* `value` - (String) Can be any value - a string, number, boolean, array, or object.

* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.

* `job_summary` - (List) The summaries of jobs that were performed on the configuration.
Nested schema for **job_summary**:
	* `apply_messages` - (Map) The messages of apply jobs on the configuration.
	* `apply_summary` - (Map) The summary of the apply jobs on the configuration.
	* `destroy_messages` - (Map) The messages of destroy jobs on the configuration.
	* `destroy_summary` - (Map) The summary of the destroy jobs on the configuration.
	* `message_summary` - (Map) The message summaries of jobs on the configuration.
	* `plan_messages` - (Map) The messages of plan jobs on the configuration.
	* `plan_summary` - (Map) The summary of the plan jobs on the configuration.

* `labels` - (List) A collection of configuration labels.
  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.

* `last_approved` - (List) The last approved metadata of the configuration.
Nested schema for **last_approved**:
	* `comment` - (String) The comment left by the user who approved the configuration.
	  * Constraints: The default value is ``.
	* `is_forced` - (Boolean) The flag that indicates whether the approval was forced approved.
	* `timestamp` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `user_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `last_deployment_job_summary` - (List) The summaries of jobs that were performed on the configuration.
Nested schema for **last_deployment_job_summary**:
	* `apply_messages` - (Map) The messages of apply jobs on the configuration.
	* `apply_summary` - (Map) The summary of the apply jobs on the configuration.
	* `destroy_messages` - (Map) The messages of destroy jobs on the configuration.
	* `destroy_summary` - (Map) The summary of the destroy jobs on the configuration.
	* `message_summary` - (Map) The message summaries of jobs on the configuration.
	* `plan_messages` - (Map) The messages of plan jobs on the configuration.
	* `plan_summary` - (Map) The summary of the plan jobs on the configuration.

* `last_save` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.

* `locator_id` - (String) A dotted value of catalogID.versionID.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.

* `name` - (String) The name of the configuration.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.

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

* `pipeline_state` - (String) The pipeline state of the configuration. It only exists after the first configuration validation.
  * Constraints: Allowable values are: `pipeline_failed`, `pipeline_running`, `pipeline_succeeded`.

* `project_config_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `setting` - (List) Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **setting**:
	* `name` - (String) The name of the configuration setting.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) The value of the configuration setting.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.

* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `deleted`, `deleting`, `deleting_failed`, `installed`, `installed_failed`, `installing`, `not_installed`, `uninstalling`, `uninstalling_failed`, `active`.

* `type` - (String) The type of a project configuration manual property.
  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.

* `updated_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.

* `version` - (Integer) The version of the configuration.

