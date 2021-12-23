---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Private DNS Permitted Networks"
description: |-
  Manages IBM Cloud infrastructure private domain name service zones permitted networks.
---

# ibm_dns_permitted_networks

Retrieve details about permitted networks for a zone that is associated with the private DNS service instance. For more information, see [managing permitted networks](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-permitted-networks).


## Example usage

```terraform
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

## Argument reference
Review the argument reference that you can specify for your data source.

- `instance_id` - (Required, String) The GUID of the private DNS service instance where you created permitted networks.
- `zone_id` - (Required, String) The ID of the zone where you added the permitted networks.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `permitted_networks`- (List) List of permitted networks-A list of all permitted networks that were created for a zone in your private DNS instance.

  Nested scheme for `permitted_networks`:
  - `created_on`- (Timestamp) The date and time when the permitted network was created.
  - `instance_id` - (String) The ID of the private DNS service instance where you created permitted networks.
  - `modified_on`- (Timestamp) The date and time when the permitted network was updated.
  - `permitted_network`- (List of VPCs) A list of VPC CRNs that are associated with the permitted network.
    
     Nested scheme for `permitted_network`:
     - `vpc_crn` - (String) The CRN of the VPC that the permitted network belongs to. 
- `permitted_network_id` - (String) The ID of the permitted network.
- `state` - (String) The state of the permitted network. 
- `type` - (String) The type of the permitted network.
- `zone_id` - (String) The ID of the zone where you added the permitted network.
