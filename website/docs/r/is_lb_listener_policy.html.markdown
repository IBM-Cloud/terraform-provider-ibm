---

subcategory: "VPC infrastructure"
page_title: "IBM : lb_listener_policy"
description: |-
  Manages IBM VPC load balancer listener policy.
---

# ibm_is_lb_listener_policy
Create, update, or delete a load balancer listener policy. For more information, about VPC load balance listener policy, see [monitoring application Load Balancer for VPC metrics](https://cloud.ibm.com/docs/vpc?topic=vpc-monitoring-metrics-alb).

## Example usage

### Sample to create a load balancer listener policy for a `redirect` action.

```terraform
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

###  Creating a load balancer listener policy for a `forward` action by using `lb` and `lb listener`.


```terraform
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
The `ibm_is_lb_listener_policy` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the load balancer listener policy is considered `failed` if no response is received for 10 minutes. 
- **delete**: The deletion of the load balancer listener policy is considered `failed` if no response is received for 10 minutes.
- **update**: The creation of the load balancer listener policy is considered `failed` if no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `action` - (Required, Forces new resource, String) The action that you want to specify for your policy. Supported values are `forward`, `redirect`, and `reject`.
- `lb` - (Required, Forces new resource, String) The ID of the load balancer for which you want to create a load balancer listener policy. 
- `listener` - (Required, Forces new resource, String) The ID of the load balancer listener.
- `name` - (Optional, String) The name for the load balancer policy. Names must be unique within a load balancer listener.
- `priority`- (Required, Integer) The priority of the load balancer policy. Low values indicate a high priority. The value must be between 1 and 10.Yes.
- `rules`- (Required, List) A list of rules that you want to apply to your load balancer policy. Note that rules can be created only. You cannot update the rules for a load balancer policy.

  Nested scheme for `rules`:
  - `condition` - (Required, String) The condition that you want to apply to your rule. Supported values are `contains`, `equals`, and `matches_regex`.
  - `type` - (Required, String) The data type where you want to apply the rule condition. Supported values are `header`, `hostname`,  and `path`.
  - `value`- (Required, Integer) The value that must be found in the HTTP header, hostname or path to apply the load balancer listener rule. The value that you define can be between 1 and 128 characters long.
  - `field`- (Required, Integer) If you selected `header` as the data type where you want to apply the rule condition, enter the name of the HTTP header that you want to check. The name of the header can be between 1 and 128 characters long.
- `target_id` - (Optional, Integer) When `action` is set to **forward**, specify the ID of the load balancer pool that the load balancer forwards network traffic to.
- `target_http_status_code` - (Optional, Integer) When `action` is set to **redirect**, specify the HTTP response code that must be returned in the redirect response. Supported values are `301`, `302`, `303`, `307`, and `308`. 
- `target_url` - (Optional, Integer) When `action` is set to **redirect**, specify the URL that is used in the redirect response.

**Note**

When action is `forward`, `target_id` should specify which pool the load balancer forwards the traffic to.
When action is `redirect`, `target_url` should specify the `url` and `target_http_status_code` to specify the code used in the redirect response.
Network load balancer does not support `ibm_is_lb_listener_policy`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the load balancer listener policy. The ID is composed of `<lb_ID>/<listener_ID>/<policy_ID>`.
- `policy_id` - (String) The ID of the load balancer listener policy.
- `status` - (String) The status of the load balancer listener policy.


## Import
The resource can be imported by using the ID. The ID is composed of `<lb_ID>/<listener_ID>/<policy_ID>`.

**Synatx**

```
$ terraform import ibm_is_lb_listener_policy.example <lb_ID>/<listener_ID>/<policy_ID>
```

**Example**

```
$ terraform import ibm_is_lb_listener_policy.example c1e3d5d3-8836-4328-b473-a90e0c9ba941/3ea13dc7-25b4-4c62-8cc7-0f7e092e7a8f/2161a3fb-123c-4a33-9a3d-b3154ef42009
```
