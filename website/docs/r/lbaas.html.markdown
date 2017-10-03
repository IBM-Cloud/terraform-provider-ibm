---
layout: "ibm"
page_title: "IBM : lbaas"
sidebar_current: "docs-ibm-resource-lbaas"
description: |-
  Manages IBM Load balancer as a service.
---

# ibm\_lbaas

The resource lbaas creates a Load Balancer as a service. Currently only one subnet is supported


 
## Example Usage

```hcl

resource "ibm_compute_vm_instance" "vm_instances" {
  count = "2"
  ....
}

resource "ibm_lbaas" "lbaas" {
  name        = "terraformLB"
  description = "delete this"
  subnets     = [1511875]
  datacenter  = "wdc04"

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

  server_instances = [
    {
      "private_ip_address" = "${ibm_compute_vm_instance.vm_instances.0.ipv4_address_private}"
    },
    {
      "private_ip_address" = "${ibm_compute_vm_instance.vm_instances.1.ipv4_address_private}"
    },
  ]
}


```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The load balancer's name.
* `description` - (Optional, string) Description of a load balancer.
* `datacenter` - (Required, string) Datacenter, where load balancer is located.
* `subnets` - (Required, array) The subnet where this Load Balancer will be provisioned.Only one subnet is supported.
* `protocols` - (Optional, array) Nested block describing protocols assigned to load balancer.
* `server_instances` - (Optional, array) Nested block describing the Server instances for this load balancer.

Nested `protocols` blocks have the following structure:

* `frontend_protocol` - (Required, string) Frontend protocol, one of 'TCP', 'HTTP', 'HTTPS'.
* `frontend_port` - (Required, int)  Frontend Protocol port number. Should be in range (1, 65535)
* `backend_protocol` - (Required, string) Backend protocol, one of 'TCP', 'HTTP', 'HTTPS'.
* `backend_port` - (Required, int)  Backend Protocol port number. Should be in range (1, 65535)
* `load_balancing_method` - (Optional, string) Load balancing algorithm: 'round_robin', 'weighted_round_robin', 'least_connection'. Default is 'round_robin'.
* `session_stickiness` - (Optional, string) Session stickness. Valid values is SOURCE_IP
* `max_conn` - (Optional, int) No. of connections the listener can accept. Should be between 1-64000 
* `tls_certificate_id` - (Optional, int) This references to SSL/TLS certificate for a protocol. Should be specified when the frontend protocol is selected as HTTPS.

Nested `server_instances` blocks have the following structure:

* `private_ip_address` - (Required, string) The Private IP address of a load balancer member.
* `weight` - (Optional, int) The weight of a load balancer member.


## Attributes Reference

The following attributes are exported:

* `id` - The id of policy created.
* `type` - Specifies if a load balancer is public or private.
* `status` - The operation status 'ONLINE' or 'OFFLINE' of a load balancer.
* `vip` - The virtual ip address of this load balancer.
* `protocol_id` - The UUID of a load balancer protocol.
* `member_id` - The UUID of a load balancer member.