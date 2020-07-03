---
layout: "ibm"
page_title: "IBM : tg_gateway"
sidebar_current: "docs-ibm-datasource-tg-gateway"
description: |-
  Manages IBM Cloud Infrastructure Transit Gateway.
---

# ibm\_tg_gateway

Import the details of an existing IBM Cloud Infrastructure transit gateway as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

resource "ibm_tg_gateway" "new_tg_gw"{
name="transit-gateway-1"
location="us-south"
global=true
resource_group="30951d2dff914dafb26455a88c0c0092"
} 

data "ibm_tg_gateway" "ds_tggateway" {
    id=ibm_tg_gateway.new_tg_gw.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the gateway.

## Attribute Reference

The following attributes are exported:

  * `created_at` - The date and time resource was created.
  * `updated_at` - The date and time resource was last updated.
  * `crn` - The CRN (Cloud Resource Name) of this gateway.
  * `global` - Gateways with global routing (true) can connect to networks outside their associated region.
  * `location` - Gateway location.
  * `id` - The unique identifier of this gateway.
  * `status` - Gateway status.
  * `resource_group` - Resource group identifier.