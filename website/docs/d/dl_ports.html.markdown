---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_ports"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Ports.
---

# ibm_dl_ports

Import the details of an existing IBM Cloud infrastructure Direct Link  ports. For more information, about Direct Link Offering Port, see [megaport ordering considerations](https://cloud.ibm.com/docs/dl?topic=dl-megaport).


## Example usage

```terraform
data "ibm_dl_ports" "ds_dlports" {
}
```

## Argument reference
Retrieve the argument reference that you need to specify for the data source. 

- `location_name` - (Optional, string) Direct Link location short name.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `ports` - (List) List of all the Direct Link Ports.
 
  Nested scheme for `ports`:
  - `direct_link_count` - (String) Count of the existing Direct Link gateways in this port account.
  - `label` - (String) The port label.
  - `location_display_name` - (String) The port location long name.
  - `location_name` - (String) The port location name.
  - `port_id` - (String) The port identifier.
  - `provider_name` - (String) The port's provider name.
  - `supported_link_speeds` - (String) The port supported speeds in megabits per second.
