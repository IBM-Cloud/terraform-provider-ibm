---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_network_gateway"
description: |-
  Manages IBM Cloud network gateway.
---

# ibm_network_gateway
Create, update, and delete an IBM Cloud network gateway appliance. A network gateway can be created in stand-alone mode and HA mode with both members, with either the same or different configurations. For more information, about Network gateway, see [viewing gateway appliance details](https://cloud.ibm.com/docs/gateway-appliance?topic=gateway-appliance-viewing-gateway-appliance-details).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer)  API Docs](http://sldn.softlayer.com/reference/services/SoftLayer_Network_Gateway).

For more information, about getting started, see the [IBM Virtual Router Appliance Docs](https://cloud.ibm.com/docs/gateway-appliance?topic=gateway-appliance-getting-started).

## Example usage

### Standalone configuration

```terraform
resource "ibm_network_gateway" "gateway" {
  name = "my-gateway"

  members {
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
  }
}

```
### HA configuration

```terraform
resource "ibm_network_gateway" "gateway" {
  name = "my-ha-gateway"

  members {
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
  }
  members {
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
  }
}

```


## Argument reference 
Review the argument references that you can specify for your resource.

- `members` - (Required, List) A nested block describes the hardware members of this network gateway.

  Nested scheme for `members`:
  - `datacenter` - (Required, Forces new resource, String) The data center in which you want to provision the member.
  - `disk_key_names` - (Optional, Forces new resource, List) Provide the disk key name. Refer to the same attribute in the `ibm_compute_bare_metal` resource.
  - `domain` - (Required, Forces new resource, String) The domain of the member.
  - `ipv6_enabled` -  (Optional, Forces new resource, Bool) Whether to enable IPv6. Default value is **true**.
  - `hostname` - (Optional, Forces new resource, String) Hostname of the member.
  - `memory` - (Required, Forces new resource, Integer) The amount of memory, expressed in megabytes, that you want to allocate.
  - `network_speed` - (Optional, Forces new resource, Integer) The connection speed (in Mbps) for the member network components. Default value is `100`.
  - `notes` - (Optional, Forces new resource, String) Descriptive text of up to 1000 characters about the member.
  - `os_key_name` - (Optional, Forces new resource, String) The os key name for member. Default value is  **OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT**. Refer to the same attribute on the `ibm_compute_bare_metal` resource.
  - `package_key_name` - (Optional, Forces new resource, String) The key name for the network gateway package. You can find available package key names in the SoftLayer API URL `https://api.softlayer.com/rest/v3/SoftLayer_Product_Package/getAllObjects?objectFilter={"type":{"keyName":{"operation":"BARE_METAL_GATEWAY"}}}` that uses your API key as the password. Default value is `NETWORK_GATEWAY_APPLIANCE`. The default value will allow order Single processor multi-core Servers. Use **2U_NETWORK_GATEWAY_APPLIANCE_1O_GBPS** for ordering Dual processor multi-core Servers.
  - `process_key_name` - (Optional, Forces new resource, String) The process key name for the member. Default value is  **INTEL_SINGLE_XEON_1270_3_40_2**. Refer to the same attribute on the **ibm_compute_bare_metal** resource.
  - `public_bandwidth` - (Optional, Forces new resource, Integer) Allowed public network traffic (in GB) per month. Default value is **20000**.
  - `private_network_only`-  (Optional,  Forces new resource, Bool)  Whether to enable a private network only. Default value is **false**.
  - `post_install_script_uri`- (Optional, Forces new resource, String) The URI of the script to be downloaded and executed on the member. Default value is `nil`.
  - `public_vlan_id` - (Optional, Forces new resource, Integer)ID of the public VLAN.
  - `private_vlan_id` - (Optional, Forces new resource, Integer) ID of the private VLAN.       **Note** If there are two members in this gateway, then both should have same value for **public_vlan_id** and **private_vlan_id**.
  - `redundant_power_supply` -  (Optional, Forces new resource, Bool) When the value is **true**, more power supply is provided. Default value is **false**.
  - `redundant_network` -  (Optional, Forces new resource, Bool) When the value is **true**, two physical network interfaces are provided with a bonding configuration. Default value is **false**.
  - `storage_groups` - (Optional, Forces new resource, List) A nested block describes the storage group for the member of the network gateway. Nested `storage_groups` blocks have the following structure:

    Nested scheme for `storage_groups`:
    - `array_type_id` - (Required, Integer) The ID of the array type.
    - `array_size` - (Optional, Integer)The size of the array. 
    - `hard_drives`- (Required, List) The list of hard disk associated with the gateway member.
    - `partition_template_id` - (Optional, Integer) The partition template ID for the member.
  - `ssh_key_ids` - (Optional, Forces new resource, List) The SSH key IDs to install on the member.
  - `tags` - (Optional, Forces new resource, Set)  Tags associated with the VM instance. Permitted characters include: A-Z, 0-9, whitespace, `_` (underscore), `- ` (hyphen), `.` (period), and `:` (colon). All other characters are removed.
  - `tcp_monitoring` -  (Optional, Forces new resource, Bool) Whether to enable TCP monitoring for the member. Default value is **false**.
  - `unbonded_network` -  (Optional, Forces new resource, Bool) When the value is **true**, two physical network interfaces are provided without a bonding configuration. Default value is **false**.
  - `user_metadata` - (Optional, Forces new resource, String) Arbitrary data to be made available to the member.
- `name` - (Required, String) The name of the gateway.
- `post_install_script_uri` - (Optional, Forces new resource, String) The URI of the script to be downloaded and executed after the gateway installation is complete. Default value is **nil**.
- `ssh_key_ids` - (Optional, Forces new resource,  List) The SSH key IDs to install on the gateway when the gateway gets created.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `associated_vlans` - (List) A nested block describes the associated VLANs for the member of the network gateway. Nested `associated_vlans` blocks export the following attributes.

  Nested scheme for `associated_vlans`:
  - `bypass` -  (String) Indicates if the VLAN is in bypass or routed mode.
  - `network_vlan_id` - (String) The ID of the VLAN that is associated.
  - `vlan_id` - (String) The VLAN ID.
- `id` - (String) The unique identifier of the network gateway.
- `members` - (List) A nested block describes the hardware members of this network gateway. Nested `members` blocks export the following attributes.

  Nested scheme for `members`:
  - `ipv6_address` -  (String) IPv6 address associated with the member.
  - `member_id` -  (String) ID of the member.
  - `public_ipv4_address` - (String) Public IPv4 address associated with the member.
  - `private_ipv4_address` - (String) Private IPv4 address associated with the member.
- `private_ipv4_address` - (String) The private IP address of the network gateway.
- `private_ip_address_id` - (String) The private IP address ID of the network gateway.
- `private_vlan_id` - (String) The private VLAN ID of the network gateway.
- `public_ip_address_id` - (String) The public IP address ID of the network gateway.
- `public_ipv6_address_id` - (String) The public IPv6 address ID for the network gateway.
- `public_vlan_id` - (String) The public VLAN ID for the network gateway.
- `public_ipv4_address` - (String) The public IP address of the network gateway.
- `status` - (String) Status of the network gateway.
