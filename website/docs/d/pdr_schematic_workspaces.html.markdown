---
layout: "ibm"
page_title: "IBM : ibm_pdr_schematic_workspaces"
description: |-
  Get information about pdr_schematic_workspaces
subcategory: "DrAutomation Service"
---

# ibm_pdr_schematic_workspaces

Provides a read-only data source to retrieve information about pdr_schematic_workspaces. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_schematic_workspaces" "pdr_schematic_workspaces" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_schematic_workspaces.
* `workspaces` - (List) List of Schematics workspaces associated with the DR automation service instance.
Nested schema for **workspaces**:
	* `catalog_ref` - (List) Reference to a catalog item associated with the DR automation workspace.
	Nested schema for **catalog_ref**:
		* `item_name` - (String) Name of the catalog item that defines the resource or configuration.
	* `created_at` - (String) Timestamp when the Schematics workspace was created, in ISO 8601 format (UTC).
	* `created_by` - (String) CRN of the user or service that created the Schematics workspace.
	* `crn` - (String) Cloud Resource Name (CRN) of the Schematics workspace.
	* `description` - (String) Detailed description of the Schematics workspace.
	* `id` - (String) Unique identifier of the Schematics workspace.
	* `location` - (String) Region where the Schematics workspace is hosted.
	* `name` - (String) Human-readable name of the Schematics workspace.
	* `status` - (String) Current lifecycle status of the Schematics workspace.

