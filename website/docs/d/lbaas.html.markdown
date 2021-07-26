---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lbaas"
description: |-
  Manages IBM Cloud load balancer as a service.
---

# ibm_lbaas
Retrieve information of an existing IBM Cloud load balancer as a read-only data source. For more information, about load balancer as a service, see [enabling auto scale for better capacity and resiliency](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-auto-scale).
 
## Example usage

```terraform
resource "ibm_lbaas" "lbaas" {
  name        = "test"
  description = "updated desc-used for terraform uat"
  subnets     = [1878778]
  datacenter  = "dal09"

  protocols {
    frontend_protocol     = "HTTP"
    frontend_port         = 80
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "round_robin"
  }

  server_instances {
    private_ip_address = "10.1.19.26"
  }
}

data "ibm_lbaas" "tfacc_lbaas" {
  name = "test"
}

```


## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the load balancer.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `active_connections`- (Integer) The number of total established connections.
- `description` - (String) A description of the load balancer.
- `datacenter` - (String) The data center where load balancer is located.
- `health_monitors` - (List of Objects) A nested block describes the health_monitors assigned to the load balancer.

  Nested scheme for `health_monitors`:
  - `interval`- (Integer) The interval in seconds to perform the health check.
  - `max_retries`- (Integer) The maximum retries before the load balancer are considered unhealthy.
  - `monitor_id` - (String) The health monitor UUID.
  - `protocol` - (String) The backend protocol.
  - `port` - (String) The backend port.
  - `timeout` - (String) The health check method.
  - `url_path` - (String)  If monitor is "HTTP", it specifies the URL path.
- `protocols` - (List of Objects) A nested block describes the protocols that are assigned to the load balancer.

  Nested scheme for `protocols`:
  - `backend_protocol` - (String) The backend protocol.
  - `backend_port`- (Integer) The backend protocol port number.
  - `frontend_protocol` - (String) The front-end protocol.
  - `frontend_port`- (Integer) The front-end protocol port number.
  - `load_balancing_method` - (String) The load-balancing algorithm.
  - `protocol_id` - (String) The UUID of a load balancer protocol.
  - `max_conn`- (Integer) The number of connections the listener can accept.
  - `session_stickiness`- (Bool) Session stickiness.
  - `tls_certificate_id` - (String) The ID of the SSL/TLS certificate used for a protocol.
- `server_instances_up`- (Integer) The number of service instances, that are in the `UP` health state.
- `server_instances_down`- (Integer) The number of service instances, that are in the `DOWN` health state.
- `status` - (String) Specifies the operation status of the load balancer as `online` or `offline`.
- `ssl_ciphers` - (Array) The list of SSL offloads.
- `type` - (String) Specifies whether a load balancer is public or private.
- `use_system_public_ip_pool` - (String) It specifies whether the public IP addresses are allocated from system public IP pool or public subnet from the account order of the load balancer.
- `vip` - (String) The virtual IP address of the load balancer.
