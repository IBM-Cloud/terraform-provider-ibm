---
layout: "ibm"
page_title: "IBM : floating_ip"
sidebar_current: "docs-ibm-resource-is-floating-ip"
description: |-
  Manages IBM Floating IP.
---

# ibm\_is_floating_ip

Provides a floatingip resource. This allows floatingip to be created, updated, and cancelled.


## Example Usage

```hcl

resource "ibm_is_instance" "testacc_instance" {
  name    = "testinstance"
  image   = "7eb4e35b-4257-56f8-d7da-326d85452591"
  profile = "b-2x8"

  primary_network_interface = {
    port_speed = "1000"
    subnet     = "70be8eae-134c-436e-a86e-04849f84cb34"
  }

  vpc  = "01eda778-b822-43a2-816d-d30713df5e13"
  zone = "us-south-1"
  keys = ["eac87f33-0c00-4da7-aa66-dc2d972148bd"]
}

resource "ibm_is_floating_ip" "testacc_floatingip" {
  name   = "testfip1"
  target = "${ibm_is_instance.testacc_instance.primary_network_interface.0.id}"
}

```

## Timeouts

ibm_is_instance provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `delete` - (Default 60 minutes) Used for deleting floating IP.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The floating ip name.
* `target` - (Optional, string) ID of the target network interface.
    **NOTE**: Conflicts with `zone`.
* `zone` - (Optional, string) Name of the target zone. 
    **NOTE**: Conflicts with `target`.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the floating ip.
* `status` - The status of the floating ip.
* `address` - The floating ip address. 
