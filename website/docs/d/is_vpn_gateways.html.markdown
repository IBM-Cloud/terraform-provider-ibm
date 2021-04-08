---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM vpn gateways.
---

# ibm\_is_vpn_gateways

Import the details of an existing IBM VPN Gateways as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_vpn_gateways" "ds_vpn_gateways" {
  
}

```

## Argument Reference

The following arguments are supported:



## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the VPN Gateway.
* `name` - VPN Gateway instance name.
* `created_at` - The date and time that this VPN gateway was created.
* `crn` - The VPN gateway's CRN.
* `members` - Collection of VPN gateway members.
  * `address` - The public IP address assigned to the VPN gateway member.
  * `role` - The high availability role assigned to the VPN gateway member.
  * `status` - The status of the VPN gateway member
* `resource_type` - The resource type(vpn_gateway)
* `status` - The status of the VPN gateway(available, deleting, failed, pending)
* `subnet` - VPNGateway subnet info
* `resource_group` - resource group identifiers(ID).
* `mode` -  VPN gateway mode(policy/route)
