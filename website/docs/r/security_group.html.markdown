---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_security_group"
description: |-
  Manages IBM Cloud Security group.
---

# ibm_security_group
Create, delete, and update a security group. Provides a networking security group resource that controls access to the public and private interfaces of a virtual server instance. To create rules for the security group, use the `security_group_rule` resource. For more information, about security group, see [managing security groups](https://cloud.ibm.com/docs/security-groups?topic=security-groups-managing-sg).

For more information, see [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_SecurityGroup).

## Example usage

```terraform
resource "ibm_security_group" "sg1" {
    name = "sg1"
    description = "allow my app traffic"
}
```

## Argument reference 
Review the argument references that you can specify for your resource.

- `name` - (Required, Forces new resource, String)The descriptive name that is used to identify the security group.
- `description`- (Optional, String) More details to describe the security group.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the security group.
