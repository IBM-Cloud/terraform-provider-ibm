---
layout: "ibm"
page_title: "IBM : tg_location"
sidebar_current: "docs-ibm-datasource-tg-location"
description: |-
  Manages IBM Cloud Infrastructure Transit Location.
---

# ibm\_tg_location

Import the details of an existing IBM Cloud Infrastructure transit location as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_tg_location" "ds_tg_location" {
  name = "us-south"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The Transit Gateway location Name.


## Attribute Reference

The following attributes are exported:

* `billing_location` - The geographical location of this location, used for billing purposes.
* `name` - Name of the Location.
* `type` - The type of the location, determining is this a multi-zone region, a single data center, or a point of presence.
* `local_connection_locations` - The set of network locations that are considered local for this Transit Gateway location.
   * `display_name` - A descriptive display name for the location.
   * `name` - The name of the location.
   * `type` - The type of the location, determining is this a multi-zone region, a single data center, or a point of presence.

