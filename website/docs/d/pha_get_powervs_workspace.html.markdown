---
layout: "ibm"
page_title: "IBM : ibm_pha_get_powervs_workspace"
description: |-
  Get information about pha_get_powervs_workspace
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_get_powervs_workspace

Provides a read-only data source to retrieve information about a pha_get_powervs_workspace. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_get_powervs_workspace" "pha_get_powervs_workspace" {
	accept_language = "en-US"
	if_none_match = "abcdef"
	location_id = "us-south"
	pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `location_id` - (Required, String) Location ID value.
  * Constraints: The maximum length is `20` characters. The minimum length is `5` characters. The value must match regular expression `/^[a-z]{2}-[a-z]+(-[0-9]+)?$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_get_powervs_workspace.
* `workspaces` - (List) Array of workspace summaries within the region.
  * Constraints: The maximum length is `16` items. The minimum length is `0` items.
Nested schema for **workspaces**:
	* `id` - (String) Unique identifier of the workspace.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `name` - (String) Name of the workspace.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

