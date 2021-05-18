---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_provider_gateways"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Provider Gateway.
---

# `ibm_dl_provider_gateways`

Import the details of an existing nfrastructure Direct Link Provider Gateway as a read-only data source.  For more information, refer to [about Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about#use-case-connect).


## Example usage

```
data "ibm_dl_provider_gateways" "ds_dlproviderGateways" {
}
```


## Argument reference
There is no Argument reference that you need to specify for the data source. 


## Attribute reference
Review the Attribute reference that you can access after your resource is created. 

- `gateways` - (String) List of all the Direct Link ports.List of all Direct Link provider gateways in the IBM Cloud Infrastructure..
    - `bgp_asn` - (String) The customer BGP ASN.
    - `bgp_cer_cidr` - (String) The BGP customer edge router CIDR.
    - `bgp_ibm_asn` - (String) The IBM BGP ASN.
    - `bgp_ibm_cidr` - (String) The IBM BGP CIDR.
    - `bgp_status` - (String) The gateway BGP status.
    - `created_at` - (String) The date and time resource was created.
    - `crn` - (String) The CRN of the gateway.
    - `global` - (Bool) The gateways with global routing set as **true** can connect to networks outside their associated region.
    - `id` - (String) The unique identifier of the gateway.
    - `name` - (String) The unique user defined name for the gateway.
    - `operational_status` - (String) The operational status of the gateway.
    - `port` - (String) The port identifier.
    - `provider_api_managed` - (String)  Indicates whether gateway was created through a provider portal. If set **true**, gateway can only be changed or deleted through the corresponding provider portal.
    - `resource_group` - (String) The resource group identifier.
    - `speed_mbps` - (String) The gateway speed in megabits per second.
    - `type` - (String) The gateway type.
    - `vlan` - (String) The VLAN allocated for the gateway. Only set for `type=connect` gateways created directly through the IBM portal.
