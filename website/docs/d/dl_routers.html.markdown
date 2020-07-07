---
layout: "ibm"
page_title: "IBM : dl_routers"
sidebar_current: "docs-ibm-datasource-dl-routers"
description: |-
  Retrieve location specific cross connect router information. Only valid for offering_type=dedicated locations.
---

# ibm\_dl_routers

Import the details of an existing IBM Cloud Infrastructure direct link location specific cross connect router information as a read-only data source. Only valid for offering_type=dedicated locations. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_dl_routers" "ds_dlrouters" {
  offering_type="dedicated"
  location_name = "dal09"
}

```

## Argument Reference

The following arguments are supported:

* `offering_type` - (Required, string) The Direct Link offering type. Only value "dedicated" is supported for this API.
* `location_name` - (Required, string) The name of the Direct Link location.

## Attribute Reference

The following attributes are exported:

* `cross_connect_routers` - List of cross connect router details
  * `router_name` - The name of the Router.
  * `total_connections` - Count of existing Direct Link Dedicated gateways on this router for this account.


