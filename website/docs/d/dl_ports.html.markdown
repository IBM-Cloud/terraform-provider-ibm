---
layout: "ibm"
page_title: "IBM : dl_ports"
sidebar_current: "docs-ibm-datasource-dl-ports"
description: |-
  Manages IBM Cloud Infrastructure Direct Link Ports.
---

# ibm\_dl_ports

Import the details of an existing IBM Cloud Infrastructure direct link ports as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_dl_ports" "ds_dlports" {
}
```
## Argument Reference

The following arguments are supported:

* `location_name` - (Optional, string) Direct Link location short name.

## Attribute Reference

The following attributes are exported:

* `ports` - List of all Direct Link ports in the IBM Cloud Infrastructure.
  * `direct_link_count` - Count of existing Direct Link gateways in this account on this port.
  * `label` - Port Label.
  * `location_display_name` - Port location long name.
  * `location_name` - Port location name identifier.
  * `port_id` - Port identifier.
  * `provider_name` - Port's provider name.
  * `supported_link_speeds` - Port's supported speeds in megabits per second.
  

