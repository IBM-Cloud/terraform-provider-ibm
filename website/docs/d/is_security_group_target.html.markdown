---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_target"
description: |-
  Manages IBM Security Group Target
---

# ibm_is_security_group_target

This request retrieves a single target for a given security group, The target must be an existing target of the security group.

## Example Usage

In the following example, you can create a security group target:

```hcl
data "ibm_is_security_group_target" "testacc_security_group_target" {
    security_group_id = ibm_is_security_group.testacc_security_group.id
    target_id = "r006-5b77aa07-7dfb-4c74-a1bd-904b23cbe198"
  }
```

## Argument Reference

The following arguments are supported:

- `security_group_id` - (Required, string,ForceNew) The security group identifier
- `target_id` - (Required, string,ForceNew) The security group target identifier

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the security group target. The id is composed of <`security_group_id`>/<`target_id`>.
- `name` - The user-defined name of the target
- `resource_type` - The resource type
- `more_info` - Link to documentation about deleted resources

