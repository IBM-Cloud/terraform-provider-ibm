---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance"
description: |-
  Manages IBM VPC instance.
---

# ibm_is_instance
Create, update, or delete a Virtual Servers for VPC instance. For more information, about managing VPC instance, see [about virtual server instances for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-advanced-virtual-servers).


## Example usage

### Sample for creating an instance in a VPC.

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "bc1-2x8"

  boot_volume {
    encryption = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
  }

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
    primary_ipv4_address = "10.240.0.6"
    allow_ip_spoofing = true
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.testacc_subnet.id
    allow_ip_spoofing = false
  }

  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

```

### Sample for creating an instance with custom security group rules.

The following example shows how you can create a virtual server instance with custom security group rules. Note that the security group, security group rules, and the virtual server instance must be created in a specific order to meet the dependencies of the individual resources. To force the creation in a specific order, you use the [`depends_on` parameter](https://www.terraform.io/docs/configuration/resources.html){: external}. If you do not provide this parameter, all resources are created at the same time which might lead to resource dependency errors during the provisioning of your virtual server, such as `The security group to attach to is not available`.

```terraform

resource "ibm_is_vpc" "testacc_vpc" {
    name = "test"
}

resource "ibm_is_security_group" "testacc_security_group" {
    name = "test"
    vpc = ibm_is_vpc.testacc_vpc.id
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_all" {
    group = ibm_is_security_group.testacc_security_group.id
    direction = "inbound"
    remote = "127.0.0.1"
    depends_on = [ibm_is_security_group.testacc_security_group]
 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp" {
    group = ibm_is_security_group.testacc_security_group.id
    direction = "inbound"
    remote = "127.0.0.1"
    icmp {
        code = 20
        type = 30
    }
    depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_all]

 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
    group = ibm_is_security_group.testacc_security_group.id
    direction = "inbound"
    remote = "127.0.0.1"
    udp {
        port_min = 805
        port_max = 807
    }
    depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_icmp]
 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp" {
    group = ibm_is_security_group.testacc_security_group.id
    direction = "outbound"
    remote = "127.0.0.1"
    tcp {
        port_min = 8080
        port_max = 8080
    }
    depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_udp]
 }

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "bc1-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
    security_groups = [ibm_is_security_group.testacc_security_group.id]
  }

  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
  depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
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

// Example to provision instance in a dedicated host
resource "ibm_is_instance" "testacc_instance1" {
  name    = "testinstance1"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "cx2-2x4"

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
    security_groups = [ibm_is_security_group.testacc_security_group.id]
  }
  dedicated_host = ibm_is_dedicated_host.is_dedicated_host.id
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
  depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

// Example to provision instance in a dedicated host that belongs to the provided dedicated host group
resource "ibm_is_instance" "testacc_instance2" {
  name    = "testinstance2"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "cx2-2x4"

  primary_network_interface {
    subnet = ibm_is_subnet.testacc_subnet.id
    security_groups = [ibm_is_security_group.testacc_security_group.id]
  }
  dedicated_host_group = ibm_is_dedicated_host_group.dh_group01.id
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
  depends_on = [ibm_is_security_group_rule.testacc_security_group_rule_tcp]

  //User can configure timeouts
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

// Example to provision instance from a snapshot, restoring boot volume from an existing snapshot

resource "ibm_is_snapshot" "testacc_snapshot" {
  name 		      	= "testsnapshot"
  source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
}

resource "ibm_is_instance" "testacc_instance_restore" {
  name    = "vsirestore"
  profile = "cx2-2x4"
  boot_volume {
    name     = "boot-restore"
    snapshot = ibm_is_snapshot.testacc_snapshot.id
  }
  primary_network_interface {
    subnet     = ibm_is_subnet.testacc_subnet.id
  }
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.testacc_sshkey.id]
  network_interfaces {
    subnet = ibm_is_subnet.testacc_subnet.id
    name   = "eth1"
  }
}

```

## Timeouts

The `ibm_is_instance` resource provides the following [[Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:


- **create**: The creation of the instance is considered failed when no response is received for 30 minutes.
- **update**: The update of the instance or the attachment of a volume to an instance is considered failed when no response is received for 30 minutes.
- **delete**: The deletion of the instance is considered failed when no response is received for 30 minutes.


## Argument reference
Review the argument references that you can specify for your resource.

- `auto_delete_volume`- (Optional, Bool) If set to **true**, automatically deletes the volumes that are attached to an instance. **Note** Setting this argument can bring some inconsistency in the volume resource, as the volumes is destroyed along with instances.
- `boot_volume`  (Optional, List) A list of boot volumes for an instance.

  Nested scheme for `boot_volume`:
  - `encryption` - (Optional, String) The type of encryption to use for the boot volume.
  - `name` - (Optional, String) The name of the boot volume.
  - `snapshot` - (Optional, Forces new resource, String) The snapshot id of the volume to be used for creating boot volume attachment
    **Note** 
    
     - `snapshot` conflicts with `image` id and `instance_template`
- `dedicated_host` - (Optional, Forces new resource, String) The placement restrictions to use the virtual server instance. Unique ID of the dedicated host where the instance id placed.
- `dedicated_host_group` - (Optional, Forces new resource, String) The placement restrictions to use for the virtual server instance. Unique ID of the dedicated host group where the instance is placed.
- `force_recovery_time` - (Optional, Integer) Define timeout (in minutes), to force the `is_instance` to recover from a perpetual "starting" state, during provisioning. And to force the is_instance to recover from a perpetual "stopping" state, during removal of user access. **Note** The force_recovery_time is used to retry multiple times until timeout.
- `image` - (Optional, String) The ID of the virtual server image that you want to use. To list supported images, run `ibmcloud is images`.
  **Note** 
    
  - `image` conflicts with `boot_volume.0.snapshot`  
- `keys` - (Optional, List) A comma-separated list of SSH keys that you want to add to your instance.
- `name` - (Optional, String) The instance name.
- `network_interfaces`  (Optional,  Forces new resource, List) A list of more network interfaces that are set up for the instance.

  Nested scheme for `network_interaces`:
  - `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on the interface. If **false**, IP spoofing is prevented on the interface. If **true**, IP spoofing is allowed on the interface.
  - `name` - (Optional, String) The name of the network interface.
  - `primary_ipv4_address` - (Optional, Forces new resource, String) The IPV4 address of the interface.
  - `subnet` - (Required, String) The ID of the subnet.
  - `security_groups`- (Optional, List of strings)A comma separated list of security groups to add to the primary network interface.
- `primary_network_interface` - (Optional, List) A nested block describes the primary network interface of this instance. Only one primary network interface can be specified for an instance.

  Nested scheme for `primary_network_interface`:
  - `allow_ip_spoofing`- (Optional, Bool) Indicates whether IP spoofing is allowed on the interface. If **false**, IP spoofing is prevented on the interface. If **true**, IP spoofing is allowed on the interface.
  - `name` - (Optional, String) The name of the network interface.
  - `port_speed` - (Deprecated, Integer) Speed of the network interface.
  - `primary_ipv4_address` - (Optional, Forces new resource, String) The IPV4 address of the interface.
  - `subnet` - (Required, String) The ID of the subnet.
  - `security_groups`-List of strings-Optional-A comma separated list of security groups to add to the primary network interface.
- `profile` - (Optional, Forces new resource, String) The name of the profile that you want to use for your instance. To list supported profiles, run `ibmcloud is instance-profiles`.
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the instance.
- `instance_template` - (Optional, String) ID of the source template.
  **Note** 
    
  - `instance_template` conflicts with `boot_volume.0.snapshot`  
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your instance. Tags can help you find your instance more easily later.
- `user_data` - (Optional, String) User data to transfer to the instance.
- `volumes`  (Optional, List) A comma separated list of volume IDs to attach to the instance.
- `vpc` - (Optional, Forces new resource, String) The ID of the VPC where you want to create the instance.
- `zone` - (Optional, Forces new resource, String) The name of the VPC zone where you want to create the instance.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `boot_volume`- (List of Strings) A list of boot volumes that the instance uses.

  Nested scheme for `boot_volume`:
  - `encryption` - (String) The type of encryption that is used for the boot volume.
  - `iops`- (Integer) The number of input and output operations per second of the volume.
  - `name` - (String) The name of the boot volume.
  - `profile` - (String) The profile of the volume.
  - `size`- (Integer) The capacity of the volume in gigabytes.
- `disks` - (List of Strings) The collection of the instance's disks. Nested `disks` blocks have the following structure:

  Nested scheme for `disks`:
  - `created_at` - (Timestamp) The date and time that the disk was created.
  - `href` - (String) The URL for the instance disk.
  - `id` - (String) The unique identifier for the instance disk.
  - `interface_type` - (String) The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing the property, check for the unknown log values. Optionally stop processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
  - `name` - (String) The user defined name for the disk.
  - `resource_type` - (String) The resource type.
  - `size` - (String) The size of the disk in GB (gigabytes).
- `gpu`- (List of Strings) A list of GPUs that are assigned to the instance.

  Nested scheme for `gpu`:
  - `cores`- (Integer) The number of cores of the GPU.
  - `count`- (Integer) The count of the GPU.
  - `manufacture` - (String) The manufacturer of the GPU.
  - `memory`- (Integer) The amount of memory of the GPU in gigabytes.
  - `model` - (String) The model of the GPU.
- `id` - (String) The ID of the instance.
- `memory`- (Integer) The amount of memory that is allocated to the instance in gigabytes.
- `network_interfaces`- (List of Strings) A list of more network interfaces that are attached to the instance.

  Nested scheme for `network_interfaces`:
  - `allow_ip_spoofing` - (String) Indicates whether IP spoofing is allowed on the interface.
  - `id` - (String) The ID of the network interface.
  - `name` - (String) The name of the network interface.
  - `subnet` - (String) The ID of the subnet.
  - `security_groups`- (List of Strings) A list of security groups that are used in the network interface.
  - `primary_ipv4_address` - (String) The primary IPv4 address.
- `primary_network_interface`- (List of Strings) A list of primary network interfaces that are attached to the instance.

  Nested scheme for `primary_network_interface`:
  - `allow_ip_spoofing` - (String) Indicates whether IP spoofing is allowed on the interface.
  - `id` - (String) The ID of the primary network interface.
  - `name` - (String) The name of the primary network interface.
  - `subnet` - (String) The ID of the subnet that the primary network interface is attached to.
  - `security_groups`-List of strings-A list of security groups that are used in the primary network interface.
  - `primary_ipv4_address` - (String) The primary IPv4 address.
- `status` - (String) The status of the instance.
- `volume_attachments`- (List of Strings) A list of volume attachments for the instance.

  Nested scheme for `volume_attachements`:
  - `id` - (String) The ID of the volume attachment.
  - `name` - (String) The name of the volume attachment.
  - `volume_id` - (String) The ID of the volume that is used in the volume attachment.
  - `volume_name` - (String) The name of the volume that is used in the volume attachment.
  - `volume_crn` - (String) The CRN of the volume that is used in the volume attachment.
- `vcpu`- (List of Strings) A list of virtual CPUs that are allocated to the instance.

  Nested scheme for `vcpu`:
  - `architecture` - (String) The architecture of the CPU.
  - `count`- (Integer) The number of virtual CPUS that are assigned to the instance.


## Import
The `ibm_is_instance` resource can be imported by using the instance ID.

**Example**

```
$ terraform import ibm_is_instance.example a1aaa111-1111-111a-1a11-a11a1a11a11a
```
