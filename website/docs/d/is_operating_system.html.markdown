---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : operating_system"
description: |-
  Manages IBM Operating System.
---

# ibm\_is_operating_system

Provides a Operating System datasource. This allows to fetch an existing operating system details.


## Example Usage

```terraform
data "ibm_is_operating_system" "testacc_dsos"{
  name = "red-8-amd64"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The globally unique name for this operating system.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The globally unique name for this operating system.
* `architecture` - The operating system architecture.
* `dedicated_host_only` - Images with this operating system can only be used on dedicated hosts or dedicated host groups.
* `display_name` - A unique, display-friendly name for the operating system.
* `family` - The name of the software family this operating system belongs to.
* `href` - The URL for this operating system.
* `name` - The globally unique name for this operating system.
* `vendor` - The vendor of the operating system.
* `version` - The major release version of this operating system.