---
layout: "ibm"
page_title: "IBM : ibm_project_config"
description: |-
  Get information about project_config
subcategory: "Project"
---

# ibm_project_config

Provides a read-only data source for project_config. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_project_config" "project_config" {
	id = ibm_project_config.project_config_instance.projectConfig_id
	project_id = ibm_project_config.project_config.project_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The unique config ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `version` - (Optional, String) The version of the configuration to return.
  * Constraints: The maximum length is `10` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(active|draft|\\d+)$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the project_config.
* `authorizations` - (List) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
Nested scheme for **authorizations**:
	* `api_key` - (String) The IBM Cloud API Key.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `method` - (String) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `trusted_profile` - (List) The trusted profile for authorizations.
	Nested scheme for **trusted_profile**:
		* `id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `target_iam_id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `compliance_profile` - (List) The profile required for compliance.
Nested scheme for **compliance_profile**:
	* `attachment_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `instance_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `instance_location` - (String) The location of the compliance instance.
	  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
	* `profile_name` - (String) The name of the compliance profile.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.

* `description` - (String) The project configuration description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.

* `input` - (List) The outputs of a Schematics template property.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **input**:
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `required` - (Boolean) Whether the variable is required or not.
	* `type` - (String) The variable type.
	  * Constraints: Allowable values are: `array`, `boolean`, `float`, `int`, `number`, `password`, `string`, `object`.
	* `value` - (String) Can be any value - a string, number, boolean, array, or object.

* `labels` - (List) A collection of configuration labels.
  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.

* `locator_id` - (String) A dotted value of catalogID.versionID.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.

* `metadata` - (List) The project configuration draft.
Nested scheme for **metadata**:
	* `pipeline_state` - (String) The pipeline state of the configuration.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(PIPELINE_RUNNING|PIPELINE_FAILED|PIPELINE_SUCCEEDED)$/`.
	* `project_id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `state` - (String) The state of the configuration draft.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(DISCARDED|MERGED|ACTIVE)$/`.
	* `version` - (Integer) The version number of the configuration.

* `name` - (String) The configuration name.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.

* `output` - (List) The outputs of a Schematics template property.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **output**:
	* `description` - (String) A short explanation of the output value.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `name` - (String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) Can be any value - a string, number, boolean, array, or object.

* `project_config_id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `setting` - (List) Schematics environment variables to use to deploy the configuration.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **setting**:
	* `name` - (String) The name of the configuration setting.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (String) The value of the configuration setting.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.

* `type` - (String) The type of a project configuration manual property.
  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

