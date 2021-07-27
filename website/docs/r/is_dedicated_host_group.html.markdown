---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_group"
description: |-
  Manages dedicated host group.
---

# ibm_is_dedicated_host_group
Create, update, delete and suspend the dedicated host resource. For more information, about dedicated host groups in your IBM Cloud VPC, see [Dedicated hosts](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-dedicated-hosts-instances).

## Example usage

```terraform
resource "ibm_is_dedicated_host_group" "is_dedicated_host_group" {
  class = "mx2"
  family = "balanced"
  zone = "us-south-1"
  name = "dh-group-name"
}
```


## Argument reference
Review the argument references that you can specify for your resource. 

- `class` - (Required, String) The dedicated host profile class for hosts in this group.
- `family` - (Required, String) The dedicated host profile family for hosts in this group.
- `name` - (Optional, String) The unique user defined name for this dedicated host group. If unspecified, the name will be a hyphenated list of randomly selected words.
- `resource_group` - (Optional, String) The unique ID of the resource group to use. If unspecified, the account's default resource group is used.
- `zone` - (Required, String) The globally unique name of the zone this dedicated host group will reside in.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `class`-  (String) The dedicated host profile class for hosts in this group.
- `family`-  (String) The dedicated host profile family for hosts in this group.
- `id`-  (String) The unique ID of the dedicated host group.
- `href`-  (String) The URL for this dedicated host group.
- `crn`-  (String) The CRN for this dedicated host group.
- `created_at`-  (String) The date and time that the dedicated host group was created.
- `dedicated_hosts`-  (String) The dedicated hosts that are in this dedicated host group.
- `name`-  (String) The unique user defined name for this dedicated host group.
- `resource_type`-  (String) The type of resource referenced.
- `resource_group`-  (String) The unique ID of the resource group for this dedicated host.
- `supported_instance_profiles`-  (String) Array of instance profiles that can be used by instances placed on this dedicated host group.
- `zone`-  (String) The zone this dedicated host resides in.


## Import
The `ibm_is_dedicated_host_group` resource can be imported by using dedicated host group ID.

**Example**

```
$ terraform import ibm_is_dedicated_host_group.example 0716-5fa4a9c4-a194-4915-854b-10101111
```
