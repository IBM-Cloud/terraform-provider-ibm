---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lbaas"
description: |-
  Manages IBM Cloud load balancer as a service.
---

# ibm_lbaas
Create, delete, and update a load balancer as a service. For more information, about load balancer as a service, see [about IBM Cloud load balancer](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-about-ibm-cloud-load-balancer). Currently, only one subnet is supported.

Cloud load balancer creation takes 5 to 10 minutes. Destroy can take up to 30 minutes. Cloud Load Balancer does not support customization of acceptable response codes. Only the range `2xx` is considered healthy. Redirects in the range 3xx are considered unhealthy. 

 
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


```

## Argument reference 
Review the argument references that you can specify for your resource. 

- `description`- (Optional, String) A description of the load balancer.
- `name` - (Required, Forces new resource, String)The load balancer's name.
- `protocols`- (Optional, List) A nested block describes the protocols that are assigned to load balancer.

  Nested scheme for `protocols`:
  - `backend_protocol` - (Required, String)The back-end protocol. Accepted values are 'TCP', 'HTTP', and 'HTTPS'.
  - `backend_port`- (Required, Integer) The back-end protocol port number. The port number must be in the range of `1-65535`.
  - `frontend_protocol` - (Required, String)The front-end protocol. Accepted values are 'TCP', 'HTTP', and 'HTTPS'. No.
  - `frontend_port`- (Required, Integer) The front-end protocol port number. The port number must be in the range of `1-65535`.
  - `load_balancing_method`- (Optional, String) The load-balancing algorithm. Accepted values are 'round_robin', 'weighted_round_robin', and 'least_connection'. The default is 'round_robin'.
  - `max_conn` - (Optional, Integer)The maximum number of connections the listener can accept. The number must be `1-64000`.
  - `session_stickiness`- (Optional, String) The SOURCE_IP for session stickiness.
  - `tls_certificate_id` - (Optional, Integer)The ID of the SSL/TLS certificate used for a protocol. This ID should be specified when `front-end protocol` has a value of `HTTPS`.
- `session_stickiness` - (Optional, String) The `SOURCE_IP` or `HTTP_COOKIE` for session stickiness.No-
- `ssl_ciphers` - (Optional, List) The comma-separated list of SSL Ciphers. You can find list of supported ciphers [SSL_offload](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-ssl-offload-with-ibm-cloud-load-balancer).
- `type`- (Optional, Forces new resource, String) Specify whether this load balancer is a public or internal facing load balancer. Accepted values are **PUBLIC** or **PRIVATE**. The default is **PUBLIC**.
- `subnets` - (Required,  Forces new resource, Array) The subnet where the load balancer will be provisioned. Only one subnet is supported.
- `use_system_public_ip_pool` -  (Optional, Bool) Applicable for public load balancer only. It specifies whether the public IP addresses are allocated from system public IP pool or public subnet from the account order of the load balancer. The default value is **true**.
- `wait_time_minutes` - (Required, Integer) The duration, expressed in minutes, to wait for the LBaaS instance to become available before declaring it as created. It is also the same amount of time waited for deletion to finish. The default value is `90`.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `datacenter` - (String) The datacenter where the load balancer is provisioned. This is based on the subnet chosen while creating load-balancer.
- `id` - (String) The unique identifier of the created policy.
- `health_monitors` - (List) A nested block describes the health_monitors assigned to the load balancer. Nested `health_monitors` blocks have the following structure.

  Nested scheme for `health_monitors`:
  - `interval` - (String) Interval in seconds to perform.
  - `max_retries` - (String) Maximum retries.
  - `monitor_id` - (String) Health Monitor UUID.
  - `protocol` - (String) The back-end protocol.
  - `port` - (String) The back-end port.
  - `timeout` - (String) Health check methods timeout in.
  - `url_path` - (String) If monitor is "HTTP", it specifies the URL path.
- `protocol_id`- (String) The UUID of a load balancer protocol.
- `status`- (String) Specifies the operation status of the load balancer as `online` or `offline`.
- `vip`- (String) The virtual IP address of the load balancer.
