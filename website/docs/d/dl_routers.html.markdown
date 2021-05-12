---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_routers"
description: |-
  Retrieve location specific cross connect router information. Only valid for offering_type=dedicated locations.
---

# `ibm_dl_routers`

Import the details of an existing IBM Cloud infrastructure Direct Link Location specific cross connect router information. For more information, about Direct Link cross connect router, see [Virtual routing and forwarding on IBM Cloud](https://cloud.ibm.com/docs/dl?topic=dl-overview-of-virtual-routing-and-forwarding-vrf-on-ibm-cloud).


## Example usage

```
data "ibm_dl_routers" "test_dl_routers" {
	offering_type = "dedicated"
	location_name = "dal09"
}
```

## Argument reference
The Argument reference that you need to specify for the data source. 

- `offering_type` - (Required, String) The Direct Link offering type. Only `dedicated` is supported in this API.
- `location_name` - (Required, String) The name of the Direct Link Location.

## Attribute reference
Review the Attribute reference that you can access after your resource is created. 

- `cross_connect_routers` - (String) List of all the cross connect router details.
	- `capabilities` - (String) Macsec and non-macsec capabilities for the router.
	- `router_name` - (String) The name of the router.
	- `total_connections` - (String) Count of existing Direct Link dedicated gateways on this router account.
