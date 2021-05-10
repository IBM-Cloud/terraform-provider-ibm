---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: istance_templates"
description: |-
  Retrives all the IBM VPC instance templates.
---

# ibm\_is_instance_templates

Retrives all the instance template in an account

## Example Usage

In the following example, you can get info of list of instance templates VPC gen-2 infrastructure.
```terraform	
data "ibm_is_instance_templates" "instancetemplates" {	   
}

```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:
* `templates` - List of templates
  * `id` - Id of the instance instance template
  * `name` - The name of the instance template.
  * `image` - The ID of the image to used to create the template.
  * `profile` - The number of instances to be created under the instance group.
  * `vpc` - The ID of VPC in which the instance templates needs to be created.
  * `zone` - Name of the zone
  * `keys` - List of ssh-key ids used to allow login user to the instances.
  * `resource_group` - Resource group ID.
  * `primary_network_interfaces` - A nested block describing the primary network interface for the template
    * `subnet` - The VPC subnet to assign to the interface. 
    * `name` - Name of the interface.
    * `primary_ipv4_address` - Pv4 address assigned to the primary network interface.
    * `security_groups` - List of security groups under the subnet.
  * `network_interfaces` - A nested block describing the network interfaces for the template.
    * `subnet` - The VPC subnet to assign to the interface. 
    * `name` - Name of the interface.
    * `primary_ipv4_address` - IPv4 address assigned to the network interface.
    * `security_groups` - List of security groups under the subnet.
  * `boot_volume` - A nested block describing the boot volume configuration for the template.
    * `encryption` - encryption key CRN to encrypt the boot volume attached. 
    * `name` - Name of the boot volume.
    * `size` - Boot volume size to configured in GB.
    * `profile` - Profile for the boot volume configured.
    * `delete_volume_on_instance_delete` - Configured to delete the boot volume to be deleted upon instance deletion.
  * `volume_attachments` - A nested block describing the storage volume configuration for the template. 
    * `name` - Name of the volume attachment
    * `volume` - Storage volume ID created under VPC.
    * `delete_volume_on_instance_delete` - Configured to delete the storage volume to be deleted upon instance deletion.
    * `volume_prototype` A nested block describing prototype for the volume
      * `iops` - The maximum I/O operations per second (IOPS) for the volume.
      * `profile` - The  globally unique name for the volume profile to use for this volume.
      * `capacity` - The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
      * `encryption_key` - The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
  * `user_data` - User data provided for the instance.
