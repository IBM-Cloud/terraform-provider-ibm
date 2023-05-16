---
layout: "ibm"
page_title: "IBM : ibm_project_instance"
description: |-
  Manages Project definition.
subcategory: "Project"
---

# ibm_project_instance

Provides a resource for Project definition. This allows Project definition to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_project" "project_instance" {
  configs {
		id = "0013790d-6cb5-4adc-8927-a725a1261d0c"
		name = "static-website-dev"
		labels = [ "env:dev", "billing:internal" ]
		description = "Website - development"
		locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.145be7c1-9ec4-4719-b586-584ee52fbed0-global"
		input {
			name = "app_repo_name"
		}
		setting {
			name = "app_repo_name"
			value = "static-website-dev-app-repo"
		}
  }
  description = "Sample static website test using the IBM catalog deployable architecture"
  name = "My static website"
  location = "us-south"
  resource_group = "Default"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `configs` - (Optional, List) The project configurations.
  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
Nested scheme for **configs**:
	* `description` - (Optional, String) The project configuration description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
	* `id` - (Optional, String) The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
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
	* `setting` - (Optional, List) Schematics environment variables to use to deploy the configuration.
	  * Constraints: The maximum length is `10000` items. The minimum length is `0` items.
	Nested scheme for **setting**:
		* `name` - (Required, String) The name of the configuration setting.
		  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
		* `value` - (Required, String) The value of the configuration setting.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
* `description` - (Optional, String) A project's descriptive text.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.
* `location` - (Required, String) The location where the project's data and tools are created.
  * Constraints: The maximum length is `12` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(us-south|us-east|eu-gb|eu-de)$/`.
* `name` - (Required, String) The project name.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]+$/`.
* `resource_group` - (Required, String) The resource group where the project's data and tools are created.
  * Constraints: The maximum length is `40` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^'"`<>{}\\x00-\\x1F]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the Project definition.
* `crn` - (String) An IBM Cloud resource name, which uniquely identifies a resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9\\-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
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

## Import

You can import the `ibm_project_instance` resource by using `id`. The unique ID of a project.

# Syntax
```
$ terraform import ibm_project_instance.project_instance <id>
```
