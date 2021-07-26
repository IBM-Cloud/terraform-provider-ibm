---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lb"
description: |-
  Manages IBM Cloud load balancer.
---

# ibm_lb
Create, update, and delete a load balancer. For more information, about load balancer, see [selecting the service and configuring basic parameters](https://cloud.ibm.com/docs/loadbalancer-service?topic=loadbalancer-service-configuring-ibm-cloud-load-balancer-basic-parameters).

## Example usage
In the following example, you can create a local load balancer:

```terraform
resource "ibm_lb" "test_lb_local" {
  connections = 1500
  datacenter  = "tok02"
  ha_enabled  = false
  dedicated   = false

  //User can increase timeouts
  timeouts {
    create = "45m"
  }
}
```

## Timeouts

The `ibm_lb` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating instance.


## Argument reference 
Review the argument references that you can specify for your resource. 

- `connections` - (Required, Integer) The number of connections for the local load balancer. Only incremental upgrade is supported. For downgrade, please open the SoftLayer support ticket.
- `datacenter` - (Required, Forces new resource, String)The data center for the local load balancer.
- `dedicated` - (Optional, Bool)  Specifies whether the local load balancer must be dedicated. The default value is **false**.
- `ha_enabled`- (Required, Forces new resource, Bool) Specifies whether the local load balancer must be HA-enabled.
- `security_certificate_id` - (Optional, Forces new resource, Integer) The ID of the security certificate associated with the local load balancer.
- `ssl_offload` - (Optional, Bool)  Specifies the local load balancer SSL offload. If **true** start SSL acceleration on all SSL virtual services. (those with a type of HTTPS) This action should be taken only after configuring an SSL certificate for the virtual IP. If **false** stop SSL acceleration on all SSL virtual services (those with a type of HTTPS). The default value is **false**.
- `tags`- (Optional, Array of Strings)Tags associated with the local load balancer instance.     **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the local load balancer.
- `ip_address`- (String) The IP address of the local load balancer.
- `hostname`- (String) The host name of the local load balancer.
- `subnet_id`- (String) The unique identifier of the subnet associated with the local load balancer.
- `ssl_enabled`- (String) The status of whether the local load balancer provides SSL capability.
