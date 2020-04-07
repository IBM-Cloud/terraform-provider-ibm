---
layout: "ibm"
page_title: "IBM : private_dns_record"
sidebar_current: "docs-ibm-resource-private-dns-record"
description: |-
  Manages IBM Private DNS Record.
---

# ibm\_dns_permitted_network

Provides a private dns record resource. This allows dns record to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_private_dns_record" "pdns-1-records" {
    instance_id = ibm_resource_instance.pdns-1.guid
    zone_id = element(split("/", ibm_dns_zone.pdns-1-zone.id),1)
    type = "A"
    ttl = 900
    ipv4_address = "1.2.3.4"
    name = "example"
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The id of the private dns instance on which record has to be created. 
* `zone_id` - (Required, string) The id of the private dns zone in which record needs to be associated.
* `type` -  (Required, string) The type of private dns record that needs to be created.
* `ttl` - (Required, number) The time to live value for private dns record (cache time).
* `ipv4_address` - (Required, string) The ip address that is to be assoicated with dns record.
* `name` - (Required, string) The name of the dns record.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private dns record.


## Import

ibm_private_dns_record can be imported using private dns record id, eg

```
$ terraform import ibm_private_dns_record.example 5ffda12064634723b079acdb018ef308
```