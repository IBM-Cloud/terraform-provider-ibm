---
layout: "ibm"
page_title: "IBM : Instance Profile"
sidebar_current: "docs-ibm-datasources-is-instance-profile"
description: |-
  Manages IBM Cloud virtual server instance profile.
---

# ibm\_is_instance_profile

Import the details of an existing IBM Cloud virtual server instance profile as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_instance_profile" "profile" {
	name = "b-2x8"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name for this virtual server instance profile.

## Attribute Reference

The following attributes are exported:

* `family` - The family of the virtual server instance profile.
* `generation` - The platform generation of the virtual server instance profile.