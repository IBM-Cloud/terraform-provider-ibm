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
  - An endpoint gateway identifier.
  
When a target is added to a security group, the security group rules are applied to the target. A request body is not required, and if supplied, is ignored. For more information, about security group target, see [required permissions](https://cloud.ibm.com/docs/vpc?topic=vpc-resource-authorizations-required-for-api-and-cli-calls).

**Note**
- IBM Cloud terraform provider currently provides both a standalone `ibm_is_security_group_target` resource and a `security_groups` block defined in-line in the `ibm_is_instance_network_interface` resource to attach security group to a network interface target. At this time you cannot use the `security_groups` block inline with `ibm_is_instance_network_interface` in conjunction with the standalone resource `ibm_is_security_group_target`. Doing so will create a conflict of security groups attaching to the network interface and will overwrite it.
- VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

  **provider.tf**

  ```terraform
  provider "ibm" {
    region = "eu-gb"
  }
  ```

## Example usage
Sample to create a security group target.

```terraform
resource "ibm_is_security_group_target" "example" {
  security_group = ibm_is_security_group.example.id
  target         = ibm_is_lb.example.id
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `security_group` - (Required, Force new resource, String) The security group identifier.
- `target` - (Required, Force new resource, String) The security group target identifier. Could be one of the below:
  - A network interface identifier.
  - An application load balancer identifier.
  - An endpoint gateway identifier.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN for this target.
- `id` - (String) The unique identifier of the security group target. The id is composed of <`security_group_id`>/<`target_id`>.
- `name` - (String) The user-defined name of the target.
- `resource_type` - (String) The resource type.

## Import

The `ibm_is_security_group_target` resource can be imported by using security group ID and target ID.

**Example**

```
$ terraform import ibm_is_security_group_target.example r006-6c6528a7-26de-4438-9685-bf2f6bbcb1ad/r006-5b77aa07-7dfb-4c74-a1bd-904123123cbe198
```
