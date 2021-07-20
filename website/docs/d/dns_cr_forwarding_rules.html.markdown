---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Forwarding Rules"
description: |-
  Manages IBM Cloud Infrastructure private domain name service Forwording Rules.
---

# ibm_dns_cr_forwarding_rules

Provides a read-only data source for Forwarding rules. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
data "dns_cr_forwarding_rules" "dns_cr_forwarding_rules" {
	instance_id = "instance_id"
	resolver_id = "resolver_id"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The unique identifier of a service instance.
* `resolver_id` - (Required, string) The unique identifier of a custom resolver.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id`- (String) The unique identifier of the forwarding rules.
* `forwarding_rules` (List of Forwarding rules) Nested `forwarding_rules` blocks have the following structure:
	* `id` - (String) Identifier of the forwarding rule.
	* `description` - (String) Descriptive text of the forwarding rule.
	* `type` - (String) Type of the forwarding rule.
	* `match` - (String) The matching zone or hostname.
	* `forward_to` - (String) The upstream DNS servers will be forwarded to.
	* `created_on` - (String) the time when a forwarding rule is created, RFC3339 format.
	* `modified_on` - (String) the recent time when a forwarding rule is modified, RFC3339 format.

