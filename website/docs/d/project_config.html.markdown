---
layout: "ibm"
page_title: "IBM : ibm_project_config"
description: |-
  Get information about project_config
subcategory: "Projects"
---

# ibm_project_config

Provides a read-only data source to retrieve information about a project_config. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_project_config" "project_config" {
	project_config_id = ibm_project_config.project_config_instance.project_config_id
	project_id = ibm_project_config.project_config_instance.project_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `project_config_id` - (Required, Forces new resource, String) The unique configuration ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the project_config.
* `approved_version` - (List) A summary of a project configuration version.
Nested schema for **approved_version**:
	* `container_state` - (String) The aggregate state from all deployabe architectures that are included in this configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working`.
	* `definition` - (List) A summary of the definition in a project configuration version.
	Nested schema for **definition**:
		* `environment_id` - (String) The ID of the project environment.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `locator_id` - (Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `href` - (String) A Url.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
	* `version` - (Integer) The version number of the configuration.
	  * Constraints: The maximum value is `10000`. The minimum value is `0`.
* `config_error` - (List) The error from config actions.
Nested schema for **config_error**:
	* `details` - (List) The error details from config actions.
	Nested schema for **details**:
	* `message` - (String) The error message from config actions.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
* `container_state` - (String) The aggregate state from all deployabe architectures that are included in this configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working`.
* `container_state_code` - (String) Computed state code clarifying the prerequisites for validation for the configuration.
  * Constraints: Allowable values are: `awaiting_input`, `awaiting_prerequisite`, `awaiting_validation`, `awaiting_member_deployment`, `awaiting_stack_setup`.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `definition` - (List) 
Nested schema for **definition**:
	* `authorizations` - (List) The authorization details. It can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (String) The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `method` - (String) The authorization method. It can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: The maximum length is `256` characters. The minimum length is `7` characters. The value must match regular expression `/^(ref:)[a-zA-Z0-9\\$\\-_\\.+%!\\*'\\(\\),=&?\/ ]+(authorizations\/method)$|^(api_key)$|^(trusted_profile)$/`.
		* `trusted_profile_id` - (String) The trusted profile ID.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `compliance_profile` - (List) The profile that is required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (String) A unique ID for the attachment to a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `id` - (String) The unique ID for the compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `instance_id` - (String) A unique ID for the instance of a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `instance_location` - (String) The location of the compliance instance.
		  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor`.
		* `profile_name` - (String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `wp_instance_id` - (String) A unique ID for the instance of a Workload Protection.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `wp_instance_location` - (String) The location of the compliance instance.
		  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor`.
		* `wp_instance_name` - (String) The name of the Workload Protection instance.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `wp_policy_id` - (String) The unique ID for the Workload Protection policy.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `wp_policy_name` - (String) The name of the Workload Protection policy.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `wp_zone_id` - (String) A unique ID for the zone to a Workload Protection policy.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `wp_zone_name` - (String) A unique ID for the zone to a Workload Protection policy.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `description` - (String) A project configuration description.
	  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `environment_id` - (String) The ID of the project environment.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `inputs` - (Map) The input variables that are used for configuration definition and environment.
	* `locator_id` - (Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `members` - (List) The member deployabe architectures that are included in the stack.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **members**:
		* `config_id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `name` - (String) The name matching the alias in the stack definition.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `name` - (String) The configuration name. It's unique within the account across projects and regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `resource_crns` - (List) The CRNs of the resources that are associated with this configuration.
	  * Constraints: The list items must match regular expression `/^(?!\\s)(?!.*\\s$)(crn)[^'"<>{}\\s\\x00-\\x1F]*$/`. The maximum length is `110` items. The minimum length is `0` items.
	* `settings` - (Map) The Schematics environment variables to use to deploy the configuration. Settings are only available if they are specified when the configuration is initially created.
	* `uses` - (List) The depending deployabe architectures that are referenced by this configuration.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **uses**:
		* `config_id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `project_id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `deployed_version` - (List) A summary of a project configuration version.
Nested schema for **deployed_version**:
	* `container_state` - (String) The aggregate state from all deployabe architectures that are included in this configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `working`.
	* `definition` - (List) A summary of the definition in a project configuration version.
	Nested schema for **definition**:
		* `environment_id` - (String) The ID of the project environment.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `locator_id` - (Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `href` - (String) A Url.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
	* `version` - (Integer) The version number of the configuration.
	  * Constraints: The maximum value is `10000`. The minimum value is `0`.
* `deployment_model` - (String) The configuration type.
  * Constraints: Allowable values are: `project_deployed`, `user_deployed`, `stack`.
* `href` - (String) A Url.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
* `last_saved_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `member_of` - (List) The stack config parent of which this configuration is a member of.
Nested schema for **member_of**:
	* `definition` - (List) The definition summary of the stack configuration.
	Nested schema for **definition**:
		* `members` - (List) The member deployable architectures that are included in the stack.
		  * Constraints: The default value is `[]`. The maximum length is `100` items. The minimum length is `0` items.
		Nested schema for **members**:
			* `config_id` - (String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `name` - (String) The name matching the alias in the stack definition.
			  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
		* `name` - (String) The configuration name. It's unique within the account across projects and regions.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `href` - (String) A Url.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `version` - (Integer) The version of the stack configuration.
	  * Constraints: The maximum value is `10000`. The minimum value is `0`.
* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `needs_attention_state` - (List) The needs attention state of a configuration.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **needs_attention_state**:
	* `action_url` - (String) An actionable Url that users can access in response to the event. This is a system generated field. For user triggered events the field is not present.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
	* `event` - (String) The name of the event.
	  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `event_id` - (String) The id of the event.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `severity` - (String) The severity of the event. This is a system generated field. For user triggered events the field is not present.
	  * Constraints: Allowable values are: `INFO`, `WARNING`, `ERROR`.
	* `target` - (String) The configuration id and version for which the event occurred. This field is only available for user generated events. For system triggered events the field is not present.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `timestamp` - (String) The time at which the event was triggered.
	* `triggered_by` - (String) The IAM id of the user that triggered the event. This field is only available for user generated events. For system triggered events the field is not present.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `outputs` - (List) The outputs of a Schematics template property.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **outputs**:
	* `description` - (String) A short explanation of the output value.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `sensitive` - (Boolean) Flag if the output value is sensitive. If not present, the output sensitivity is indeterminate.
	* `value` - (String) This property can be any value - a string, number, boolean, array, or object.
* `project` - (List) The project that is referenced by this resource.
Nested schema for **project**:
	* `crn` - (String) An IBM Cloud resource name that uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)(crn)[^'"<>{}\\s\\x00-\\x1F]*$/`.
	* `definition` - (List) The definition of the project reference.
	Nested schema for **definition**:
		* `name` - (String) The name of the project.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A Url.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `references` - (List) The resolved references that are used by the configuration.
Nested schema for **references**:
* `schematics` - (List) A Schematics workspace that is associated to a project configuration, with scripts.
Nested schema for **schematics**:
	* `deploy_post_script` - (List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **deploy_post_script**:
		* `path` - (String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `deploy_pre_script` - (List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **deploy_pre_script**:
		* `path` - (String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `undeploy_post_script` - (List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **undeploy_post_script**:
		* `path` - (String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `undeploy_pre_script` - (List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **undeploy_pre_script**:
		* `path` - (String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `validate_post_script` - (List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **validate_post_script**:
		* `path` - (String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `validate_pre_script` - (List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **validate_pre_script**:
		* `path` - (String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `workspace_crn` - (String) An IBM Cloud resource name that uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)(crn)[^'"<>{}\\s\\x00-\\x1F]*$/`.
* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
* `state_code` - (String) Computed state code clarifying the prerequisites for validation for the configuration.
  * Constraints: Allowable values are: `awaiting_input`, `awaiting_prerequisite`, `awaiting_validation`, `awaiting_member_deployment`, `awaiting_stack_setup`.
* `template_id` - (String) The stack definition identifier.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.
* `version` - (Integer) The version of the configuration.
  * Constraints: The maximum value is `10000`. The minimum value is `0`.

