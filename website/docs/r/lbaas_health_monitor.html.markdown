---
layout: "ibm"
page_title: "IBM : lbaas_health_monitor"
sidebar_current: "docs-ibm-resource-lbaas-health-monitor"
description: |-
  Manages health monitors of a IBM Load balancer as a service.
---

# ibm\_lbaas\_health\_monitor

Provides a resource for the health monitors of IBM Lbaas. This allows to update a health monitor configuration of the load balancer. Health monitors are created and deleted by the creation and deletion of lbaas protocols.
 
## Example Usage

```hcl

resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]

  protocols = [{
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
    tls_certificate_id    = 11670
  },
    {
      frontend_protocol     = "HTTP"
      frontend_port         = 80
      backend_protocol      = "HTTP"
      backend_port          = 80
      load_balancing_method = "round_robin"
    },
  ]
}

resource "ibm_lbaas_health_monitor" "lbaas_hm" {
  protocol = "${ibm_lbaas.lbaas.health_monitors.0.protocol}"
  port = "${ibm_lbaas.lbaas.health_monitors.0.port}"
  timeout = 3
  interval = 5
  max_retries = 6
  url_path = "/"
  lbaas_id = "${ibm_lbaas.lbaas.id}"
  monitor_id = "${ibm_lbaas.lbaas.health_monitors.0.monitor_id}"
}

```

## Argument Reference

The following arguments are supported:

* `monitor_id` - (Required,string) Health Monitor unique identifier. The monitor id can be imported from either the ibm_lbaas resource or datasource.
ex: ibm_lbaas.lbaas.health_monitors.X.monitor_id or data.ibm_lbaas.lbaas.health_monitors.X.monitor_id
* `lbaas_id` - (Required,string) Lbaas unique identifier
* `protocol` - (Required, string) Backends protocol
* `port` - (Required, int) Backends port
* `interval` - (Optional,int) Interval in seconds to perform 
* `max_retries` - (Optional,int) Maximum retries
* `timeout` - (Optional,int) Health check methods timeout in 
* `url_path` - (Optional,string) If monitor is "HTTP", this specifies URL path

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the lbaas health monitor resource. The id is composed of \<lbaas_id\>/\<monitor_id/>

## Import

ibm_lbaas_health_monitor can be imported using lbaas_id and monitor_id, eg

```
$ terraform import ibm_lbaas_health_monitor.example 988-454f-45vf-454542/d343f-f44r-wer3-fe3