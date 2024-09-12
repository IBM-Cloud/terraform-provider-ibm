---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : VPN-gateway-connection"
description: |-
  Manages IBM VPN gateway connection.
---

# ibm_is_vpn_gateway_connection
Create, update, or delete a VPN gateway connection. For more information, about VPN gateway, see [adding connections to a VPN gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-adding-connections).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```


## Example usage
The following example creates a VPN gateway:

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_subnet" "example2" {
  name            = "example-subnet2"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.68.0/24"
}

resource "ibm_is_vpn_gateway" "example" {
  name   = "example-vpn-gateway"
  subnet = ibm_is_subnet.example.id
  mode   = "route"
}

resource "ibm_is_vpn_gateway_connection" "example" {
  name          = "example-vpn-gateway-connection"
  vpn_gateway   = ibm_is_vpn_gateway.example.id
  preshared_key = "VPNDemoPassword"
  # peer_address  = ibm_is_vpn_gateway.example.public_ip_address # deprecated, replaced with peer block
  # peer_cidrs    = [ibm_is_subnet.example2.ipv4_cidr_block] # deprecated, replaced with peer block
  peer {
      address  = ibm_is_vpn_gateway.example.public_ip_address
      cidrs    = [ibm_is_subnet.example2.ipv4_cidr_block]   
  }
  # local_cidrs   = [ibm_is_subnet.example.ipv4_cidr_block] # deprecated, replaced with local block
  local {
      cidrs   = [ibm_is_subnet.example.ipv4_cidr_block]
  }
}

```
## Example usage ( policy mode with an active peer VPN gateway )
The following example creates a VPN gateway:

```terraform
resource "ibm_is_vpn_gateway" "example" {
  name   = "example-vpn-gateway"
  subnet = ibm_is_subnet.example.id
  mode   = "policy"
}

resource "ibm_is_vpn_gateway_connection" "example" {
  name          = "example-vpn-gateway-connection"
  vpn_gateway   = ibm_is_vpn_gateway.example.id
  preshared_key = "VPNDemoPassword"

  # local_cidrs   = [ibm_is_subnet.example.ipv4_cidr_block] # depcreated
  local {
    cidrs = [ibm_is_subnet.example.ipv4_cidr_block]
  }
  # peer_cidrs    = [ibm_is_subnet.example2.ipv4_cidr_block] # depcreated
  # peer_address  = ibm_is_vpn_gateway.example.public_ip_address != "0.0.0.0" ? ibm_is_vpn_gateway.example.public_ip_address : ibm_is_vpn_gateway.example.public_ip_address2 # depcreated
  peer {
    address   = ibm_is_vpn_gateway.example.public_ip_address != "0.0.0.0" ? ibm_is_vpn_gateway.example.public_ip_address : ibm_is_vpn_gateway.example.public_ip_address2
    cidrs     = [ibm_is_subnet.example2.ipv4_cidr_block]
  }
}

```

## Timeouts
The `ibm_is_vpn_gateway_connection` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 10 minutes) Used for deleting instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `action` - (Optional, String)  Dead peer detection actions. Supported values are **restart**, **clear**, **hold**, or **none**. Default value is `restart`.
- `admin_state_up` - (Optional, Bool) The VPN gateway connection status. Default value is **false**. If set to false, the VPN gateway connection is shut down.
- `establish_mode` - (Optional, String) The establish mode of the VPN gateway connection:- `bidirectional`: Either side of the VPN gateway can initiate IKE protocol   negotiations or rekeying processes.- `peer_only`: Only the peer can initiate IKE protocol negotiations for this VPN gateway   connection. Additionally, the peer is responsible for initiating the rekeying process   after the connection is established. If rekeying does not occur, the VPN gateway   connection will be brought down after its lifetime expires.
- `ike_policy` - (Optional, String) The ID of the IKE policy. Updating value from ID to `""` or making it `null` or removing it  will remove the existing policy.
- `interval` - (Optional, Integer) Dead peer detection interval in seconds. Default value is 2.
- `ipsec_policy` - (Optional, String) The ID of the IPSec policy. Updating value from ID to `""` or making it `null` or removing it  will remove the existing policy.
- `local` - (Optional, List) 
  Nested schema for **local**:
	- `ike_identities` - (Required, List) The local IKE identities.A VPN gateway in static route mode consists of two members in active-active mode. The first identity applies to the first member, and the second identity applies to the second member.
	  Nested schema for **ike_identities**:
		- `type` - (Required, String) The IKE identity type. [ **fqdn**, **hostname**, **ipv4_address**, **key_id** ] The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the backup policy on which the unexpected property value was encountered.
		- `value` - (Optional, String) The IKE identity FQDN value.
- `local_cidrs` - (Optional, DEPRECATED, Forces new resource, List) List of local CIDRs for this resource. `local_cidrs` is deprecated and use `local` block instead. 
- `name` - (Required, String) The name of the VPN gateway connection.
- `peer` - (Optional, List) 
  Nested schema for **peer**:
	- `address` - (Optional, String) The IP address of the peer VPN gateway for this connection.
	- `fqdn` - (Optional, String) The FQDN of the peer VPN gateway for this connection.
	- `ike_identity` - (Required, List) The peer IKE identity.
	  Nested schema for **ike_identity**:
		- `type` - (Required, String) The IKE identity type. [ **fqdn**, **hostname**, **ipv4_address**, **key_id** ] The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the backup policy on which the unexpected property value was encountered.
		- `value` - (Optional, String) The IKE identity FQDN value.
	- `type` - (Computed, String) Indicates whether `peer.address` or `peer.fqdn` is used.

- `peer_cidrs` - (Optional, DEPRECATED, Forces new resource, List) List of peer CIDRs for this resource. `peer_cidrs` is deprecated and use `peer` block instead.
- `peer_address` - (Optional, DEPRECATED, String) The IP address of the peer VPN gateway. `peer_address` is deprecated and use `peer` block instead.
- `preshared_key` - (Required, Forces new resource, String) The preshared key.
- `timeout` - (Optional, Integer) Dead peer detection timeout in seconds. Default value is 10.
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
- `status_reasons` - (List) Array of reasons for the current status (if any).

  Nested `status_reasons`:
    - `code` - (String) The status reason code.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) Link to documentation about this status reason
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
