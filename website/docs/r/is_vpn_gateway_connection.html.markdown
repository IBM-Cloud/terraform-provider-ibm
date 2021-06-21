---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : VPN-gateway-connection"
description: |-
  Manages IBM VPN gateway connection.
---

# ibm_is_vpn_gateway_connection
Create, update, or delete a VPN gateway connection. For more information, about VPN gateway, see [adding connections to a VPN gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-adding-connections).


## Example usage
The following example creates a VPN gateway:

```terraform
resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection" {
  name          = "test2"
  vpn_gateway   = ibm_is_vpn_gateway.testacc_VPNGateway2.id
  peer_address  = ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address
  preshared_key = "VPNDemoPassword"
  local_cidrs = [ibm_is_subnet.testacc_subnet2.ipv4_cidr_block]
  peer_cidrs = [ibm_is_subnet.testacc_subnet1.ipv4_cidr_block]
}

```

## Timeouts
The `ibm_is_vpn_gateway_connection` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 10 minutes) Used for deleting instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `action` - (Optional, String)  Dead peer detection actions. Supported values are **restart**, **clear**, **hold**, or **none**. Default value is `none`.
- `admin_state_up` - (Optional, Bool) The VPN gateway connection status. Default value is **false**. If set to false, the VPN gateway connection is shut down.
- `ike_policy` - (Optional, String) The ID of the IKE policy.
- `interval` - (Optional, Integer) Dead peer detection interval in seconds. Default value is 30.
- `ipsec_policy` - (Optional, String) The ID of the IPSec policy.
- `local_cidrs` - (Optional, Forces new resource, List) List of local CIDRs for this resource.
- `name` - (Required, String) The name of the VPN gateway connection.
- `peer_cidrs` - (Optional, Forces new resource, List) List of peer CIDRs for this resource.
- `peer_address` - (Required, String) The IP address of the peer VPN gateway.
- `preshared_key` - (Required, Forces new resource, String) The preshared key.
- `timeout` - (Optional, Integer) Dead peer detection timeout in seconds. Default value is 120.
- `vpn_gateway` - (Required, Forces new resource, String) The unique identifier of the VPN gateway.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `authentication_mode` - (String) The authentication mode, only `psk` is supported.
- `created_at`-  (Timestamp) The date and time that VPN gateway connection was created.
- `crn` - (String) The `VPN Gateway information ID`.
- `gateway_connection` - The unique identifier for this VPN gateway connection.
- `id` - (String) The unique identifier of the VPN gateway connection. The ID is composed of `<vpn_gateway_id>/<vpn_gateway_connection_id>`.
- `mode` -  (String) The mode of the `VPN gateway` either **policy** or **route**.
- `resource_type` -  (String) The resource type (vpn_gateway_connection).
- `status` -  (String) The status of a VPN gateway connection either `down` or `up`.
- `tunnels` -  (List) The VPN tunnel configuration for the VPN gateway connection (in static route mode).

  Nested scheme for `tunnels`
  - `address`-  (String) The IP address of the VPN gateway member in which the tunnel resides.
  - `resource_type`-  (String) The status of the VPN tunnel.


## Import
The `ibm_is_vpn_gateway_connection` resource can be imported by using the VPN gateway ID and the VPN gateway connection ID. 

**Syntax**

```
$ terraform import ibm_is_vpn_gateway_connection.example <vpn_gateway_ID>/<vpn_gateway_connection_ID>
```

**Example**

```
$ terraform import ibm_is_vpn_gateway_connection.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
