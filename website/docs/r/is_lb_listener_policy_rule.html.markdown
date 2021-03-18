---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : lb_listener_policy_rule"
description: |-
  Manages IBM VPC load balancer listener policy rule.
---

# ibm\_is_lb_listener_policy

Provides a load balancer listener policy rule resource. This allows load balancer listener policy rule to be created, updated, and cancelled.

## Example Usage

In the following example, you can create a load balancer listener policy rule, along with lb and lb listener:

```hcl
resource "ibm_is_lb" "lb2"{
  name    = "mylb"
  subnets = ["35860fed-c911-4936-8c94-f0d8577dbe5b"]
}

resource "ibm_is_lb_listener" "lb_listener2"{
  lb       = ibm_is_lb.lb2.id
  port     = "9086"
  protocol = "http"
}
resource "ibm_is_lb_listener_policy" "lb_listener_policy" {
  lb = ibm_is_lb.lb2.id
  listener = ibm_is_lb_listener.lb_listener2.listener_id
  action = "redirect"
  priority = 2
  name = "mylistener8"
  target_http_status_code = 302
  target_url = "https://www.redirect.com"
  rules{
      condition = "contains"
      type = "header"
      field = "1"
      value = "2"
  }
}

resource "ibm_is_lb_listener_policy_rule" "lb_listener_policy_rule" {
  lb        = ibm_is_lb.lb2.id
  listener  = ibm_is_lb_listener.lb_listener2.listener_id
  policy    = ibm_is_lb_listener_policy.lb_listener_policy.policy_id
  condition = "equals"
  type      = "header"
  field     = "MY-APP-HEADER"
  value     = "New-value"
}
```
## Timeouts

ibm_is_lb_listener_policy rule provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 10 minutes) Used for deleting Instance.


## Argument Reference

The following arguments are supported:

* `lb` - (Required, Forces new resource, string) Unique Load Balancer ID
* `listener` - (Required, Forces new resource, string) Unique Load Balancer Listener ID
* `policy` - (Required, Forces new resource, string) Unique Load Balancer listener policy ID
* `condition` - (Required, string). The condition of the rule. Allowable values: [contains,equals,matches_regex].
* `type` - (Required, string) The type of the rule.Allowable values: [header,hostname,path].
* `value` - (Required, string) Value to be matched for rule condition. Constraints: 1 ≤ length ≤ 128
* `field` - (Optional, string) HTTP header field. This is only applicable to "header" rule type. 


`Note :`

Network load balancer does not support ibm_is_lb_listener_policy_rule.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the load balancer listener Policy. Its a combination of lb, listener and policyID delimited by /.
* `status` - The status of load balancer listener.
* `rule` - Rule ID

## Import

ibm_is_lb_listener_policy_rule can be imported using lbID, listenerID, policyID and ruleID eg

```
$ terraform import ibm_is_lb_listener_policy.example c1e3d5d3-8836-4328-b473-a90e0c9ba941/3ea13dc7-25b4-4c62-8cc7-0f7e092e7a8f/2161a3fb-123c-4a33-9a3d-b3154ef42009/356789523dc7-25b4-4c62-8cc7-0f7e092e7a8f
```
