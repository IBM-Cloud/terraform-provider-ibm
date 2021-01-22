---
layout: "ibm"
page_title: "IBM : Volume Profile"
sidebar_current: "docs-ibm-datasources-is-volume-profile"
description: |-
  Manages IBM Cloud virtual server volume profile.
---

# ibm\_is_volume_profile

Import the details of an existing IBM Cloud virtual server volume profile as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_volume_profile" "volprofile"{
  name = "general-purpose"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name for this virtual server volume profile.

## Attribute Reference

The following attributes are exported:

* `family` - The family of the virtual server volume profile.