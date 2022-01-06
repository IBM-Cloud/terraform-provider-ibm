---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_security_group_targets"
description: |-
  Manages IBM Cloud security group targets.
---

# ibm_is_security_group_targets
Retrieve information of an existing security group targets as a read only data source. For more information, about security group targets, see [required permissions](https://cloud.ibm.com/docs/vpc?topic=vpc-resource-authorizations-required-for-api-and-cli-calls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can create a security group target:

```terraform
data "ibm_is_security_group_targets" "example" {
  security_group = ibm_is_security_group.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `security_group` - (Required, String) The security group identifier

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The unique identifier of the security group target <`security_group`>
- `targets` - (List) Collection of security group target references

  Nested scheme for `targets`:
  - `crn` - (String) The CRN for this target.
  - `target` - (String) The unique identifier for this load balancer/network interface.
  - `name` - (String) The user-defined name of the target.
  - `resource_type` - (String) The resource type.
  - `more_info` - (String) Link to documentation about deleted resources.
