---
layout: "ibm"
page_title: "IBM: ibm_security_group"
sidebar_current: "docs-ibm-resource-security-group"
description: |-
  Manages IBM Security Groups
---

# ibm\_security_group

Provide a Security Group resource. This allows Groups to be created, updated and deleted.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_SecurityGroup).

## Example Usage

```
resource "ibm_security_group" "sg1" {
    name = "sg1"
    description = "allow my app traffic"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The descriptive name used to identify the security group.
* `description` - (Optional, string) A description for the security group.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new security group
