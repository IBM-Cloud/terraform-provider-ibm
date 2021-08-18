---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Forwarding Rule"
description: |-
  Manages Forwarding Rule.
---

# ibm_dns_custom_resolver_forwarding_rule

Provides a resource for ibm_dns_custom_resolver_forwarding_rule. This allows Forwarding Rule to be created, updated and deleted.For more information, about Forwarding Rules, see [create-forwarding-rule](https://cloud.ibm.com/apidocs/dns-svcs#create-forwarding-rule).

## Example Usage

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
  locations {
    subnet_crn  = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-6c3a997d-72b2-47f6-8788-6bd95e1bdb03"
    enabled     = true
  }
}

resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  description = "test forward rule"
  type = "zone"
  match = "test.example.com"
  forward_to = ["168.20.22.122"]
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

* `id` - (String) The unique identifier of the ibm_dns_custom_resolver_forwarding_rule.
* `created_on` - (String) the time when a forwarding rule is created, RFC3339 format.
* `modified_on` -(String) the recent time when a forwarding rule is modified, RFC3339 format.

## Import

You can import the `ibm_dns_custom_resolver_forwarding_rule` resource by using `id`.
The `id` property can be formed from `instance_id`, `resolver_id`, and `rule_id` in the following format:

```
<instance_id>/<resolver_id>/<rule_id>
```
* `instance_id`: A string. The GUID of the private DNS service instance.
* `resolver_id`: A string. The unique identifier of a custom resolver.
* `rule_id`: A string. The unique identifier of a forwarding rule.

```
$ terraform import ibm_dns_custom_resolver_forwarding_rule.ibm_dns_custom_resolver_forwarding_rule <instance_id>/<resolver_id>/<rule_id>
```
