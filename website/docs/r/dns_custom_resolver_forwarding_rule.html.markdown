---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Forwarding Rule"
description: |-
  Manages Forwarding Rule.
---

# ibm_dns_custom_resolver_forwarding_rule

Provides a resource for dns_custom_resolver_forwarding_rule. This allows Forwarding Rule to be created, updated and deleted.For more information, about Forwarding Rules, see [create-forwarding-rule](https://cloud.ibm.com/apidocs/dns-svcs#create-forwarding-rule).

## Example Usage

```terraform
resource "dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  description = "forwarding rule"
  type = "zone"
  match = "example.com"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, string) The GUID of the private DNS service instance.
* `resolver_id` - (Required, string) The unique identifier of a custom resolver.
* `description` - (Optional, string) Descriptive text of the forwarding rule.
* `type` - (Optional, string) Type of the forwarding rule.
  * Constraints: Allowable values are: zone, hostname
* `match` - (Optional, string) The matching zone or hostname.
* `forward_to` - (Optional, List) The upstream DNS servers will be forwarded to.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `id` - (String) The unique identifier of the dns_custom_resolver_forwarding_rule.
* `created_on` - (String) the time when a forwarding rule is created, RFC3339 format.
* `modified_on` -(String) the recent time when a forwarding rule is modified, RFC3339 format.

## Import

You can import the `dns_custom_resolver_forwarding_rule` resource by using `id`.
The `id` property can be formed from `instance_id`, `resolver_id`, and `rule_id` in the following format:

```
<instance_id>/<resolver_id>/<rule_id>
```
* `instance_id`: A string. The GUID of the private DNS service instance.
* `resolver_id`: A string. The unique identifier of a custom resolver.
* `rule_id`: A string. The unique identifier of a forwarding rule.

```
$ terraform import dns_custom_resolver_forwarding_rule.dns_custom_resolver_forwarding_rule <instance_id>/<resolver_id>/<rule_id>
```
