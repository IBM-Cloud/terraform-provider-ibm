---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lb_service"
description: |-
  Manages IBM Cloud local load balancer service.
---

# ibm_lb_service
Create, update, and delete a local load balancer service. For more information, about local load balancer service, see [load balancer basics](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-ibm-cloud-load-balancer-basics).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_Service).

## Example usage

In the following example, you can create a local load balancer service:

```terraform
resource "ibm_lb_service" "test_lb_local_service" {
    port = 80
    enabled = true
    service_group_id = ibm_lb_service_group.test_service_group.service_group_id
    weight = 1
    health_check_type = "DNS"
    ip_address_id = ibm_compute_vm_instance.test_server.ip_address_id
}

```

## Argument reference 
Review the argument references that you can specify for your resource. 

- `ip_address_id` - (Required, Forces new resource,Integer) The ID of the virtual server.Yes.
- `enabled` - (Required, Integer) Specifies whether you want to enable the load balancer service. The default value is **false**.
- `health_check_type` - (Required, String)The health check type for the load balancer service.No.
- `port` - (Required, Integer) The port for the local load balancer service.
- `service_group_id` - (Required, Forces new resource,Integer) The ID of the local load balancer service group.
- `tags`- (Optional, Array of Strings) Tags associated with the local load balancer service instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `weight` - (Required, Integer) The weight for the load balancer service.
