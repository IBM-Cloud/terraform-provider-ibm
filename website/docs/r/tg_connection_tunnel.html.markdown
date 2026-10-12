---

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
  - `connection_id` - (Required, Forces new resource, String) The unique identifier of the gateway connection.
  - `name` - (Required, String) The user-defined name for this tunnel connection. This is the only attribute that can be updated after creation.
  - `local_gateway_ip` - (Required, Forces new resource, String) The local gateway IP address.
  - `local_tunnel_ip` - (Required, Forces new resource, String) The local tunnel IP address.
  - `remote_gateway_ip` - (Required, Forces new resource, String) The remote gateway IP address.
  - `remote_tunnel_ip` - (Required, Forces new resource, String) The remote tunnel IP address.
  - `zone` - (Required, Forces new resource, String) The location of the GRE tunnel.
  - `local_bgp_asn` - (Optional, Computed, Integer) The local network BGP ASN (will be generated for the connection if not specified).
  - `remote_bgp_asn` - (Optional, Computed, Integer) The remote network BGP ASN (will be generated for the connection if not specified).
  - `base_network_type` - (Optional, Computed, String) The type of the base network for the tunnel. For example, `classic` or `vpc`.
  - `network_account_id` - (Optional, Computed, Forces new resource, String) The ID of the account that owns the network being connected.
  - `network_id` - (Optional, Computed, Forces new resource, String) The ID of the network being connected via this tunnel.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

   - `created_at` - (Timestamp) The date and time the connection tunnel was created.
   - `id` - (String) The unique identifier of the connection tunnel resource.
   - `mtu` - (Integer) GRE tunnel MTU.
   - `status` - (String) The configuration status of the connection tunnel, such as **attached**, **failed**, **pending**, **deleting**, **detaching**, **detached**.
   - `tunnel_id` - (String) The Transit Gateway tunnel identifier.
   - `updated_at` - (Timestamp) Last updated date and time of the connection tunnel.

**Note**

The resource does not wait for the available status if you are provisioning a cross-account gateway or connection. You must complete the manual approval process for provisioning.

## Timeouts
The following timeouts are available for `ibm_tg_connection_rgre_tunnel`:

- **Create**: Default 10 minutes
- **Delete**: Default 10 minutes
- **Update**: Default 10 minutes

## Import
The `ibm_tg_connection_rgre_tunnel` resource can be imported by using transit gateway ID and connection ID and tunnel ID.

**Example**

---
```
$ terraform import ibm_tg_connection_rgre_tunnel.example 5ffda12064634723b079acdb018ef308/cea6651a-bd0a-4438-9f8a-a0770bbf3ebb

```
---