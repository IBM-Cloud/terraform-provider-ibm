---
layout: "ibm"
page_title: "IBM : ibm_pha_powervs_workspace"
description: |-
  Get information about pha_powervs_workspace
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_powervs_workspace

Retrieves the power virtual server workspaces for primary and standby vms based on location id.

## Example Usage

```hcl
data "ibm_pha_powervs_workspace" "pha_powervs_workspace" {
	accept_language = "en-US"
	if_none_match = "abcdef"
	instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
	location_id = "us-south"
}
```
### Path Parameters

* `pha_instance_id` - (Required, Forces new resource, String) The unique identifier of the powerha service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `6` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `location_id` - (Required, String) Location ID value.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^.+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_powervs_workspace.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `workspaces` - (List) Array of workspace summaries within the region.
  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
Nested schema for **workspaces**:
	* `id` - (String) Unique identifier of the workspace.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
	* `name` - (String) Name of the workspace.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

