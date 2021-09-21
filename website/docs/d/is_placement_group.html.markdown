---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Placement_Group"
description: |-
  Get information about PlacementGroup
---

# ibm_is_placement_group

Provides a read-only data source for PlacementGroup. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "is_placement_group" "is_placement_group" {
	id = "id"
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required, String) The placement group identifier.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the PlacementGroup.
- `created_at` - The date and time that the placement group was created.
- `crn` - The CRN for this placement group.
- `href` - The URL for this placement group.
- `lifecycle_state` - The lifecycle state of the placement group.
- `name` - The user-defined name for this placement group.
- `resource_group` - The unique identifier of this resource group for this placement group. 
- `resource_type` - The resource type.
- `strategy` - The strategy for this placement group- `host_spread`: place on different compute hosts- `power_spread`: place on compute hosts that use different power sourcesThe enumerated values for this property may expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the placement group on which the unexpected strategy was encountered.
- `access_tags`  - (String) Access management tags associated to the placement group.
- `tags`  - (String) Usertags associated to the placement group.

