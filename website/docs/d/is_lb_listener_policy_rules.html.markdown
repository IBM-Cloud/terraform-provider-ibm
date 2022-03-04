---
subcategory: "VPC infrastructure"
page_title: "IBM : ibm_is_lb_listener_policy_rules"
description: |-
  Get information about LoadBalancerListenerPolicyRuleCollection
---

# ibm_is_lb_listener_policy_rules

Provides a read-only data source for LoadBalancerListenerPolicyRuleCollection. For more information, about load balancer listener policy and rules, see [layer 7 load balancing policies and rules](https://cloud.ibm.com/docs/vpc?topic=vpc-layer-7-load-balancing).
## Example Usage

```terraform
data "ibm_is_lb_listener_policy_rules" "example" {
  listener = ibm_is_lb_listener.example.listener_id
  lb       = ibm_is_lb.example.id
  policy   = ibm_is_lb_listener_policy.example.policy_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `listener` - (Required, String) The listener identifier.
- `lb` - (Required, String) The load balancer identifier.
- `policy` - (Required, String) The policy identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the LoadBalancerListenerPolicyRuleCollection.
- `rules` - (List) Collection of rules.
Nested scheme for `rules`:
	- `condition` - (String) The condition of the rule.
	- `created_at` - (String) The date and time that this rule was created.
	- `field` - (Optional, String) The field. This is applicable to `header`, `query`, and `body` rule types.If the rule type is `header`, this property is required.If the rule type is `query`, this is optional. If specified and the rule condition is not`matches_regex`, the value must be percent-encoded.If the rule type is `body`, this is optional.
	- `href` - (String) The rule's canonical URL.
	- `id` - (String) The rule's unique identifier.
	- `provisioning_status` - (String) The provisioning status of this rule.
	- `type` - (String) The type of the rule.Body rules are applied to form-encoded request bodies using the `UTF-8` character set.
	- `value` - (String) Value to be matched for rule condition.If the rule type is `query` and the rule condition is not `matches_regex`, the value must be percent-encoded.