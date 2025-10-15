---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : lb_listener_policy_rule"
description: |-
  Manages IBM VPC load balancer listener policy rule.
---

# ibm_is_lb_listener_policy_rule
Create, update, or delete a VPC load balancer listener policy rule. For more information, about load balancer listener policy and rules, see [layer 7 load balancing policies and rules](https://cloud.ibm.com/docs/vpc?topic=vpc-layer-7-load-balancing).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
Sample to create a load balancer listener policy rule, along with `lb` and `lb listener`.

```terraform
resource "ibm_is_lb" "example" {
  name    = "example-lb"
  subnets = [ibm_is_subnet.example.id]
}

resource "ibm_is_lb_listener" "example" {
  lb       = ibm_is_lb.example.id
  port     = "9086"
  protocol = "http"
}
resource "ibm_is_lb_listener_policy" "example" {
  lb                      = ibm_is_lb.example.id
  listener                = ibm_is_lb_listener.example.listener_id
  action                  = "redirect"
  priority                = 2
  name                    = "example-listener"
  target_http_status_code = 302
  target_url              = "https://www.redirect.com"
  rules {
    condition = "contains"
    type      = "header"
    field     = "1"
    value     = "2"
  }
}

resource "ibm_is_lb_listener_policy_rule" "example" {
  lb        = ibm_is_lb.example.id
  listener  = ibm_is_lb_listener.example.listener_id
  policy    = ibm_is_lb_listener_policy.example.policy_id
  condition = "equals"
  type      = "header"
  field     = "MY-APP-HEADER"
  value     = "New-value"
}
```

### Create load load balancer listener policy rule , with sni_hostname .

```terraform
resource "ibm_is_lb" "example" {
  name    = "example-lb"
  subnets = [ibm_is_subnet.example.id]
}

resource "ibm_is_lb_listener" "example" {
  lb       = ibm_is_lb.example.id
  port     = "9086"
  protocol = "http"
}
resource "ibm_is_lb_listener_policy" "example" {
  lb                      = ibm_is_lb.example.id
  listener                = ibm_is_lb_listener.example.listener_id
  action                  = "redirect"
  priority                = 2
  name                    = "example-listener"
  target_http_status_code = 302
  target_url              = "https://www.redirect.com"
  rules {
    condition = "contains"
    type      = "tcp"
    field     = "1"
    value     = "2"
  }
}

resource "ibm_is_lb_listener_policy_rule" "example" {
  lb        = ibm_is_lb.example.id
  listener  = ibm_is_lb_listener.example.listener_id
  policy    = ibm_is_lb_listener_policy.example.policy_id
  condition = "equals"
  type      = "sni_hostname"
  field     = "MY-APP-HEADER"
  value     = "New-value"
}
```



## Timeouts
The `ibm_is_lb_listener_policy_rule` rule provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create**: The creation of the resource is considered failed if no response is received for 10 minutes. 
- **Update**: The update of the resource is considered failed if no response is received for 10 minutes. 
- **Delete**: The deletion of the resource is considered failed if no response is received for 10 minutes. 

## Argument reference
Review the argument references that you can specify for your resource. 

- `condition` - (Required, String) The condition that you want to apply to your rule. Supported values are `contains`, `equals`, and `matches_regex`.
- `field` - (Optional, String) If you set `type` to `header`, enter the HTTP header field where you want to apply the rule condition.
- `lb` - (Required, Forces new resource, String) The ID of the load balancer for which you want to create a listener policy rule.
- `listener` - (Required, Forces new resource, String) The ID of the load balancer listener for which you want to create a policy rule. 
- `policy` - (Required, Forces new resource, String) The ID of the load balancer listener policy for which you want to create a policy rule. 
- `type` - (Required, String) The object where you want to apply the rule. Supported values are `body`, `header`, `hostname`, `path`, `query`, and `sni_hostname`.
  The content the rule applies to:
    - `body`: The UTF-8 form-encoded HTTP request body
    - `header`: The HTTP header
    - `hostname`: The fully-qualified domain name of the server specified in the Host HTTP request header
    - `path`: The path of the HTTP request
    - `query`: The query of the HTTP request URL
    - `sni_hostname`: The fully-qualified domain name of the server provided in the "server name indicator" extension during TLS negotiation
    - For listeners with protocol http or https, any type may be specified.
    - For listeners with protocol tcp, only type `sni_hostname` may be specified.

- `value` - (Required, String) The value that must match the rule condition. The value can be between 1 and 128 characters long. No.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the load balancer listener policy rule. The ID is composed of ` <loadbalancer_ID>/<listener_ID>/<policy_ID>/<rule_ID>`.
- `rule` - (String) The ID of the rule.
- `status` - (String) The status of the load balancer listener policy rule.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the **listener policy rule** using `lb_ID`, `listener_ID`, `policy_ID`, and `rule_ID`.

**Example**

```hcl
import {
  to = ibm_is_lb_listener_policy_rule.example
  id = "c1e3d5d3-8836-4328-b473-a90e0c9ba941/3ea13dc7-25b4-4c62-8cc7-0f7e092e7a8f/2161a3fb-123c-4a33-9a3d-b3154ef42009/356789523dc7-25b4-4c62-8cc7-0f7e092e7a8f"
}
```

For versions prior to Terraform v1.5.0, use the `terraform import` command:

**Syntax**

```console
$ terraform import ibm_is_lb_listener_policy_rule.example <loadbalancer_ID>/<listener_ID>/<policy_ID>/<rule_ID>
```

**Example**

```console
$ terraform import ibm_is_lb_listener_policy_rule.example c1e3d5d3-8836-4328-b473-a90e0c9ba941/3ea13dc7-25b4-4c62-8cc7-0f7e092e7a8f/2161a3fb-123c-4a33-9a3d-b3154ef42009/356789523dc7-25b4-4c62-8cc7-0f7e092e7a8f
```
