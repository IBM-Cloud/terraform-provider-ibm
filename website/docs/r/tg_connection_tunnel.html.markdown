
subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_connection_rgre_tunnel"
description: |-
  Manages IBM Transit Gateway connection tunnel.
---

# ibm_tg_connection_rgre_tunnel
Create, update and delete for the transit gateway's connection tunnel resource. For more information, about Transit Gateway connection, see [adding a cross-account connection](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-edit-gateway#adding-cross-account-connections).

## Example usage

---
```terraform

resource "ibm_tg_connection_rgre_tunnel" "test_ibm_tg_connection_tunnel" {
  gateway = ibm_tg_gateway.test_tg_gateway.id
  connection_id = ibm_tg_connection.test_ibm_tg_connection.connection_id
  local_gateway_ip = "192.139.200.1"
  local_tunnel_ip = "192.178.239.2"
  name =  "tunnel_name"
  remote_gateway_ip = "10.186.203.4"
  remote_tunnel_ip = "192.178.239.1"
  zone =  "us-south-3"
}    
```
---
## Argument reference
Review the argument references that you can specify for your resource. 
 
  - `gateway` - (Required, Forces new resource, String) Enter the transit gateway identifier.
  - `connection_id` - (Required, String) The unique identifier of the gateway connection
  - `name` - (Required, String) The user-defined name for this tunnel connection.
  - `local_gateway_ip` - (Required, String)  The local gateway IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `local_tunnel_ip` - (Required, String) The local tunnel IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `remote_gateway_ip` - (Required, String) The remote gateway IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `remote_tunnel_ip` - (Required, String) The remote tunnel IP address. This field only applies to network type 'gre_tunnel' and 'unbound_gre_tunnel' connections.
  - `zone` - (Optional, Forces new resource, String) - The location of the GRE tunnel. This field only applies to network type `gre_tunnel` and `unbound_gre_tunnel` connections.
  - `remote_bgp_asn` - (Optional, Forces new resource, Integer) - The remote network BGP ASN (will be generated for the connection if not specified). This field only applies to network type`gre_tunnel` and `unbound_gre_tunnel` connections.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.


   - `created_at` -  (Timestamp) The date and time the connection  tunnel was created. 
   - `id` - (String) The unique identifier of the connection tunnel ID resource.
   - `mtu` - (Integer) GRE tunnel MTU.
   - `status` - (String) The configuration status of the connection tunnel, such as **attached**, **failed**,
   - `updated_at` - (Timestamp) Last updated date and time of the connection tunnel.
   - `local_bgp_asn` - (Integer) The local network BGP ASN.
 

**Note**

The resource do not wait for the available status, if you are provisioning the cross account gateway or connection. You need to complete the manual approval process for provisioning.


## Import
The `ibm_tg_connection_rgre_tunnel` resource can be imported by using transit gateway ID and connection ID and tunnel ID.

**Example**

---
```
$ terraform import ibm_tg_connection_rgre_tunnel.example 5ffda12064634723b079acdb018ef308/cea6651a-bd0a-4438-9f8a-a0770bbf3ebb

```
---
