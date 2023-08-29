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
data "ibm_project_config" "project_config" {
	id = ibm_project_config.project_config_instance.projectConfigCanonical_id
	project_id = ibm_project_config.project_config.project_id
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
* `active_draft` - (List) The project configuration version.
Nested schema for **active_draft**:
	* `href` - (String) A relative URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `pipeline_state` - (String) The pipeline state of the configuration. It only exists after the first configuration validation.
	  * Constraints: Allowable values are: `pipeline_failed`, `pipeline_running`, `pipeline_succeeded`.
	* `state` - (String) The state of the configuration draft.
	  * Constraints: Allowable values are: `discarded`, `merged`, `active`.
	* `version` - (Integer) The version number of the configuration.

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
	  * Constraints: The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
	* `profile_name` - (String) The name of the compliance profile.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.

* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.

* `definition` - (List) The project configuration definition.
Nested schema for **definition**:
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
		  * Constraints: The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
		* `profile_name` - (String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `description` - (String) The description of the project configuration.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `input` - (List) The inputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **input**:
		* `name` - (String) The variable name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (String) Can be any value - a string, number, boolean, array, or object.
	* `labels` - (List) A collection of configuration labels.
	  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
	* `locator_id` - (String) A dotted value of catalogID.versionID.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (String) The name of the configuration.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `output` - (List) The outputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **output**:
		* `description` - (String) A short explanation of the output value.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
		* `name` - (String) The variable name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (String) Can be any value - a string, number, boolean, array, or object.
	* `setting` - (List) Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested schema for **setting**:
		* `name` - (String) The name of the configuration setting.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (String) The value of the configuration setting.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `type` - (String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

* `description` - (String) The description of the project configuration.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.

* `href` - (String) A relative URL.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.

* `input` - (List) The inputs of a Schematics template property.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **input**:
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) Can be any value - a string, number, boolean, array, or object.

* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.

* `labels` - (List) A collection of configuration labels.
  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.

* `last_approved` - (List) The last approved metadata of the configuration.
Nested schema for **last_approved**:
	* `comment` - (String) The comment left by the user who approved the configuration.
	* `is_forced` - (Boolean) The flag that indicates whether the approval was forced approved.
	* `timestamp` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `user_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

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

* `project_config_canonical_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `setting` - (List) Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **setting**:
	* `name` - (String) The name of the configuration setting.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) The value of the configuration setting.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.

* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `installed`, `installed_failed`, `installing`, `superceded`, `uninstalling`, `uninstalling_failed`, `validated`, `validating`, `validating_failed`.

* `type` - (String) The type of a project configuration manual property.
  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.

* `updated_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.

* `version` - (Integer) The version of the configuration.

