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
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "test-pdns-custom-resolver-vpc"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "test-pdns-cr-subnet1"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "us-south-1"
		ipv4_cidr_block	= "10.240.0.0/24"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet2" {
		name			= "test-pdns-cr-subnet2"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "us-south-1"
		ipv4_cidr_block	= "10.240.64.0/24"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name		= "test-customresolver"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "new test CR - TF"
		enabled 	= true
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled		= true
		}
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet2.crn
			enabled     = true
		}
	}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `instance_id` - (Required, String) The GUID of the private DNS service instance.
- `name`- (Required, String) The name of the custom resolver.
- `enabled`- (Optional, Bool) To make custom resolver enabled/disable.
- `description` - (Optional, String) Descriptive text of the custom resolver.
- `high_availability` - (Optional, Bool) High Availability is enabled by Default, Need to add two or more locations.
- `locations`- (Optional, Set) The list of locations where this custom resolver is deployed. There is no update for location argument in resolver resource.

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
  - `enabled`- (Bool) Whether the location is enabled.
  - `location_id`- (String) The location ID.

 Nested scheme for `rules`:
 - `rule_id` - (String) The rule ID is unique identifier of the custom resolver forwarding rule.
 - `description`- (String) Descriptive text of the forwarding rule.
 - `type` - (String) Type of the forwarding rule.Constraints: Allowable values are: `zone`, `hostname`.
 - `match` - (String) The matching zone or hostname.
 - `forward_to` - (List) The upstream DNS servers will be forwarded to.

## Import
The `ibm_dns_custom_resolver` can be imported by using private DNS instance ID, Custom Resolver ID.
The `id` property can be formed from `custom_resolver_id` and `instance_id` in the following format:

```
<custom_resolver_id>:<instance_id>
```

**Example**

```
$ terraform import ibm_dns_custom_resolver.example 270edfad-8574-4ce0-86bf-5c158d3e38fe:345ca2c4-83bf-4c04-bb09-5d8ec4d425a8
```
