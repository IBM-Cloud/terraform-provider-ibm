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
```hcl	
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
    * `name` - Name of the boot volume.
    * `volume` - Storage volume ID created under VPC.
    * `delete_volume_on_instance_delete` - Configured to delete the storage volume to be deleted upon instance deletion.
  * `user_data` - User data provided for the instance.
