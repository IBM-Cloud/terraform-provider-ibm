---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_hosts"
description: |-
  Get information about dedicate hosts.
---

# ibm_is_dedicated_hosts
Retrieve the dedicated hosts. For more information, about dedicated hosts in the IBM Cloud VPC, see [dedicated hosts](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-dedicated-hosts-instances).

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
data "ibm_is_dedicated_hosts" "example" {
  host_group = ibm_is_dedicated_host_group.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `host_group` - (Optional, String) The unique identifier of the dedicated host group.
- `resource_group` (Optional, String) The ID of the Resource group this dedicated host belongs to.
- `name` (Optional, String) The name of the dedicated host

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` -  (String) The unique identifier of the dedicated host collection.
- `dedicated_hosts` -  (List) Collection of dedicated hosts. Nested dedicated_hosts blocks have the following structure.

  Nested scheme for `dedicated_hosts`:
  - `access_tags`  - (List) Access management tags associated for dedicated hosts.
  - `available_memory` -  (String) The amount of memory in GB that is currently available for 
  - `available_vcpu` -  (List) The available `VCPU` for the dedicated host. Nested available_vcpu blocks have the following structure.

      Nested scheme for `available_vcpu`:
      - `architecture` -  (String) The `VCPU` architecture.
      - `count` -  (String) The number of `VCPUs` assigned.
  - `created_at` -  (Timestamp) The date and time that the dedicated host was created.
  - `crn` -  (String) The CRN for this dedicated host.
  - `disks` - (List) Collection of the dedicated host's disks.

    Nested scheme for `disks`:
    - `available` - (String) The remaining space left for instance placement in GB (gigabytes).
    - `created_at` - (Timestamp) The date and time that the disk was created.
    - `href` - (String) The URL for this disk.
    - `id` - (String) The unique identifier for this disk.
    - `instance_disks` - (List) Instance disks that are on this dedicated host disk. 
      
      Nested scheme for `instance_disks`:
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.

         Nested scheme for `deleted`:
         - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this instance disk.
      - `id` - (String) The unique identifier for this instance disk.
      - `name` - (String) The user-defined name for this disk.
      - `resource_type` - (String) The resource type.
    - `interface_type` - (String) The disk interface used for attaching the disk. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
    - `lifecycle_state` - (String) The lifecycle state of this dedicated host disk.
    - `name` - (String) The user-defined or system-provided name for this disk.
    - `provisionable` - (String) Indicates whether this dedicated host disk is available for instance disk creation.
    - `resource_type` - (String) The type of resource referenced.
    - `size` - (String) The size of the disk in GB (gigabytes).
    - `supported_instance_interface_types` - (String) The instance disk interfaces supported for this dedicated host disk.
  - `host_group` -  (String) The unique identifier of the dedicated host group this dedicated host is in.
  - `href` -  (String) The URL for this dedicated host.
  - `id` -  (String) The unique identifier for this dedicated host.
  - `instance_placement_enabled` -  (String) If set to **true**, instances can be placed on this dedicated host.
  - `instances` -  (List) Array of instances that are allocated to this dedicated host. Nested instances blocks have the following structure.

      Nested scheme for `instances`:
      - `crn` -  (String) The CRN for this virtual server instance.
      - `deleted` -  (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information. Nested deleted blocks have the following structure.

        Nested scheme for `deleted`:
        - `more_info` -  (String) Link to documentation about deleted resources.
      - `href` -  (String) The URL for this virtual server instance.
      - `id` -  (String) The unique identifier for this virtual server instance.
      - `name` -  (String) The user-defined name for this virtual server instance (and default system hostname).
  - `lifecycle_state` -  (String) The lifecycle state of the dedicated host resource.
  - `memory` -  (String) The total amount of memory in GB for this host.
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
  - `resource_group` -  (List) The resource group for this dedicated host. 

      Nested scheme for `resource_group`:
      - `href` -  (String) The URL for this resource group.
      - `id` -  (String) The unique identifier for this resource group.
      - `name` -  (String) The user defined name for this resource group.
  - `resource_type` -  (String) The type of resource referenced.
  - `socket_count` -  (String) The total number of sockets for this host.
  - `state` -  (String) The administrative state of the dedicated host. The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally, halt processing and surface the error, or bypass the dedicated host on which the unexpected property value was encountered.
  - `supported_instance_profiles` -  (List) Array of instance profiles that can be used by instances placed on this dedicated host. 

      Nested scheme for `supported_instance_profiles`:
      - `href` -  (String) The URL for this virtual server instance profile.
      - `name` -  (String) The globally unique name for this virtual server instance profile.
  - `vcpu` -  (List) The total `VCPU` of the dedicated host.

      Nested scheme for `vcpu`:
      - `architecture` -  (String) The `VCPU` architecture.
      - `count` -  (String) The number of `VCPUs` assigned.
  - `zone` -  (String) The globally unique name of the zone this dedicated host resides in.
- `total_count` -  (String) The total number of resources across all pages.

