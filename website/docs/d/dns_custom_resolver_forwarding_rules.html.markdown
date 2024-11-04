---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : Forwarding Rules"
description: |-
  Manages IBM Cloud Infrastructure private domain name service forwarding rules.
---

# ibm_dns_custom_resolver_forwarding_rules

Provides a read-only data source for forwarding rules. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information about forwarding rules, refer to [list-forwarding-rules](https://cloud.ibm.com/apidocs/dns-svcs#list-forwarding-rules)

## Example usage

```terraform
data "ibm_dns_custom_resolver_forwarding_rules" "test-fr" {
  instance_id = ibm_dns_custom_resolver.test.instance_id
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `instance_id` - (Required, String) The GUID of the private DNS service instance.
- `resolver_id` - (Required, String) The unique identifier of a custom resolver.

## Attribute reference

In addition to the argument references list, you can access the following attribute references after your data source are created.

- `forwarding_rules` (List) List of forwarding rules.

  Nested scheme for `forwarding_rules`:
  - `description` - (String) Descriptive text of the forwarding rule.
  - `forward_to` - (List) List of the upstream DNS servers that the matching DNS queries will be forwarded to.
  - `match` - (String) The matching zone or hostname.
  - `rule_id` - (String) Identifier of the forwarding rule.
  - `type` - (String) Type of the forwarding rule.
  - `views` (List) List of views attached to the custom resolver.

    Nested scheme for `views`:
    - `name` - (String) Name of the view.
    - `description` - (String) Description of the view.
    - `expression` - (String) Expression of the view.
    - `forward_to` - (List) List of upstream DNS servers that the matching DNS queries will be forwarded to.
