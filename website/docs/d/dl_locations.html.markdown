---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_locations"
description: |-
  Manages IBM Cloud Infrastructure Directlink Offering Locations.
---

# ibm_dl_locations
Retrieve the list of valid locations for the specified Direct Link offering.
## Example Usage

```hcl
  data "ibm_dl_locations" "test_dl_locations"{
		offering_type = "dedicated"
	 }
```

## Argument Reference

The following arguments are supported:

* `offering_type` - (Required, string) The Direct Link offering type. Current supported values are "dedicated" and "connect".
Allowable values: [dedicated,connect]. Example: dedicated
## Attribute Reference

The following attributes are exported:

* `locations` - List of Direct Link locations
  * `billing_location` - Billing location.
  * `building_colocation_owner` - Building colocation owner. Only present for offering_type=dedicated locations.
  * `name` - Location short name
  * `display_name` - Location long name.
  * `location_type` - Location type.
  * `market` - Location market.
  * `market_geography` - Location geography.
  * `mzr` - Is location a multi-zone region (MZR)
  * `vpc_region` - Location's VPC region.
 