---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_locations"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Offering Locations.
---

# `ibm_dl_locations`

Import the details of valid locations for the specified Direct Link Offering Locations. For more information, about IBM Cloud Direct Link Offerings, see [about IBM Cloud Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about).

## Example usage

```
  data "ibm_dl_locations" "test_dl_locations"{
		offering_type = "dedicated"
	 }
```

## Argument reference
Retrieve the Argument reference that you need to specify for the data source. 

- `offering_type` - (Required, String) The Direct Link offering type. Possible values are `dedicated`,`connect`.| 

## Attribute reference
Review the Attribute reference that you can access after your resource is created. 

- `locations` - (String) List of all the Direct Link Locations in the IBM Cloud infrastructure.
	- `billing_location` - (String) The billing location.
	- `building_colocation_owner` - (String) The building co-location owner. Only present for dedicated offering type. 
	- `display_name` - (String) The location long name.
	- `location_type` - (String) The location type.
	- `market` - (String) The market location.
	- `market_geography` - (String) The location geography.
	- `mzr` - (Bool) Is location a multi-zone region.
	- `name` - (String) The location short name.
	- `vpc_region` - (String) The location VPC region.
