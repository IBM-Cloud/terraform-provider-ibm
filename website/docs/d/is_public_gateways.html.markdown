---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_gateways"
description: |-
  Manages IBM Public Gateways.
---

# ibm\_is_public_gateways

Import the details of an existing IBM Cloud Infrastructure Public Gateways as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.



## Example Usage

```hcl
data "ibm_is_public_gateways" "testacc_dspgw"{
}

```

## Attribute Reference

The following attributes are exported:
* `public_gateways` - List of all Public Gateways in the IBM Cloud Infrastructure region.
  * `id` - The id of the public gateway.
  * `status` - The status of the gateway.
  * `vpc` - The vpc id of gateway.
  * `zone` - The gateway zone name.
  * `tags` - Tags associated with the Public gateway.
  * `name` - The name of the public gateway.
  * `floating_ip` - A nested block describing the floating IP of this gateway.
  Nested `floating_ip` blocks have the following structure:
    * `id` - ID of the floating ip bound to the public gateway.
    * `address` - IP address of the floating ip bound to the public gateway.