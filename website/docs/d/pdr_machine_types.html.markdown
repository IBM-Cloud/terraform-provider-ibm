---
layout: "ibm"
page_title: "IBM : ibm_pdr_machine_types"
description: |-
  Get information about pdr_machine_types
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pdr_machine_types

Retrieves the list of machines types available based on the provided workspaces for the specified service instance.

## Example Usage

```hcl
data "ibm_pdr_machine_types" "pdr_machine_types" {
	accept_language = "en-US"
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
	primary_workspace_name = "Test-ws-wdc06"
	standby_workspace_name = "Test-workspace-wdc07"
}
```

### Path Parameters

* `instance_id` - (Required, Forces new resource, String) The unique identifier of the DR service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.
* `primary_workspace_name` - (Required, String) The primary Power virtual server workspace name.
  * Constraints: The maximum length is `63` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.
* `standby_workspace_name` - (Optional, String) The standby Power virtual server workspace name.
  * Constraints: The maximum length is `63` characters. The minimum length is `0` characters. The value must match regular expression `/^.*$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (String) Unique identifier of the API key.
  * Constraints: The maximum length is `512` characters. The minimum length is `20` characters. The value must match regular expression `/^crn:v1:[a-zA-Z0-9\\-_]+:public:power-dr-automation:[a-zA-Z0-9\\-_]+:[a-zA-Z0-9\\-_\/]+:[a-zA-Z0-9\\-_]+::$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `workspaces` - (Map) The Map of workspace IDs to lists of machine types.

Nested schema for **workspaces**:
	* `machine_type` - (List) The list of machine types supported by workspace.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items. Each list item has a maximum length of `32` characters. Each list item has a minimum length of `1` character. Each list item must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.

Nested schema for **workspaces**:
	* `workspace_id` - (List) The list of machine types supported by workspace.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items. Each list item has a maximum length of `32` characters. Each list item has a minimum length of `1` character. Each list item must match regular expression `/^[a-zA-Z0-9\\-_]+$/`.