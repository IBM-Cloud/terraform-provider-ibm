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

data "ibm_resource_group" "default" {
  name = "Default" ///give your resource grp
}

resource "ibm_is_dedicated_host_group" "dh_group01" {
  family = "compute"
  class = "cx2"
  zone = "us-south-1"
  name = "my-dh-group-01"
  resource_group = data.ibm_resource_group.default.id
}

resource "ibm_is_dedicated_host" "is_dedicated_host" {
  profile = "bx2d-host-152x608"
  name = "my-dedicated-host-01"
	host_group = ibm_is_dedicated_host_group.dh_group01.id
  resource_group = data.ibm_resource_group.default.id
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

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testtemplate1"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
    allow_ip_spoofing = true
  }

  dedicated_host_group = ibm_is_dedicated_host_group.dh_group01.id
  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]

  boot_volume {
    name                             = "testbootvol"
    delete_volume_on_instance_delete = true
  }
}

resource "ibm_is_instance_template" "instancetemplate2" {
  name    = "testtemplat2"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
    allow_ip_spoofing = true
  }

  dedicated_host = "7eb4e35b-4257-56f8-d7da-326d85452592"
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
* `dedicated_host` - (Optional, string, ForceNew) The placement restrictions to use for the virtual server instance. Unique Identifier of the Dedicated Host where the instance will be placed
* `dedicated_host_group` - (Optional, string, ForceNew) The placement restrictions to use for the virtual server instance. Unique Identifier of the Dedicated Host Group where the instance will be placed
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

**NOTE**: `volume_attachments` Volume attachments are not supported for instance groups, so if you plan to use this instance template with an instance group do not include volume attachment for Data volumes. Otherwise, you can add one or more secondary data volumes to be included in the template to be used when you provision the instances with the Template

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id of the instance template
* `name` - Instance Template name
* `volume_attachments` - Collection of volume attachments. Nested `volume_attachments` blocks have the following structure:
	* `delete_volume_on_instance_delete` - If set to true, when deleting the instance the volume will also be deleted.
	* `name` - The user-defined name for this volume attachment.
	* `volume` - The identity of the volume to attach to the instance, or a prototype object for a newvolume.
* `primary_network_interfaces` - A nested block describing the primary network interface for the template
  * `allow_ip_spoofing` - Indicates whether source IP spoofing is allowed on this interface.
  * `subnet` - The VPC subnet to assign to the interface. 
  * `name` - Name of the interface.
  * `primary_ipv4_address` - Pv4 address assigned to the primary network interface.
  * `security_groups` - List of security groups under the subnet.
* `network_interfaces` - A nested block describing the network interfaces for the template.
  * `allow_ip_spoofing` - Indicates whether source IP spoofing is allowed on this interface.
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
* `resource_group` - Resource group ID.
* `user_data` - User data provided for the instance.
* `image` - The ID of the image to used to create the template.
* `placement_target` - The placement restrictions to use for the virtual server instance.
  * `id` - The unique identifier for this dedicated host
  * `crn` - The CRN for this dedicated host.
  * `href` - The URL for this dedicated host.

## Import

`ibm_is_instance_template` can be imported using instance template ID, eg ibm_is_instance_template.template

```
$ terraform import ibm_is_instance_template.template r006-14140f94-fcc4-1349-96e7-a72734715115
```
