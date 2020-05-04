---
layout: "ibm"
page_title: "IBM : dns_permitted_network"
sidebar_current: "docs-ibm-resource-dns-permitted-network"
description: |-
  Manages IBM Private DNS Permitted Network.
---

# ibm\_dns_permitted_network

Provides a private dns permitted network resource. This allows dns permitted network to be created, and deleted.

## Example Usage

```hcl

resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
    instance_id = ibm_resource_instance.test-pdns-instance.guid
    zone_id = ibm_dns_zone.test-pdns-zone.zone_id
    vpc_crn = ibm_is_vpc.test_pdns_vpc.crn
    type = "vpc"
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The id of the private DNS on which zone has to be created.
* `zone_id` - (Required, string) The id of the private DNS zone in which the network needs to be associated.
* `vpc_crn` -  (Required, string) The CRN of VPC instance.
* `type` - (Required, string) The permitted network type. Valid values: "vpc".

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the private DNS zone. The id is composed of <instance_id>/<zone_id>/<permitted_network_id>.
* `created_on` - The time (Created On) of the DNS permitted network. 
* `modified_on` - The time (Modified On) of the DNS permitted network.

## Import

ibm_dns_permitted_network can be imported using private DNS instance ID, zone ID and permitted network ID, eg

```
$ terraform import ibm_dns_permitted_network.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```