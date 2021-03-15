---
layout: "ibm"
page_title: "IBM : is_dedicated_host_group"
sidebar_current: "docs-ibm-resource-is-dedicated-host-group"
description: |-
  Manages DedicatedHostGroup.
---

# ibm\_is_dedicated_host_group

Provides a resource for DedicatedHostGroup. This allows DedicatedHostGroup to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_is_dedicated_host_group" "is_dedicated_host_group" {
  class = "mx2"
  family = "balanced"
  zone = "us-south-1"
  name = "dh-group-name"
}
```

## Argument Reference

The following arguments are supported:

* `class` - (Required, string) The dedicated host profile class for hosts in this group.
* `family` - (Required, string) The dedicated host profile family for hosts in this group.
* `name` - (Optional, string) The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.
* `resource_group` - (Optional, string) The unique identifier of the resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
* `zone` - (Required, List) The globally unique name of the zone this dedicated host group will reside in.

## Attribute Reference

The following attributes are exported:

* `class` - The dedicated host profile class for hosts in this group.
* `family` - The dedicated host profile family for hosts in this group.
* `id` - The unique identifier of the DedicatedHostGroup.
* `href` - The URL for this dedicated host group.
* `crn` - The CRN for this dedicated host group.
* `created_at` - The date and time that the dedicated host group was created.
* `dedicated_hosts` - The dedicated hosts that are in this dedicated host group.
* `name` - The unique user-defined name for this dedicated host group.
* `resource_type` - The type of resource referenced.
* `resource_group` - The unique identifier of the resource group for this dedicated host.
* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on this dedicated host group.
* `zone` - The zone this dedicated host resides in.
