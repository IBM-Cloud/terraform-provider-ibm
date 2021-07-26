---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lb_vpx_service"
description: |-
  Manages IBM VPX load balancer services.
---

# ibm_lb_vpx_service
Create, update, and delete a VPX load balancer service. For more information, about VPC load balancer virtual IP address, see [about Citrix Netscaler VPX](https://cloud.ibm.com/docs/citrix-netscaler-vpx?topic=citrix-netscaler-vpx-about-citrix-netscaler-vpx).

**Note** 

If you use Netscaler VPX 10.5, Terraform uses Netscaler's [NITRO REST API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) to manage the resource. Terraform can only access the NITRO API in the IBM Cloud Classic Infrastructure (SoftLayer) private network, so connect to the private network when running  Terraform. You can also use the [SSL VPN](https://www.ibm.com/cloud/vpn-access) to access a private network connection.

## Example usage
In the following example, you can create a VPX load balancer:

```terraform
resource "ibm_lb_vpx_service" "test_service" {
  name = "test_load_balancer_service"
  vip_id = ibm_lb_vpx_vip.testacc_vip.id
  destination_ip_address = ibm_compute_vm_instance.test_server.ipv4_address
  destination_port = 80
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
  usip = "NO"
}
```

## Argument reference 
Review the argument references that you can specify for your resource. 

- `connection_limit`- (Required, Integer) The connection limit for this service. Acceptable values are `0`- `4294967294`. See the [Citrix Netscaler Docs](https://docs.citrix.com/en-us/netscaler/11/reference/netscaler-command-reference/basic/service.html) for details.
- `destination_ip_address` - (Required, Forces new resource, String)The IP address of the server to which traffic directs. If you use Netscaler VPX 10.1, you must indicate a public IP address in an IBM Cloud Classic Infrastructure (SoftLayer) account. If you use Netscaler VPX 10.5, you can use any IP address.
- `destination_port`- (Required, Integer) The destination port of the server to which traffic directs.
- `health_check`- (Required, String) The health check type. See the [IBM Cloud Classic Infrastructure (SoftLayer) API Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service) for details.
- `name`- (Required, Forces new resource, String) The ID of the VPX load balancer service.
- `tags`- (Optional, Array of string)  Tags associated with the VPX load balancer service instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `usip`- (Optional, Forces new resource, String) Whether the service reports the source IP of the client to the service being load balance. Acceptable values are **YES** or **NO**. The default value is **NO**. **Note** This argument is only available for VPX 10.5.
- `vip_id`- (Required, Forces new resource, String) The ID of the VPX load balancer virtual IP address to which the service is assigned.
- `weight` - (Required, Integer) The percentage of the total connection limit allocated to the load balancer between all your services. See the [IBM Cloud Classic Infrastructure (SoftLayer)  API Docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service) for details.   **Note** If you use Netscaler VPX 10.5, the weight value is ignored.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the VPX load balancer service.
