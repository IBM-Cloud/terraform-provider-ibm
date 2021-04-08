---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_provider_ports"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Provider Ports.
---

# ibm\_dl_ports

Import the details of an existing IBM Cloud Infrastructure directlink provider ports (associated with the caller) as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_dl_provider_ports" "ds_dl_provider_ports" {
}
```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `ports` - List of all Direct Link ports in the IBM Cloud Infrastructure.
  * `label` - Port Label.
  * `location_display_name` - Port location long name.
  * `location_name` - Port location name identifier.
  * `port_id` - Port identifier.
  * `provider_name` - Port's provider name.
  * `supported_link_speeds` - Port's supported speeds in megabits per second.
  

