---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_cr_locations"
description: |-
  Manages IBM Private DNS Custom Resolver Locations.
---

# ibm_dns_cr_locations

Provides a private DNS Custom Resolver Locations resource. This allows DNS Custom Resolver Location to create, update, and delete. 


## Example usage

```terraform
resource "ibm_dns_cr_locations" "test" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  subnet_crn  = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-03d54d71-b438-4d20-b943-76d3d2a1a590"
  enabled     = false
}

```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The unique identifier of a service instance.
* `resolver_id` - (Required, string) The unique identifier of a custom resolver loactions.
* `location_id` - (Computed, string) Type of the custom resolver loaction id.
* `subnet_crn` - (Required, string) The subnet crn of the VPC.
* `enabled` - (Optional, Bool) The Custom resolver location will enable.
* `healthy` - (Computed, Bool) The Custom resolver location will enable.
* `dns_server_ip` - (Computed, string) Custom resolver location server ip.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - (String) The unique identifier of the ibm_dns_cr_locations.
* `created_on` - (String) the time when a forwarding rule is created, RFC3339 format.
* `modified_on` -(String) the recent time when a forwarding rule is modified, RFC3339 format.

## Import

You can import the `ibm_dns_cr_locations` resource by using `id`.
The `id` property can be formed from `instance_id`, `resolver_id`, and `location_id` in the following format:

```
<instance_id>/<resolver_id>/<rule_id>
```
* `instance_id`: A string. The unique identifier of a service instance.
* `resolver_id`: A string. The unique identifier of a custom resolver locations.
* `location_id`: A string. The unique identifier of a custom resolver locations.

```
$ terraform import ibm_dns_cr_locations <instance_id>/<resolver_id>/<location_id>
```