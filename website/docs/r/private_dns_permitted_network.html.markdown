---
layout: "ibm"
page_title: "IBM : dns_permitted_network"
sidebar_current: "docs-ibm-resource-dns-permitted-network"
description: |-
  Manages IBM Private DNS Permitted Network.
---

# ibm\_dns_permitted_network

Provides a private dns permitted network resource. This allows dns permitted network to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_dns_permitted_network" "pdns-1-permitted-network" {
    instance_id = ibm_resource_instance.pdns-1.guid
    zone_id = element(split("/", ibm_dns_zone.pdns-1-zone.id),1)
    vpc_crn = ibm_is_vpc.test_vpc.resource_crn
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The id of the private dns instance. 
* `zone_id` - (Required, string) The id of the private dns zone in which network needs to be associated.
* `vpc_crn` -  (Required, string) The crn number of the network that needs to be associated.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private dns permitted network association.


## Import

ibm_dns_permitted_network can be imported using permitted network association id, eg

```
$ terraform import ibm_dns_permitted_network.example 5ffda12064634723b079acdb018ef308
```