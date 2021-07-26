---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lb_service_group"
description: |-
  Manages IBM local load balancer service group.
---

# ibm_lb_service_group
Create, update, and delete a [local load balancer service group](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-ibm-cloud-load-balancer-basics). 

**Note**

For more information,  see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service_Group).

## Example usage

In the following example, you can create a local load balancer service group:

```terraform
resource "ibm_lb_service_group" "test_service_group" {
    port = 82
    routing_method = "CONSISTENT_HASH_IP"
    routing_type = "HTTP"
    load_balancer_id = ibm_lb.test_lb_local.id
    allocation = 100
}
```

## Argument reference 
Review the argument references that you can specify for your resource. 

- `allocation` - (Required, Integer) The connection allocation for the load balancer service group.
- `load_balancer_id` - (Required, Forces new resource,Integer) The ID of the local load balancer.
- `port` - (Required, Integer) The port for the local load balancer service group.
- `routing_method` - (Required, String) The routing method for the load balancer group. For example, `CONSISTENT_HASH_IP`.
- `routing_type`- (Required, String) The routing type for the group.
- `timeout`- (Optional, Integer) The timeout value for connections from remote clients to the load balancer. Timeout values are only valid for HTTP service groups.
- `tags`- (Optional, Array of Strings) Tags associated with the local load balancer service group instance.  **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `service_group_id`- (String) The unique identifier of the load balancer service group.
- `virtual_server_id` - (String) The unique identifier of the virtual server.
