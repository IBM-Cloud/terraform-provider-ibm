---
layout: "ibm"
page_title: "IBM : ibm_project_event_notification"
description: |-
  Get information about Project definition
subcategory: "Project"
---

# ibm_project_event_notification

Provides a read-only data source for Project definition. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_project_event_notification" "project_event_notification" {
	project_id = "072b70cb-4db7-4c5c-bf1b-8d93d422537d"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the Project definition.
* `configs` - (List) The project configurations.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **configs**:
	* `description` - (String) The project configuration description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `id` - (String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
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
	* `setting` - (List) Schematics environment variables to use to deploy the configuration.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **setting**:
		* `name` - (String) The name of the configuration setting.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (String) The value of the configuration setting.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
	* `type` - (String) The type of a project configuration manual property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `description` - (String) A project descriptive text.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.

* `metadata` - (List) The metadata of the project.
Nested scheme for **metadata**:
	* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.
	* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `cumulative_needs_attention_view` - (List) The cumulative list of needs attention items for a project.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **cumulative_needs_attention_view**:
		* `config_id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `config_version` - (Integer) The version number of the configuration.
		* `event` - (String) The event name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
		* `event_id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `cumulative_needs_attention_view_err` - (String) \"True\" indicates that the fetch of the needs attention items failed.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `event_notifications_crn` - (String) The CRN of the event notifications instance if one is connected to this project.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `location` - (String) The IBM Cloud location where a resource is deployed.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `resource_group` - (String) The resource group where the project's data and tools are created.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.
	* `state` - (String) The project status value.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(CREATING|CREATING_FAILED|UPDATING|UPDATING_FAILED|READY)$/`.

* `name` - (String) The project name.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.

