---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Private DNS Permitted Networks"
description: |-
Manages IBM Cloud Infrastructure Private Domain Name Service Zones Permitted Networks.

---

# ibm_pdns_permitted_networks

Import the details of an existing IBM Cloud Infrastructure private domain name service zones permitted networks as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_resource_group" "rg" {
  name = "default"
}

resource "ibm_is_vpc" "test_pdns_vpc" {
  name           = "test-pdns-vpc"
  resource_group = data.ibm_resource_group.rg.id
}

resource "ibm_resource_instance" "test-pdns-instance" {
  name              = "test-pdns"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

resource "ibm_dns_zone" "test-pdns-zone" {
  name        = "test.com"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description = "testdescription"
  label       = "testlabel-updated"
}

resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  vpc_crn     = ibm_is_vpc.test_pdns_vpc.crn
}

data "ibm_dns_permitted_networks" "test" {
  instance_id = ibm_dns_permitted_network.test-pdns-permitted-network-nw.instance_id
  zone_id     = ibm_dns_permitted_network.test-pdns-permitted-network-nw.zone_id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The resource instance id of the private DNS on which permitted networks were created.
* `zone_id` - (Required, string) The resource zone id of the private DNS on which permitted networks were created.

## Attribute Reference

The following attributes are exported:

* `permitted_networks` - List of all private domain name service zones permitted networks in the IBM Cloud Infrastructure.
  * `created_on` - The created time of the Private DNS zone.
  * `instance_id` - The resource instance id of the Private DNS on which zones were created.
  * `modified_on` - The modified time of the Private DNS zone.
  * `permitted_network` - The permitted networks crn detail.
    * `vpc_crn` - The VPC CRN number.
  * `permitted_network_id` - The unique identifier for this instance
  * `state` - The state of the Private DNS zone.
  * `type` - The type of Private DNS.
  * `zone_id` - The unique identifier of the private DNS zone.
