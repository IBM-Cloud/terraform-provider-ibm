---
layout: "ibm"
page_title: "IBM : ibm_project"
description: |-
  Get information about project
subcategory: "Projects"
---

# ibm_project

Provides a read-only data source to retrieve information about a project. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_project" "project" {
	project_id = ibm_project.project_instance.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the project.
* `configs` - (List) The project configurations. These configurations are only included in the response of creating a project if a configuration array is specified in the request payload.
  * Constraints: The default value is `[]`. The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **configs**:
	* `approved_version` - (List) A summary of a project configuration version.
	Nested schema for **approved_version**:
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
	* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
	* `definition` - (List) The description of a project configuration.
	Nested schema for **definition**:
		* `description` - (String) A project configuration description.
		  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `locator_id` - (Forces new resource, String) A unique concatenation of the catalog ID and the version ID that identify the deployable architecture in the catalog. If importing from an existing Schematics workspace that is not backed by cart, a `locator_id` is required. If using a Schematics workspace that is backed by cart, a `locator_id` is not necessary because the Schematics workspace has one.> There are 3 scenarios:> 1. If only a `locator_id` is specified, a new Schematics workspace is instantiated with that `locator_id`.> 2. If only a schematics `workspace_crn` is specified, a `400` is returned if a `locator_id` is not found in the existing schematics workspace.> 3. If both a Schematics `workspace_crn` and a `locator_id` are specified, a `400` message is returned if the specified `locator_id` does not agree with the `locator_id` in the existing Schematics workspace.> For more information of creating a Schematics workspace, see [Creating workspaces and importing the Terraform template](/docs/schematics?topic=schematics-sch-create-wks).
		  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[\\.0-9a-z-A-Z_-]+$/`.
		* `name` - (String) The configuration name. It's unique within the account across projects and regions.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `deployed_version` - (List) A summary of a project configuration version.
	Nested schema for **deployed_version**:
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
	* `id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
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
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
	* `version` - (Integer) The version of the configuration.
	  * Constraints: The maximum value is `10000`. The minimum value is `0`.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `crn` - (String) An IBM Cloud resource name that uniquely identifies a resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)(crn)[^'"<>{}\\s\\x00-\\x1F]*$/`.
* `cumulative_needs_attention_view` - (List) The cumulative list of needs attention items for a project. If the view is successfully retrieved, an empty or nonempty array is returned.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **cumulative_needs_attention_view**:
	* `config_id` - (String) A unique ID for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `config_version` - (Integer) The version number of the configuration.
	  * Constraints: The maximum value is `10000`. The minimum value is `0`.
	* `event` - (String) The event name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]*$/`.
	* `event_id` - (String) A unique ID for this individual event.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `cumulative_needs_attention_view_error` - (Boolean) A value of `true` indicates that the fetch of the needs attention items failed. This property only exists if there was an error while retrieving the cumulative needs attention view.
  * Constraints: The default value is `false`.
* `definition` - (List) The definition of the project.
Nested schema for **definition**:
	* `auto_deploy` - (Boolean) A boolean flag to enable deploying configurations automatically.
	  * Constraints: The default value is `false`.
	* `auto_deploy_mode` - (String) This is an advanced setting to auto deploy to tell how auto deploy should behave when it is enabled. There are 2 options:> 1. `auto_approval` will automatically approve the configuration after validated without user confirmation.> 2. `manual_approval` will require user confirmation to approve the configuration after validated before deploying the configuration starts.
	  * Constraints: The default value is `manual_approval`. Allowable values are: `auto_approval`, `manual_approval`.
	* `description` - (String) A brief explanation of the project's use in the configuration of a deployable architecture. A project can be created without providing a description.
	  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `destroy_on_delete` - (Boolean) The policy that indicates whether the resources are undeployed or not when a project is deleted.
	  * Constraints: The default value is `true`.
	* `monitoring_enabled` - (Boolean) A boolean flag to enable automatic drift detection. Use this field to run a daily check to compare the configurations to those deployed resources to detect any difference.
	  * Constraints: The default value is `false`.
	* `name` - (String) The name of the project.  It's unique within the account across regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"\/<>{}\\x00-\\x1F]+$/`.
	* `store` - (List) The details required to custom store project configs.
	Nested schema for **store**:
		* `config_directory` - (String) The directory where project configs are stored.
		  * Constraints: The default value is `''`. The maximum length is `255` characters. The minimum length is `0` characters. The value must match regular expression `/^\/?[^\/]*(?:\/[^\/]*)*\/?$/`.
		* `token` - (String) The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `type` - (String) The type of store used for the project.
		  * Constraints: The default value is `gh`. Allowable values are: `gh`, `ghe`, `gitlab`.
		* `url` - (String) A Url.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `terraform_engine` - (List) Experimental schema - this is for prototyping purposes.
	Nested schema for **terraform_engine**:
		* `id` - (String) The identifier of the Terraform engine.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
		* `type` - (String) The type of the engine.
		  * Constraints: Allowable values are: `terraform-enterprise`, `schematics`.
* `environments` - (List) The project environment. These environments are only included in the response if project environments were created on the project.
  * Constraints: The default value is `[]`. The maximum length is `20` items. The minimum length is `0` items.
Nested schema for **environments**:
	* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
	* `definition` - (List) The environment definition that is used in the project collection.
	Nested schema for **definition**:
		* `description` - (String) The description of the environment.
		  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `name` - (String) The name of the environment. It's unique within the account across projects and regions.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A Url.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `id` - (String) The environment ID as a friendly name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
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
* `event_notifications_crn` - (String) The CRN of the Event Notifications instance if one is connected to this project.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `href` - (String) A Url.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
* `location` - (Forces new resource, String) The IBM Cloud location where a resource is deployed.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]*$/`.
* `resource_group` - (Forces new resource, String) The resource group name where the project's data and tools are created.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]*$/`.
* `resource_group_id` - (String) The resource group ID where the project's data and tools are created.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^[0-9a-zA-Z]+$/`.
* `state` - (String) The project status value.
  * Constraints: Allowable values are: `ready`, `deleting`, `deleting_failed`.

