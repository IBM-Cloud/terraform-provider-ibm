---
layout: "ibm"
page_title: "IBM: ibm_network_gateway"
sidebar_current: "docs-ibm-resource-network-gateway"
description: |-
  Manages IBM Network Gateway.
---

# ibm\_network_gateway

Provides a resource for IBM network gateway appliance. This allows network gateway to be created, updated, and deleted.
Currently gateway can be created with standalone mode and HA mode with both members either same or different configurations.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Network_Gateway).

Getting started documentation can be found [here](https://console.bluemix.net/docs/infrastructure/virtual-router-appliance/getting-started.html#getting-started)

## Example Usage

```hcl
resource "ibm_network_gateway" "gateway" {
  name = "my-gateway"

  members {
    hostname             = "host-name"
    domain               = "terraformuat1.ibm.com"
    datacenter           = "ams01"
    network_speed        = 100
    private_network_only = false
    public_vlan_id       = 1234
    private_vlan_id      = 4567
    tcp_monitoring       = true
    ssh_key_ids          = [1234]
    tags                 = ["gateway"]
    notes                = "my gateway"
    process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
    public_bandwidth     = 20000
    memory               = 4
    ipv6_enabled         = true
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the gateway.
* `ssh_key_ids` - (Optional, list) The SSH key IDs to install on the gateway when the gateway gets created.
* `post_install_script_uri` - (Optional, string) Default `nil`. The URI of the script to be downloaded and executed after gateway installation is complete.
* `members` - (Required, list)A nested block describing the hardware members of this network Gateway.
Nested `members` blocks have the following structure:
  * `hostname` - (Optional, string) Hostname of the member.
  * `domain` - (Required, string) The domain of the member
  * `notes` - (Optional, string) Descriptive text of up to 1000 characters about the member.
  * `datacenter` - (Required, string) The datacenter in which you want to provision the member.
  * `network_speed` - (Optional, interger) The connection speed (in Mbps) for the member network components. The default value is `100`.
  * `tcp_monitoring` - (Optional, boolean) Whether to enable tcp monitoring for the member. Default is `false`.
  * `process_key_name` - (Optional, string) The process key name for the member. Default is `INTEL_SINGLE_XEON_1270_3_40_2`. Refer to the same attribute on the `ibm_compute_bare_metal` to know more about this field.
  * `os_key_name` - (Optional, string) The os key name for member. The default is `OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT`. Refer to the same attribute on the `ibm_compute_bare_metal` to know more about this field.
  * `redundant_network` - (Optional, boolean) When the value is `true`, two physical network interfaces are provided with a bonding configuration. The default value is `false`.
  * `unbonded_network` - (Optional, boolean) When the value is `true`, two physical network interfaces are provided without a bonding configuration. The default value is `false`.
  * `tags` - (Optional, set) Tags associated with the VM instance. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.
  * `public_bandwidth` - (Optional, integer) Allowed public network traffic(GB) per month. Default value is `20000`.
  * `memory` - (Required, integer) The amount of memory, expressed in megabytes, that you want to allocate.
  * `storage_groups` - (Optional, list) A nested block describing the strorage group for the member of network gateway. Nested `storage_groups` blocks have the following structure:
      * `array_type_id` - (Required, integer) The ID of the array type.
      * `hard_drives` - (Required, list) The list og hard drive associated with gateway member.
      * `array_size` - (Optional, integer) The size of the array.
      * `partition_template_id` - (Optional, integer) The partition template id for the member.
  * `ssh_key_ids` - (Optional, list) The SSH key IDs to install on the member.
  * `post_install_script_uri` - (Optional, string) The URI of the script to be downloaded and executed on member.The Default value is `nil`.
  * `user_metadata` - (Optional, string) Arbitrary data to be made available to the member.
  * `disk_key_names` - (Optional, list) Provide the disk key name. Refer to the same attribute on the `ibm_compute_bare_metal` to know more about this field.
  * `public_vlan_id` - (Optional, integer) ID of the public vlan.
  * `private_vlan_id` - (Optional, integer) ID of the private vlan.
    **NOTE** : If there are two members in this gateway then both should have same value for `public_vlan_id` and `private_vlan_id`.
  * `ipv6_enabled` - (Optional, boolean) Whether to enable ipv6 . Default is `true`.
  * `private_network_only` - (Optional, integer) Whether to enable private network only.The default vale is `false`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the Network gateway.
* `private_ip_address_id` - The private ip address of the network gateway.
* `private_vlan_id` - The private vlan id of the network gateway.
* `public_ip_address_id` - The public ip address of the network gateway.
* `public_ipv6_address_id` - The public ipv6 address for the network gateway.
* `public_vlan_id` - The public vlan id for the network gateway.
* `status` - Status of the network gateway.
* `associated_vlans` - A nested block describing the associated vlans for the member of network gateway. Nested `associated_vlans` blocks export the following attributes:
  * `vlan_id` - Vlan ID.
  * `network_vlan_id` - The Identifier of the VLAN which is associated.
  * `bypass` -  Indicates if the VLAN is in bypass or routed modes
* `members` - A nested block describing the hardware members of this network Gateway.
Nested `members` blocks export the following attributes:
  * `member_id` -  ID of the member.
  * `public_ipv4_address` - Public IPv4 address associated with the member.
  * `private_ipv4_address` - Private IPv4 address associated with the member.
  * `ipv6_address` -  IPv6 address associated with the member.

