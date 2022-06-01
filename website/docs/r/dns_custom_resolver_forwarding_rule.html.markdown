---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Forwarding Rule"
description: |-
  Manages forwarding rule.
---

# ibm_dns_custom_resolver_forwarding_rule

Provides a resource for ibm_dns_custom_resolver_forwarding_rule. This allows Forwarding Rule to be created, updated and deleted.For more information, about Forwarding Rules, see [create-forwarding-rule](https://cloud.ibm.com/apidocs/dns-svcs#create-forwarding-rule).

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
	resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
		instance_id		= ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id		= ibm_dns_custom_resolver.test.custom_resolver_id
		description		= "Test Fw Rule"
		type			= "zone"
		match			= "test.example.com"
		forward_to		= ["168.20.22.122"]
	}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) The GUID of the private DNS service instance.
* `resolver_id` - (Required, String) The unique identifier of a custom resolver.
* `description` - (Optional, String) Descriptive text of the forwarding rule.
* `type` - (Optional, String) Type of the forwarding rule.
  * Constraints: Allowable values are: `zone`, `hostname`,`Default`.
* `match` - (Optional, String) The matching zone or hostname.
* `forward_to` - (Optional, List) The upstream DNS servers will be forwarded to.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `id` - (String) The unique identifier of the DNS custom resolver forwarding rule.
* `created_on` - (String) The time when a forwarding rule is created, RFC3339 format.
* `modified_on` -(String) The recent time when a forwarding rule is modified, RFC3339 format.
* `rule_id` - (String) The rule ID is unique identifier of the custom resolver forwarding rule.

## Import

You can import the `ibm_dns_custom_resolver_forwarding_rule` resource by using `id`.
The `id` property can be formed from `rule_id`, `resolver_id`, and `instance_id` in the following format:

```
<rule_id>:<resolver_id>:<instance_id>
```
* `rule_id`: A String. The unique identifier of a forwarding rule.
* `resolver_id`: A String. The unique identifier of a custom resolver.
* `instance_id`: A String. The GUID of the private DNS service instance.

```
$ terraform import ibm_dns_custom_resolver_forwarding_rule.ibm_dns_custom_resolver_forwarding_rule <rule_id>:<resolver_id>:<instance_id>
```
