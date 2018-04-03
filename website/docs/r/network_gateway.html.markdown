---
layout: "ibm"
page_title: "IBM: ibm_network_gateway"
sidebar_current: "docs-ibm-resource-network-gateway"
description: |-
  Manages IBM Network Gateway.
---

# ibm\_network_gateway

Provides a resource for an IBM Cloud network gateway appliance. This resource allows a network gateway to be created, updated, and deleted.  

A network gateway can be created in standalone mode and HA mode with both members, with either the same or different configurations.

For additional details, see the [IBM Cloud infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Network_Gateway).

For more information about getting started, see the [IBM Virtual Router Appliance docs](https://console.bluemix.net/docs/infrastructure/virtual-router-appliance/getting-started.html#getting-started).

## Example Usage

### Standalone configuration

```hcl
resource "ibm_network_gateway" "gateway" {
  name = "my-gateway"

  members = [{
    hostname             = "host-name"
    domain               = "ibm.com"
    datacenter           = "ams01"
    network_speed        = 100
    private_network_only = false
    tcp_monitoring       = true
    process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
    redundant_network    = false
    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
    public_bandwidth     = 20000
    memory               = 8
    tags                 = ["gateway tags 1", "terraform test tags 1"]
    notes                = "gateway notes 1"
    ipv6_enabled         = true
  },
  ]
}
```
### HA configuration

```hcl
resource "ibm_network_gateway" "gateway" {
  name = "my-ha-gateway"

  members = [{
    hostname             = "host-name-1"
    domain               = "ibm.com"
    datacenter           = "ams01"
    network_speed        = 100
    private_network_only = false
    tcp_monitoring       = true
    process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
    redundant_network    = false
    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
    public_bandwidth     = 20000
    memory               = 8
    tags                 = ["gateway tags", "terraform test tags 1"]
    notes                = "gateway notes"
    ipv6_enabled         = true
  },
    {
      hostname             = "host-name-2"
      domain               = "ibm.com"
      datacenter           = "ams01"
      network_speed        = 100
      private_network_only = false
      tcp_monitoring       = true
      process_key_name     = "INTEL_SINGLE_XEON_1270_3_50"
      os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
      redundant_network    = false
      disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
      public_bandwidth     = 20000
      memory               = 8
      tags                 = ["gateway tags 1", "terraform test tags 1"]
      notes                = "my ha mode gateway"
      ipv6_enabled         = true
    },
  ]
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the gateway.
* `ssh_key_ids` - (Optional, list) The SSH key IDs to install on the gateway when the gateway gets created.
* `post_install_script_uri` - (Optional, string) The URI of the script to be downloaded and executed after the gateway installation is complete. Default value: `nil`. 
* `members` - (Required, list) A nested block describing the hardware members of this network gateway.
Nested `members` blocks have the following structure:
  * `hostname` - (Optional, string) Hostname of the member.
  * `domain` - (Required, string) The domain of the member
  * `notes` - (Optional, string) Descriptive text of up to 1000 characters about the member.
  * `datacenter` - (Required, string) The data center in which you want to provision the member.
  * `network_speed` - (Optional, integer) The connection speed (in Mbps) for the member network components. Default value: `100`.
  * `redundant_power_supply` - (Optional, boolean) When the value is `true`, an additional power supply is provided. Default value: `false`
  * `tcp_monitoring` - (Optional, boolean) Whether to enable TCP monitoring for the member. Default value: `false`.
  * `process_key_name` - (Optional, string) The process key name for the member. Default value:  `INTEL_SINGLE_XEON_1270_3_40_2`. Refer to the same attribute on the `ibm_compute_bare_metal` resource.
  * `package_key_name` - (Optional, string) The key name for the network gateway package. You can find available package key names in the [Softlayer API](https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_GATEWAY"}}}), using your API key as the password. Default value: `NETWORK_GATEWAY_APPLIANCE`. The default will allow to order Single Processor Multi-Core Servers. Use `2U_NETWORK_GATEWAY_APPLIANCE_1O_GBPS` for ordering Dual Processor Multi-Core Servers.
  * `os_key_name` - (Optional, string) The os key name for member. Default value:  `OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT`. Refer to the same attribute on the `ibm_compute_bare_metal` resource.
  * `redundant_network` - (Optional, boolean) When the value is `true`, two physical network interfaces are provided with a bonding configuration. Default value: `false`. 
  * `unbonded_network` - (Optional, boolean) When the value is `true`, two physical network interfaces are provided without a bonding configuration. Default value: `false`. 

  * `tags` - (Optional, set) Tags associated with the VM instance. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.
  * `public_bandwidth` - (Optional, integer) Allowed public network traffic (in GB) per month. Default value: `20000`.
  * `memory` - (Required, integer) The amount of memory, expressed in megabytes, that you want to allocate.
  * `storage_groups` - (Optional, list) A nested block describing the storage group for the member of the network gateway. Nested `storage_groups` blocks have the following structure:
      * `array_type_id` - (Required, integer) The ID of the array type.
      * `hard_drives` - (Required, list) The list of hard drive associated with the gateway member.
      * `array_size` - (Optional, integer) The size of the array.
      * `partition_template_id` - (Optional, integer) The partition template ID for the member.
  * `ssh_key_ids` - (Optional, list) The SSH key IDs to install on the member.
  * `post_install_script_uri` - (Optional, string) The URI of the script to be downloaded and executed on the member. Default value: `nil`.
  * `user_metadata` - (Optional, string) Arbitrary data to be made available to the member.
  * `disk_key_names` - (Optional, list) Provide the disk key name. Refer to the same attribute in the `ibm_compute_bare_metal` resource.
  * `public_vlan_id` - (Optional, integer) ID of the public VLAN.
  * `private_vlan_id` - (Optional, integer) ID of the private VLAN.  
    **NOTE**: If there are two members in this gateway, then both should have same value for `public_vlan_id` and `private_vlan_id`.
    
  * `ipv6_enabled` - (Optional, boolean) Whether to enable IPv6. Default value: `true`.
  * `private_network_only` - (Optional, boolean) Whether to enable a private network only. Default value: `false`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the network gateway.
* `public_ipv4_address` - The public IP address of the network gateway.
* `private_ipv4_address` - The private IP address of the network gateway.
* `private_ip_address_id` - The private IP address ID of the network gateway.
* `private_vlan_id` - The private VLAN ID of the network gateway.
* `public_ip_address_id` - The public IP address ID of the network gateway.
* `public_ipv6_address_id` - The public IPv6 address ID for the network gateway.
* `public_vlan_id` - The public VLAN ID for the network gateway.
* `status` - Status of the network gateway.
* `associated_vlans` - A nested block describing the associated VLANs for the member of the network gateway. Nested `associated_vlans` blocks export the following attributes:
  * `vlan_id` - The VLAN ID.
  * `network_vlan_id` - The ID of the VLAN that is associated.
  * `bypass` -  Indicates if the VLAN is in bypass or routed mode.
* `members` - A nested block describing the hardware members of this network gateway.
Nested `members` blocks export the following attributes:
  * `member_id` -  ID of the member.
  * `public_ipv4_address` - Public IPv4 address associated with the member.
  * `private_ipv4_address` - Private IPv4 address associated with the member.
  * `ipv6_address` -  IPv6 address associated with the member.
