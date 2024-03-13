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
		* `api_key` - (Optional, String) The IBM Cloud API Key.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `method` - (Optional, String) The authorization method. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: Allowable values are: `api_key`, `trusted_profile`.
		* `trusted_profile_id` - (Optional, String) The trusted profile ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `compliance_profile` - (Optional, List) The profile required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (Optional, String) A unique ID for the attachment to a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `id` - (Optional, String) The unique ID for that compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_id` - (Optional, String) A unique ID for an instance of a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_location` - (Optional, String) The location of the compliance instance.
		  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
		* `profile_name` - (Optional, String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `description` - (Optional, String) A project configuration description.
	  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `environment_id` - (Optional, String) The ID of the project environment.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `inputs` - (Optional, Map) The input variables for configuration definition and environment.
	* `locator_id` - (Optional, Forces new resource, String) A unique concatenation of catalogID.versionID that identifies the DA in the catalog. Either schematics.workspace_crn, definition.locator_id, or both must be specified.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (Required, String) The configuration name. It is unique within the account across projects and regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `resource_crns` - (Optional, List) The CRNs of resources associated with this configuration.
	  * Constraints: The list items must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"<>{}\\s\\x00-\\x1F]*/`. The maximum length is `110` items. The minimum length is `0` items.
	* `settings` - (Optional, Map) Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `schematics` - (Optional, List) A schematics workspace associated to a project configuration, with scripts.
Nested schema for **schematics**:
	* `deploy_post_script` - (Optional, List) A script to be run as part of a Project configuration, for a given stage (pre, post) and action (validate, deploy, undeploy).
	Nested schema for **deploy_post_script**:
		* `path` - (Computed, String) The path to this script within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `deploy_pre_script` - (Optional, List) A script to be run as part of a Project configuration, for a given stage (pre, post) and action (validate, deploy, undeploy).
	Nested schema for **deploy_pre_script**:
		* `path` - (Computed, String) The path to this script within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `undeploy_post_script` - (Optional, List) A script to be run as part of a Project configuration, for a given stage (pre, post) and action (validate, deploy, undeploy).
	Nested schema for **undeploy_post_script**:
		* `path` - (Computed, String) The path to this script within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `undeploy_pre_script` - (Optional, List) A script to be run as part of a Project configuration, for a given stage (pre, post) and action (validate, deploy, undeploy).
	Nested schema for **undeploy_pre_script**:
		* `path` - (Computed, String) The path to this script within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `validate_post_script` - (Optional, List) A script to be run as part of a Project configuration, for a given stage (pre, post) and action (validate, deploy, undeploy).
	Nested schema for **validate_post_script**:
		* `path` - (Computed, String) The path to this script within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `validate_pre_script` - (Optional, List) A script to be run as part of a Project configuration, for a given stage (pre, post) and action (validate, deploy, undeploy).
	Nested schema for **validate_pre_script**:
		* `path` - (Computed, String) The path to this script within the current version source.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^\\S+$/`.
		* `short_description` - (Computed, String) The short description for this script.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `type` - (Computed, String) The type of the script.
		  * Constraints: The maximum length is `7` characters. The minimum length is `7` characters. The value must match regular expression `/^(ansible)$/`.
	* `workspace_crn` - (Optional, String) An IBM Cloud resource name, which uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"<>{}\\s\\x00-\\x1F]*/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project_config.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `href` - (String) A URL.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
* `last_saved_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `needs_attention_state` - (List) The needs attention state of a configuration.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
* `outputs` - (List) The outputs of a Schematics template property.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **outputs**:
	* `description` - (String) A short explanation of the output value.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (Map) Can be any value - a string, number, boolean, array, or object.
* `project` - (List) The project referenced by this resource.
Nested schema for **project**:
	* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
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
