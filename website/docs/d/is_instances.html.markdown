---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instances"
description: |-
  Manages IBM Cloud virtual server instances.
---

# ibm_is_instances
Retrieve information of an existing  IBM Cloud virtual server instances as a read-only data source. For more information, about virtual server instances, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).


## Example usage

```terraform

data "ibm_is_instances" "ds_instances" {
}

```

```terraform

data "ibm_is_instances" "ds_instances1" {
  vpc_name = "testacc_vpc"
}

```

## Argument reference
The input parameters that you need to specify for the data source. 

- `resource_group` - (optional, string) Resource Group ID to filter the instances attached to it.
- `vpc` - (Optional, String) The VPC ID to filter the instances attached.
- `vpc_crn` - (optional, string) VPC CRN to filter the instances attached to it.
- `vpc_name` - (Optional, String) The name of the VPC to filter the instances attached.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `instances`- (List of Object) A list of Virtual Servers for VPC instances that exist in your account.
   
   Nested scheme for `instances`:
	- `boot_volume`- (List) A list of boot volumes that were created for the instance.

	  Nested scheme for `boot_volume`:
		- `device` - (String) The name of the device that is associated with the boot volume.
		- `id` - (String) The ID of the boot volume attachment.
		- `name` - (String) The name of the boot volume.
		- `volume_id` - (String) The ID of the volume that is associated with the boot volume attachment.
		- `volume_crn` - (String) The CRN of the volume that is associated with the boot volume attachment.
	- `disks` - (List) Collection of the instance's disks. Nested `disks` blocks has the following structure:

	  Nested scheme for `disks`:
		- `created_at` - (Timestamp) The date and time that the disk was created.
	  	- `href` - (String) The URL for this instance disk.
	  	- `id` - (String) The unique identifier for this instance disk.
	  	- `interface_type` - (String) The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  	- `name` - (String) The user-defined name for this disk.
	  	- `resource_type` - (String) The resource type.
	  	- `size` - (String) The size of the disk in GB (gigabytes).
	- `id` - (String) The ID that was assigned to the Virtual Servers for VPC instance.
	- `image` - (String) The ID of the virtual server image that is used in the instance.
	- `memory`- (Integer) The amount of memory that was allocated to the instance.
	- `network_interfaces`- (List) A list of more network interfaces that the instance uses.

	  Nested scheme for `network_interfaces`:
		- `id` - (String) The ID of the more network interface.
		- `name` - (String) The name of the more network interface.
		- `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses.
		- `subnet` - (String) The ID of the subnet that is used in the more network interface.
		- `security_groups` (List)A list of security groups that were created for the interface.
	- `primary_network_interface`- (List) A list of primary network interfaces that were created for the instance. 

	  Nested scheme for `primary_network_interface`:
		- `id` - (String) The ID of the primary network interface.
		- `name` - (String) The name of the primary network interface.
		- `subnet` - (String) The ID of the subnet that is used in the primary network interface.
		- `security_groups` (List)A list of security groups that were created for the interface.
		- `primary_ipv4_address` - (String) The IPv4 address range that the subnet uses.- `resource_group` - (String) The name of the resource group where the instance was created.
	- `status` - (String) The status of the instance.
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

