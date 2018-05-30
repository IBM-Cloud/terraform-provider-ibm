---
layout: "ibm"
page_title: "IBM: ibm_bare_metal"
sidebar_current: "docs-ibm-datasource-compute-bare-metal"
description: |-
  Get information on a IBM Compute Bare Metal
---

# ibm\_compute_bare_metal

Import the details of an existing bare metal as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_compute_bare_metal" "bare_metal" {
  hostname    = "jumpbox"
  domain      = "example.com"
  most_recent = true
}

data "ibm_compute_bare_metal" "bare_metal" {
  global_identifier = "a471e9a6-82e7-41a7-ac8d-39ced672c0ed"
}
```

## Argument Reference

The following arguments are supported:

* `hostname` - (Optional, string) The hostname of the bare metal server.  
  **NOTE**: Conflicts with `global_identifier`.
* `domain` - (Optional, string) The domain of the bare metal server.  
  **NOTE**: Conflicts with `global_identifier`.
* `global_identifier` - (Optional, string) The unique global identifier of the bare metal server. To see global identifier, log in to the [IBM Cloud Infrastructure (SoftLayer) API](https://api.softlayer.com/rest/v3.1/SoftLayer_Account/getHardware.json), using your API key as the password.  
  **NOTE**: Conflicts with `hostname`, `domain`, `most_recent`.
* `most_recent` - (Optional, boolean) If there are multiple bare metals, you can set this argument to `true` to import only the most recently created server.  
   **NOTE**: Conflicts with `global_identifier`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the bare metal server.
* `datacenter` - The data center in which the bare metal server is deployed.
* `network_speed` - The connection speed, expressed in Mbps,  for the server network components.
* `public_bandwidth` - The amount of public network traffic, allowed per month.
* `public_ipv4_address` - The public IPv4 address of the bare metal server.
* `public_ipv4_address_id` - The unique identifier for the public IPv4 address of the bare metal server.
* `private_ipv4_address` - The private IPv4 address of the bare metal server.
* `private_ipv4_address_id` - The unique identifier for the private IPv4 address of the bare metal server.
* `public_vlan_id` - The public VLAN used for the public network interface of the server. 
* `private_vlan_id` - The private VLAN used for the private network interface of the server. 
* `public_subnet` - The public subnet used for the public network interface of the server. 
* `private_subnet` - The private subnet used for the private network interface of the server. 
* `hourly_billing` -  The billing type of the server.
* `private_network_only` - Specifies whether the server only has access to the private network.
* `user_metadata` - Arbitrary data available to the computing server.
* `notes` -  Notes associated with the server.
* `memory` - The amount of memory in gigabytes, for the server.
* `redundant_power_supply` -  When the value is `true`, it indicates additional power supply is provided.
* `redundant_network` - When the value is `true`, two physical network interfaces are provided with a bonding configuration.
* `unbonded_network` - When the value is `true`, two physical network interfaces are provided without a bonding configuration.
* `os_reference_code` - An operating system reference code that provisioned the computing server.
*  `tags` - Tags associated with this bare metal server.
* `block_storage_ids` - Block storage to which this computing server have access.
* `file_storage_ids` - File storage to which this computing server have access.
* `ipv6_enabled` - Indicates whether the public IPv6 address enabled or not.
* `ipv6_address` - The public IPv6 address of the bare metal server.
* `ipv6_address_id` - The unique identifier for the public IPv6 address of the bare metal server.
* `secondary_ip_count` - The number of secondary IPv4 addresses of the bare metal server.
* `secondary_ip_addresses` - The public secondary IPv4 addresses of the bare metal server.
