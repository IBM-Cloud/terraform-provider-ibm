---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_targets"
description: |-
  Manages IBM Security Group Targets
---

# ibm_is_security_group_target

This request lists all targets associated with a security group, to which the rules in the security group are applied.

## Example Usage

In the following example, you can create a security group target:

```hcl
data "ibm_is_security_group_targets" "testacc_security_group_targets" {
    security_group = ibm_is_security_group.testacc_security_group.id
  }
```

## Argument Reference

The following arguments are supported:

- `security_group` - (Required, string) The security group identifier

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the security group target <`security_group`>
- `targets` - Collection of security group target references
    - `target` - The unique identifier for this load balancer/network interface
    - `name` - The user-defined name of the target
    - `resource_type` - The resource type
    - `more_info` - Link to documentation about deleted resources
