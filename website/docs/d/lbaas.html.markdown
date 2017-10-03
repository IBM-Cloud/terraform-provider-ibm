---
layout: "ibm"
page_title: "IBM : lbaas"
sidebar_current: "docs-ibm-datasource-lbaas"
description: |-
  Manages IBM Load Balancer As A Service.
---

# ibm\_lbaas

 
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

* `description` -  Description of a load balancer.
* `datacenter` -  Datacenter, where load balancer is located.
* `protocols` -  Nested block describing protocols assigned to load balancer.
* `server_instances` -  Nested block describing the Server instances for this load balancer.
* `type` - Specifies if a load balancer is public or private.
* `status` - The operation status 'ONLINE' or 'OFFLINE' of a load balancer.
* `vip` - The virtual ip address of this load balancer.
* `server_instances_up` - The number of service instances which are `UP` health state.
* `server_instances_down` - The number of service instances which are `DOWN` health state.
* `active_connections` - The number of total established connections.
* `active_sessions` - The number of total current sessions.

Nested `protocols` blocks have the following structure:

* `frontend_protocol` -  Frontend protocol, one of 'TCP', 'HTTP', 'HTTPS'.
* `frontend_port` -  Frontend Protocol port number.
* `backend_protocol` - Backend protocol, one of 'TCP', 'HTTP'.
* `backend_port` -  Backend Protocol port number.
* `load_balancing_method` - Load balancing algorithm: 'round_robin', 'weighted_round_robin', 'least_connection'.
* `session_stickiness` - Session stickness. Valid values is SOURCE_IP
* `max_conn` - No. of connections the listener can accept. 
* `tls_certificate_id` - This references to SSL/TLS certificate for a protocol.
* `protocol_id` - The UUID of a load balancer protocol.

Nested `server_instances` blocks have the following structure:

* `private_ip_address` - The Private IP address of a load balancer member.
* `weight` - The weight of a load balancer member.
* `status` - The status of a load balancer member whether it is UP/DOWN.
* `member_id` - The UUID of a load balancer member.


    