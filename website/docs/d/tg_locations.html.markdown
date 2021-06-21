---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_locations"
description: |-
  Manages IBM Cloud Infrastructure Transit locations.
---

# ibm_tg_locations
Retrieves information of an existing IBM Cloud infrastructure transit location as a read only data source. For more information, about transit location, see [IBM Cloud Transit Gateway locations](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-tg-locations).

## Example usage

```terraform
data "ibm_tg_locations" "ds_tg_locations" {
}
```

## Argument reference
There is no argument reference for `ibm_tg_locations`.

## Attribute reference
You can access the following attribute references after your data source is created. 

- `locations` - (String) List of all locations that supports transit gateways.

  Nested scheme for `locations`:
 - `billing_location` - (String) The geographical location of the location, used for billing purposes.
 - `name` - (String) The name of the location.
 - `type` - (String) The type of the location, determining a `multi-zone region`, a `single data center`, or a `point of presence`.
