---
layout: "ibm"
page_title: "IBM : lb_listener_policy"
sidebar_current: "docs-ibm-resource-is-lb-listener_policy"
description: |-
  Manages IBM VPC load balancer listener policy.
---

# ibm\_is_lb_listener_policy

Provides a load balancer listener policy resource. This allows load balancer listener policy to be created, updated, and cancelled.

## Example Usage

In the following example, you can create a load balancer listener policy when action is redirect, along with lb and lb listener:

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
```
In the following example, you can create a load balancer listener policy when action is forward, along with lb and lb listener:

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
  action = "forward"
  priority = 2
  name = "mylistener8"
  target_id = "r006-beafdff0-4fe0-4db4-8f0c-b0b4ad828712"
  rules{
      condition = "contains"
      type = "header"
      field = "1"
      value = "2"
  }
}
```
## Timeouts

ibm_is_lb_listener_policy provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 10 minutes) Used for deleting Instance.


## Argument Reference

The following arguments are supported:

* `lb` - (Required, Forces new resource, string) Unique Load Balancer ID
* `listener` - (Required, Forces new resource, string) Unique Load Balancer Listener ID
* `action` - (Required, Forces new resource, string) The policy action. Allowable values: [forward,redirect,reject] 
* `priority` - (Required, Forces new resource, integer). Priority of the policy. Lower value indicates higher priority.
* `name` - (Optional, string) The user-defined name for this policy. Names must be unique within the load balancer listener the policy resides in.
* Nested `rules` block have the following structure:
	*	`condition` : Allowable values: [contains,equals,matches_regex]
	*	`type` : Allowable values: [header,hostname,path]
	*	`value` : Constraints: 1 ≤ length ≤ 128
	*	`filed`: Constraints: 1 ≤ length ≤ 128

    `Note` - As of now we rules can’t be updated, will be supported in upcoming release as a new resource

* `target_id` - (Optional, integer) The unique identifier for this load balancer pool, specified with 'forward' action.
* `target_http_status_code` -(Optional, integer) The http status code in the redirect response, one of [301, 302, 303, 307, 308], specified with 'redirect' action.
* `target_url` - (Optional, integer) The redirect target URL, specified with 'redirect' action


`Note :`

When action is forward, target_id should specify which pool the load balancer forwards the traffic to.
When action is redirect,target_url should specify the url and target_http_status_code to specify the code used in the redirect response.
Network load balancer does not support ibm_is_lb_listener_policy.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the load balancer listener Policy. Its a combination of lb, listener and policyID delimited by /.
* `status` - The status of load balancer listener.
* `policy_id` - Policy ID

## Import

ibm_is_lb_listener_policy can be imported using lbID, listenerID and policyID, eg

```
$ terraform import ibm_is_lb_listener_policy.example c1e3d5d3-8836-4328-b473-a90e0c9ba941/3ea13dc7-25b4-4c62-8cc7-0f7e092e7a8f/2161a3fb-123c-4a33-9a3d-b3154ef42009
```
