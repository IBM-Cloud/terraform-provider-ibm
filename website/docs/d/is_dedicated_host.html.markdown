---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host"
description: |-
  Get information about dedicated host
---

# ibm_is_dedicated_host
Retrieve the dedicated host data sources. For more information, about dedicated host in your IBM Cloud VPC, see [Dedicated hosts](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-dedicated-hosts-instances).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example retrieves information about the dedicated host data sources.

```terraform
data "ibm_is_dedicated_host" "example" {
  host_group = ibm_is_dedicated_host_group.example.id
  name       = "example-dedicated-host"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `host_group` - (Required, String) The unique identifier of the dedicated host group.
- `name` - (Required, String) The unique name of this dedicated host.
- `resource_group` - (Optional, String) The unique identifier of the resource group.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `access_tags`  - (List) Access management tags associated for dedicated host.
- `available_memory` -  (String) The amount of memory in `GB` that is currently available for instances.
- `available_vcpu` -  (List) The available `VCPU` for the dedicated host. 

  Nested scheme for `available_vcpu`:
  - `architecture` -  (String) The `VCPU` architecture.
  - `manufacturer` -  (String) The `VCPU` manufacturer.
  - `count` -  (String) The number of `VCPUs` assigned.
- `created_at` -  (String) The date and time that the dedicated host was created.
- `crn` -  (String) The CRN for this dedicated host.
- `disks` - (List) The Collection of the dedicated host's disks. 

  Nested scheme for `disks`:
  - `available` - (String) The remaining space left for instance placement in GB (gigabytes).
  - `created_at` - (String) The creation date and time of the disk.
  - `href` - (String) The URL for the disk.
  - `id` - (String) The unique identifier for the disk.
  - `instance_disks` - (List) Instance disks that are on the dedicated host disk. 

     Nested scheme for `instance_disks`:
     - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information. 

       Nested scheme for `deleted`:
       - `more_info` - (List) Link to documentation about deleted resources.
     - `href` - (String) The URL for this instance disk.
     - `id` - (String) The unique identifier for this instance disk.
     - `name` - (String) The user-defined name for this disk.
     - `resource_type` - (String) The resource type.
   - `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for the property are expected to expand in the future. When processing the property, you can check for and log unknown values. Optionally, halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
   - `lifecycle_state` - (String) The lifecycle state of this dedicated host disk.
   - `name` - (String) The user-defined or system-provided name for the disk.
   - `provisionable` - (String)  Indicates whether the dedicated host disk is available for instance disk creation.
   - `resource_type` - (String) The type of resource referenced.
   - `size` - (String) The size of the disk in GB (gigabytes).
   - `supported_instance_interface_types` - (String) The instance disk interfaces supported for this dedicated host disk.
- `host_group` -  (String) The unique identifier of the dedicated host group this dedicated host is in.
- `href` -  (String) The URL for this dedicated host.
- `id` -  (String) The unique identifier of the dedicated host.
- `instance_placement_enabled` -  (String) If set to **true**, instances can be placed on this dedicated host.
- `instances` -  (List) Array of instances that are allocated to this dedicated host. 

  Nested scheme for `instances`:
  - `crn` -  (String) The CRN for this virtual server instance.
  - `deleted` -  (List) If present, this property indicates the referenced resource has been deleted and provides supplementary information. 

    Nested scheme for `deleted`:
    - `more_info` -  (String) Link to documentation about deleted resources.
  - `href` -  (String) The URL for this virtual server instance.
  - `id` -  (String) The unique identifier for this virtual server instance.
  - `name` -  (String) The user defined name for this virtual server instance (and default system hostname).
- `lifecycle_state` -  (String) The lifecycle state of the dedicated host resource.
- `memory` -  (String) The total amount of memory in `GB`` for this host.
- `name` -  (String) The unique user defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `numa` - The dedicated host NUMA configuration.
    
      Nested scheme for `numa`:
      - `count` - (Integer) The total number of NUMA nodes for this dedicated host.
      - `nodes` - (List) The NUMA nodes for this dedicated host.
      
          Nested scheme for `nodes`:
          - `available_vcpu` - (Integer) The available VCPU for this NUMA node.
          - `vcpu` - (Integer) The total VCPU capacity for this NUMA node.
- `profile` -  (List) The profile this dedicated host uses. 

  Nested scheme for `profile`:
  - `href` -  (String) The URL for this dedicated host.
  - `name` -  (String) The globally unique name for this dedicated host profile.
- `provisionable` -  (String) Indicates whether this dedicated host is available for instance creation.
- `resource_group` -  (String) The unique identifier of the resource group.
- `resource_type` -  (String) The type of resource referenced.
- `socket_count` -  (String) The total number of sockets for this host.
- `state` -  (String) The administrative state of the dedicated host.
- `supported_instance_profiles` -  (List) Array of instance profiles that can be used by instances placed on this dedicated host. 

  Nested scheme for `supported_instance_profiles`:
  - `href` -  (String) The URL for this virtual server instance profile.
  - `name` -  (String) The globally unique name for this virtual server instance profile.
- `vcpu` -  (List) The total `VCPU` of the dedicated host. 

  Nested scheme for `vcpu`:
  - `architecture` -  (String) The `VCPU` architecture.
  - `manufacturer` -  (String) The `VCPU` manufacturer.
  - `count` -  (String) The number of `VCPUs` assigned.
- `zone` -  (String) The globally unique name of the zone this dedicated host resides in.
