---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lbaas_health_monitor"
description: |-
  Manages health monitors of a IBM load balancer as a service.
---

# ibm_lbaas_health_monitor
Create, delete, and update a health monitor configuration of the load balancer. For more information, about health monitor of a load balancer as a service, see [monitoring and managing your service](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-monitoring-and-managing-your-service).

## Example usage

```terraform

resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]

  protocols {
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
    tls_certificate_id    = 11670
  }
  protocols {
    frontend_protocol     = "HTTP"
    frontend_port         = 80
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
  }
}

resource "ibm_lbaas_health_monitor" "lbaas_hm" {
  protocol    = ibm_lbaas.lbaas.health_monitors[0].protocol
  port        = ibm_lbaas.lbaas.health_monitors[0].port
  timeout     = 3
  interval    = 5
  max_retries = 6
  url_path    = "/"
  lbaas_id    = ibm_lbaas.lbaas.id
  monitor_id  = ibm_lbaas.lbaas.health_monitors[0].monitor_id
}

```

## Argument reference 
Review the argument references that you can specify for your resource.

- `interval` - (Optional, Integer) Interval in seconds to perform.
- `lbaas_id` - (Required, Forces new resource, String) LBaaS unique identifier.
- `max_retries` - (Optional, Integer) Maximum retries.
- `monitor_id` - (Required, Forces new resource, String)Health Monitor unique identifier. The monitor ID can be imported from either the `ibm_lbaas` resource or datasource. For example, `ibm_lbaas.lbaas.health_monitors.X.monitor_id` or `data.ibm_lbaas.lbaas.health_monitors.X.monitor_id`.
- `protocol` - (Required, String)The back-end protocol.
- `port` - (Required, Integer) The back-end port.
- `timeout` - (Optional, Integer) Health check methods timeout in.
- `url_path`- (Optional, String) If monitor is **HTTP**, it specifies the URL path.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the LBaaS health monitor resource. The ID is composed of `<lbaas_id>/<monitor_id>`.


## Import
The `ibm_lbaas_health_monitor` resource can be imported by using LBaaS ID and monitor ID.

**Example**

```
$ terraform import ibm_lbaas_health_monitor.example 988-454f-45vf-454542/d343f-f44r-wer3-fe
```
