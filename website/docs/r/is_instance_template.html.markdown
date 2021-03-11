---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_template"
description: |-
  Manages IBM VPC instance template.
---

# ibm\_is_instance_template

Create, Update or delete an instance template on VPC

## Example Usage

In the following example, you can create a instance template VPC gen-2 infrastructure.
```hcl
provider "ibm" {
  generation = 2
}

resource "ibm_is_vpc" "vpc2" {
  name = "vpc2test"
}

resource "ibm_is_subnet" "subnet2" {
  name            = "subnet2"
  vpc             = ibm_is_vpc.vpc2.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = "ssh1"
  public_key = "SSH KEY"
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testtemplate"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
    allow_ip_spoofing = true
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]

  boot_volume {
    name                             = "testbootvol"
    delete_volume_on_instance_delete = true
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the instance template.
* `image` - (Required, string) The ID of the image to used to create the template.
* `profile` - (Required, string) The number of instances to be created under the instance group.
* `vpc` - (Required, string) The ID of VPC in which the instance templates needs to be created.
* `zone` - (Required, string) Name of the zone
* `keys` - (Required, list) List of ssh-key ids used to allow login user to the instances.
* `resource_group` - (Optional, Forces new resource, string) Resource group ID.
* `primary_network_interfaces` - (Required, list) A nested block describing the primary network interface for the template. Nested  primary_network_interface block have the following structure:
  * `subnet` - (Required, Forces new resource, string) The VPC subnet to assign to the interface. 
  * `name` - (Optional, string) Name of the interface.
  * `primary_ipv4_address` - (Optional, string) IPv4 address assigned to the primary network interface.
  * `security_groups` - (Optional, list) List of security groups under the subnet.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `network_interfaces` - (Optional, list) A nested block describing the network interfaces for the template. Nested  network_interfaces blocks have the following structure:
  * `subnet` - (Required, Forces new resource, string) The VPC subnet to assign to the interface. 
  * `name` - (Optional, string) Name of the interface.
  * `primary_ipv4_address` - (Optional, string) IPv4 address assigned to the network interface.
  * `security_groups` - (Optional, list) List of security groups under the subnet.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `boot_volume` - (Optional, block) A nested block describing the boot volume configuration for the template. Nested  boot_volume blocks have the following structure:
  * `encryption` - (Optional, string) encryption key CRN to encrypt the boot volume attached. 
  * `name` - (Optional, string) Name of the boot volume.
  * `delete_volume_on_instance_delete` - (Optional, bool) Configured to delete the boot volume to be deleted upon instance deletion.
* `volume_attachments` - (Optional, list) A nested block describing the storage volume configuration for the template. Nested volume_attachments blocks have the following structure: 
  * `name` - (Required, string) Name of the boot volume.
  * `volume` - (Required, string) Storage volume ID created under VPC.
  * `delete_volume_on_instance_delete` - (Required, bool) Configured to delete the storage volume to be deleted upon instance deletion.
* `user_data` - (Optional, string) User data provided for the instance.

## Attribute Reference

The following attributes are exported:

* `id` - Id of the instance template

## Import

`ibm_is_instance_template` can be imported using instance template ID, eg ibm_is_instance_template.template

```
$ terraform import ibm_is_instance_template.template r006-14140f94-fcc4-1349-96e7-a72734715115
```
