---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_gateway"
description: |-
  Manages IBM Cloud Infrastructure Transit Gateway.
---

# ibm_tg_gateway
Retrieve information of an existing IBM Cloud infrastructure transit gateway as a read only data source. For more information, about Transit Gateway, see [getting started with IBM Cloud Transit Gateway](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-getting-started).


## Example usage

```terraform
resource "ibm_tg_gateway" "new_tg_gw" {
  name           = "transit-gateway-1"
  location       = "us-south"
  global         = true
  resource_group = "30951d2dff914dafb26455a88c0c0092"
}

data "ibm_tg_gateway" "ds_tggateway" {
  name = ibm_tg_gateway.new_tg_gw.name
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the gateway.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `created_at` - (Timestamp) The date and time resource is created.
- `crn` - (String) The CRN of the gateway.
- `global` - (String) The gateways with global routing true to connect to the networks outside the associated region.
- `id` - (String) The unique identifier of this gateway.
- `location` - (String) The gateway location.
- `resource_group` - (String) The resource group identifier.

  Nested scheme for `resource_group`:
	- `created_at` - (String) The date and time the connection is created.
	- `id` - (String) The unique identifier for the transit gateway connection to network either `VPC` or `classic`).
	- `name` - (String) The user-defined name for the transit gateway connection.
	- `network_type` - (String) The type of network connected with the connection. Possible values are `classic` or `VPC`.
	- `network_account_id` - (String) The ID of the network connected account. This is used if the network is in a different account than the gateway.
	- `network_id` - (String) The ID of the network being connected with the connection.
	- `status` - (String) The current configuration state of the connection. Possible values are `attached`, `failed,` `pending`, `deleting`.
	- `updated_at` - (String) The date and time the connection is last updated.
- `status` - (String) The gateway status.
- `updated_at` - (Timestamp) The date and time resource is last updated.
