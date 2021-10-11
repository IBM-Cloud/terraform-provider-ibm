---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_custom_resolver_location"
description: |-
  Manages IBM Private DNS custom resolver locations.
---

# ibm_dns_custom_resolver_location

Provides a private DNS custom resolver locations resource. This allows DNS custom resolver location to create, update, and delete. For more information, about custom resolver locations, see [add-custom-resolver-location](https://cloud.ibm.com/apidocs/dns-svcs#add-custom-resolver-location).


## Example usage

```terraform
data "ibm_resource_group" "rg" {
  name = "default"
}

resource "ibm_resource_instance" "test-pdns-instance" {
  name              = "test-pdns"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

resource "ibm_dns_custom_resolver" "test" {
  name        = "testCR-TF"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description = "testdescription-CR"
}

resource "ibm_dns_custom_resolver_location" "loc1" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  subnet_crn  = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
  enabled     = true
}

resource "ibm_dns_custom_resolver_location" "loc2" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  subnet_crn  = "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-202fe4570c83"
  enabled     = true
}

```

## Argument reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) The GUID of the private DNS service instance.
* `resolver_id` - (Required, String) The unique identifier of a custom resolver.
* `subnet_crn` - (Required, String) The subnet CRN of the VPC.
* `enabled` - (Optional, Bool) The custom resolver location will enable.


## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `dns_server_ip` - (Computed, String) Custom resolver location server ip.
* `healthy` - (Computed, Bool) The Custom resolver location will enable.
* `id` - (String) The unique identifier of the IBM DNS custom resolver location.
* `location_id` - (Computed, String) Type of the custom resolver loaction ID.

