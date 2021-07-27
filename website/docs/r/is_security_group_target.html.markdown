---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_target"
description: |-
  Manages IBM security group target.
---

# ibm_is_security_group_target
This request adds a resource to an existing security group. The supplied target identifier can be:
  - A network interface identifier.
  - An application load balancer identifier.
When a target is added to a security group, the security group rules are applied to the target. A request body is not required, and if supplied, is ignored. For more information, about security group target, see [required permissions](https://cloud.ibm.com/docs/vpc?topic=vpc-resource-authorizations-required-for-api-and-cli-calls).

## Example usage
Sample to create a security group target.

```terraform
resource "ibm_is_security_group_target" "testacc_security_group_target" {
    security_group = ibm_is_security_group.testacc_security_group.id
    target = "r006-5b77aa07-7dfb-4c74-a1bd-904b23cbe198"
  }
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `security_group` - (Required, Force new resource, String) The security group identifier.
- `target` - (Required, Force new resource, String) The security group target identifier.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the security group target. The id is composed of <`security_group_id`>/<`target_id`>.
- `name` - (String) The user-defined name of the target.
- `resource_type` - (String) The resource type.

## Import

The `ibm_is_security_group_target` resource can be imported by using security group ID and target ID.

**Example**

```
$ terraform import ibm_is_security_group_target.example r006-6c6528a7-26de-4438-9685-bf2f6bbcb1ad/r006-5b77aa07-7dfb-4c74-a1bd-904123123cbe198

```
