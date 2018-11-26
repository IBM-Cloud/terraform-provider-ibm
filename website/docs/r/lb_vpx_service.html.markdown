---
layout: "ibm"
page_title: "IBM : lb_vpx_service"
sidebar_current: "docs-ibm-resource-lb-vpx-service"
description: |-
  Manages IBM VPX load balancer services
---

# ibm\_lb_vpx_service

Provides a resource for VPX load balancer services. This allows VPX load balancer services to be created, updated, and deleted. For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service).  

**NOTE**: If you use NetScaler VPX 10.5, Terraform uses NetScaler's [NITRO REST API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) to manage the resource.  Terraform can only access the NITRO API in the IBM Cloud Infrastructure (SoftLayer) private network, so connect to the private network when running Terraform. You can also use the [SSL VPN](http://www.softlayer.com/VPN-Access) to access a private network connection.

## Example Usage

In the following example, you can create a VPX load balancer:

```hcl
resource "ibm_lb_vpx_service" "test_service" {
  name = "test_load_balancer_service"
  vip_id = "${ibm_lb_vpx_vip.testacc_vip.id}"
  destination_ip_address = "${ibm_compute_vm_instance.test_server.ipv4_address}"
  destination_port = 80
  weight = 55
  connection_limit = 5000
  health_check = "HTTP"
  usip = "NO"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The ID of the VPX load balancer service.
* `vip_id` - (Required, string) The ID of the VPX load balancer virtual IP address to which the service is assigned.
* `destination_ip_address` - (Required, string) The IP address of the server to which traffic directs. If you use NetScaler VPX 10.1, you must indicate a public IP address in an IBM Cloud Infrastructure (SoftLayer) account. If you use NetScaler VPX 10.5, you can use any IP address.
* `destination_port` - (Required, integer) The destination port of the server to which traffic directs.
* `weight` - (Required, integer) The percentage of the total connection limit allocated to the load balancer between all your services. See the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service) for details.  
    **NOTE**: If you use NetScaler VPX 10.5, the weight value is ignored.
* `connection_limit` - (Required, integer) The connection limit for this service. Acceptable values are `0` - `4294967294`. See the [Citrix NetScaler docs](https://docs.citrix.com/en-us/netscaler/11/reference/netscaler-command-reference/basic/service.html) for details.
* `health_check` - (Required, string) The health check type. See the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service) for details.
* `usip` - (Optional, string) Whether the service reports the source IP of the client to the service being load balanced. Acceptable values are `YES` or `NO`. The default value is `NO`. See the [Citrix NetScaler docs](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service/service) for more details.  
    **NOTE**: This argument is only available for VPX 10.5.
* `tags` - (Optional, array of strings) Tags associated with the VPX load balancer service instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPX load balancer service.
