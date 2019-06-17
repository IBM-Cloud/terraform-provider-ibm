---
layout: "ibm"
page_title: "IBM: ibm_cse_instance"
sidebar_current: "docs-ibm-resource-cse-instance"
description: |-
  Provides a IBM CSE instance resource.
---

# ibm_cse_instance

Provides a IBM CSE instance resource, which can help customer to connect to service 
via private network.

## Example Usage

```hcl
# Add a CSE instance
resource "ibm_cse_instance" "instance1" {
  service = "terraform-1"
  customer = "customer1"
  service_addresses = ["10.102.33.131", "10.102.33.133"]
  region = "us-south"
  data_centers = ["dal10", "dal13"]
  tcp_ports = [8080, 80]
  udp_ports = [12345]
  tcp_range = "8000-9000"
  udp_range = "50000-51000"
  max_speed = "1g"
  estado_proto = "http"
  estado_port = 80
  estado_path = "/healthcheck"
  dedicated = 0
  multi_tenant = 1
  acl = ["10.10.10.1/24"]
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Required) The service name which this instance is created for.
* `customer` - (Required) The customer name which this instance is created for.
* `service_addresses` - (Required) The service addresses to which the traffic will be sent to from this instance.
* `region` - (Required) In which Region this instance will be put.
* `data_centers` - (Required) In which data centers this instance will be put.
* `tcp_ports` - (Required) To which ports the tcp traffic will be forwarded.
* `udp_ports` - (Optional) To which ports the tcp traffic will be forwarded.
* `tcp_range` - (Optional) In which port range the tcp traffic will be forwarded.
* `udp_range` - (Optional) In which port range the udp traffic will be forwarded.
* `max_speed` - (Required) Netweek speed. 1g or 20g.
* `estado_proto` - (Optional) The protocal used by service health check.
* `estado_port` - (Optional) The port used by service health check.
* `estado_path` - (Optional) The path used by service health check.
* `dedicated` - (Optional) Whether to put this instance in a dedicated machine or not. 1 means putting
to a dedicated machine. 0 not.
* `multi_tenant` - (Optional) Whether to put this instance with other instances together. 1 means together while 0 means not.
* `acl` - (Optional) To specify source addresses from which traffic can be forwarded.


## Attributes Reference

The following attributes are exported:

* `id` - The CSE instance ID
* `url` - The url to access this CSE instance
* `static_addresses` - The ip addreses to access this CSE instance
