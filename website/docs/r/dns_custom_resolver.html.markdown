---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_custom_resolver"
description: |-
  Manages IBM Private DNS custom resolver.
---

# ibm_dns_custom_resolver

Provides a private DNS custom resolver resource. This allows DNS custom resolver to create, update, and delete. For more information, about customer resolver, see [create-custom-resolver](https://cloud.ibm.com/apidocs/dns-svcs#create-custom-resolver).


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
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.
- `name`- (Required, String) The name of the custom resolver.
- `description` - (Optional, String) Descriptive text of the custom resolver.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time (created On) of the DNS Custom Resolver. 
- `custom_resolver_id` - (String) The unique ID of the private DNS custom resolver.
- `modified_on` - (Timestamp) The time (modified On) of the DNS Custom Resolver.
- `health`- (String) The status of DNS Custom Resolver's health. Possible values are `DEGRADED`, `CRITICAL`, `HEALTHY`.
- `locations` - (Set) Locations on which the custom resolver will be running.

  Nested scheme for `locations`:
  - `healthy`- (String) The health status.
  - `dns_server_ip`- (String) The DNS server IP.
  - `enabled`- (String) Whether the location is enabled.
  - `location_id`- (String) The location ID.

## Import
The `ibm_dns_custom_resolver` can be imported by using private DNS instance ID, Custom Resolver ID.
The `id` property can be formed from `custom resolver id` and `instance_id` in the following format:
<custom_resolver_id>:<instance_id>

**Example**

```
$ terraform import ibm_dns_custom_resolver.example 270edfad-8574-4ce0-86bf-5c158d3e38fe:345ca2c4-83bf-4c04-bb09-5d8ec4d425a8
```
