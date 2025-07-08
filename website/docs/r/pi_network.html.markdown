---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_network"
description: |-
  Manages networks in the IBM Power Virtual Server cloud.
---

# ibm_pi_network

Create, update, or delete a network connection for your Power Systems Virtual Server instance. For more information, about power virtual server instance network, see [setting up an IBM network install server](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-configuring-subnet).

## Example Usage

The following example creates a network connection for your Power Systems Virtual Server instance.

```terraform
resource "ibm_pi_network" "power_networks" {
  count                = 1
  pi_network_name      = "power-network"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_network_type      = "vlan"
  pi_cidr              = "<Network in CIDR notation (192.168.0.0/24)>"
  pi_dns               = [<"DNS Servers">]
  pi_gateway           = "192.168.0.1"
  pi_ipaddress_range {
    pi_starting_ip_address  = "192.168.0.2"
    pi_ending_ip_address    = "192.168.0.254"
  }
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Timeouts

The `ibm_pi_network` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for creating a network.
- **update** - (Default 60 minutes) Used for updating a network.
- **delete** - (Default 60 minutes) Used for deleting a network.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_advertise` - (Optional, String) Enable the network to be advertised. Only supported for `vlan` network type.
- `pi_arp_broadcast` - (Optional, String) Enable ARP Broadcast. Only supported for `vlan` network type.
- `pi_cidr` - (Optional, String) The network CIDR. Required for `vlan` network type.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_dns` - (Optional, Set of String) The DNS Servers for the network. If not specified, default is 127.0.0.1 for 'vlan' (private network) and 9.9.9.9 for 'pub-vlan' (public network). A maximum of one DNS server can be specified for private networks in Power Edge Router workspaces.
- `pi_gateway` - (Optional, String) The gateway ip address.
- `pi_ipaddress_range` - (Optional, List of Map) List of one or more ip address range(s). The `pi_ipaddress_range` object structure is documented below. The `pi_ipaddress_range` block supports:
  - `pi_ending_ip_address` - (Required, String) The ending ip address.
  - `pi_starting_ip_address` - (Required, String) The staring ip address. **Note** if the `pi_gateway` or `pi_ipaddress_range` is not provided, it will calculate the value based on CIDR respectively.
- `pi_network_mtu` - (Optional, Integer) Maximum Transmission Unit option of the network. Minimum is 1450 and maximum is 9000.
- `pi_network_name` - (Required, String) The name of the network.
- `pi_network_type` - (Required, String) The type of network that you want to create. Valid values are `pub-vlan`, `vlan` and `dhcp-vlan`.
- `pi_network_peer` - (Optional, List) Network peer information (for on-prem locations only). Max items: 1.

  Nested schema for `pi_network_peer`:
  - `id` - (Required, String) ID of the network peer.
  - `network_address_translation` - (Deprecated, Optional, List) Contains the Network Address Translation Details. Max items: 1.

      Nested schema for `network_address_translation`:
        - `source_ip` - (Deprecated, Optional, String) source IP address, required if network peer type is `L3BGP` or `L3STATIC` and if NAT is enabled.
  - `type` - (Deprecated, Optional, String) Type of the network peer. Allowable values are: `L2`, `L3BGP`, `L3Static`.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of this resource.
- `id` - (String) The unique identifier of the network. The ID is composed of `<pi_cloud_instance_id>/<network_id>`.
- `network_address_translation` - (Deprecated, List) Contains the network address translation details (for on-prem locations only).

    Nested schema for  `network_address_translation`:
      - `source_ip` - (Deprecated, String) source IP address.
- `network_id` - (String) The unique identifier of the network.
- `peer_id` - (Deprecated, String) Network peer ID (for on-prem locations only).
- `vlan_id` - (Integer) The ID of the VLAN that your network is attached to.

## Import

The `ibm_pi_network` resource can be imported by using `pi_cloud_instance_id` and `network_id`.

### Example

```bash
terraform import ibm_pi_network.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
