---
layout: "ibm"
page_title: "IBM : lb_service_group"
sidebar_current: "docs-ibm-resource-lb-service-group"
description: |-
  Manages IBM local load balancer service group.
---

# ibm\_lb_service_group

Provides a resource for local load balancer service groups. This allows local load balancer service groups to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group).

## Example Usage

In the following example, you can create a local load balancer service group:

```hcl
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

* `load_balancer_id` - (Required, integer) The ID of the local load balancer.
* `allocation` - (Required, integer) The connection allocation for the load balancer service group.
* `port` - (Required, integer) The port for the local load balancer service group.
* `routing_method` - (Required, string) The routing method for the load balancer group. For example, `CONSISTENT_HASH_IP`.
* `routing_type` - (Required, string) The routing type for the group.
* `timeout` - (Optional, int) The timeout value for connections from remote clients to the load balancer. Timeout values are only valid for HTTP service groups. 
* `tags` - (Optional, array of strings) Tags associated with the local load balancer service group instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `virtual_server_id` - The unique identifier of the virtual server.
* `service_group_id` - The unique identifier of the load balancer service group.
