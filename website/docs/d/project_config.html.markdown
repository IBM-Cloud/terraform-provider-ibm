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

* `project_config_id` - (Required, Forces new resource, String) The unique config ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the project_config.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.

* `definition` - (List) The type and output of a project configuration.
Nested schema for **definition**:
	* `authorizations` - (List) The authorization details. You can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (String) The IBM Cloud API Key.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
		* `method` - (String) The authorization method. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: Allowable values are: `api_key`, `trusted_profile`.
		* `trusted_profile_id` - (String) The trusted profile ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `compliance_profile` - (List) The profile required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_location` - (String) The location of the compliance instance.
		  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
		* `profile_name` - (String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `description` - (String) A project configuration description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `environment` - (String) The ID of the project environment.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `inputs` - (List) The input variables for configuration definition and environment.
	Nested schema for **inputs**:
	* `locator_id` - (Forces new resource, String) A dotted value of catalogID.versionID.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
	* `name` - (String) The configuration name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `settings` - (List) Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.
	Nested schema for **settings**:
	* `type` - (String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

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

* `schematics` - (List) A schematics workspace associated to a project configuration.
Nested schema for **schematics**:
	* `workspace_crn` - (String) An existing schematics workspace CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `state` - (String) The state of the configuration.
  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`.

* `update_available` - (Boolean) The flag that indicates whether a configuration update is available.

* `version` - (Integer) The version of the configuration.

