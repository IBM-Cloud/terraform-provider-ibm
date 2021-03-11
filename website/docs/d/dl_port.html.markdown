---
subcategory: "Direct Link Gateway"
layout: "ibm"
page_title: "IBM : dl_port"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Port.
---

# ibm\_dl_port

Import the details of an existing IBM Cloud Infrastructure direct link port as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_dl_port" "ds_dlport" {
    port_id = "dl_port_id"
}
```

## Argument Reference

The following arguments are supported:

* `port_id` - (Required, string) The unique port_id for this dl port.

## Attribute Reference

The following attributes are exported:

* `direct_link_count` - Count of existing Direct Link gateways in this account on this port.
* `label` - Port Label.
* `location_display_name` - Port location long name.
* `location_name` - Port location name identifier.
* `provider_name` - Port's provider name.
* `supported_link_speeds` - Port's supported speeds in megabits per second.
  

