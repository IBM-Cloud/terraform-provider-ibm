---
layout: "ibm"
page_title: "IBM : ibm_project"
description: |-
  Get information about Project definition
subcategory: "Project"
---

# ibm_project

Provides a read-only data source for Project definition. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_project" "project" {
	id = ibm_project.project_instance.project_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the project.
* `description` - (String) A project descriptive text.
  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s).*\\S$/`.

* `destroy_on_delete` - (Boolean) The policy that indicates whether the resources are destroyed or not when a project is deleted.
  * Constraints: The default value is `true`.

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
	* `cumulative_needs_attention_view_err` - (String) True indicates that the fetch of the needs attention items failed.
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

