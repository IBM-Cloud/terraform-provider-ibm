---
layout: "ibm"
page_title: "IBM : lb_listener"
sidebar_current: "docs-ibm-resource-is-lb-listener"
description: |-
  Manages IBM load balancer listener.
---

# ibm\_is_lb_listener

Provides a load balancer listener resource. This allows load balancer listener to be created, updated, and cancelled.

**Note**: When provisioning the load balancer listener along with load balancer pool or pool member, Use explicit depends on the resources or perform the terraform apply with parallelism 1. For more information on explicit dependencies refer [here](https://learn.hashicorp.com/terraform/getting-started/dependencies#implicit-and-explicit-dependencies)

## Example Usage

In the following example, you can create a load balancer listener along with pool and pool member:

```hcl
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
  depends_on = ["ibm_is_lb_listener.testacc_lb_listener"]
}

resource "ibm_is_lb_pool_member" "webapptier-lb-pool-member-zone1" {
  count = "2"
  lb    = "8898e627-f61f-4ac8-be85-9db9d8bfd345"
  pool  = "${element(split("/",ibm_is_lb_pool.webapptier-lb-pool.id),1)}"
  port  = "80"
  target_address = "192.168.0.1"
  depends_on = ["ibm_is_lb_listener.testacc_lb_listener"]
}


```

## Timeouts

ibm_is_lb_listener provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating Instance.
* `update` - (Default 60 minutes) Used for updating Instance.
* `delete` - (Default 60 minutes) Used for deleting Instance.


## Argument Reference

The following arguments are supported:

* `lb` - (Required, Forces new resource, string) The load balancer unique identifier.
* `port` - (Required, int) The listener port number. Valid range 1 to 65535.
* `protocol` - (Required, string) The listener protocol. Enumeration type: http, tcp, https.
* `default_pool` - (Optional, string) The load balancer pool unique identifier.
* `certificate_instance` - (Optional, string) CRN of the certificate instance.
* `connection_limit` - (Optional, int) The connection limit of the listener. Valid range  1 to 15000.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the load balancer listener.
* `status` - The status of load balancer listener.

## Import

ibm_is_lb_listener can be imported using lbID and listenerID, eg

```
$ terraform import ibm_is_lb_listener.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
