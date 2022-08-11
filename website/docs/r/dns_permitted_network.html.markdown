---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_permitted_network"
description: |-
  Manages IBM Private DNS Permitted Network.
---

# ibm_dns_permitted_network

Create or delete a DNS permitted network. For more information, see [Managing permitted networks](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-managing-permitted-networks).

You can add a VPC as a permitted network to a DNS entry only. 


## Example usage

```terraform
resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
    instance_id = ibm_resource_instance.test-pdns-instance.guid
    zone_id = ibm_dns_zone.test-pdns-zone.zone_id
    vpc_crn = ibm_is_vpc.test_pdns_vpc.crn
    type = "vpc"
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The GUID of the IBM Cloud DNS service instance where you want to add a permitted network.
- `type` - (Required, String) The type of permitted network that you want to add. Supported values are `vpc`.
- `vpc_crn` - (Required, String) The CRN of the VPC that you want to add as a permitted network.
- `zone_id` - (Required, String) The ID of the private DNS zone where you want to add the permitted network.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time when the permitted network was added to the DNS.
- `id` - (String) The unique identifier of the DNS private network. The ID is composed of `<instance_ID>/<zone_ID>/<permitted_network_ID>`.
- `modified_on` - (Timestamp) The time when the permitted network was modified.


## Import

The  `ibm_dns_permitted_network` can be imported by using private DNS instance ID, zone ID and permitted network ID.

**Example**

```
$ terraform import ibm_dns_permitted_network.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
