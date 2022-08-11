---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_location"
description: |-
  Manages IBM Cloud Infrastructure Transit location.
---

# ibm_tg_location
Retreive information of an existing IBM Cloud infrastructure transit location as a read only data source. For more information, about transit location, see [about IBM Cloud Transit Gateway](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-about).


## Example usage

```terraform
data "ibm_tg_location" "ds_tg_location" {
  name = "us-south"
}
```
## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the transit gateway location.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `billing_location` - (String) The geographical location of the location, used for billing purposes.
- `name` - (String) The name of the location.
- `type` - (String) The type of the location, determining a `multi-zone region`, a `single data center`, or a `point of presence`.
- `local_connection_locations` - (String) The set of network locations that are considered local for the transit gateway location.

  Nested scheme for `local_connection_locations`:
  - `display_name` - (String) The descriptive display name for the location.
  - `name` - (String) The name of the location.
  - `type` - (String) The type of the location, determining a `multi-zone region`, a `single data center`, or a `point of presence`.
