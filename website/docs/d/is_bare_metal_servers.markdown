---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Servers"
description: |-
  Manages IBM Cloud Bare Metal Servers.
---

# ibm\_is_bare_metal_servers

Import the details of an existing IBM Cloud vBare Metal Server collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_bare_metal_servers" "ds_bmservers" {
}

```

## Attribute Reference

The following attributes are exported:

* `servers` - List of all instances in the IBM Cloud Infrastructure.
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
  * `boot_volume` - A nested block describing the boot volume.
  
