---
layout: "ibm"
page_title: "IBM : lb_service_group"
sidebar_current: "docs-ibm-resource-lb-service-group"
description: |-
  Manages IBM local load balancer service group.
---

# ibm\_lb_service_group

Provides a resource for local load balancer groups. This allows local load balancer groups to be created, updated, and deleted.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group).

## Example Usage

```hcl
# Create a local load balancer service group
resource "ibm_lb_service_group" "test_service_group" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = "${ibm_lb.test_lb_local.id}"
    allocation = 100
}
```

## Argument Reference

The following arguments are supported:

* `load_balancer_id` - (Required, integer) Set the ID of the local load balancer.
* `allocation` - (Required, integer) Set the allocation field for the load balancer service group.
* `port` - (Required, integer) Set the port for the local load balancer service group.
* `routing_method` - (Required, string) Set the routing method for the load balancer group. For example, `CONSISTENT_HASH_IP`.
* `routing_type` - (Required, string) Set the routing type for the group.

## Attributes Reference

The following attributes are exported:

* `virtual_server_id` - ID of the virtual server.
* `service_group_id` - ID of the load balancer service group.
