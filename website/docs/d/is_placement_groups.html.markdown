---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Placement_Groups"
description: |-
  Get information about placement groups
---

# ibm_is_placement_groups

Retrieve information of a placement groups as a read-only data source. For more information, about placement groups, see [managing placement groups](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-placement-group&interface=ui).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_placement_groups" "example" {
}
```

## Argument reference

The following arguments are supported:


## Attribute reference

The following attributes are exported:

- `placement_groups` - Collection of placement groups. Nested `placement_groups` blocks have the following structure:
	- `created_at` - The date and time that the placement group was created.
	- `crn` - The CRN for this placement group.
	- `href` - The URL for this placement group.
	- `id` - The unique identifier for this placement group.
	- `lifecycle_state` - The lifecycle state of the placement group.
	- `name` - The user-defined name for this placement group.
	- `resource_group` - The unique identifier of the resource group for this placement group. 
	- `resource_type` - The resource type.
	- `strategy` - The strategy for this placement group- `host_spread`: place on different compute hosts- `power_spread`: place on compute hosts that use different power sourcesThe enumerated values for this property may expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the placement group on which the unexpected strategy was encountered.

- `total_count` - The total number of resources across all pages.

