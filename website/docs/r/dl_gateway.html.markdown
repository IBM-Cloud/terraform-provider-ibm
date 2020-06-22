---
layout: "ibm"
page_title: "IBM : dl_gateway"
sidebar_current: "docs-ibm-resource-dl-gateway"
description: |-
  Manages IBM Direct Link Gateway.
---

# ibm\_dl_gateway

Provides a direct link gateway resource. This allows direct link gateway to be created, and updated and deleted.

## Example Usage

```hcl
resource ibm_dl_gateway test_dl_gateway {
  bgp_asn =  64999
  bgp_base_cidr =  "169.254.0.0/16"
  bgp_ibm_cidr =  "169.254.0.29/30"
  bgp_cer_cidr =  "169.254.0.30/30"
  global = true 
  metered = false
  name = "Gateway1"
  resource_group = "bf823d4f45b64ceaa4671bee0479346e"
  speed_mbps = 1000 
  loa_reject_reason = "The port mentioned was incorrect"
  operational_status = "loa_accepted"
  type =  "dedicated" 
  cross_connect_router = "LAB-xcr01.dal09"
  location_name = "dal09"
  customer_name = "Customer1" 
  carrier_name = "Carrier1"

}   

```

## Argument Reference

The following arguments are supported:

* `bgp_asn` - (Required, Forces new resource, integer) The BGP ASN of the Gateway to be created. Example: 64999
* `bgp_base_cidr` - (Required, Forces new resource, string) The BGP base CIDR of the Gateway to be created. Example: 10.254.30.76/30 
* `global` - (Required, boolean) Gateways with global routing (true) can connect to networks outside their associated region.
* `metered` -  (Required, boolean) Metered billing option. When true gateway usage is billed per gigabyte. When false there is no per gigabyte usage charge, instead a flat rate is charged for the gateway.
* `name` - (Required, boolean) The unique user-defined name for this gateway. Example: myGateway
* `speed_mbps` - (Required, integer) Gateway speed in megabits per second. Example: 10.254.30.78/30
* `type` - (Required, Forces new resource, string) Gateway type. Allowable values: [dedicated,connect]. 
* `bgp_cer_cidr` - (Optional, Forces new resource, string) BGP customer edge router CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is 169.254.0.0/16, this field can be ommitted and a CIDR will be selected automatically. Example: 10.254.30.78/30
* `bgp_ibm_cidr` - (Optional, Forces new resource, string) BGP IBM CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is 169.254.0.0/16, this field can be ommitted and a CIDR will be selected automatically. Example: 10.254.30.77/30 
* `resource_group` - (Optional, Forces new resource, string) Resource group for this resource. If unspecified, the account's default resource group is used. 
* `carrier_name` - (Required, Forces new resource, string) Carrier name. Constraints: 1 ≤ length ≤ 128, Value must match regular expression ^[a-z][A-Z][0-9][ -_]$. Example: myCarrierName
* `cross_connect_router` - (Required, Forces new resource, string) Cross connect router. Example: xcr01.dal03
* `customer_name` - (Required, Forces new resource, string) Customer name. Constraints: 1 ≤ length ≤ 128, Value must match regular expression ^[a-z][A-Z][0-9][ -_]$. Example: newCustomerName
* `location_name` - (Required, Forces new resource, string) Gateway location. Example: dal03
* `loa_reject_reason` - (Optional, string) Use this field during LOA rejection to provide the reason for the rejection. Example: The port mentioned was incorrect
* `operational_status` - (Optional, string) Gateway operational status. For gateways pending LOA approval, patch operational_status to the appropriate value to approve or reject its LOA. When rejecting an LOA, provide reject reasoning in loa_reject_reason. Allowable values: [loa_accepted,loa_rejected]. Example: loa_accepted


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of this gateway. 
* `name` - The unique user-defined name for this gateway. 
* `crn` - The CRN (Cloud Resource Name) of this gateway. 
* `created_at` - The date and time resource was created.
* `location_display_name` - Gateway location long name. 
* `resource_group` - Resource group reference
* `bgp_asn` - IBM BGP ASN.
* `bgp_status` - Gateway BGP status.
* `completion_notice_reject_reason` - Reason for completion notice rejection. 
* `link_status` - Gateway link status. Only included on type=dedicated gateways. Example: down, up.
* `port` - gateway port for type=connect gateways
* `vlan` - VLAN allocated for this gateway. Only set for type=connect gateways created directly through the IBM portal. 
* `provider_api_managed` - Indicates whether gateway changes must be made via a provider portal.

## Import

ibm_dl_gateway can be imported using gateway id, eg

```
$ terraform import ibm_dl_gateway.example 5ffda12064634723b079acdb018ef308
```
