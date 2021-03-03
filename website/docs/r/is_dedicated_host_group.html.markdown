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
resource "is_dedicated_host_group" "is_dedicated_host_group" {
  class = "mx2"
  family = "balanced"
  zone = {"name":"us-south-1"}
}
```

## Argument Reference

The following arguments are supported:

* `class` - (Optional, string) The dedicated host profile class for hosts in this group.
* `family` - (Optional, string) The dedicated host profile family for hosts in this group.
* `name` - (Optional, string) The unique user-defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly-selected words.
* `resource_group` - (Optional, List) The resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
  * `id` - (Optional, string) The unique identifier for this resource group.
* `zone` - (Optional, List) The zone this dedicated host group will reside in.
  * `name` - (Optional, string) The globally unique name for this zone.
  * `href` - (Optional, string) The URL for this zone.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHostGroup.
* `created_at` - The date and time that the dedicated host group was created.
* `crn` - The CRN for this dedicated host group.
* `dedicated_hosts` - The dedicated hosts that are in this dedicated host group.
* `href` - The URL for this dedicated host group.
* `resource_type` - The type of resource referenced.
* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on this dedicated host group.
