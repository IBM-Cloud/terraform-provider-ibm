---
layout: "ibm"
page_title: "IBM : ibm_project_config"
description: |-
  Manages project_config.
subcategory: "Projects"
---

# ibm_project_config

Create, update, and delete project_configs with this resource.

## Example Usage

```hcl
resource "ibm_project_config" "project_config_instance" {
  definition {
    name = "static-website-dev"
    description = "Website - development"
    authorizations {
      method = "api_key"
      api_key = "<your_apikey_here>"
    }
    locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.145be7c1-9ec4-4719-b586-584ee52fbed0-global"
    inputs = {
      app_repo_name = "static-website-repo"
    }
  }
  project_id = ibm_project.project_instance.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `definition` - (Required, List) 
Nested schema for **definition**:
	* `authorizations` - (Optional, List) The authorization details. You can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (Optional, String) The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `method` - (Optional, String) The authorization method. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: Allowable values are: `api_key`, `trusted_profile`.
		* `trusted_profile_id` - (Optional, String) The trusted profile ID.
		  * Constraints: The maximum length is `512` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `compliance_profile` - (Optional, List) The profile that is required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (Optional, String) A unique ID for the attachment to a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `id` - (Optional, String) The unique ID for the compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_id` - (Optional, String) A unique ID for the instance of a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_location` - (Optional, String) The location of the compliance instance.
		  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor`.
		* `profile_name` - (Optional, String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `description` - (Optional, String) A project configuration description.
	  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `environment_id` - (Optional, String) The ID of the project environment.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `inputs` - (Optional, Map) The input variables that are used for configuration definition and environment.
	* `locator_id` - (Optional, Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. I you're importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If you're using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing your Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `members` - (Optional, List) The member deployabe architectures that are included in your stack.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **members**:
		* `config_id` - (Required, String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `name` - (Required, String) The name matching the alias in the stack definition.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `name` - (Optional, String) The configuration name. It's unique within the account across projects and regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `resource_crns` - (Optional, List) The CRNs of the resources that are associated with this configuration.
	  * Constraints: The list items must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"<>{}\\s\\x00-\\x1F]*/`. The maximum length is `110` items. The minimum length is `0` items.
	* `settings` - (Optional, Map) The Schematics environment variables to use to deploy the configuration. Settings are only available if they are specified when the configuration is initially created.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `schematics` - (Optional, List) A Schematics workspace that is associated to a project configuration, with scripts.
Nested schema for **schematics**:
	* `deploy_post_script` - (Optional, List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **deploy_post_script**:
		* `path` - (Computed, String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `deploy_pre_script` - (Optional, List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **deploy_pre_script**:
		* `path` - (Computed, String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `undeploy_post_script` - (Optional, List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **undeploy_post_script**:
		* `path` - (Computed, String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `undeploy_pre_script` - (Optional, List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **undeploy_pre_script**:
		* `path` - (Computed, String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `validate_post_script` - (Optional, List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **validate_post_script**:
		* `path` - (Computed, String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `validate_pre_script` - (Optional, List) A script to be run as part of a project configuration for a specific stage (pre or post) and action (validate, deploy, or undeploy).
	Nested schema for **validate_pre_script**:
		* `path` - (Computed, String) The path to this script is within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `workspace_crn` - (Optional, String) An IBM Cloud resource name that uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"<>{}\\s\\x00-\\x1F]*/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project_config.
* `approved_version` - (List) A summary of a project configuration version.
Nested schema for **approved_version**:
	* `definition` - (List) A summary of the definition in a project configuration version.
	Nested schema for **definition**:
		* `environment_id` - (String) The ID of the project environment.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `locator_id` - (Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. I you're importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If you're using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing your Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
	* `state_code` - (String) Computed state code clarifying the prerequisites for validation for the configuration.
	  * Constraints: Allowable values are: `awaiting_input`, `awaiting_prerequisite`, `awaiting_validation`, `awaiting_member_deployment`, `awaiting_stack_setup`.
	* `version` - (Integer) The version number of the configuration.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `deployed_version` - (List) A summary of a project configuration version.
Nested schema for **deployed_version**:
	* `definition` - (List) A summary of the definition in a project configuration version.
	Nested schema for **definition**:
		* `environment_id` - (String) The ID of the project environment.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `locator_id` - (Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. I you're importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If you're using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing your Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
	* `state_code` - (String) Computed state code clarifying the prerequisites for validation for the configuration.
	  * Constraints: Allowable values are: `awaiting_input`, `awaiting_prerequisite`, `awaiting_validation`, `awaiting_member_deployment`, `awaiting_stack_setup`.
	* `version` - (Integer) The version number of the configuration.
* `deployment_model` - (String) The configuration type.
  * Constraints: Allowable values are: `project_deployed`, `user_deployed`, `stack`.
* `href` - (String) A URL.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
* `last_saved_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `member_of` - (List) The stack config parent of which this configuration is a member of.
Nested schema for **member_of**:
	* `definition` - (List) The definition summary of the stack configuration.
	Nested schema for **definition**:
		* `members` - (List) The member deployabe architectures that are included in your stack.
		  * Constraints: The default value is `[]`. The maximum length is `100` items. The minimum length is `0` items.
		Nested schema for **members**:
			* `config_id` - (String) The unique ID.
			  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
			* `name` - (String) The name matching the alias in the stack definition.
			  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
		* `name` - (String) The configuration name. It's unique within the account across projects and regions.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `version` - (Integer) The version of the stack configuration.
* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `needs_attention_state` - (List) The needs attention state of a configuration.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **needs_attention_state**:
	* `action_url` - (String) An actionable URL that users can access in response to the event. This is a system generated field. For user triggered events the field is not present.
	* `event` - (String) The name of the event.
	* `event_id` - (String) The id of the event.
	* `severity` - (String) The severity of the event. This is a system generated field. For user triggered events the field is not present.
	  * Constraints: Allowable values are: `INFO`, `WARNING`, `ERROR`.
	* `target` - (String) The configuration id and version for which the event occurred. This field is only available for user generated events. For system triggered events the field is not present.
	* `timestamp` - (String) The timestamp of the event.
	* `triggered_by` - (String) The IAM id of the user that triggered the event. This field is only available for user generated events. For system triggered events the field is not present.
* `outputs` - (List) The outputs of a Schematics template property.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **outputs**:
	* `description` - (String) A short explanation of the output value.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) This property can be any value - a string, number, boolean, array, or object.
* `project` - (List) The project that is referenced by this resource.
Nested schema for **project**:
	* `crn` - (String) An IBM Cloud resource name that uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"<>{}\\s\\x00-\\x1F]*/`.
	* `definition` - (List) The definition of the project reference.
	Nested schema for **definition**:
		* `name` - (String) The name of the project.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_config_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
* `state_code` - (String) Computed state code clarifying the prerequisites for validation for the configuration.
  * Constraints: Allowable values are: `awaiting_input`, `awaiting_prerequisite`, `awaiting_validation`, `awaiting_member_deployment`, `awaiting_stack_setup`.
* `template_id` - (String) The stack definition identifier.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.
* `version` - (Integer) The version of the configuration.


## Import

You can import the `ibm_project_config` resource by using `id`.
The `id` property can be formed from `project_id`, and `project_config_id` in the following format:

<pre>
&lt;project_id&gt;/&lt;project_config_id&gt;
</pre>
* `project_id`: A string. The unique project ID.
* `project_config_id`: A string. The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.

# Syntax
<pre>
$ terraform import ibm_project_config.project_config &lt;project_id&gt;/&lt;project_config_id&gt;
</pre>
