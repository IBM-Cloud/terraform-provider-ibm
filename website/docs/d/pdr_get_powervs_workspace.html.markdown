---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_powervs_workspace"
description: |-
  Get information about pdr_get_powervs_workspace
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_powervs_workspace

Provides a read-only data source to retrieve information about a pdr_get_powervs_workspace. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_get_powervs_workspace" "pdr_get_powervs_workspace" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
	location_id = "location_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `location_id` - (Required, String) Location ID value.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_powervs_workspace.
* `dr_standby_workspace_description` - (String) Description of Standby Workspace.
* `dr_standby_workspaces` - (List) List of standby disaster recovery workspaces.
Nested schema for **dr_standby_workspaces**:
	* `details` - (List) Detailed information of the standby DR workspace.
	Nested schema for **details**:
		* `crn` - (String) Cloud Resource Name (CRN) of the DR workspace.
	* `id` - (String) Unique identifier of the standby workspace.
	* `location` - (List) Location information of the standby workspace.
	Nested schema for **location**:
		* `region` - (String) The region identifier of the DR location.
		* `type` - (String) The type of location (e.g., data-center, cloud-region).
		* `url` - (String) The URL endpoint to access the DR location.
	* `name` - (String) Name of the standby workspace.
	* `status` - (String) Current status of the standby workspace.
* `dr_workspace_description` - (String) Description of Workspace.
* `dr_workspaces` - (List) List of primary disaster recovery workspaces.
Nested schema for **dr_workspaces**:
	* `default` - (Boolean) Indicates if this is the default DR workspace.
	* `details` - (List) Detailed information about the DR workspace.
	Nested schema for **details**:
		* `crn` - (String) Cloud Resource Name (CRN) of the DR workspace.
	* `id` - (String) Unique identifier of the DR workspace.
	* `location` - (List) Location information of the DR workspace.
	Nested schema for **location**:
		* `region` - (String) The region identifier of the DR location.
		* `type` - (String) The type of location (e.g., data-center, cloud-region).
		* `url` - (String) The URL endpoint to access the DR location.
	* `name` - (String) Name of the DR workspace.
	* `status` - (String) Current status of the DR workspace.

