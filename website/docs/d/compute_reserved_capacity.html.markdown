---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : ibm_compute_reserved_group"
description: |-
  Get information on a IBM Cloud compute reserved group resources
---

# ibm_compute_reserved_group
Retrieve information of an existing reserved group as a read-only data source.

## Example usage

```terraform
data "ibm_compute_reserved_capacity" "reservedcapacityds" {
  name = "reservedgroup"
  most_recent = true
}
```

The following example shows how you can use this data source to reference the reserved group ID in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```terraform
resource "ibm_compute_vm_instance" "vm1" {
    reserved_capacity_id = data.ibm_compute_placement_group.group.id
}
```
## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the reserved capacity.
- `most_recent` - (Optional, Bool) For multiple VM instances, you can set this argument to **true** to import only the most recently created instance.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `datacenter`- (String) The data center in which reserved capacity resides.
- `id` - (String) The unique identifier of the reserved capacity.
- `pod` - (String) The pod in which reserved capacity resides.
- `instances` - (int) Number of VSI instances this capacity reservation can support.
- `virtual_guests` - (List of Objects) A nested block describes the VSIs attached to the reserved capacity.

  Nested scheme for `virtual_guests`:
	- `id` - (String) The ID of the virtual guest.
	- `domain` - (String) The domain of the virtual guest.
	- `hostname` - (String) The hostname of the virtual guest.
