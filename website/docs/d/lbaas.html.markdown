---
layout: "ibm"
page_title: "IBM : lbaas"
sidebar_current: "docs-ibm-datasource-lbaas"
description: |-
  Manages IBM Load Balancer As A Service.
---

# ibm\_lbaas

Import the details of an existing IBM Cloud load balancer as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.
 
## Example Usage

```hcl
resource "ibm_lbaas" "lbaas" {
  name        = "test"
  description = "updated desc-used for terraform uat"
  subnets     = [1878778]
  datacenter  = "dal09"

  protocols = [{
    "frontend_protocol" = "HTTP"
    "frontend_port" = 80
    "backend_protocol" = "HTTP"
    "backend_port" = 80
    "load_balancing_method" = "round_robin"
  }]

  server_instances = [{
    "private_ip_address" = "10.1.19.26",
  },
  ]
}
    data "ibm_lbaas" "tfacc_lbaas" {
    name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The load balancer's name.

## Attributes Reference

The following attributes are exported:

* `description` - A description of the load balancer.
* `datacenter` - The datacenter where load balancer is located.
* `protocols` - A nested block describing the protocols assigned to the load balancer. Nested `protocols` blocks have the following structure:
  * `frontend_protocol` - The frontend protocol.
  * `frontend_port` - The frontend protocol port number.
  * `backend_protocol` - The backend protocol.
  * `backend_port` - The backend protocol port number.
  * `load_balancing_method` - The load balancing algorithm.
  * `session_stickiness` - Session stickiness.
  * `max_conn` - The number of connections the listener can accept.
  * `tls_certificate_id` - The ID of the SSL/TLS certificate being used for a protocol.
  * `protocol_id` - The UUID of a load balancer protocol.
* `server_instances` - A nested block describing the server instances for this load balancer. Nested `server_instances` blocks have the following structure:
  * `private_ip_address` - The private IP address of a load balancer member.
  * `weight` - The weight of a load balancer member.
  * `status` - Specifies the status of a load balancer member as `UP` or `DOWN`.
  * `member_id` - The UUID of a load balancer member.
* `health_monitors` - A nested block describing the health_monitors assigned to the load balancer. Nested `health_monitors` blocks have the following structure:
  * `protocol` - Backends protocol
  * `port` - Backends port
  * `interval` - Interval in seconds to perform 
  * `max_retries` - Maximum retries
  * `timeout` - Health check methods timeout in 
  * `url_path` - If monitor is "HTTP", this specifies URL path
  * `monitor_id` - Health Monitor UUID
* `type` - Specifies whether a load balancer is public or private.
* `status` - Specifies the operation status of the load balancer as 'ONLINE' or 'OFFLINE'.
* `vip` - The virtual IP address of the load balancer.
* `server_instances_up` - The number of service instances which are in the `UP` health state.
* `server_instances_down` - The number of service instances which are in the `DOWN` health state.
* `active_connections` - The number of total established connections.
