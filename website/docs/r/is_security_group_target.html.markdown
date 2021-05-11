---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_target"
description: |-
  Manages IBM Security Group Target
---

# ibm_is_security_group_target

This request adds a resource to an existing security group. The supplied target identifier can be:
  - A network interface identifier
  - An application load balancer identifier
When a target is added to a security group, the security group rules are applied to the target. A request body is not required, and if supplied, is ignored.

## Example Usage

In the following example, you can create a security group target:

```hcl
resource "ibm_is_security_group_target" "testacc_security_group_target" {
    security_group = ibm_is_security_group.testacc_security_group.id
    target = "r006-5b77aa07-7dfb-4c74-a1bd-904b23cbe198"
  }
```

## Argument Reference

The following arguments are supported:

- `security_group` - (Required, string,ForceNew) The security group identifier
- `target` - (Required, string,ForceNew) The security group target identifier

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the security group target. The id is composed of <`security_group_id`>/<`target_id`>.
- `name` - The user-defined name of the target
- `resource_type` - The resource type

## Import

ibm_is_security_group_target can be imported using security group ID and target ID, eg

```
$ terraform import ibm_is_security_group_target.example r006-6c6528a7-26de-4438-9685-bf2f6bbcb1ad/r006-5b77aa07-7dfb-4c74-a1bd-904b23cbe198

```