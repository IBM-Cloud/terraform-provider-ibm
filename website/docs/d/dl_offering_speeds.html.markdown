---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_offering_speeds"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Offering Speeds.
---

# ibm_dl_offering_speeds

Import the details of an existing IBM Cloud Infrastructure Direct Link offering speed options. For more information, about Direct Link Offering speed, see [arranging for Direct Link connectivity](https://cloud.ibm.com/docs/dl?topic=dl-pricing-for-ibm-cloud-dl#arranging-for-dl-conectivity).


## Example usage

```terraform
data "ibm_dl_offering_speeds" "ds_dlspeedoptions" {
  offering_type="dedicated"
}
```

## Argument reference
Retrieve the argument reference that you need to specify for the data source. 

- `offering_type` - (Required, String) The Direct Link offering type. Possible values are `dedicated`,`connect`.| 

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `offering_speeds` - (String) List of all the Direct Link offering speeds in the IBM Cloud infrastructure.

  Nested scheme for `offering_speeds`:
  - `capabilities` - (String) The capabilities for billing option.
  - `link_speed` - (String) The link speed in megabits per second.
