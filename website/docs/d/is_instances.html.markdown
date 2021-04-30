---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instances"
description: |-
  Manages IBM Cloud virtual server instances.
---

# ibm\_is_instances

Import the details of an existing IBM Cloud virtual server instances as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_instances" "ds_instances" {
}

```

```hcl

data "ibm_is_instances" "ds_instances1" {
  vpc_name = "testacc_vpc"
}

```
## Argument Reference

The following arguments are supported:

* `vpc_name` - (optional, string) Name of the vpc to filter the instances attached to it.
* `vpc` - (optional, string) VPC ID to filter the instances attached to it.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `instances` - List of all instances in the IBM Cloud Infrastructure.
  * `id` - The id of the instance.
  * `memory` - Memory of the instance.
  * `status` - Status of the instance.
  * `image` - Image used in the instance.
  * `disks` - Collection of the instance's disks. Nested `disks` blocks have the following structure:
	  * `created_at` - The date and time that the disk was created.
	  * `href` - The URL for this instance disk.
	  * `id` - The unique identifier for this instance disk.
	  * `interface_type` - The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	  * `name` - The user-defined name for this disk.
	  * `resource_type` - The resource type.
	  * `size` - The size of the disk in GB (gigabytes).
  * `zone` - zone of the instance.
  * `vpc` - vpc id of the instance.
  * `resource_group` - resource group id of the instance.
  * `vcpu` - A nested block describing the VCPU configuration of this instance.
  Nested `vcpu` blocks have the following structure:
    * `architecture` - The architecture of the instance.
    * `count` - The number of VCPUs assigned to the instance.
  * `primary_network_interface` - A nested block describing the primary network interface of this instance.
  Nested `primary_network_interface` blocks have the following structure:
    * `id` - The id of the network interface.
    * `name` - The name of the network interface.
    * `subnet` -  ID of the subnet.
    * `security_groups` -  List of security groups.
    * `primary_ipv4_address` - The primary IPv4 address.
  * `network_interfaces` - A nested block describing the additional network interface of this instance.
  Nested `network_interfaces` blocks have the following structure:
    * `id` - The id of the network interface.
    * `name` - The name of the network interface.
    * `subnet` -  ID of the subnet.
    * `security_groups` -  List of security groups.
    * `primary_ipv4_address` - The primary IPv4 address.
  * `boot_volume` - A nested block describing the boot volume.
  Nested `boot_volume` blocks have the following structure:
    * `id` -  The id of the boot volume attachment.
    * `name` - The name of the boot volume.
    * `device` -  The boot volume device Name.
    * `volume_id` - The id of the boot volume attachment's volume
    * `volume_crn` - The CRN/encryption of the boot volume attachment's volume
  * `volume_attachments` - A nested block describing the volume attachments.  
  Nested `volume_attachments` block have the following structure:
    * `id` - The id of the volume attachment
    * `name` -  The name of the volume attachment
    * `volume_id` - The id of the volume attachment's volume
    * `volume_name` -  The name of the volume attachment's volume
    * `volume_crn` -  The CRN of the volume attachment's volume
