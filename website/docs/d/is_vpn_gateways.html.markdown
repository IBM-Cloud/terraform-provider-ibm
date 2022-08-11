---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpn_gateways"
description: |-
  Manages IBM Cloud VPN gateways.
---

# ibm_is_vpn_gateways
Retrieve information of an existing VPN gateways. For more information, about IBM Cloud VPN gateways, see [configuring ACLs and security groups for use with VPN](https://cloud.ibm.com/docs/vpc?topic=vpc-acls-security-groups-vpn).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform

data "ibm_is_vpn_gateways" "example" {
}

```

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `vpn_gateways` - (List) Collection of VPN Gateways.

  Nested scheme for `vpn_gateways`:
  - `crn` - (String) The VPN gateway's CRN.
  - `created_at`- (Timestamp) The date and time the VPN gateway was created.
  - `id` - (String) The ID of the VPN gateway.
  - `name`-  (String) The VPN gateway instance name.
  - `members` - (List) Collection of VPN gateway members.

    Nested scheme for `members`:
	  - `address` - (String) The public IP address assigned to the VPN gateway member.
	  - `role`-  (String) The high availability role assigned to the VPN gateway member.
    - `private_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
      Nested scheme for `private_ip`:
      - `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `href`- (String) The URL for this reserved IP
      - `name`- (String) The user-defined or system-provided name for this reserved IP
      - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_type`- (String) The resource type.
	  - `private_address` - (String) The private IP address assigned to the VPN gateway member. Same as `private_ip.0.address`
	  - `status` - (String) The status of the VPN gateway member.
  - `resource_type` - (String) The resource type, supported value is `vpn_gateway`.
  - `status` - (String) The status of the VPN gateway, supported values are **available**, **deleting**, **failed**, **pending**.
  - `subnet` - (String) The VPN gateway subnet information.
  - `resource_group` - (String) The resource group ID.
  - `mode` - (String) The VPN gateway mode, supported values are `policy` and `route`.

