---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_target"
description: |-
  Manages IBM Cloud security group target.
---

# ibm_is_security_group_target
Retrieve information of an existing security group target as a read only data source. For more information, about security group target, see [required permissions](https://cloud.ibm.com/docs/vpc?topic=vpc-resource-authorizations-required-for-api-and-cli-calls).


## Example usage
In the following example, you can create a security group target:

```terraform
data "ibm_is_security_group_target" "testacc_security_group_target" {
    security_group = "r006-5b77aa07-7dfb-4c74-a1bd-904b23cbe198"
    name = "securitygrouptargetname"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `security_group` - (Required, String) The security group identifier.
- `name` - (Required, String) The user defined name of the target.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The unique identifier of the security group target. The ID is composed of <`security_group_id`>/<`target_id`>.
- `name` - (String) The user defined name of the target.
- `resource_type` - (String) The resource type.
- `more_info` - (String) Link to documentation about deleted resources.
