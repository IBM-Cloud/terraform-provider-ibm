---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_gateways"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateways.
---

# `ibm_dl_gateways`

Import the details of an existing IBM Cloud Infrastructure Direct Link Gateways.  For more information, about IBM Cloud Direct Link, see [Getting started with IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-get-started-with-ibm-cloud-dl).


## Example usage

```
data "ibm_dl_gateways" "ds_dlgateways" {
}
     
```

## Argument reference
There is no Argument reference that you need to specify for the data source. 


## Attribute reference
Review the Attribute reference that you can access after your resource is created. 

- `gateways` - (String) List of all the Direct Link Gateways in the IBM Cloud infrastructure.
    - `bgp_asn` - (String) Customer BGP ASN.
    - `bgp_base_cidr` - (String) The BGP base CIDR.
    - `bgp_cer_cidr` - (String) The BGP customer edge router CIDR.
    - `bgp_ibm_asn` - (String) The IBM BGP ASN.
    - `bgp_ibm_cidr` - (String) The IBM BGP  CIDR.
    - `bgp_status` - (String) The gateway BGP status.
    - `completion_notice_reject_reason` - (String) The reason for completion notice rejection. Only included on a dedicated gateways type with a rejected completion notice.
    - `cross_connect_router` - (String) The cross connect router. Only included on a dedicated gateways type..
    |`link_status` |String| The gateway link status. Only included on a dedicated gateways type.
    - `created_at` - (String) The date and time resource is created.
    - `crn` - (String) The CRN of the gateway.
    - `global` - (Bool) Gateway with global routing as **true** can connect networks outside your associated region.
    - `id` - (String) The unique identifier of the gateway.
    - `location_display_name` - (String) Long name of the gateway location.
    - `location_name` - (String) The location name of the gateway.
    - `metered` - (String) Metered billing option. If set **true** gateway usage is billed per GB. Otherwise, flat rate is charged for the gateway.
    - `name` - (String) The unique user defined name of the gateway.
    - `operational_status` - (String) The gateway operational status.
    - `port` - (Integer) The port identifier.
    - `provider_api_managed` - (Bool) Indicates the gateway is created through a provider portal. If set **true**, gateway can only be changed. If set **false**, gateway is deleted through the corresponding provider portal.
    - `resource_group` - (String) The resource group identifier.
    - `speed_mbps` - (String) The gateway speed in MBPS.
    - `type` - (String) The gateway type.
    - `vlan` - (String) The VLAN allocated for the gateway. Only set for `type=connect` gateways created directly through the IBM portal.
