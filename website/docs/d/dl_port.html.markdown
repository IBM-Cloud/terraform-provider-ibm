---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_port"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Port.
---

# `ibm_dl_port`

Import the details of an existing IBM Cloud Infrastructure Direct Link Offering Port. For more information, about Direct Link Offering Port, see [Megaport ordering considerations](https://cloud.ibm.com/docs/dl?topic=dl-megaport).


## Example usage

```
  data "ibm_dl_port" "ds_dlport" {
      port_id = "dl_port_id"
   }
```

## Argument reference
Retrieve the Argument reference that you need to specify for the data source. 

- `port_id` - (Required, String) The unique ID for the Direct Link port.

## Attribute reference
Review the Attribute reference that you can access after your resource is created. 

- `direct_link_count` - (String) The count of the existing Direct Link gateways on the port.
- `label` - (String) The port label.
- `location_display_name` - (String) The port location long name.
- `location_name` - (String) The port location name.
- `provider_name` - (String) The port's provider name.
- `supported_link_speeds` - (String) The port supported speeds in megabits per second.
