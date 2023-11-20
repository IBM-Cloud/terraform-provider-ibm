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
    labels = [ "env:dev", "billing:internal" ]
    description = "Website - development"
    authorizations {
      method = "api_key"
      api_key = "<your_apikey_here>"
    }
    locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.145be7c1-9ec4-4719-b586-584ee52fbed0-global"
    input {
      name = "app_repo_name"
    }
    setting {
      name = "app_repo_name"
      value = "static-website-dev-app-repo"
    }
  }
  project_id = ibm_project.project_instance.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `definition` - (Required, List) The type and output of a project configuration.
Nested schema for **definition**:
	* `authorizations` - (Optional, List) The authorization details. You can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (Optional, String) The IBM Cloud API Key.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
		* `method` - (Optional, String) The authorization method. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: Allowable values are: `api_key`, `trusted_profile`.
		* `trusted_profile_id` - (Optional, String) The trusted profile ID.
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
	* `description` - (Optional, String) A project configuration description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `environment` - (Optional, String) The ID of the project environment.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `inputs` - (Optional, List) The input variables for configuration definition and environment.
	Nested schema for **inputs**:
	* `locator_id` - (Required, Forces new resource, String) A dotted value of catalogID.versionID.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (Required, String) The configuration name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `settings` - (Optional, List) Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.
	Nested schema for **settings**:
	* `type` - (Computed, String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `schematics` - (Optional, List) A schematics workspace associated to a project configuration.
Nested schema for **schematics**:
	* `workspace_crn` - (Optional, String) An existing schematics workspace CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project_config.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `is_draft` - (Boolean) The flag that indicates whether the version of the configuration is draft, or active.
* `last_saved_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
* `needs_attention_state` - (List) The needs attention state of a configuration.
  * Constraints: The default value is `[]`. The maximum length is `10000` items. The minimum length is `0` items.
* `outputs` - (List) The outputs of a Schematics template property.
  * Constraints: The default value is `[]`. The maximum length is `10000` items. The minimum length is `0` items.
Nested schema for **outputs**:
	* `description` - (String) A short explanation of the output value.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) Can be any value - a string, number, boolean, array, or object.
* `project` - (List) The project referenced by this resource.
Nested schema for **project**:
	* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `definition` - (List) The definition of the project reference.
	Nested schema for **definition**:
		* `name` - (String) The name of the project.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_config_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`.
* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.
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
