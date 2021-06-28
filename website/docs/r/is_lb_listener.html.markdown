---

subcategory: "VPC infrastructure"
page_title: "IBM : lb_listener"
description: |-
  Manages IBM load balancer listener.
---

# ibm_is_lb_listener
Create, update, or delete a listener for a VPC load balancer. For more information, about load balancer listener, see [working with listeners](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-listeners).

**Note**

When provisioning the load balancer listener along with load balancer pool or pool member, Use explicit depends on the resources or perform the terraform apply with parallelism. For more information, about explicit dependencies, see [create resource dependencies](https://learn.hashicorp.com/terraform/getting-started/dependencies#implicit-and-explicit-dependencies).

## Example usage
An example, to create a load balancer listener along with the pool and pool member.

```terraform
resource "ibm_is_lb_listener" "testacc_lb_listener" {
  lb       = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  port     = "9080"
  protocol = "http"
}

resource "ibm_is_lb_pool" "webapptier-lb-pool" {
  lb                 = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  name               = "a-webapptier-lb-pool"
  protocol           = "http"
  algorithm          = "round_robin"
  health_delay       = "5"
  health_retries     = "2"
  health_timeout     = "2"
  health_type        = "http"
  health_monitor_url = "/"
  depends_on         = [ibm_is_lb_listener.testacc_lb_listener]
}

resource "ibm_is_lb_pool_member" "webapptier-lb-pool-member-zone1" {
  count          = "2"
  lb             = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  pool           = element(split("/", ibm_is_lb_pool.webapptier-lb-pool.id), 1)
  port           = "80"
  target_address = "192.168.0.1"
  depends_on     = [ibm_is_lb_listener.testacc_lb_listener]
}
```

## Timeouts
The `ibm_is_lb_listener` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for updating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `accept_proxy_protocol`- (Optional, Bool)  If set to **true**, listener forwards proxy protocol information that are supported by load balancers in the application family. Default value is **false**.
- `lb` - (Required, Forces new resource, String) The load balancer unique identifier.
- `port`- (Required, Integer) The listener port number. Valid range 1 to 65535.
- `protocol` - (Required, String) The listener protocol. Enumeration type are `http`, `tcp`, and `https`. Network load balancer supports only `tcp` protocol.
- `default_pool` - (Optional, String) The load balancer pool unique identifier.
- `certificate_instance` - (Optional, String) The CRN of the certificate instance, it is applicable(mandatory) only to https protocol.
- `connection_limit` - (Optional, Integer) The connection limit of the listener. Valid range is **1 to 15000**. Network load balancer do not support `connection_limit` argument.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer listener.
- `status` - (String) The status of load balancer listener.

## Import
The `ibm_is_lb_listener` resource can be imported by using the load balancer ID and listener ID.

**Syntax**

```
$ terraform import ibm_is_lb_listener.example <loadbalancer_ID>/<listener_ID>
```

**Example**

```
$ terraform import ibm_is_lb_listener.example d7bec597-4726-451f-8a63-e61212c32c/cea6651a-bc0a-4438-9f8a-44444f3ebb
```
