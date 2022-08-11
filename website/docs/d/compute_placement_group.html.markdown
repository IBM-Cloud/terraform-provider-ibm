---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : ibm_compute_placement_group"
description: |-
  Get information on a IBM Cloud compute placement group resources
---

# ibm_compute_placement_group
Retrieve information of an existing placement group as a read-only data source. For more information, about compute placement group resource, see [workload Placement for virtual servers](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-introduction#workload-placement-for-virtual-servers).

## Example usage

```terraform
data "ibm_compute_placement_group" "group" {
    name = "demo"
}
```

The following example shows how you can use this data source to reference the placement group ID in the `ibm_compute_vm_instance` resource because the numeric IDs are often unknown.

```terraform
resource "ibm_compute_vm_instance" "vm1" {
    placement_group_id = data.ibm_compute_placement_group.group.id
}
```
## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the placement group.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `datacenter`- (String) The data center in which placement group resides.
- `id` - (String) The unique identifier of the placement group.
- `pod` - (String) The pod in which placement group resides.
- `rule` - (String) The rule of the placement group.
- `virtual_guests` - (List of Objects) A nested block describes the VSIs attached to the placement group.

  Nested scheme for `virtual_guests`:
	- `id` - (String) The ID of the virtual guest.
	- `domain` - (String) The domain of the virtual guest.
	- `hostname` - (String) The hostname of the virtual guest.
