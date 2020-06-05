---
layout: "ibm"
page_title: "IBM : Instance templates"
sidebar_current: "docs-ibm-datasources-is-instance-templates"
description: |-
  Manages IBM Cloud virtual server instance templates.
---

# ibm\_is_instance_templates

Import the details of an existing IBM Cloud virtual server instance templates as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_instance_templates" "ds_instance_templates" {
}

```

## Attribute Reference

The following attributes are exported:

* `templates` - List of all virtual server instance templates in the region.
  * `id` - The id of the instance template.
  * `vpc` - The vpc id.
  * `zone` - Name of the zone.
  * `profile` - The profile name.
  * `image` - ID of the image.
  * `keys` - Comma separated IDs of ssh keys.
  * `user_data` - User data to transfer to the server instance.
  * `resource_group` - The resource group ID for the instance.
  * `vcpu` - A nested block describing the VCPU configuration of this instance.
  Nested `vcpu` blocks have the following structure:
    * `architecture` - The architecture of the instance.
    * `count` - The number of VCPUs assigned to the instance.
    * `gpu` - A nested block describing the gpu of this instance.
  Nested `gpu` blocks have the following structure:
    * `cores` - The cores of the gpu.
    * `count` - Count of the gpu.
    * `manufacture` - Manufacture of the gpu.
    * `memory` - Memory of the gpu.
    * `model` - Model of the gpu.
    * `primary_network_interface` - A nested block describing the primary network interface of this instance.
  Nested `primary_network_interface` blocks have the following structure:
    * `id` - The id of the network interface.
    * `name` - The name of the network interface.
    * `subnet` -  ID of the subnet.
    * `security_groups` -  List of security group ids.
    * `primary_ipv4_address` - The primary IPv4 address.
    * `network_interfaces` - A nested block describing the additional network interface of this instance.
  Nested `network_interfaces` blocks have the following structure:
    * `id` - The id of the network interface.
    * `name` - The name of the network interface.
    * `subnet` -  ID of the subnet.
    * `security_groups` -  List of security group ids.
    * `primary_ipv4_address` - The primary IPv4 address.
    * `boot_volume` - A nested block describing the boot volume.
  Nested `boot_volume` blocks have the following structure:
    * `name` - The name of the boot volume.
    * `size` -  Capacity of the volume in GB.
    * `iops` -  Input/Output Operations Per Second for the volume.
    * `profile` - The profile of the volume.
    * `encryption` - The encryption of the boot volume.
  * `volume_attachments` - A nested block describing the volume attachments.
  Nested `volume_attachments` block have the following structure:
    * `id` - The id of the volume attachment
    * `name` -  The name of the volume attachment
    * `volume_id` - The id of the volume attachment's volume
    * `volume_name` -  The name of the volume attachment's volume
    * `volume_crn` -  The CRN of the volume attachment's volume
