---
layout: "ibm"
page_title: "IBM : compute_placement_group"
sidebar_current: "docs-ibm-resource-compute-placement-group"
description: |-
  Manages IBM Compute Placement Group.
---


# ibm\_compute_placement_group

Provides provisioning placement groups. This allows placement groups to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](https://softlayer.github.io/reference/datatypes/SoftLayer_Virtual_PlacementGroup).

## Example Usage

```hcl
resource "ibm_compute_placement_group" "test_placement_group" {
    name = "test"
    pod = "pod01"
    datacenter = "dal05"  
}
```

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:  
  * `delete` - (Defaults to 10 mins) Used when deleting the placement group. There might be Virtual Guest resources on the placement group. The placement group delete request is issued once there are no Virtual Guests on the placement group.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify a placement group.
* `datacenter` - (Required, string) The datacenter in which you want to provision the placement group.
* `pod` - (Required, string) The pod in which you want to provision the placement group.
* `rule` - (Optional, string) The rule of the placement group. Default `SPREAD`. 
* `tags` - (Optional, array of strings) Tags associated with the placement group.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new placement group.
