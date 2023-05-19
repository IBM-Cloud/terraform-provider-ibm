---
layout: "ibm"
page_title: "IBM : ibm_project_config"
description: |-
  Manages project_config.
subcategory: "Projects API"
---

# ibm_project_config

Provides a resource for project_config. This allows project_config to be created, updated and deleted.

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
  input {
		name = "name"
		value = "anything as a string"
  }
  locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.018edf04-e772-4ca2-9785-03e8e03bef72-global"
  name = "env-stage"
  project_id = ibm_project.project_instance.id
  setting {
		name = "name"
		value = "value"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `authorizations` - (Optional, List) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
Nested scheme for **authorizations**:
	* `api_key` - (Optional, String) The IBM Cloud API Key.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
	* `method` - (Optional, String) The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `trusted_profile` - (Optional, List) The trusted profile for authorizations.
	Nested scheme for **trusted_profile**:
		* `id` - (Optional, String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `target_iam_id` - (Optional, String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `compliance_profile` - (Optional, List) The profile required for compliance.
Nested scheme for **compliance_profile**:
	* `attachment_id` - (Optional, String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `id` - (Optional, String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `instance_id` - (Optional, String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `instance_location` - (Optional, String) The location of the compliance instance.
	  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
	* `profile_name` - (Optional, String) The name of the compliance profile.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^`<>\\x00-\\x1F]*$/`.
* `description` - (Optional, String) The project configuration description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
* `input` - (Optional, List) The input values to use to deploy the configuration.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **input**:
	* `name` - (Required, String) The variable name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (Optional, String) Can be any value - a string, number, boolean, array, or object.
* `labels` - (Optional, List) A collection of configuration labels.
  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
* `locator_id` - (Required, String) A dotted value of catalogID.versionID.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
* `name` - (Required, String) The configuration name.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **setting**:
	* `name` - (Required, String) The name of the configuration setting.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `value` - (Required, String) The value of the configuration setting.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the project_config.
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
* `type` - (String) The type of a project configuration manual property.
  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.


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
