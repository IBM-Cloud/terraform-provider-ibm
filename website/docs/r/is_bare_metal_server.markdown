---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server"
description: |-
  Manages IBM bare metal sever.
---

# ibm\_is_bare_metal_server

Create, update, or delete a Bare Metal Server for VPC. For more information, about managing VPC Bare Metal Server, see [About Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-bare-metal-servers).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

In the following example, you can create a Bare Metal Server:

### Basic Example Using AMI Lookup
```terraform

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.vpc.id
  zone            = "us-south-3"
  ipv4_cidr_block = "10.240.129.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_bare_metal_server" "example" {
  profile = "mx2d-metal-32x192"
  name    = "example-bms"
  image   = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
  zone    = "us-south-3"
  keys    = [ibm_is_ssh_key.example.id]
  primary_network_interface {
    subnet     = ibm_is_subnet.example.id
  }
  vpc   = ibm_is_vpc.example.id
}

```
### Reserved ip example
```terraform
resource "ibm_is_bare_metal_server" "bms" {
  profile = "mx2d-metal-32x192"
  name    = "example-bms"
  image   = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
  zone    = "us-south-3"
  keys    = [ibm_is_ssh_key.example.id]
  primary_network_interface {
    subnet     = ibm_is_subnet.example.id
    primary_ip {
      auto_delete = true
      name        = "example-reserved-ip"
      address     = "${replace(ibm_is_subnet.example.ipv4_cidr_block, "0/28", "14")}"
    }
  }
  vpc   = ibm_is_vpc.example.id
}

```

## Timeouts

ibm_is_bare-metal_server provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- `create` - (Default 30 minutes) Used for creating bare metal server.
- `update` - (Default 30 minutes) Used for updating bare metal server or while attaching it with volume attachments or interfaces.
- `delete` - (Default 30 minutes) Used for deleting bare metal server.

## Argument Reference

Review the argument references that you can specify for your resource. 

- `delete_type` - (Optional, String) Type of deletion on destroy. **soft** signals running operating system to quiesce and shutdown cleanly, **hard** immediately stop the server. By default its `hard`.
- `image` - (Required, String) ID of the image.
- `keys` - (Required, List) Comma separated IDs of ssh keys.  
- `name` - (Optional, String) The bare metal server name.

  -> **NOTE:**
    a bare metal server can take up to 30 mins to clean up on delete, replacement/re-creation using the same name will return error

- `primary_network_interface` - (Required, List) A nested block describing the primary network interface of this bare metal server. We can have only one primary network interface.
  
  Nested scheme for `primary_network_interface`:
    - `allow_ip_spoofing` - (Optional, Boolean) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface. [default : `false`]
    - `allowed_vlans` - (Optional, Array) Comma separated VLANs, Indicates what VLAN IDs (for VLAN type only) can use this physical (`PCI` type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.
    - `enable_infrastructure_nat` - (Optional, Boolean) If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations. [default : `true`]
    - `name` - (Optional, String) The name of the network interface.
    - `primary_ip` - (Optional, List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
        - `address` - (Optional, String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
        - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
        - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP. `reserved_ip` is mutually exclusive with rest of the `primary_ip` attributes.
        - `name`- (Optional, String) The user-defined or system-provided name for this reserved IP
        
    - `security_groups` - (Optional, Array) Comma separated IDs of security groups.
    - `subnet` -  (Required, String) ID of the subnet to associate with.

- `profile` - (Required, Forces new resource, String) The name the profile to use for this bare metal server. 
- `resource_group` - (Optional, Forces new resource, String) The resource group ID for this bare metal server.
- `user_data` - (Optional, String) User data to transfer to the server bare metal server.
- `vpc` - (Required, Forces new resource, String) The VPC ID of the bare metal server is to be a part of. It must match the VPC tied to the subnets of the server's network interfaces.
- `zone` - (Required, Forces new resource, String) Name of the zone in which this bare metal server will reside in.


## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `bandwidth` - (Integer) The total bandwidth (in megabits per second) shared across the bare metal server's network interfaces.
- `boot_target` - (String) The unique identifier for this bare metal server disk.
- `crn` - (String) The CRN for this bare metal server
- `cpu` - (String) A nested block describing the CPU configuration of this bare metal server.
  
  Nested scheme for `cpu`:
    - `architecture` - (String) The architecture of the bare metal server.
    - `core_count` - (Integer) The total number of cores
    - `socket_count` - (Integer) The total number of CPU sockets
    - `threads_per_core` - (Integer) The total number of hardware threads per core
- `href` - (String) The URL for this bare metal server
- `id` - (String) The unique identifier for this bare metal server
- `memory` - (Integer) The amount of memory, truncated to whole gibibytes
- `network_interfaces` - (List) The additional network interfaces to create for the bare metal server to this bare metal server. Use `ibm_is_bare_metal_server_network_interface` resource for network interfaces.
  
  Nested scheme for `network_interfaces`:
    - `allow_ip_spoofing` - (Boolean) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface. [default : `false`]
    - `allowed_vlans` - (Array) Comma separated VLANs, Indicates what VLAN IDs (for VLAN type only) can use this physical (`PCI` type) interface. A given VLAN can only be in the `allowed_vlans` array for one PCI type adapter per bare metal server.  [ conflicts with `vlan`]
    - `enable_infrastructure_nat` - (Boolean) If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations. [default : `true`]
    - `name` - (String) The name of the network interface.
    - `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

      Nested scheme for `primary_ip`:
        - `address` - (String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
        - `reserved_ip`- (String) The unique identifier for this reserved IP.
        - `name`- (String) The user-defined or system-provided name for this reserved IP
      
    - `security_groups` - (Array) Comma separated IDs of security groups.
    - `subnet` -  (String) ID of the subnet to associate with.
    - `vlan` -  (Integer) Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface. [ conflicts with `allowed_vlans`]

- `resource_type` - (String) The type of resource.
- `status` - (String) The status of the bare metal server :[ **failed**, **pending**, **restarting**, **running**, **starting**, **stopped**, **stopping** ]
- `status_reasons` - (List) Array of reasons for the current status (if any).

  Nested `status_reasons`:
    - `code` - (String) The status reason code
    - `message` - (String) An explanation of the status reason
    - `more_info` - (String) Link to documentation about this status reason


## Import

The `ibm_is_bare_metal_server` can be imported using Bare Metal Server ID


## Syntax
```
$ terraform import ibm_is_bare_metal_server.example <bare_metal_server_id>
```

**Example**

```
$ terraform import ibm_is_bare_metal_server.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
