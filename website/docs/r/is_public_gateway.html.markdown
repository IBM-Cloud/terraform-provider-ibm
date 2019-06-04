---
layout: "ibm"
page_title: "IBM : public_gateway"
sidebar_current: "docs-ibm-resource-is-public-gateway"
description: |-
  Manages IBM Public Gateway.
---

# ibm\_is_public_gateway

Provides a public gateway resource. This allows gateway to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
	name = "test"
}

resource "ibm_is_public_gateway" "testacc_gateway" {
	name = "test_gateway"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	zone = "us-south-1"

	//User can configure timeouts
  	timeouts {
      	create = "90m"
    }
}
```

## Timeouts

ibm_is_public_gateway provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating public gateway.
* `delete` - (Default 60 minutes) Used for deleting public gateway.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the gateway.
* `vpc` - (Required, string) The vpc id.
* `zone` - (Required, string) The gateway zone name.
* `floating_ip` - (Optional, string) A nested block describing the floating IP of this gateway.
Nested `floating_ip` blocks have the following structure:
  * `id` - (Optional, string) ID of the floating ip bound to the public gateway.
  * `address` - (Optional, string) IP address of the floating ip bound to the public gateway. 

## Attribute Reference

The following attributes are exported:

* `id` - The id of the gateway.
* `status` - The status of the gateway.

## Import

ibm_is_public_gateway can be imported using ID, eg

```
$ terraform import ibm_is_public_gateway.example d7bec597-4726-451f-8a63-e62e6f19c32c
```