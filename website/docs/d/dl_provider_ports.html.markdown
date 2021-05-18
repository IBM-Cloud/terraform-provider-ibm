---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_provider_ports"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Provider Ports.
---

# `ibm_dl_provider_ports`

Import the details of an existing IBM Cloud Infrastructure Direct Link Provider Ports. For more information, about Direct Link Provider Ports, see [About Direct Link](https://cloud.ibm.com/docs/dl?topic=dl-dl-about).


## Example usage

```
data "ibm_dl_provider_ports" "ds_dl_provider_ports" {
}
```


## Argument reference
There is no Argument reference that you need to specify for the data source. 


## Attribute reference
Review the Attribute reference that you can access after your resource is created. 

- `ports` - (String) List of all the Direct Link ports in the IBM Cloud infrastructure.
    - `label` - (String) The port label.
    - `location_display_name` - (String) The port location long name.
    - `location_name` - (String) The port location name.
    - `port_id` - (String) The port identifier.
    - `provider_name` - (String) The port's provider name.
    - `supported_link_speeds` - (String) The port supported speeds in megabits per second.
