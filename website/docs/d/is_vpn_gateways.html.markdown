---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpn_gateways"
description: |-
  Manages IBM Cloud VPN gateways.
---

# ibm_is_vpn_gateways
Retrieve information of an existing VPN gateways. For more information, about IBM Cloud VPN gateways, see [configuring ACLs and security groups for use with VPN](https://cloud.ibm.com/docs/vpc?topic=vpc-acls-security-groups-vpn).


## Example usage

```terraform

data "ibm_is_vpn_gateways" "ds_vpn_gateways" {
  
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
	  - `private_address` - (String) The private IP address assigned to the VPN gateway member.
	  - `status` - (String) The status of the VPN gateway member.
  - `resource_type` - (String) The resource type, supported value is `vpn_gateway`.
  - `status` - (String) The status of the VPN gateway, supported values are **available**, **deleting**, **failed**, **pending**.
  - `subnet` - (String) The VPN gateway subnet information.
  - `resource_group` - (String) The resource group ID.
  - `mode` - (String) The VPN gateway mode, supported values are `policy` and `route`.

