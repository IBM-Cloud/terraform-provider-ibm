---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Forwarding Rule"
description: |-
  Manages forwarding rule.
---

# ibm_dns_custom_resolver_forwarding_rule

Provides a resource for DNS custom resolver forwarding rule. This allows forwarding rule to be created, updated and deleted. For more information, about forwarding rules, see [create-forwarding-rule](https://cloud.ibm.com/apidocs/dns-svcs#create-forwarding-rule).

## Example usage

```terraform
resource "dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  description = "forwarding rule"
  type = "zone"
  match = "example.com"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) The GUID of the private DNS service instance.
* `resolver_id` - (Required, String) The unique identifier of a custom resolver.
* `description` - (Optional, String) Descriptive text of the forwarding rule.
* `type` - (Optional, String) Type of the forwarding rule.
  * Constraints: Allowable values are: `zone`, `hostname`.
* `match` - (Optional, String) The matching zone or hostname.
* `forward_to` - (Optional, List) The upstream DNS servers will be forwarded to.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `id` - (String) The unique identifier of the DNS custom resolver forwarding rule.
* `created_on` - (String) The time when a forwarding rule is created, RFC3339 format.
* `modified_on` -(String) The recent time when a forwarding rule is modified, RFC3339 format.

## Import

You can import the `dns_custom_resolver_forwarding_rule` resource by using `id`.
The `id` property can be formed from `instance_id`, `resolver_id`, and `rule_id` in the following format:

```
<instance_id>/<resolver_id>/<rule_id>
```
* `instance_id`: A String. The GUID of the private DNS service instance.
* `resolver_id`: A String. The unique identifier of a custom resolver.
* `rule_id`: A String. The unique identifier of a forwarding rule.

```
$ terraform import dns_custom_resolver_forwarding_rule.dns_custom_resolver_forwarding_rule <instance_id>/<resolver_id>/<rule_id>
```
