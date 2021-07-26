---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : compute_placement_group"
description: |-
  Manages IBM Cloud compute placement group.
---


# ibm_compute_placement_group
Create, update, or delete a placement group for your virtual server instance. With placement groups, you can control the physical host your virtual server instance is deployed to. For more information, about the placement group, see [placement groups](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-placement-groups). 

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](https://softlayer.github.io/reference/datatypes/SoftLayer_Virtual_PlacementGroup).

## Example usage

```terraform
resource "ibm_compute_placement_group" "test_placement_group" {
    name = "test"
    pod = "pod01"
    datacenter = "dal05"  
}
```

## Timeouts
The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) for certain actions:  

- `delete`- (Defaults to 10 mins) Used when you delete the placement group. There might be Virtual Guest resources on the placement group. The placement group delete request is issued once there are no Virtual Guests on the placement group.

## Argument reference
Review the argument references that you can specify for your resource. 

- `datacenter` - (Required, Forces new resource, String)The datacenter in which you want to provision the placement group.
- `name` - (Required, String) The descriptive that is used to identify a placement group.
- `pod` - (Required, Forces new resource, String) The data center pod where you want to create the placement group. To find the pod, run `ibmcloud sl placement-group create-options` and select one of the **Back-end Router IDs** for the data center where you want to create the placement group.
- `rule`- (Optional, Forces new resource, String) The rule of the placement group. Default `SPREAD`.
- `tags`- (Optional, Array of Strings) Tags associated with the placement group. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the new placement group.
