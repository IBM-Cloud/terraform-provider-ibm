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
* `configs` - (List) The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.
  * Constraints: The default value is `[]`. The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **configs**:
	* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `definition` - (List) The name and description of a project configuration.
	Nested schema for **definition**:
		* `description` - (String) A project configuration description.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `name` - (String) The configuration name. It is unique within the account across projects and regions.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$/`.
	* `deployment_model` - (String) The configuration type.
	  * Constraints: Allowable values are: `project_deployed`, `user_deployed`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `project` - (List) The project referenced by this resource.
	Nested schema for **project**:
		* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
		  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"`<>{}\\s\\x00-\\x1F]*/`.
		* `definition` - (List) The definition of the project reference.
		Nested schema for **definition**:
			* `name` - (String) The name of the project.
			  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
		* `href` - (String) A URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `state` - (String) The state of the configuration.
	  * Constraints: Allowable values are: `approved`, `deleted`, `deleting`, `deleting_failed`, `discarded`, `draft`, `deployed`, `deploying_failed`, `deploying`, `superseded`, `undeploying`, `undeploying_failed`, `validated`, `validating`, `validating_failed`, `applied`, `apply_failed`.
	* `version` - (Integer) The version of the configuration.

* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.

* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"`<>{}\\s\\x00-\\x1F]*/`.

* `cumulative_needs_attention_view` - (List) The cumulative list of needs attention items for a project. If the view is successfully retrieved, an array which could be empty is returned.
  * Constraints: The default value is `[]`. The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **cumulative_needs_attention_view**:
	* `config_id` - (String) A unique ID for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `config_version` - (Integer) The version number of the configuration.
	* `event` - (String) The event name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `event_id` - (String) A unique ID for that individual event.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `cumulative_needs_attention_view_error` - (Boolean) True indicates that the fetch of the needs attention items failed. It only exists if there was an error while retrieving the cumulative needs attention view.
  * Constraints: The default value is `false`.

* `definition` - (List) The definition of the project.
Nested schema for **definition**:
	* `description` - (String) A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `destroy_on_delete` - (Boolean) The policy that indicates whether the resources are destroyed or not when a project is deleted.
	* `name` - (String) The name of the project.  It is unique within the account across regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.

* `environments` - (List) The project environments. These environments are only included in the response if project environments were created on the project.
  * Constraints: The default value is `[]`. The maximum length is `20` items. The minimum length is `0` items.
Nested schema for **environments**:
	* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `definition` - (List) The environment definition used in the project collection.
	Nested schema for **definition**:
		* `description` - (String) The description of the environment.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
		* `name` - (String) The name of the environment.  It is unique within the account across projects and regions.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The environment id as a friendly name.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `project` - (List) The project referenced by this resource.
	Nested schema for **project**:
		* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
		  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"`<>{}\\s\\x00-\\x1F]*/`.
		* `definition` - (List) The definition of the project reference.
		Nested schema for **definition**:
			* `name` - (String) The name of the project.
			  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
		* `href` - (String) A URL.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
		* `id` - (String) The unique ID.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

* `event_notifications_crn` - (String) The CRN of the event notifications instance if one is connected to this project.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `href` - (String) A URL.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.

* `location` - (Forces new resource, String) The IBM Cloud location where a resource is deployed.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.

* `resource_group` - (Forces new resource, String) The resource group name where the project's data and tools are created.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.

* `resource_group_id` - (String) The resource group id where the project's data and tools are created.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^[0-9a-zA-Z]+$/`.

* `state` - (String) The project status value.
  * Constraints: Allowable values are: `ready`, `deleting`, `deleting_failed`.

