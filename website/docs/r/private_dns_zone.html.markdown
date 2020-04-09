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

* `name` - (Required, string) The name of the DNS zone to be created.
* `instance_id` - (Required, string) The id of the private DNS on which zone has to be created. 
* `description` - (Optional, string) The text describing the purpose of a DNS zone.
* `label` -  (Optional, string) The label of a DNS zone.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private DNS zone.
* `created_on` - The time (Created On) of the DNS zone. 
* `modified_on` - The time (Modified On) of the DNS zone.
* `instance_id` - The instance id of the DNS instance on which zone created. 
* `name` - The name of the DNS zone created.
* `description` - The text describing the purpose of a DNS zone.
* `state` - The state of the DNS zone.
* `label` - The label of a DNS zone.

## Import

ibm_dns_zone can be imported using zone id, eg

```
$ terraform import ibm_dns_zone.example 5ffda12064634723b079acdb018ef308
```