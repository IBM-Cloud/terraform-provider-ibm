
subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_connection"
description: |-
  Manages IBM Transit Gateway connection.
---

# ibm_tg_connection
Create, update and delete for the transit gateway's connection resource. For more information, about Transit Gateway connection, see [adding a cross-account connection](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-edit-gateway#adding-cross-account-connections).

## Example usage

---
```terraform
resource "ibm_tg_connection" "test_ibm_tg_connection" {
  gateway      = ibm_tg_gateway.test_tg_gateway.id
  network_type = "vpc"
  name         = "myconnection"
  network_id   = ibm_is_vpc.test_tg_vpc.resource_crn
}

resource "ibm_tg_connection" "test_ibm_tg_rgre_connection" {
  gateway      = ibm_tg_gateway.test_tg_gateway.id
  name = redundant_ugre_vpc
  network_type = redundant_gre
  base_network_type = vpc
  network_id = ibm_is_vpc.test_tg_vpc.resource_crn
  tunnels {
           local_gateway_ip = "192.129.200.1"
           local_tunnel_ip = "192.158.239.2"
           name =  "tunne1_testtgw1461"
           remote_gateway_ip = "10.186.203.4"
           remote_tunnel_ip = "192.158.239.1"
           zone =  "us-south-1"
        }    
 tunnels {
             local_gateway_ip = "192.129.220.1"
             local_tunnel_ip = "192.158.249.2"
             name =  "tunne2_testtgw1462"
             remote_gateway_ip = "10.186.203.4"
             remote_tunnel_ip = "192.158.249.1"
             zone =  "us-south-1"
         }  
}
  
```
---

## Argument reference
Review the argument references that you can specify for your resource. 
 
- `base_connection_id` - (Optional, Forces new resource, String) - The ID of a network_type 'classic' connection a tunnel is configured over.  This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
- `base_network_type` - (Optional, String) - The type of network the unbound gre tunnel is targeting. This field is required for network type `unbound_gre_tunnel`.
- `gateway` - (Required, Forces new resource, String) Enter the transit gateway identifier.
- `local_gateway_ip` - (Optional, Forces new resource, String) - The local gateway IP address.  This field is required for and only applicable to `gre_tunnel` connection types.
- `local_tunnel_ip` - (Optional, Forces new resource, String) - The local tunnel IP address. This field is required for and only applicable to type gre_tunnel connections.
- `name` -  (Optional, String) Enter a name. If the name is not given, the default name is provided based on the network type, such as `vpc` for network type VPC and `classic` for network type classic.
- `network_account_id` - (Optional, Forces new resource, String) The ID of the network connected account. This is used if the network is in a different account than the gateway.
- `network_type` - (Required, Forces new resource, String) Enter the network type. Allowed values are `classic`, `directlink`, `gre_tunnel`, `unbound_gre_tunnel`,  `vpc`, `vpn_gateway`, and `power_virtual_server`.
- `network_id` -  (Optional, Forces new resource, String) Enter the ID of the network being connected through this connection. This parameter is required for network type `vpc` and `directlink`, the CRN of the VPC or direct link gateway to be connected. This field is required to be unspecified for network type `classic`. For example, `crn:v1:bluemix:public:is:us-south:a/123456::vpc:4727d842-f94f-4a2d-824a-9bc9b02c523b`.
- `remote_bgp_asn` - (Optional, Forces new resource, Integer) - The remote network BGP ASN (will be generated for the connection if not specified). This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
- `remote_gateway_ip` - (Optional, Forces new resource, String) - The remote gateway IP address. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
- `remote_tunnel_ip` - (Optional, Forces new resource, String) - The remote tunnel IP address. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
- `zone` - (Optional, Forces new resource, String) - The location of the GRE tunnel. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
- `cidr` - (Optional, String) - network_type `vpn_gateway` connections use `cidr` to specify the CIDR to use for the VPN GRE tunnels.
- `tunnels` - (Optional, List) List of GRE tunnels for a transit gateway redundant GRE tunnel connection. This field is required for 'redundant_gre' connections.
Nested scheme for `tunnel`:
  - `name` - (Required, String) The user-defined name for this tunnel connection.
  - `local_gateway_ip` - (Required, String)  The local gateway IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `local_tunnel_ip` - (Required, String) The local tunnel IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `remote_gateway_ip` - (Required, String) The remote gateway IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `remote_tunnel_ip` - (Required, String) The remote tunnel IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `zone` - (Optional, Forces new resource, String) - The location of the GRE tunnel. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
  - `remote_bgp_asn` - (Optional, Forces new resource, Integer) - The remote network BGP ASN (will be generated for the connection if not specified). This field only applies to network type`gre_tunnel` and `unbound_gre_tunnel` connections.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `connection_id` - (String) The unique identifier for transit gateway connection to network.
- `created_at` -  (Timestamp) The date and time the connection was created. 
- `id` - (String) The unique identifier of the gateway ID or connection ID resource.
- `local_bgp_asn` - (Integer) The local network BGP ASN. This field only applies to network type `gre_tunnel` connections.
- `mtu` - (Integer) GRE tunnel MTU. This field only applies to network type `gre_tunnel` connections.
- `status` - (String) The configuration status of the connection, such as **attached**, **failed**, **pending**, **deleting**.
- `updated_at` - (Timestamp) Last updated date and time of the connection.
-  `tunnels` - (Optional, List) List of GRE tunnels for a transit gateway redundant GRE tunnel connection. This field is required for 'redundant_gre' connections.
   Nested scheme for `tunnel`:
   - `created_at` -  (Timestamp) The date and time the connection  tunnel was created. 
   - `id` - (String) The unique identifier of the connection tunnel ID resource.
   - `mtu` - (Integer) GRE tunnel MTU.
   - `status` - (String) The configuration status of the connection tunnel, such as **attached**, **failed**,
   - `updated_at` - (Timestamp) Last updated date and time of the connection tunnel.
   - `local_bgp_asn` - (Integer) The local network BGP ASN.
 

**Note**

The resource do not wait for the available status, if you are provisioning the cross account gateway or connection. You need to complete the manual approval process for provisioning.


## Import
The `ibm_tg_connection` resource can be imported by using transit gateway ID and connection ID.

**Example**

---
```
$ terraform import ibm_tg_connection.example 5ffda12064634723b079acdb018ef308/cea6651a-bd0a-4438-9f8a-a0770bbf3ebb

```
---
