
subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_gateway"
description: |-
  Manages IBM Cloud Infrastructure Transit Gateway.
---

# ibm_tg_gateway
Retrieve information of an existing IBM Cloud infrastructure transit gateway as a read only data source. For more information, about Transit Gateway, see [getting started with IBM Cloud Transit Gateway](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-getting-started).


## Example usage

---
```terraform
resource "ibm_tg_gateway" "new_tg_gw" {
  name           = "transit-gateway-1"
  location       = "us-south"
  global         = true
  resource_group = "30951d2dff914dafb26455a88c0c0092"
}

data "ibm_tg_gateway" "ds_tggateway" {
  name = ibm_tg_gateway.new_tg_gw.name
}
```
---

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the gateway.
- `gre_enhanced_route_propagation` - (Optional, Bool) Allows route propagation across all GRE connections on the same Transit Gateway (redundant_gre, unbound_gre_tunnel, and gre_tunnel).

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `created_at` - (Timestamp) The date and time resource is created.
- `crn` - (String) The CRN of the gateway.
- `global` - (String) The gateways with global routing true to connect to the networks outside the associated region.
- `id` - (String) The unique identifier of this gateway.
- `location` - (String) The gateway location.
- `resource_group` - (String) The resource group identifier.
- `connections` - (String) A list of connections in the gateway

  Nested scheme for `connections`:
	- `created_at` - (String) The date and time the connection is created.
	- `id` - (String) The unique identifier for the transit gateway connection to network either `vpc`,  `classic`.
  - `base_connection_id` - (String) The ID of a network_type `classic` connection a tunnel is configured over.  This field applies to network type `gre_tunnel` or `unbound_gre_tunnel` connections.
  - `base_network_type` - (String) The type of network the unbound gre tunnel is targeting. This field is required for network type `unbound_gre_tunnel`.
  - `name` - (String) The user-defined name for the transit gateway connection.
  - `network_type` - (String) The type of network connected with the connection. Possible values are `classic`, `directlink`, `vpc`, `gre_tunnel`,  `unbound_gre_tunnel`, or `power_virtual_server`.
  - `network_account_id` - (String) The ID of the network connected account. This is used if the network is in a different account than the gateway.
  - `network_id` - (String) The ID of the network being connected with the connection.
  - `local_bgp_asn` - (Integer) The local network BGP ASN. This field only applies to network type '`gre_tunnel` connections.
  - `local_gateway_ip` - (String) The local gateway IP address.  This field is required for and only applicable to `gre_tunnel` connection types.
  - `local_tunnel_ip` - (String) The local tunnel IP address. This field is required for and only applicable to type gre_tunnel connections.
  - `mtu` - (Integer) GRE tunnel MTU. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
  - `remote_bgp_asn` - (Integer) The remote network BGP ASN (will be generated for the connection if not specified). This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
  - `remote_gateway_ip` - (String) The remote gateway IP address. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
  - `remote_tunnel_ip` - (String) The remote tunnel IP address. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
	- `status` - (String) The current configuration state of the connection. Possible values are `attached`, `failed,` `pending`, `deleting`.
	- `updated_at` - (String) The date and time the connection is last updated.
  - `zone` - (String) The location of the GRE tunnel. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
- `status` - (String) The gateway status.
- `updated_at` - (Timestamp) The date and time resource is last updated.
- `tunnels` - (List) List of GRE tunnels for a transit gateway redundant GRE tunnel connection. This field is required for 'redundant_gre' connections.
          
        Nested scheme for `tunnel`:
  - `name` - (Required, String) The user-defined name for this tunnel connection.
  - `local_gateway_ip` - (String)  The local gateway IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `local_tunnel_ip` - (String) The local tunnel IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `remote_gateway_ip` - (String) The remote gateway IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `remote_tunnel_ip` - (String) The remote tunnel IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `zone` - (String) - The location of the GRE tunnel. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
  - `remote_bgp_asn` - (Integer) - The remote network BGP ASN (will be generated for the connection if not specified). This field only applies to network type`gre_tunnel` and `unbound_gre_tunnel` connections.
  - `created_at` -  (Timestamp) The date and time the connection  tunnel was created. 
  - `id` - (String) The unique identifier of the connection tunnel ID resource.
  - `mtu` - (Integer) GRE tunnel MTU.
  - `status` - (String) The configuration status of the connection tunnel, such as **attached**, **failed**,
  - `updated_at` - (Timestamp) Last updated date and time of the connection tunnel.
  - `local_bgp_asn` - (Integer) The local network BGP ASN.
