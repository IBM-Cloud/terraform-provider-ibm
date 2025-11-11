---
layout: "ibm"
page_title: "IBM : ibm_pdr_validate_workspace"
description: |-
  Get information about pdr_validate_workspace
subcategory: "DrAutomation Service"
---

# ibm_pdr_validate_workspace

Provides a read-only data source to retrieve information about a pdr_validate_workspace. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_validate_workspace" "pdr_validate_workspace" {
	crn = "crn:v1:bluemix:public:power-iaas:dal10:a/094f4214c75941f991da601b001df1fe:75cbf05b-78f6-406e-afe7-a904f646d798::"
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
	location_url = "https://us-south.power-iaas.cloud.ibm.com"
	workspace_id = "75cbf05b-78f6-406e-afe7-a904f646d798"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `crn` - (Required, String) crn value.
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `location_url` - (Required, String) schematic_workspace_id value.
* `workspace_id` - (Required, String) standBy workspaceID value.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_validate_workspace.
* `description` - (String) Human-readable message describing the validation result.
* `status` - (String) Status of the workspace validation (for example, Valid, Invalid, or Pending).

