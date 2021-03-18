---

subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_zone"
description: |-
  Manages IBM Private DNS Zone.
---

# ibm\_dns_zone

Provides a private dns zone resource. This allows dns zones to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_dns_zone" "pdns-1-zone" {
    name = "test.com"
    instance_id = p-dns-instance-guid
    description = "testdescription"
    label = "testlabel"
}


```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the DNS zone to be created.
* `instance_id` - (Required, string) The guid of the private DNS on which zone has to be created. 
* `description` - (Optional, string) The text describing the purpose of a DNS zone.
* `label` -  (Optional, string) The label of a DNS zone.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private DNS zone. The id is composed of <instance_guid>/<zone_id>.
* `zone_id` - The unique identifier of the private DNS zone.
* `created_on` - The time (Created On) of the DNS zone. 
* `modified_on` - The time (Modified On) of the DNS zone.
* `state` - The state of the DNS zone.

## Import

ibm_dns_zone can be imported using private DNS instance ID and zone ID, eg

```
$ terraform import ibm_dns_zone.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308
```
