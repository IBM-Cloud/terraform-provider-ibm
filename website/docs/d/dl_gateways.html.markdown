---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_gateways"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Gateways.
---

# ibm_dl_gateways

Import the details of an existing IBM Cloud Infrastructure Direct Link Gateways.  For more information, about IBM Cloud Direct Link, see [getting started with IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-get-started-with-ibm-cloud-dl).


## Example usage

---
```terraform
data "ibm_dl_gateways" "ds_dlgateways" {
}
     
```
---
## Argument reference
There is no argument reference that you need to specify for the data source. 


## Attribute reference
You can access the following attribute references after your data source is created.

- `gateways` - (String) List of all the Direct Link Gateways in the IBM Cloud infrastructure.

  Nested scheme for `gateways`:
  - `as_prepends` - (List) List of AS Prepend configuration information
    Nested scheme for `as_prepend`:
     - `created_at`- (String) The date and time AS Prepend was created.
     - `id` - (String) The unique identifier for this AS Prepend.
     - `length` - (Integer) Number of times the ASN to appended to the AS Path.
     - `policy` - (String) Route type this AS Prepend applies to. Possible values are `import` and `export`.
     - `prefix` - (Deprecated, String) Comma separated list of prefixes this AS Prepend applies to. Maximum of 10 prefixes. If not specified, this AS Prepend applies to all prefixes.
     - `specific_prefixes` - (Array of Strings) Array of prefixes this AS Prepend applies to. 
     - `updated_at`- (String) The date and time AS Prepend was updated   
  - `authentication_key` - (String) BGP MD5 authentication key.
  - `bfd_interval` - (String) Minimum interval in milliseconds at which the local routing device transmits hello packets and then expects to receive a reply from a neighbor with which it has established a BFD session.
  - `bfd_multiplier` - (String) The number of hello packets not received by a neighbor that causes the originating interface to be declared down.
  - `bfd_status` - (String) Gateway BFD status.
  - `bfd_status_updated_at` - (String) Date and time bfd status was updated.
  - `bgp_asn` - (String) Customer BGP ASN.
  - `bgp_base_cidr` - (String) The BGP base CIDR.
  - `bgp_cer_cidr` - (String) The BGP customer edge router CIDR.
  - `bgp_ibm_asn` - (String) The IBM BGP ASN.
  - `bgp_ibm_cidr` - (String) The IBM BGP  CIDR.
  - `bgp_status` - (String) The gateway BGP status.
  - `bgp_status_updated_at` - (String) Date and time bgp status was updated.
  - `default_export_route_filter` - (String) The default directional route filter action    that applies to routes that do not match any directional route filters. 
  - `default_import_route_filter` - (String) The default directional route filter action that applies to routes that do not match any directional route filters.
  - `completion_notice_reject_reason` - (String) The reason for completion notice rejection. Only included on a dedicated gateways type with a rejected completion notice.
  - `connection_mode` - (String) Type of network connection that you want to bind to your direct link.
  - `cross_connect_router` - (String) The cross connect router. Only included on a dedicated gateways type..
  - `link_status` - (String) The gateway link status. Only included on a dedicated gateways type.
  - `link_status_updated_at` - (String) Date and time link status was updated.
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
  - `vlan` - (String) The VLAN allocated for the gateway.
