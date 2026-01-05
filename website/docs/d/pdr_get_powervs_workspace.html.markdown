---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_powervs_workspace"
description: |-
  Get information about pdr_get_powervs_workspace
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_powervs_workspace

Retrieves the power virtual server workspaces for primary and standby orchestrator based on location id.

## Example Usage

```hcl
data "ibm_pdr_get_powervs_workspace" "pdr_get_powervs_workspace" {
	instance_id = "123456d3-1122-3344-b67d-4389b44b7bf9"
	location_id = "syd04"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) ID of the service instance.
* `location_id` - (Required, String) Location ID value. You can use datsource ibm_pdr_get_dr_locations to fetch location id.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_powervs_workspace.
* `dr_standby_workspace_description` - (String) Description of Standby Workspace.
* `dr_standby_workspaces` - (List) The list of standby disaster recovery workspaces.
Nested schema for **dr_standby_workspaces**:
	* `details` - (List) The detailed information of the standby DR workspace.
	Nested schema for **details**:
		* `crn` - (String) Cloud Resource Name (CRN) of the DR workspace.
	* `id` - (String) The unique identifier of the standby workspace.
	* `location` - (List) The location information of the standby workspace.
	Nested schema for **location**:
		* `region` - (String) The region identifier of the DR location.
		* `type` - (String) The type of location (e.g., data-center, cloud-region).
		* `url` - (String) The URL endpoint to access the DR location.
	* `name` - (String) The name of the standby workspace.
	* `status` - (String) The status of the standby workspace.
* `dr_workspace_description` - (String) Description of Workspace.
* `dr_workspaces` - (List) The list of primary disaster recovery workspaces.
Nested schema for **dr_workspaces**:
	* `default` - (Boolean) Indicates if this is the default DR workspace.
	* `details` - (List) The detailed information about the DR workspace.
	Nested schema for **details**:
		* `crn` - (String) Cloud Resource Name (CRN) of the DR workspace.
	* `id` - (String) The unique identifier of the DR workspace.
	* `location` - (List) The location information of the DR workspace.
	Nested schema for **location**:
		* `region` - (String) The region identifier of the DR location.
		* `type` - (String) The type of location (e.g., data-center, cloud-region).
		* `url` - (String) The URL endpoint to access the DR location.
	* `name` - (String) The name of the DR workspace.
	* `status` - (String) The status of the DR workspace.
