---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_custom_resolver_location"
description: |-
  Manages IBM Private DNS Custom Resolver Locations.
---

# ibm_dns_custom_resolver_location

Provides a private DNS Custom Resolver Locations resource. This allows DNS Custom Resolver Location to create, update, and delete. For more information, about Customer Resolver Locations, see [add-custom-resolver-location](https://cloud.ibm.com/apidocs/dns-svcs#add-custom-resolver-location).


## Example usage

```terraform
resource "ibm_dns_custom_resolver_location" "test" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  subnet_crn  = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-03d54d71-b438-4d20-b943-76d3d2a1a590"
  enabled     = false
}

```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, string) The GUID of the private DNS service instance.
* `resolver_id` - (Required, string) The unique identifier of a custom resolver.
* `subnet_crn` - (Required, string) The subnet crn of the VPC.
* `enabled` - (Optional, Bool) The Custom resolver location will enable.


## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `dns_server_ip` - (Computed, string) Custom resolver location server ip.
* `healthy` - (Computed, Bool) The Custom resolver location will enable.
* `id` - (String) The unique identifier of the ibm_dns_custom_resolver_location.
* `location_id` - (Computed, string) Type of the custom resolver loaction id.

