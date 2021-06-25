---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare metal server"
description: |-
  Manages IBM bare metal sever.
---

# ibm\_is_bare_metal_server

Provides a Bare Metal Server resource. This allows Bare Metal Server to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a Bare Metal Server:

```hcl
resource "ibm_is_bare_metal_server" "bms" {
    profile = "mx2d-metal-32x192"
    name = "my-bms"
    image = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
    zone = "us-south-3"
    keys = [ibm_is_ssh_key.sshkey.id]
    primary_network_interface {
      subnet     = ibm_is_subnet.subnet1.id
    }
    vpc = ibm_is_vpc.vpc1.id
}

```

## Timeouts

ibm_is_bare-metal_server provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 30 minutes) Used for creating bare metal server.
* `update` - (Default 30 minutes) Used for updating bare metal server or while attaching it with volume attachments or interfaces.
* `delete` - (Default 30 minutes) Used for deleting bare metal server.

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The bare metal server name.
* `vpc` - (Required, Forces new resource, string) The vpc id. 
* `zone` - (Required, Forces new resource, string) Name of the zone. 
* `profile` - (Required, Forces new resource, string) The profile name. 
* `image` - (Required, string) ID of the image.
* `keys` - (Required, list) Comma separated IDs of ssh keys.  
* `primary_network_interface` - (Required, list) A nested block describing the primary network interface of this bare metal server. We can have only one primary network interface.
Nested `primary_network_interface` block have the following structure:
  * `name` - (Optional, string) The name of the network interface.
  * `port_speed` - (Deprecated, int) Speed of the network interface.
  * `primary_ipv4_address` - (Optional, Forces new resource, string) The IPV4 address of the interface
  * `subnet` -  (Required, string) ID of the subnet.
  * `security_groups` - (Optional, list) Comma separated IDs of security groups.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `network_interfaces` - (Optional, Forces new resource, list) A nested block describing the additional network interface of this bare metal server.
Nested `network_interfaces` block have the following structure:
  * `name` - (Optional, string) The name of the network interface.
  * `primary_ipv4_address` - (Optional, Forces new resource, string) The IPV4 address of the interface
  * `subnet` -  (Required, string) ID of the subnet.
  * `security_groups` - (Optional, list) Comma separated IDs of security groups.
  * `allow_ip_spoofing` - (Optional, bool) Indicates whether IP spoofing is allowed on this interface. If false, IP spoofing is prevented on this interface. If true, IP spoofing is allowed on this interface.
* `user_data` - (Optional, string) User data to transfer to the server bare metal server.
* `resource_group` - (Optional, Forces new resource, string) The resource group ID for this bare metal server.

## Attribute Reference

The following attributes are exported:

* `name` - The name of the bare metal server.
* `id` - The unique identifier for this bare metal server
* `bandwidth` - The total bandwidth (in megabits per second) shared across the bare metal server's network interfaces.
* `boot_target` - The unique identifier for this bare metal server disk.
* `image` - Image used in the bare metal server.
* `zone` - zone of the bare metal server.
* `vpc` - vpc id of the bare metal server.
* `resource_group` - resource group id of the bare metal server.
* `trusted_platform_module` - the trusted platform module of the bare metal server.
  Nested `trusted_platform_module` blocks have the following structure:
  * `enabled` - Indicates whether the trusted platform module (TPM) is enabled. If enabled, mode will also be set.
  * `mode` - The mode for the trusted platform module (TPM) : [ tpm_2, tpm_2_with_txt ]
* `status` - The status of the bare metal server :[ failed, pending, restarting, running, starting, stopped, stopping ]
* `status_reasons` - Array of reasons for the current status (if any).
Nested `status_reasons` blocks have the following structure:
  * `code` - The status reason code
  * `message` - An explanation of the status reason
* `resource_type` - The type of resource referenced
* `profile` - The name for this bare metal server profile
* `memory` - The amount of memory, truncated to whole gibibytes
* `href` - The URL for this bare metal server
* `enable_secure_boot` - Indicates whether secure boot is enabled. If enabled, the image must support secure boot or the server will fail to boot.
* `crn` - The CRN for this bare metal server
* `cpu` - A nested block describing the CPU configuration of this bare metal server.
Nested `cpu` blocks have the following structure:
  * `architecture` - The architecture of the bare metal server.
  * `core_count` - The total number of cores
  * `socket_count` - The total number of CPU sockets
  * `threads_per_core` - The total number of hardware threads per core
* `primary_network_interface` - A nested block describing the primary network interface of this bare metal server.
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


## Import

ibm_is_bare_metal_server can be imported using bare metal server ID , eg

```
$ terraform import ibm_is_bare_metal_server.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
