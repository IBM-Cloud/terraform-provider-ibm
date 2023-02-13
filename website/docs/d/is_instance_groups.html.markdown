---
layout: "ibm"
page_title: "IBM : ibm_is_instance_groups"
description: |-
  Get information about InstanceGroupCollection
subcategory: "VPC infrastructure"
---

# ibm_is_instance_groups

Provides a read-only data source for InstanceGroupCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about instance groups , see [Managing instance groups](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-instance-group&mhsrc=ibmsearch_a&mhq=instance+group).

**Note:** 
VPC infrastructure services are regional specific and by default targets to `us-south`. If VPC service is created in a region other than `us-south`, please make sure to target the region in the provider block as shown in the `provider.tf` file, .

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```hcl
data "ibm_is_instance_groups" "is_instance_groups" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the InstanceGroupCollection.
- `instance_groups` - (List) Collection of instance groups.

Nested scheme for `instance_groups`:
  - `access_tags`  - (List) Access management tags associated for the instance group.
  - `application_port` - (Integer) Required if specifying a load balancer pool only. Used by the instance group when scaling up instances to supply the port for the load balancer pool member.
  - `created_at` - (String) The date and time that the instance group was created.
  - `crn` - (String) The CRN for this instance group.
  - `href` - (String) The URL for this instance group.
  - `id` - (String) The unique identifier for this instance group.
  - `instance_template` - (List) The template used to create new instances for this group.
  Nested scheme for `instance_template`:
    - `crn` - (String) The CRN for this instance template.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this instance template.
    - `id` - (String) The unique identifier for this instance template.
    - `name` - (String) The unique user-defined name for this instance template.
  - `load_balancer_pool` - (List) The load balancer pool managed by this group. Instances createdby this group will have a new load balancer pool member in thatpool created.
    Nested scheme for `load_balancer_pool`:
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The pool's canonical URL.
    - `id` - (String) The unique identifier for this load balancer pool.
    - `name` - (String) The user-defined name for this load balancer pool.
  - `managers` - (List) The managers for the instance group.
    Nested scheme for `managers`:
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
    Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this instance group manager.
    - `id` - (String) The unique identifier for this instance group manager.
    - `name` - (String) The user-defined name for this instance group manager.
  - `membership_count` - (Integer) The number of instances in the instance group.
  - `name` - (String) The user-defined name for this instance group.
  - `resource_group` - (List)
    Nested scheme for `resource_group`:
    - `href` - (String) The URL for this resource group.
    - `id` - (String) The unique identifier for this resource group.
    - `name` - (String) The user-defined name for this resource group.
  - `status` - (String) The status of the instance group- `deleting`: Group is being deleted- `healthy`: Group has `membership_count` instances- `scaling`: Instances in the group are being created or deleted to reach             `membership_count`- `unhealthy`: Group is unable to reach `membership_count` instances.
  - `subnets` - (List) The subnets to use when creating new instances.
    Nested scheme for `subnets`:
    - `crn` - (String) The CRN for this subnet.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this subnet.
    - `id` - (String) The unique identifier for this subnet.
    - `name` - (String) The user-defined name for this subnet.
    - `resource_type` - (String) The resource type.
  - `updated_at` - (String) The date and time that the instance group was updated.
  - `vpc` - (List) The VPC the instance group resides in.
    Nested scheme for `vpc`:
    - `crn` - (String) The CRN for this VPC.
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this VPC.
    - `id` - (String) The unique identifier for this VPC.
    - `name` - (String) The unique user-defined name for this VPC.
    - `resource_type` - (String) The resource type.
  