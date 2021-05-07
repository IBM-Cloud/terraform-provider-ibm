---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_gateway"
description: |-
  Manages IBM Public Gateway.
---

# ibm\_is_public_gateway

Provides a public gateway datasource. This allows to fetch an existing gateway.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_public_gateway" "testacc_gateway" {
  name = "test-gateway"
  vpc  = ibm_is_vpc.testacc_vpc.id
  zone = "us-south-1"
}

data "ibm_is_public_gateway" "testacc_dspgw"{
  name = ibm_is_public_gateway.testacc_public_gateway.name
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the gateway.
* `resource_group` - (Optional, string) The resource group ID of the Public gateway. (This argument is supported only for Generation `2` infrastructure)

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id of the public gateway.
* `status` - The status of the gateway.
* `vpc` - The vpc id of gateway.
* `zone` - The gateway zone name.
* `tags` - Tags associated with the Public gateway.
* `name` - The name of the public gateway
* `floating_ip` - A nested block describing the floating IP of this gateway.
Nested `floating_ip` blocks have the following structure:
  * `id` - ID of the floating ip bound to the public gateway.
  * `address` - IP address of the floating ip bound to the public gateway.