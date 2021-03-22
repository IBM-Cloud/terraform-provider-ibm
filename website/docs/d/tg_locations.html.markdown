---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_locations"
description: |-
  Manages IBM Cloud Infrastructure Transit Locations.
---

# ibm\_tg_locations

Import the details of an existing IBM Cloud Infrastructure transit locations as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_tg_locations" "ds_tg_locations" {
}
```

## Attribute Reference

The following attributes are exported:

* `locations` - List of all locations that support Transit Gateways in the IBM Cloud Infrastructure.
  * `billing_location` - The geographical location of this location, used for billing purposes.
  * `name` - Name of the Location.
  * `type` - The type of the location, determining is this a multi-zone region, a single data center, or a point of presence.
