---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_routers"
description: |-
  Retrieve location specific cross connect router information. Only valid for offering_type=dedicated locations.
---

# ibm_dl_routers

Import the details of an existing IBM Cloud infrastructure Direct Link Location specific cross connect router information. For more information, about Direct Link cross connect router, see [virtual routing and forwarding on IBM Cloud](https://cloud.ibm.com/docs/dl?topic=dl-overview-of-virtual-routing-and-forwarding-vrf-on-ibm-cloud).


## Example usage

```terraform
data "ibm_dl_routers" "test_dl_routers" {
	offering_type = "dedicated"
	location_name = "dal09"
}
```

## Argument reference
The argument reference that you need to specify for the data source. 

- `offering_type` - (Required, String) The Direct Link offering type. Only `dedicated` is supported in this API.
- `location_name` - (Required, String) The name of the Direct Link Location.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `cross_connect_routers` - (List) List of all the cross connect router details.

  Nested scheme for `cross_connect_routers`:
  - `capabilities` - (String) Macsec and non-macsec capabilities for the router.
  - `router_name` - (String) The name of the router.
  - `total_connections` - (String) Count of existing Direct Link dedicated gateways on this router account.
