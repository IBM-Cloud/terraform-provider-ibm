---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_location_nlb_dns"
description: |-
  Managed satellite location nlb dns.
---

# ibm_satellite\_location\_nlb\_dns

Provides a resource to register public ip address to satellite dns records. This allows satellite dns register to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_satellite_location_nlb_dns" "satellite_dns" {
  location = "satellite-ibm"
  ips      = ["52.116.125.50","169.62.17.178","169.63.178.155"]
}
```

## Argument reference

The following arguments are supported:

* `ips` - (Required, Forces new resource, List) Public IP address of satellite location DNS records.
* `location` - (Required, Forces new resource, string) The name or ID of the Satellite location.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the ibm_satellite_location_nlb_dns.

## Import

The import functionality is not supported for this resource.
