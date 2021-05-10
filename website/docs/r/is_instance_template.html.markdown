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

resource "ibm_is_volume" "datavol" {
  name           = "datavol1"
  resource_group = data.ibm_resource_group.default.id
  zone           = "us-south-2"

  profile  = "general-purpose"
  capacity = 50
}

data "ibm_resource_group" "default" {
  is_default = true
}

// Create a new volume with the volume attachment. This template format can be used with instance groups
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
   volume_attachments {
        delete_volume_on_instance_delete = true
        name                             = "volatt-01"
        new_volume {
            iops = 3000
            profile = "general-purpose"
            capacity = 200
        }
    }
}

// Template with volume attachment that attaches exisiting storage volume. This template cannot be used with instance groups
resource "ibm_is_instance_template" "instancetemplate2" {
  name    = "testtemplate1"
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
   volume_attachments {
        delete_volume_on_instance_delete = true
        name                             = "volatt-01"
        volume                           = ibm_is_volume.datavol.id
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
  * `name` - (Required, string) Name of the volume attachment.
  * `volume` - (Optional, string, ForceNew) Storage volume ID created under VPC.
  * `delete_volume_on_instance_delete` - (Required, bool) Configured to delete the storage volume to be deleted upon instance deletion.
  * `new_volume` - (Optional, list, ForceNew)
    * `iops` - (Optional, int) The maximum I/O operations per second (IOPS) for the volume.
    * `profile` - (Optional, string) The  globally unique name for the volume profile to use for this volume.
    * `capacity` - (Optional, int) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
    * `encryption_key` - (Optional, string) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
* `user_data` - (Optional, string) User data provided for the instance.

**NOTE**: `volume_attachments` Provide either 'volume'  with a storage volume id, or 'new_volume' to create a new volume. If you plan to use this template with instance group, provide the 'new_volume'. Instance group does not support template with existing storage volume IDs.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id of the instance template
* `primary_network_interfaces` - (Required, list) A nested block describing the primary network interface for the template. Nested  primary_network_interface block have the following structure:
  * `subnet` - (Required, Forces new resource, string) The VPC subnet to assign to the interface. 
  * `name` - (Optional, string) Name of the interface.
  * `primary_ipv4_address` - (Optional, string) IPv4 address assigned to the primary network interface.
  * `security_groups` - (Optional, list) List of security groups under the subnet.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `volume_attachments` - (Optional, list) A nested block describing the storage volume configuration for the template. Nested volume_attachments blocks have the following structure: 
  * `name` - (Required, string) Name of the volume attachment.
  * `volume` - (Optional, string, ForceNew) Storage volume ID created under VPC.
  * `delete_volume_on_instance_delete` - (Required, bool) Configured to delete the storage volume to be deleted upon instance deletion.
  * `new_volume` - (Optional, list, ForceNew)
    * `iops` - (Optional, int) The maximum I/O operations per second (IOPS) for the volume.
    * `profile` - (Optional, string) The  globally unique name for the volume profile to use for this volume.
    * `capacity` - (Optional, int) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
    * `encryption_key` - (Optional, string) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
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
* `resource_group` - (Optional, Forces new resource, string) Resource group ID.

## Import

`ibm_is_instance_template` can be imported using instance template ID, eg ibm_is_instance_template.template

```
$ terraform import ibm_is_instance_template.template r006-14140f94-fcc4-1349-96e7-a72734715115
```
