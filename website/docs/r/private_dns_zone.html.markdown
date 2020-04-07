---
layout: "ibm"
page_title: "IBM : dns_zone"
sidebar_current: "docs-ibm-resource-dns-zone"
description: |-
  Manages IBM Private DNS Zone.
---

# ibm\_dns_zone

Provides a private dns zone resource. This allows dns zones to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_dns_zone" "pdns-1-zone" {
    name = "test.com"
    instance_id = p-dns-instance-id
    description = "testdescription"
    label = "testlabel"
}


```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string)The name of the pdns zone to be created.
* `instance_id` - (Required, string) The id of the private dns instance on which zone has to be created. 
* `description` - (Optional, string) description of the zone to be created.
* `label` -  (Optional, string) label of the zone to be created.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private dns zone.


## Import

ibm_dns_zone can be imported using zone id, eg

```
$ terraform import ibm_dns_zone.example 5ffda12064634723b079acdb018ef308
```