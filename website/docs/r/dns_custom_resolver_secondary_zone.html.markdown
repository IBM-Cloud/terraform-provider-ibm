---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_custom_resolver_secondary_zone"
description: |-
  Manages IBM Private DNS custom resolver secondary zone.
---

# ibm_dns_custom_resolver_secondary_zone

The DNS custom resolver secondary zone resource allows users to request and manage secondary zones for a given custom resolver. By creating and enabling a secondary zone resource for a custom resolver, DNS records for a given DNS zone will be transferred from a user provided primary DNS server (on premise) to a private DNS custom resolver hosted on an IBM Cloud VPC. This framework will improve the availability, speed, and security of DNS queries for a given DNS zone.


## Example usage

```
data "ibm_resource_group" "rg" {
  is_default = true
}

# create a VPC for the subnets
resource "ibm_is_vpc" "test-pdns-cr-vpc" {
  depends_on     = [data.ibm_resource_group.rg]
  name           = "moises-sdk-testing-cross-account-vpc"
  resource_group = data.ibm_resource_group.rg.id
}

# create subnets for the custom resolver locations
resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
  name            = "moises-sdk-testing-cross-account-subnet1"
  vpc             = ibm_is_vpc.test-pdns-cr-vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
  resource_group  = data.ibm_resource_group.rg.id
}

resource "ibm_is_subnet" "test-pdns-cr-subnet2" {
  name            = "moises-sdk-testing-cross-account-subnet2"
  vpc             = ibm_is_vpc.test-pdns-cr-vpc.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/24"
  resource_group  = data.ibm_resource_group.rg.id
}

# create a DNS instance
resource "ibm_resource_instance" "test-pdns-cr-instance" {
  name              = "moises-sdk-testing-cross-account-dns1"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

# create a custom resolver
resource "ibm_dns_custom_resolver" "test" {
  name        = "msz-test-cr2"
  instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
  description = "new test CR - TF"
  enabled     = true
  locations {
    subnet_crn = ibm_is_subnet.test-pdns-cr-subnet1.crn
    enabled    = true
  }
  locations {
    subnet_crn = ibm_is_subnet.test-pdns-cr-subnet2.crn
    enabled    = true
  }
}

resource "ibm_dns_zone" "pdns-1-zone" {
  name        = "moises-zone3.com"
  instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
  description = "testdescription"
  label       = "testlabel"
}

resource "ibm_dns_custom_resolver_secondary_zone" "test" {
  instance_id   = ibm_resource_instance.test-pdns-cr-instance.guid
  resolver_id   = ibm_dns_custom_resolver.test.custom_resolver_id
  zone          = "moises-zone4.com"
  enabled       = true
  transfer_from = ["10.0.0.8"]
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.
- `resolver_id` - (Required, String) The GUID of the custom resolver.
- `zone` - (Required, String) DNS records associated with this DNS zone will be transferred to the custom resolver.
- `enabled`- (Required, Bool) To enable or disable a secondary zone transfer rule. 
- `transfer_from`- (Required, List of Strings) List of IP addresses. DNS records will be transferred from the primary DNS servers associated with the IP addresses in this list to custom resolvers hosted in an IBM Cloud VPC.
- `description` - (Optional, String) Descriptive text of the secondary zone.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time (created On) of the Secondary Zone. 
- `modified_on` - (Timestamp) The time (modified On) of the Secondary Zone.
- `secondary_zone_id` - (String) The unique ID of the private DNS custom resolver secondary zone.

## Import
The `ibm_dns_custom_resolver_secondary_zone` can be imported by using private DNS instance ID, Custom Resolver ID, and Secondary Zone ID.
The `id` property can be formed from `instance_id`, `custom_resolver_id` and `secondary_zone_id` in the following format:

```
<instance_id>/<custom_resolver_id>/<secondary_zone_id>
```

**Example**

```
terraform import ibm_dns_custom_resolver_secondary_zone.sample "d10e6956-377a-43fb-a5a6-54763a6b1dc2/63481bef-3759-4b5e-99df-73be7ba40a8a/bd2d4867-f606-45da-93b4-02dc69635d5e"
```
