---
layout: "ibm"
page_title: "IBM : lb_service"
sidebar_current: "docs-ibm-resource-lb-service"
description: |-
  Manages IBM local load balancer service.
---

# ibm\_lb_service

Provides a resource for local load balancer services. This allows local load balancer services to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service).

## Example Usage

In the following example, you can create a local load balancer service:

```hcl
resource "ibm_lb_service" "test_lb_local_service" {
    port = 80
    enabled = true
    service_group_id = "${ibm_lb_service_group.test_service_group.service_group_id}"
    weight = 1
    health_check_type = "DNS"
    ip_address_id = "${ibm_compute_vm_instance.test_server.ip_address_id}"
}

```

## Argument Reference

The following arguments are supported:

* `service_group_id` - (Required, integer) The ID of the local load balancer service group.
* `ip_address_id` - (Required, integer) The ID of the virtual server.
* `port` - (Required, integer) The port for the local load balancer service.
* `enabled` - (Required, boolean) Specifies whether you want to enable the load balancer service. The default value is `false`.
* `health_check_type` - (Required, string) The health check type for the load balancer service.
* `weight` - (Required, integer) The weight for the load balancer service.
* `tags` - (Optional, array of strings) Tags associated with the local load balancer service instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.
