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
		high_availability = false
		enabled 	= true
	}
	resource "ibm_dns_custom_resolver_location" "test1" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet1.crn
		enabled     = true
		cr_enabled	= true
	}
	resource "ibm_dns_custom_resolver_location" "test2" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet2.crn
		enabled     = true
		cr_enabled	= true
	}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) The GUID of the private DNS service instance.
* `resolver_id` - (Required, String) The unique identifier of a custom resolver.
* `subnet_crn` - (Required, String) The subnet CRN of the VPC.
* `enabled` - (Optional, Bool) The custom resolver location will enabled or disable.
* `cr_enabled` - (Optional, Bool) Indicates whether to enable or disable the customer resolver. Default is 'true'



## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `dns_server_ip` - (Computed, String) Custom resolver location server ip.
* `healthy` - (Computed, Bool) The Custom resolver location will enable.
* `id` - (String) The unique identifier of the IBM DNS custom resolver location.
* `location_id` - (Computed, String) Type of the custom resolver loaction ID.

