---
layout: "ibm"
page_title: "IBM : dl_gateways"
sidebar_current: "docs-ibm-datasource-dl-gateways"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateway.
---

# ibm\_dl_gateways

Import the details of an existing IBM Cloud Infrastructure direct link gateway as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_dl_gateways" "ds_dlgateways" {
}
```

## Attribute Reference

The following attributes are exported:

* `gateways` - List of all Direct Link gateways in the IBM Cloud Infrastructure.
  * `bgp_asn` - Customer BGP ASN.
  * `created_at` - The date and time resource was created.
  * `crn` - The CRN (Cloud Resource Name) of this gateway.
  * `global` - Gateways with global routing (true) can connect to networks outside their associated region.
  * `id` - The unique identifier of this gateway.
  * `location_display_name` - Gateway location long name.
  * `location_name` - Gateway location.
  * `metered` - Metered billing option. When true gateway usage is billed per gigabyte. When false there is no per gigabyte usage charge, instead a flat rate is charged for the gateway.
  * `name` - The unique user-defined name for this gateway.
  * `operational_status` - Gateway operational status.
  * `resource_group` - Resource group identifier.
  * `speed_mbps` - Gateway speed in megabits per second.
  * `type` - Gateway type.
  * `bgp_base_cidr` - BGP base CIDR.
  * `bgp_cer_cidr` - BGP customer edge router CIDR.
  * `bgp_ibm_asn` - IBM BGP ASN.
  * `bgp_ibm_cidr` - BGP IBM CIDR.
  * `bgp_status` - Gateway BGP status.
  * `completion_notice_reject_reason` - Reason for completion notice rejection. Only included on type=dedicated gateways with a rejected completion notice.
  * `cross_connect_router` - Cross connect router. Only included on type=dedicated gateways.
  * `link_status` - Gateway link status. Only included on type=dedicated gateways.
  * `port` - Port Identifier.
  * `provider_api_managed` - Indicates whether gateway was created through a provider portal. If true, gateway can only be changed or deleted through the corresponding provider portal.
  * `vlan` - VLAN allocated for this gateway. Only set for type=connect gateways created directly through the IBM portal.

