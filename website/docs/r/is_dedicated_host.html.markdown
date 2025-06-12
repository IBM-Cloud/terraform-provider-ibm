---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host"
description: |-
  Manages IBM DedicatedHost.
---

# ibm_is_dedicated_host
Create, update, delete and suspend the dedicated host resource. For more information, about dedicated host in your IBM Cloud VPC, see [Dedicated hosts](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-dedicated-hosts-instances).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_dedicated_host_group" "example" {
  family = "memory"
  class  = "beta"
  zone   = "us-south-1"
}
data "ibm_is_dedicated_host_group" "example" {
  name = ibm_is_dedicated_host_group.example.name
}
resource "ibm_is_dedicated_host" "example" {
  profile    = "dh2-56x464"
  host_group = ibm_is_dedicated_host_group.example.id
  name       = "example-dh-host"
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the dedicated host.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `host_group` - (Required, String)The unique ID of the dedicated host group for this dedicated host.
- `instance_placement_enabled`- (Optional, Bool) If set to **true** instances can be placed on the dedicated host.
- `name` - (Optional, String) The unique user-defined name for the dedicated host. If unspecified, the name will be a hyphenated list of randomly selected words.
- `profile`-  (String)  Required - The globally unique name of the dedicated host profile to use for the dedicated host.
- `resource_group`- (Optional, String) The unique ID of the resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`-  (String) The unique ID of the dedicated host.
- `available_memory`-  (String) The amount of memory in `GB` that is currently available for instances.
- `available_vcpu`-  (String) The available `VCPU` for the dedicated host.
- `created_at`-  (String) The date and time that the dedicated host was created.
- `crn`-  (String) The CRN for this dedicated host.

- `disks` - (List) Collection of the dedicated host's disks. Nested `disks` blocks have the following structure:

  Nested scheme for `disks`:
  - `available` - (String) The remaining space left for instance placement in GB (gigabytes).
  - `created_at` - (String) The date and time that the disk was created.
  - `href` - (String) The URL for this disk.
  - `id` - (String) The unique identifier for this disk.
  - `instance_disks` - (List) Instance disks that are on this dedicated host disk. Nested `instance_disks` blocks have the following structure:
  
      Nested scheme for `instance_disks`:
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
          
          Nested scheme for `deleted`:
          - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this instance disk.
      - `id` - (String) The unique identifier for this instance disk.
      - `name` - (String) The user defined name for this disk.
      - `resource_type` - (String) The resource type.
 - `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
 - `lifecycle_state` - (String) The lifecycle state of this dedicated host disk.
 - `name` - (String) The user defined or system provided name for this disk.
 - `provisionable` - (String) Indicates whether this dedicated host disk is available for instance disk creation.
 - `resource_type` - (String) The type of resource reference.
 - `size` - (String) The size of the disk in GB (gigabytes).
 - `supported_instance_interface_types` - (String) The instance disk interfaces supported for this dedicated host disk.
- `host_group`-  (String) The unique ID of the dedicated host group this dedicated host is in.
- `href`-  (String) The URL for this dedicated host.
- `instance_placement_enabled`-  (String) If set to **true**, instances can be placed on this dedicated host.
- `instances`-  (List)  Array of instances that are allocated to this dedicated host.

  Nested scheme for `instance`:
  - `crn`-  (List) The CRN for the VSI.

    Nested scheme for `crn`:
    - `deleted`-  (List) If present, this property indicates the referenced resource has been deleted and provides supplementary information. Nested deleted blocks have the following structure.

      Nested scheme for `deleted`:
      - `more_info`-  (String) Link to documentation about deleted resources.
    - `href`-  (String) The URL for this VSI.
    - `id`-  (String) The unique ID for this virtual server instance.
    - `name`-  (String) The user defined name for the VSI and is the default system hostname.
- `lifecycle_state`-  (String) The lifecycle state of the dedicated host resource.
- `memory`-  (String) The total amount of memory in `GB` for this host.
- `name`-  (String) The unique user defined name for this dedicated host.
- `numa` - The NUMA configuration for this dedicated host.
  
  Nested scheme for `numa`:
    - `count` - (Integer) The total number of NUMA nodes for this dedicated host.
    - `nodes` - (List) The NUMA nodes for this dedicated host.
      
      Nested scheme for `nodes`:
        - `available_vcpu` - (Integer) The available VCPU for this NUMA node.
        - `vcpu` - (Integer) The total VCPU capacity for this NUMA node.
- `profile`-  (String) The profile this dedicated host uses.
- `provisionable`-  (String) Indicates whether this dedicated host is available for instance creation.
- `resource_group`-  (String) The unique identifier of the resource group for this dedicated host.
- `resource_type`-  (String) The type of resource referenced.
- `socket_count`-  (String) The total number of sockets for this host.
- `state`-  (String) The administrative state of the dedicated host. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.
- `supported_instance_profiles`-  (String) Array of instance profiles that can be used by instances placed on this dedicated host.
- `vcpu`-  (String) The total `VCPU` of the dedicated host.
- `zone`-  (String) The zone this dedicated host resides in.


## Import
The `ibm_is_dedicated_host` resource can be imported by using dedicated host ID.

**Example**

```
$ terraform import ibm_is_dedicated_host.example 0716-1c372bb2-decc-4555111a6-1010101
```
