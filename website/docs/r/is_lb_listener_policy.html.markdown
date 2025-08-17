---

subcategory: "VPC infrastructure"
page_title: "IBM : lb_listener_policy"
description: |-
  Manages IBM VPC load balancer listener policy.
---

# ibm_is_lb_listener_policy
Create, update, or delete a load balancer listener policy. For more information, about VPC load balance listener policy, see [monitoring application Load Balancer for VPC metrics](https://cloud.ibm.com/docs/vpc?topic=vpc-monitoring-metrics-alb).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Sample to create a load balancer listener policy for a `redirect` action.

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
  lb       = ibm_is_lb.example.id
  listener = ibm_is_lb_listener.example.listener_id
  action   = "redirect"
  priority = 4
  name     = "example-listener-policy"
  target {
    http_status_code = 302
    url              = "https://www.example.com"
  }
}
```

### Sample to create a load balancer listener policy for a `redirect` action with parameterized url.

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
  lb       = ibm_is_lb.example.id
  listener = ibm_is_lb_listener.example.listener_id
  action   = "redirect"
  priority = 4
  name     = "example-listener-policy"
  target {
    http_status_code = 302
    url              = "https://{host}:8080/{port}/{host}/{path}"
  }
}
```

### Sample to create a load balancer listener policy for a `https_redirect` action.

```terraform
resource "ibm_is_lb" "example" {
  name    = "example-lb"
  subnets = [ibm_is_subnet.example.id]
}

resource "ibm_is_lb_listener" "example_http_source" {
  lb       = ibm_is_lb.example.id
  port     = "9080"
  protocol = "http"
}

resource "ibm_is_lb_listener" "example_https_target" {
  lb                   = ibm_is_lb.example.id
  port                 = "9086"
  protocol             = "https"
  certificate_instance = "crn:v1:staging:public:cloudcerts:us-south:a2d1bace7b46e4815a81e52c6ffeba5cf:af925157-b125-4db2-b642-adacb8b9c7f5:certificate:c81627a1bf6f766379cc4b98fd2a44ed"
}

resource "ibm_is_lb_listener_policy" "example" {
  lb                                = ibm_is_lb.example.id
  listener                          = ibm_is_lb_listener.example_http_source.listener_id
  action                            = "https_redirect"
  priority                          = 2
  name                              = "example-listener"
  target {
    http_status_code = 302
    listener {
      id = ibm_is_lb_listener.example_https_target.listener_id
    }
    uri = "/example?doc=get"
  }
  rules {
    condition = "contains"
    type      = "header"
    field     = "1"
    value     = "2"
  }
}
```

###  Creating a load balancer listener policy for a `forward_to_pool` action.


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

resource "ibm_is_lb_pool" "example" {
  name           = "example-lb-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "http"
}

resource "ibm_is_lb_listener_policy" "example" {
  lb       = ibm_is_lb.example.id
  listener = ibm_is_lb_listener.example.listener_id
  action   = "forward_to_pool"
  priority = 3
  name     = "example-listener"
  target {
    id = ibm_is_lb_pool.example.pool_id
  }
}
```

###  Creating a load balancer listener policy for a `forward_to_listener` action.


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

resource "ibm_is_lb_listener" "example1" {
  lb       = ibm_is_lb.example.id
  port     = "9087"
  protocol = "tcp"
}

resource "ibm_is_lb_listener_policy" "example" {
  lb       = ibm_is_lb.example.id
  listener = ibm_is_lb_listener.example.listener_id
  action   = "forward_to_listener"
  priority = 2
  name     = "example-listener"
  target {
    id = ibm_is_lb_listener.example1.listener_id
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

- `action` - (Required, Forces new resource, String) The action that you want to specify for your policy. Supported values are `forward_to_pool`,`forward_to_listener`, `redirect`, `reject`, and `https_redirect`.
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
- `target_id` - (Optional, Integer) When `action` is set to **forward_to_pool**, specify the ID of the load balancer pool that the load balancer forwards network traffic to. or When `action` is set to **forward_to_listener**, specify the ID of the load balancer listener that the load balancer forwards network traffic to.
- `target_http_status_code` - (Optional, Integer) When `action` is set to **redirect**, specify the HTTP response code that must be returned in the redirect response. Supported values are `301`, `302`, `303`, `307`, and `308`. 
- `target_url` - (Optional, Integer) When `action` is set to **redirect**, specify the URL that is used in the redirect response.
- `target_https_redirect_listener` - (Optional, String) When `action` is set to **https_redirect**, specify the ID of the listener that will be set as http redirect target.
- `target_https_redirect_status_code` - (Optional, Integer) When `action` is set to **https_redirect**, specify the HTTP status code to be returned in the redirect response. Supported values are `301`, `302`, `303`, `307`, `308`.
- `target_https_redirect_uri` - (Optional, String) When `action` is set to **https_redirect**, specify the target URI where traffic will be redirected.

~> **Note:** `target_id`, `target_http_status_code`, `target_url`, `target_https_redirect_listener`, `target_https_redirect_status_code`, `target_https_redirect_uri` are deprecated and will be removed soon. Please use `target` instead.

- `target` - (Optional, List) - If `action` is `forward_to_pool`, the response is a `LoadBalancerPoolReference`-If `action` is `forward_to_listener`, specify a `LoadBalancerListenerIdentity` in this load balancer to forward to.- If `action` is `redirect`, the response is a `LoadBalancerListenerPolicyRedirectURL`- If `action` is `https_redirect`, the response is a `LoadBalancerListenerHTTPSRedirect`.
    Nested schema for **target**:
	- `deleted` - (Computed, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (Computed, String) Link to documentation about deleted resources.
	- `href` - (Optional, String) The pool's canonical URL.
	- `http_status_code` - (Optional, Integer) The HTTP status code for this redirect. Allowable values are: `301`, `302`, `303`, `307`, `308`.
	- `id` - (Optional, String) The unique identifier for this load balancer pool.
	- `listener` - (Optional, List)
	  Nested schema for **listener**:
		- `deleted` - (Computed, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (Computed, String) Link to documentation about deleted resources.
		- `href` - (Optional, String) The listener's canonical URL.
		- `id` - (Optional, String) The unique identifier for this load balancer listener.
	- `name` - (Computed, String) The name for this load balancer pool. The name is unique across all pools for the load balancer.
	- `url` - (Optional, String) The redirect target URL. The URL supports [RFC 6570 level 1 expressions](https://datatracker.ietf.org/doc/html/rfc6570#section-1.2) for the following variables which expand to values from the originally requested URL (or the indicated defaults if the request did not include them):

      **&#x2022;** protocol </br>
      **&#x2022;** host </br>
      **&#x2022;** port (default: 80 for HTTP requests, 443 for HTTPS requests) </br>
      **&#x2022;** path (default: '/') </br>
      **&#x2022;** query (default: '') </br>

	  
~> **Note:** When `action` is set to **forward_to_pool**, specify the ID of the load balancer pool that the load balancer forwards network traffic to. or When `action` is set to **forward_to_listener**, specify the ID of the load balancer listener that the load balancer forwards network traffic to.
When action is `redirect`, `target.url` should specify the `url` and `target.http_status_code` to specify the code used in the redirect response.
When action is `https_redirect`, `target.listener.id` should specify the ID of the listener, `target.http_status_code` to specify the code used in the redirect response and `target.uri` to specify the target URI where traffic will be redirected.
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
