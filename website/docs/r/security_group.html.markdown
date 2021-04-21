---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_security_group"
description: |-
  Manages IBM Security Groups
---

# ibm\_security_group

Provides a networking security group resource that controls access to the public and private interfaces of a virtual server instance. This resource allows security groups to be created, updated, and deleted. To create rules for the security group, use the `security_group_rule` resource.

For additional details, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_SecurityGroup).

## Example Usage

```
resource "ibm_security_group" "sg1" {
    name = "sg1"
    description = "allow my app traffic"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The descriptive name used to identify the security group.
* `description` - (Optional, string) Additional details to describe the security group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the security group.
