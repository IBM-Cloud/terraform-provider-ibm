---
layout: "ibm"
page_title: "IBM : dns_secondary_zone"
sidebar_current: "docs-ibm-resource-dns-secondary-zone"
description: |-
  Manages IBM DNS Secondary Zone services
---

# ibm\_dns_secondary_zone

SoftLayer provides Secondary DNS standard to all customers to cache primary DNS Zones in the event of a loss of data. While maintaining a Secondary DNS Zone is not mandatory, it is strongly encouraged for users with multiple domains.

## Example Usage

In the following example, you can create a DNS Secondary Zone:

```hcl
resource "ibm_dns_secondary" "dns-secondary-zone-1" {
   zone_name = "new-secondary-zone1.com"
   transfer_frequency = "10"
   master_ip_address = "172.12.10.1"
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required, string) The domain’s name, including the top-level domain. For example, “example.com”.
* `transfer_frequency` - (Required, string) The transfer time (in minutes) for which the primary DNS Zone will transfer to the Secondary DNS Zone in the Transfer Interval field.
* `master_ip_address` - (Required, string) The primary target IP address to which the domain resolves. When the domain is created, an A record with a host value of @ and a data-target value of the IP address are provided and associated with the new domain.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the Dns Secondary Zone.
* `status_id` - The status identifier of the Dns Secondary Zone.
* `status_text` - The status text of the Dns Secondary Zone.

