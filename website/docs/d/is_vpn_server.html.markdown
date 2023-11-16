---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server"
description: |-
  Get information about VPNServer
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_server

Provides a read-only data source for VPNServer. For more information, about VPN Server, see [Creating a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-create-server&interface=ui).

## Example Usage

```terraform
data "ibm_is_vpn_server" "example" {
	identifier = ibm_is_vpn_server.example.id
}
```

## Argument Reference
Review the argument reference that you can specify for your data source.

- `identifier` - (Optional, String) The ID of the VPN server.
- `name` - (Optional, String) The name of the VPN server.

  ~> **NOTE**
    `identifier` and `name` are mutually exclusive.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `access_tags`  - (List) Access management tags associated for the vpn server.
- `identifier` - The unique identifier of the VPNServer.
- `certificate` - (List) The certificate instance for this VPN server.

	Nested scheme for `certificate`:
	- `crn` - (String) The CRN for this certificate instance.

- `client_authentication` - (List) The methods used to authenticate VPN clients to this VPN server. VPN clients must authenticate against all provided methods.
  
  Nested scheme for `client_authentication`:
	- `method` - (String) The type of authentication.
	- `identity_provider` - (String) The type of identity provider to be used by VPN client.
	- `client_ca` - (String) The certificate instance used for the VPN client certificate authority (CA).

- `client_auto_delete` - (Boolean) If set to `true`, disconnected VPN clients will be automatically deleted after the `client_auto_delete_timeout` time has passed.

- `client_auto_delete_timeout` - (Integer) Hours after which disconnected VPN clients will be automatically deleted. If `0`, disconnected VPN clients will be deleted immediately.

- `client_dns_server_ips` - (List) The DNS server addresses that will be provided to VPN clients that are connected to this VPN server.
	
	Nested scheme for `client_dns_server_ips`:
	- `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.

- `client_idle_timeout` - (Integer) The seconds a VPN client can be idle before this VPN server will disconnect it.  If `0`, the server will not disconnect idle clients.

- `client_ip_pool` - (String) The VPN client IPv4 address pool, expressed in CIDR format.

- `created_at` - (String) The date and time that the VPN server was created.

- `crn` - (String) The CRN for this VPN server.

- `enable_split_tunneling` - (Boolean) Indicates whether the split tunneling is enabled on this VPN server.

- `health_reasons` - (List) The reasons for the current health_state (if any).

  Nested scheme for `health_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) Link to documentation about the reason for this health state.

- `health_state` - (String) The health of this resource.

  -> **Supported health_state values:** 
    </br>&#x2022; `ok`: Healthy
    </br>&#x2022; `degraded`: Suffering from compromised performance, capacity, or connectivity
    </br>&#x2022; `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated
    </br>&#x2022; `inapplicable`: The health state does not apply because of the current lifecycle state. 
      **Note:** A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.
- `hostname` - (String) Fully qualified domain name assigned to this VPN server.

- `href` - (String) The URL for this VPN server.

- `lifecycle_reasons` - (List) The reasons for the current lifecycle_reasons (if any).

  Nested scheme for `lifecycle_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle reason.
  - `message` - (String) An explanation of the reason for this lifecycle reason.
  - `more_info` - (String) Link to documentation about the reason for this lifecycle reason.
  
- `lifecycle_state` - (String) The lifecycle state of the VPN server.

- `name` - (String) The unique user-defined name for this VPN server.

- `port` - (Integer) The port number used by this VPN server.

- `private_ips` - (List) The reserved IPs bound to this VPN server.

	Nested scheme for `private_ips`:
	- `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this reserved IP.
	- `id` - (String) The unique identifier for this reserved IP.
	- `name` - (String) The user-defined or system-provided name for this reserved IP.
	- `resource_type` - (String) The resource type.

- `protocol` - (String) The transport protocol used by this VPN server.

- `resource_group` - (List) The resource group object, for this VPN server.

	Nested scheme for `resource_group`:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The user-defined name for this resource group.

- `resource_type` - (String) The type of resource referenced.

- `security_groups` - (List) The security groups targeting this VPN server.

	Nested scheme for `security_groups`:
	- `crn` - (String) The security group's CRN.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The security group's canonical URL.
	- `id` - (String) The unique identifier for this security group.
	- `name` - (String) The user-defined name for this security group. Names must be unique within the VPC the security group resides in.

- `subnets` - (List) The subnets this VPN server is part of.

	Nested scheme for `subnets`:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.

		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.
	- `resource_type` - (String) The resource type.

- `vpc` - (List) The VPC this VPN server resides in.

	Nested scheme for `vpc`:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this vpc.
	- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
