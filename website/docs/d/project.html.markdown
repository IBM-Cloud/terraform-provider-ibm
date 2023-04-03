---
layout: "ibm"
page_title: "IBM : ibm_project"
description: |-
  Get information about project
subcategory: "Projects API Specification"
---

# ibm_project

Provides a read-only data source for project. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_project" "project" {
	id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `complete` - (Optional, Boolean) The flag to determine if full metadata should be returned.
  * Constraints: The default value is `false`.
* `exclude_configs` - (Optional, Boolean) Only return with the active configuration, no drafts.
  * Constraints: The default value is `false`.
* `id` - (Required, Forces new resource, String) The ID of the project, which uniquely identifies it.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the project.
* `configs` - (List) The project configurations.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **configs**:
	* `description` - (String) A project configuration description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s).*\\S$/`.
	* `id` - (String) The unique ID of a project.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `input` - (List) The outputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **input**:
		* `name` - (String) The variable name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.
		* `required` - (Boolean) Whether the variable is required or not.
		* `type` - (String) The variable type.
		  * Constraints: Allowable values are: `array`, `boolean`, `float`, `int`, `number`, `password`, `string`, `object`.
	* `labels` - (List) A collection of configuration labels.
	  * Constraints: The list items must match regular expression `/^[_\\-a-z0-9:\/=]+$/`. The maximum length is `10000` items. The minimum length is `0` items.
	* `locator_id` - (String) The location ID of a Project configuration manual property.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.
	* `name` - (String) The configuration name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.
	* `output` - (List) The outputs of a Schematics template property.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **output**:
		* `description` - (String) A short explanation of the output value.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
		* `name` - (String) The variable name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.
		* `value` - (List) The output value.
		  * Constraints: The list items must match regular expression `/^(?!\\s).+\\S$/`. The maximum length is `10000` items. The minimum length is `0` items.
	* `setting` - (List) An optional setting object That is passed to the cart API.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **setting**:
		* `name` - (String) The name of the configuration setting.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.
		* `value` - (String) The value of a the configuration setting.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.
	* `type` - (String) The type of a Project Config Manual Property.
	  * Constraints: Allowable values are: `terraform_template`, `schematics_blueprint`.

* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

* `description` - (String) A project descriptive text.

* `metadata` - (List) Metadata of the project.
Nested scheme for **metadata**:
	* `created_at` - (String) A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339.
	* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `cumulative_needs_attention_view` - (List) The cumulative list of needs attention items of a project.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **cumulative_needs_attention_view**:
		* `config_id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `config_version` - (Integer) The version number of the configuration.
		* `event` - (String) The event name.
		* `event_id` - (String) The unique ID of a project.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `cumulative_needs_attention_view_err` - (String) True to indicate the fetch of needs attention items that failed.
	* `event_notifications_crn` - (String) The CRN of the event notifications instance if one is connected to this project.
	* `location` - (String) The location of where the project was created.
	* `resource_group` - (String) The resource group of where the project was created.
	* `state` - (String) The project status value.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(CREATING|CREATING_FAILED|UPDATING|UPDATING_FAILED|READY)$/`.

* `name` - (String) The project name.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s).+\\S$/`.

