---

subcategory: "VPC infrastructure"
page_title: "IBM : lb_listener"
description: |-
  Manages IBM load balancer listener.
---

# ibm_is_lb_listener
Create, update, or delete a listener for a VPC load balancer. For more information, about load balancer listener, see [working with listeners](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-listeners).

**Note**
- When provisioning the load balancer listener along with load balancer pool or pool member, Use explicit depends on the resources or perform the terraform apply with parallelism. For more information, about explicit dependencies, see [create resource dependencies](https://learn.hashicorp.com/terraform/getting-started/dependencies#implicit-and-explicit-dependencies).
- VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

  **provider.tf**

  ```terraform
  provider "ibm" {
    region = "eu-gb"
  }
  ```
  
## Example usage
An example, to create a load balancer listener along with the pool and pool member.

```terraform

resource "ibm_is_lb_listener" "example" {
  lb                         = ibm_is_lb.example.id
  port                       = "9080"
  protocol                   = "http"
  https_redirect_listener    = ibm_is_lb_listener.example.listener_id
  https_redirect_status_code = 301
  https_redirect_uri         = "/example?doc=get"
}

resource "ibm_is_lb_pool" "example" {
  lb                 = ibm_is_lb.example.id
  name               = "example-lb-pool"
  protocol           = "http"
  algorithm          = "round_robin"
  health_delay       = "5"
  health_retries     = "2"
  health_timeout     = "2"
  health_type        = "http"
  health_monitor_url = "/"
  depends_on         = [ibm_is_lb_listener.example]
}

resource "ibm_is_lb_pool_member" "example" {
  count          = "2"
  lb             = ibm_is_lb.example.id
  pool           = element(split("/", ibm_is_lb_pool.example.id), 1)
  port           = "80"
  target_address = "192.168.0.1"
  depends_on     = [ibm_is_lb_listener.example]
}
```

### Sample to create a load balancer listener with `https_redirect`.

```terraform
resource "ibm_is_lb" "example2" {
  name    = "example-lb"
  subnets = [ibm_is_subnet.example.id]
}

resource "ibm_is_lb_listener" "example1" {
  lb                   = ibm_is_lb.example2.id
  port                 = "9086"
  protocol             = "https"
  certificate_instance = "crn:v1:bluemix:public:cloudcerts:us-south:a2d1bace7b46e4815a81e52c6ffeba5cf:af925157-b125-4db2-b642-adacb8b9c7f5:certificate:c81627a1bf6f766379cc4b98fd2a44ed"
}

resource "ibm_is_lb_listener" "example2" {
  lb                         = ibm_is_lb.example2.id
  port                       = "9087"
  protocol                   = "http"
  idle_connection_timeout = 60 
  https_redirect {
    http_status_code = 302
    listener {
        id = ibm_is_lb_listener.example1.listener_id
    }
    uri = "/example?doc=get"
  }
}
```

### Sample to create a load balancer listener for a route mode enabled private network load balancer.

```terraform

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.68.0/24"
}

resource "ibm_is_lb" "example" {
  name       = "example-lb"
  subnets    = [ibm_is_subnet.example.id]
  profile    = "network-fixed"
  type       = "private"
  route_mode = "true"
}

resource "ibm_is_lb_listener" "example" {
  lb       = ibm_is_lb.example.id
  protocol = "tcp"
  default_pool = ibm_is_lb_pool.example.id
}
```

### Sample to create a public load balancer listener with range of ports.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.68.0/24"
}

resource "ibm_is_lb" "example" {
  name       = "example-lb"
  subnets    = [ibm_is_subnet.example.id]
  profile    = "network-fixed"
  type       = "public"
}

resource "ibm_is_lb_listener" "example1" {
  lb        = ibm_is_lb.example.id
  protocol  = "tcp"
  port_min 	= 100
  port_max 	= 200
}
resource "ibm_is_lb_listener" "example2" {
  lb        = ibm_is_lb.example.id
  protocol  = "tcp"
  port_min 	= 300
  port_max 	= 400
}
```
### Example to create a listener with 
## Timeouts
The `ibm_is_lb_listener` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for updating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `accept_proxy_protocol`- (Optional, Bool)  If set to **true**, listener forwards proxy protocol information that are supported by load balancers in the application family. Default value is **false**.
- `lb` - (Required, Forces new resource, String) The load balancer unique identifier.

- `port`- (Optional, Integer) The listener port number. Valid range `1` to `65535`.

  ~> **NOTE**
    Private network load balancers with `route_mode` enabled don't support `port`, they support only one port range from `port_min (1)` - `port_max (65535)`. Only accepted value of `port` for `route_mode` enabled private network load balancer is `1`. Any other value will show change or update-in-place and returns an error.

  ~> **NOTE**
    Either use `port` or (`port_min`-`port_max`) for public network load balancers 
- `port_min`- (Optional, Integer) The inclusive lower bound of the range of ports used by this listener.

  ~> **NOTE**
    Only load balancers in the `network` family support more than one port per listener. When route mode is enabled, only a value of `1` is supported for `port_min`.

- `port_max`- (Optional, Integer) The inclusive upper bound of the range of ports used by this listener.

  ~> **NOTE**
    Only load balancers in the `network` family support more than one port per listener. When `route mode` is enabled, only a value of `65535` is supported for port_max.

- `protocol` - (Required, String) The listener protocol. Enumeration type are `http`, `tcp`, `https` and `udp`. Network load balancer supports only `tcp` and `udp` protocol.
- `default_pool` - (Optional, String) The load balancer pool unique identifier.

    ~> **The specified pool must**
      </br>&#x2022; Belong to this load balancer
      </br>&#x2022; Have the  same protocol as this listener, or have a compatible protocol. At present, the compatible protocols are http and https.
      </br>&#x2022; Not already be the default_pool for another listener 

- `certificate_instance` - (Optional, String) The CRN of the certificate instance, it is applicable(mandatory) only to https protocol.

  !> **Removal Notification** Certificate Manager support is removed, please use Secrets Manager.

- `connection_limit` - (Optional, Integer) The connection limit of the listener. Valid range is **1 to 15000**. Network load balancer do not support `connection_limit` argument.
- `https_redirect_listener` - (Optional, String) ID of the listener that will be set as http redirect target.
- `https_redirect_status_code` - (Optional, Integer) The HTTP status code to be returned in the redirect response, one of [301, 302, 303, 307, 308].
- `https_redirect_uri` - (Optional, String) Target URI where traffic will be redirected.

~**Note** 
  `https_redirect_listener`, `https_redirect_status_code` and `https_redirect_uri` are deprecated and will be removed in future. Please use `https_redirect` instead

- `https_redirect` - (Optional, List) If present, the target listener that requests are redirected to. Removing `https_redirect` would update the load balancer listener and disable the `https_redirect`
  
  Nested schema for **https_redirect**:
	- `http_status_code` - (Required, Integer) The HTTP status code for this redirect. Allowable values are: `301`, `302`, `303`, `307`, `308`.
	- `listener` - (Required, List)
	  
      Nested schema for **listener**:
		- `id` - (Required, String) The unique identifier for this load balancer listener.
	- `uri` - (Optional, String) The redirect relative target URI. Removing `uri` would update the load balancer listener and remove the `uri` from `https_redirect`
- `idle_connection_timeout` - (Optional, Integer) The idle connection timeout of the listener in seconds. Supported for load balancers in the `application` family. Default value is `50`, allowed value is between `50` - `7200`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer listener.
- `status` - (String) The status of load balancer listener.
- `https_redirect` - (List) If present, the target listener that requests are redirected to.
  
  Nested schema for **https_redirect**:
	- `http_status_code` - (Integer) The HTTP status code for this redirect.
	- `listener` - (List)
	  
      Nested schema for **listener**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  
          Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The listener's canonical URL.
		- `id` - (String) The unique identifier for this load balancer listener.
	- `uri` - (String) The redirect relative target URI.
## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_lb_listener` resource by using `id`.
The `id` property can be formed from `load balancer ID`, and `listener ID`. For example:

```terraform
import {
  to = ibm_is_lb_listener.example
  id = "<loadbalancer_ID>/<listener_ID>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_lb_listener.example <loadbalancer_ID>/<listener_ID>
```