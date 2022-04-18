---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instances"
description: |-
  Manages IBM Cloud virtual server instances.
---

# ibm_is_instances
Retrieve information of an existing  IBM Cloud virtual server instances as a read-only data source. For more information, about virtual server instances, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).

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

data "ibm_is_instances" "example" {
}

```

```terraform

data "ibm_is_instances" "example" {
  vpc_name = "example-vpc"
}

```

## Argument reference
The input parameters that you need to specify for the data source. 

- `resource_group` - (optional, String) Resource Group ID to filter the instances attached to it.
- `vpc` - (Optional, String) The VPC ID to filter the instances attached.
- `vpc_crn` - (optional, String) VPC CRN to filter the instances attached to it.
- `vpc_name` - (Optional, String) The name of the VPC to filter the instances attached.
- `instance_group` - (Optional, String) Instance group ID to filter the instances attached to it.
- `instance_group_name` - (Optional, String) Instance group name to filter the instances attached to it.
- `dedicated_host_name` - (Optional, String) Dedicated host name to filter the instances attached to it.
- `dedicated_host` - (Optional, String) Dedicated host ID to filter the instances attached to it.
- `placement_group_name` - (Optional, String) Placement group name to filter the instances attached to it.
- `placement_group` - (Optional, String) Placement group ID to filter the instances attached to it.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `instances`- (List of Object) A list of Virtual Servers for VPC instances that exist in your account.
   
   Nested scheme for `instances`:
    - `availability_policy_host_failure` - (String) The availability policy for this virtual server instance. The action to perform if the compute host experiences a failure. 
    - `bandwidth` - (Integer) The total bandwidth (in megabits per second) shared across the instance's network interfaces and storage volumes
	- `boot_volume`- (List) A list of boot volumes that were created for the instance.

	  Nested scheme for `boot_volume`:
		- `device` - (String) The name of the device that is associated with the boot volume.
		- `id` - (String) The ID of the boot volume attachment.
		- `name` - (String) The name of the boot volume.
		- `volume_id` - (String) The ID of the volume that is associated with the boot volume attachment.
		- `volume_crn` - (String) The CRN of the volume that is associated with the boot volume attachment.
	- `crn` - (String) The CRN of the instance.
	- `disks` - (List) Collection of the instance's disks. Nested `disks` blocks has the following structure:

	  Nested scheme for `disks`:
		- `created_at` - (Timestamp) The date and time that the disk was created.
	  	- `href` - (String) The URL for this instance disk.
	  	- `id` - (String) The unique identifier for this instance disk.
	  	- `interface_type` - (String) The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  	- `name` - (String) The user-defined name for this disk.
	  	- `resource_type` - (String) The resource type.
	  	- `size` - (String) The size of the disk in GB (gigabytes).
	- `gpu` - A nested block describing the gpu of this instance.
      Nested `gpu` blocks have the following structure:
        - `count` - Count of the gpu.
        - `manufacture` - Manufacture of the gpu.
        - `memory` - Memory of the gpu.
        - `model` - Model of the gpu.
	- `id` - (String) The ID that was assigned to the Virtual Servers for VPC instance.
	- `image` - (String) The ID of the virtual server image that is used in the instance.
	- `memory`- (Integer) The amount of memory that was allocated to the instance.
	- `metadata_service_enabled` - (Boolean) Indicates whether the metadata service endpoint is available to the virtual server instance.
	- `network_interfaces`- (List) A list of more network interfaces that the instance uses.

	  Nested scheme for `network_interfaces`:
		- `id` - (String) The ID of the more network interface.
		- `name` - (String) The name of the more network interface.
		- `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

			Nested scheme for `primary_ip`:
			- `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
			- `href`- (String) The URL for this reserved IP
			- `name`- (String) The user-defined or system-provided name for this reserved IP
			- `reserved_ip`- (String) The unique identifier for this reserved IP
			- `resource_type`- (String) The resource type.
		- `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses. Same as `primary_ip.0.address`
		- `subnet` - (String) The ID of the subnet that is used in the more network interface.
		- `security_groups` (List)A list of security groups that were created for the interface.
	- `placement_target`- (List) The placement restrictions for the virtual server instance.

	  Nested scheme for `placement_target`: 
		- `crn` - (String) The CRN for this placement target resource.
		- `deleted` - (String) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			- `more_info` -  (String) Link to documentation about deleted resources. 
		- `href` - (String) The URL for this placement target resource.
		- `id` - (String) The unique identifier for this placement target resource.
		- `name` - (String) The unique user-defined name for this placement target resource. If unspecified, the name will be a hyphenated list of randomly-selected words.
		- `resource_type` - (String) The type of resource referenced.
	- `primary_network_interface`- (List) A list of primary network interfaces that were created for the instance. 

	  Nested scheme for `primary_network_interface`:
		- `id` - (String) The ID of the primary network interface.
		- `name` - (String) The name of the primary network interface.
		- `subnet` - (String) The ID of the subnet that is used in the primary network interface.
		- `security_groups` (List)A list of security groups that were created for the interface.
		- `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

			Nested scheme for `primary_ip`:
			- `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
			- `href`- (String) The URL for this reserved IP
			- `name`- (String) The user-defined or system-provided name for this reserved IP
			- `reserved_ip`- (String) The unique identifier for this reserved IP
			- `resource_type`- (String) The resource type.
		- `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses. Same as `primary_ip.0.address`
		- `resource_group` - (String) The name of the resource group where the instance was created.
	- `status` - (String) The status of the instance.
	- `status_reasons` - (List) Array of reasons for the current status. 

		Nested scheme for `status_reasons`:
		- `code` - (String)  A snake case string identifying the status reason.
		- `message` - (String)  An explanation of the status reason
		- `more_info` - (String) Link to documentation about this status reason
	- `total_volume_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
    - `total_network_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance network interfaces.
	- `volume_attachments`- (List) A list of volume attachments that were created for the instance.

	  Nested scheme for `volume_attachments`: 
		- `crn` - (String) The CRN of the volume that is associated with the volume attachment.
		- `id` - (String) The ID of the volume attachment.
		- `name` - (String) The name of the volume attachment.
		- `volume_id` - (String) The ID of the volume that is associated with the volume attachment.
		- `volume_name` - (String) The name of the volume that is associated with the volume attachment.
	- `vcpu`- (List) A list of virtual CPUs that were allocated to the instance.

	  Nested scheme for `vcpu`:
		- `architecture` - (String) The architecture of the virtual CPU.
		- `count`- (Integer) The number of virtual CPUs that are allocated to the instance.
	- `vpc` - (String) The ID of the VPC that the instance belongs to.
	- `zone` - (String) The zone where the instance was created.
