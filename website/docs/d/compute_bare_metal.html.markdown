---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_bare_metal"
description: |-
  Get information of an IBM Cloud compute bare metal
---

# ibm_compute_bare_metal
Retrieve information of an existing bare metal as a read-only data source. For more details, about compute bare metal, see [compute services](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-compute).

## Example usage

```terraform
data "ibm_compute_bare_metal" "bare_metal" {
  hostname    = "jumpbox"
  domain      = "example.com"
  most_recent = true
}

data "ibm_compute_bare_metal" "bare_metal" {
  global_identifier = "a471e9a6-82e7-41a7-ac8d-39ced672c0ed"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `domain` - (Optional, String) The domain of the Bare Metal server. If you specify this option, do not specify `global_identifier` at the same time.
- `global_identifier` - (Optional, String) The unique global identifier of the Bare Metal server. To see global identifier, log in to the [IBM Cloud Classic Infrastructure API](https://api.softlayer.com/rest/v3.1/SoftLayer_Account/getHardware.json), that uses your API key as the password. If you specify this option, do not specify `hostname`, `domain`, or `most_recent` at the same time.
- `hostname` - (Optional, String) The hostname of the Bare Metal server. If you specify the `hostname`, do not specify `global_identifier` at the same time.
- `most_recent` - (Optional, Bool) For multiple Bare Metal services, you can set this argument to **true** to import only the most recently created server. If you specify this option, do not specify `global_identifier` at the same time.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `block_storage_ids`- (List of string) Block storage to which this computing server has access.
- `datacenter` - (String) The data center in which the Bare Metal server is deployed.
- `file_storage_ids`- (List of string) File storage to which this computing server has access.
- `hourly_billing` - (String) The billing type of the server.
- `id` - (String) The unique identifier of the Bare Metal server.
- `ipv6_enabled`- (Bool) Indicates whether the public IPv6 address is enabled or not.
- `ipv6_address` - (String) The public IPv6 address of the Bare Metal server.
- `ipv6_address_id` - (String) The unique identifier for the public IPv6 address of the Bare Metal server.
- `memory`- (Integer) The amount of memory in gigabytes, for the server.
- `network_speed` - (String) The connection speed, expressed in Mbps, for the server network components.
- `notes`-  (String) Notes associated with the server.
- `os_reference_code` - (String) An operating system reference code that provisioned the computing server.
- `public_bandwidth` - (String) The amount of public network traffic, allowed per month.
- `public_ipv4_address` - (String) The public IPv4 address of the Bare Metal server.
- `public_ipv4_address_id` - (String) The unique identifier for the public IPv4 address of the Bare Metal server.
- `private_ipv4_address` - (String) The private IPv4 address of the Bare Metal server.
- `private_ipv4_address_id` - (String) The unique identifier for the private IPv4 address of the Bare Metal server.
- `public_vlan_id` - (String) The public VLAN used for the public network interface of the server.
- `private_vlan_id` - (String) The private VLAN used for the private network interface of the server.
- `public_subnet` - (String) The public subnet used for the public network interface of the server.
- `private_subnet` - (String) The private subnet used for the private network interface of the server.
- `private_network_only` - (String) Specifies whether the server has only access to the private network.
- `redundant_power_supply`-  (Bool) When the value is **true**, it indicates that more power supply is provided.
- `redundant_network`- (Bool) When the value is **true**, two physical network interfaces are provided with a bonding configuration.
- `secondary_ip_count`- (Integer) The number of secondary IPv4 addresses of the Bare Metal server.
- `secondary_ip_addresses` - (String) The public secondary IPv4 addresses of the Bare Metal server.
- `tags`- (List of string) Tags associated with this Bare Metal server.
- `user_metadata` - (String) Arbitrary data available to the computing server.
- `unbonded_network`- (Bool) When the value is **true**, two physical network interfaces are provided without a bonding configuration.
