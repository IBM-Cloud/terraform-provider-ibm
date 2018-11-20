---
layout: "ibm"
page_title: "IBM : ibm_compute_placement_group"
sidebar_current: "docs-ibm-datasource-compute-placement-group"
description: |-
  Get information on a IBM Compute Placement Group resources
---

# ibm\_compute_placement_group

Import the details of an existing placement group as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_compute_placement_group" "group" {
    name = "demo"
}
```

The following example shows how you can use this data source to reference the placement group ID in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```hcl
resource "ibm_compute_vm_instance" "vm1" {
    ...
    placement_group_id = "${data.ibm_compute_placement_group.group.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the placement group.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the placement group.
* `datacenter` - The datacenter in which placement group resides.
* `pod` - The pod in which placement group resides.
* `rule` - The rule of the placement group.
* `virtual_guests` - A nested block describing the VSIs attached to the placement group. Nested `virtual_guests` blocks have the following structure:
  * `id` - The ID of the virtual guest.
  * `domain` - The domain of the virtual guest.
  * `hostname` - The hostname of the virtual guest.



