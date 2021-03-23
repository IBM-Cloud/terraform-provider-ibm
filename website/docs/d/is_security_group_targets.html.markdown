---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_target"
description: |-
  Manages IBM Security Group Target
---

# ibm_is_security_group_target

This request lists all targets associated with a security group, to which the rules in the security group are applied.

## Example Usage

In the following example, you can create a security group target:

```hcl
data "ibm_is_security_group_targets" "testacc_security_group_targets" {
    security_group_id = ibm_is_security_group.testacc_security_group.id
  }
```

## Argument Reference

The following arguments are supported:

- `security_group_id` - (Required, string,ForceNew) The security group identifier
- `limit` - (Optional) The number of resources to return on a page, Default : 50


## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the security group target <`security_group_id`>
- `limit` - The maximum number of resources that can be returned by the request(Constraints: 1 ≤ value ≤ 100)
- `total_count` - The total number of resources across all pages(Constraints: value ≥ 0)
- `targets` - Collection of security group target references
    - `target_id` - The unique identifier for this load balancer/network interface
    - `name` - The user-defined name of the target
    - `resource_type` - The resource type
    - `more_info` - Link to documentation about deleted resources

