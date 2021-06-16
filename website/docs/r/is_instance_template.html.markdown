---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_template"
description: |-
  Manages IBM VPC instance template.
---

# ibm_is_instance_template
Create, update, or delete an instance template on VPC.

## Example usage
The following example creates an instance template in a VPC generation-2 infrastructure.

```terraform
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
        volume_prototype {
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `boot_volume` - (Optional, List) A nested block describes the boot volume configuration for the template.

  Nested scheme for `boot_volume`:
	- `delete_volume_on_instance_delete` - (Optional, Bool) You can configure to delete the boot volume based on instance deletion.
	- `encryption` - (Optional, String) The encryption key CRN to encrypt the boot volume attached.
	- `name` - (Optional, String) The name of the boot volume.
- `dedicated_host` - (Optional, Force new resource,String) The placement restrictions to use for the virtual server instance. Unique Identifier of the dedicated host where the instance is placed.
- `dedicated_host_group` - (Optional, Force new resource, String) The placement restrictions to use for the virtual server instance. Unique Identifier of the dedicated host group where the instance is placed.
- `image` - (Required, String) The ID of the image to create the template.
- `keys` - (Required, List) List of SSH key IDs used to allow log in user to the instances.
- `name` - (Required, String) The name of the instance template.
- `profile` - (Required, String) The number of instances created in the instance group.
- `primary_network_interfaces` (Required, List) A nested block describes the primary network interface for the template.

  Nested scheme for `primary_network_interfaces`:
	- `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on this interface. If set to **false** IP spoofing is prevented on the interface. If set to **true**, IP spoofing is allowed on the interface.
	- `name` - (Optional, String) The name of the interface.
	- `primary_ipv4_address` - (Optional, String) The IPv4 address assigned to the primary network interface.
  - `security_groups`- (Optional, List) List of security groups of the subnet.
  - `subnet` - (Required, Force new resource, String) The VPC subnet to assign to the interface.
- `network_interfaces` - (Optional, List) A nested block describes the network interfaces for the template.

  Nested scheme for `network_interfaces`:
	- `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on this interface. If set to **false** IP spoofing is prevented on the interface. If set to **true**, IP spoofing is allowed on the interface.
	- `name` - (Optional, String) The name of the interface.
	- `primary_ipv4_address` - (Optional, String) The IPv4 address assigned to the network interface.
  - `security_groups` - (Optional, List) List of security groups of the subnet.
  - `subnet` - (Required, Forces new resource, String) The VPC subnet to assign to the interface.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID.
- `volume_attachments` - (Optional, List) A nested block describes the storage volume configuration for the template.

  Nested scheme for `volume_attachments`:
	- `name` - (Required, String) The name of the boot volume.
	- `volume` - (Required, String) The storage volume ID created in VPC.
  - `delete_volume_on_instance_delete`- (Required, Bool) You can configure to delete the storage volume to delete based on instance deletion.
  - `volume_prototype` - (Optional, Force new resource, List)

    Nested scheme for `volume_prototype`:
    - `capacity` - (Optional, Integer) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
    - `encryption_key` - (Optional, String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for the resource.
    - `iops` - (Optional, Integer) The maximum input and output operations per second (IOPS) for the volume.
    - `profile` - (Optional, String) The global unique name for the volume profile to use for the volume.
    
    **Note** 
    
    `volume_attachments` provides either `volume` with a storage volume ID, or `volume_prototype` to create a new volume. If you plan to use this template with instance group, provide the `volume_prototype`. Instance group does not support template with existing storage volume IDs.
- `vpc` - (Required, String) The VPC ID that the instance templates needs to be created.
- `user_data` -  (Optional, String) The user data provided for the instance.
- `zone` - (Required, String) The name of the zone.

## Attribute reference
In addition to all arguments listed, you can access the following attribute references after your resource is created.

- `id` - (String) The ID of an instance template.

## Import
The `ibm_is_instance_template` resource can be imported by using instance template ID.

**Example**

```
$ terraform import ibm_is_instance_template.template r006-14140f94-fcc4-1349-96e7-a71212125115
```
