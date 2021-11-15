---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_reserved_capacity"
description: |-
  Manages IBM Cloud compute reserved capacity
---


# ibm_compute_reserved_capacity
Create, or update reserved capacity for your Virtual Server Instance. With reserved capacity, you can control the physical host your Virtual Server Instance is deployed to. For more information, about the reserved capacity, see [reserved groups](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-about-reserved-virtual-servers). 

**Note**

For more information, see [IBM Cloud Classic Infrastructure (SoftLayer) API docs](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_ReservedCapacityGroup).


terraform destroy does not remove the reserved capacity but only clears the state file. We cannot cancel reserved capacity. For more information see [FAQS](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-faqs-reserved-capacity-and-instances#what-happens-if-i-don-t-need-my-reserved-virtual-server-instances-anymore-)


## Example usage

```terraform
resource "ibm_compute_reserved_capacity" "reserved" {
    datacenter = "lon02"
    pod = "pod01"
    instances = 6
    name = "reservedinstanceterraformupdate"
    flavor = "B1_4X16_1_YEAR_TERM"
}
```

## Timeouts
The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) for certain actions:  

- `create`- (Defaults to 10 mins) Used when you create the reserved capacity.

## Argument reference
Review the argument references that you can specify for your resource. 

- `datacenter` - (Required, Forces new resource, String) The datacenter in which you want to provision the reserved capacity.
- `flavor`- (Required, String) Capacity keyname. For example, C1_2X2_1_YEAR_TERM.
- `force_create` - (Optional, Boolean) Force the creation of reserved capacity with same name.
- `instances`- (Required, Forces new resource, Integer) Number of VSI instances this capacity reservation can support.
- `name` - (Required, String) The descriptive that is used to identify a reserved capacity.
- `pod` - (Required, Forces new resource, String) The data center pod where you want to create the reserved capacity.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the new reserved capacity.

## Import
The `ibm_compute_reserved_capacity` resource can be imported by using reserved capacity ID.

**Example**

```
$ terraform import ibm_compute_reserved_capacity.example 88205074
```