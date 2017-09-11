---
layout: "ibm"
page_title: "IBM : lb_vpx_service"
sidebar_current: "docs-ibm-resource-lb-vpx-service"
description: |-
  Manages IBM VPX load balancer services
---

# ibm\_lb_vpx_service

Create, update, and delete VPX load balancer services. For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service).

**NOTE**: If NetScaler VPX 10.5 is used, Terraform uses NetScaler's [NITRO REST API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html)) to manage the resource. The NITRO API is only accessible in the Bluemix Infrastructure (SoftLayer) private network, so you need to be connected to the private network when running Terraform. You can also use the [SSL VPN](http://www.softlayer.com/VPN-Access) to access a private network connection.
 
## Example Usage

Create a VPX load balancer:

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

* `name` - (Required, string) The unique identifier for the VPX load balancer service.
* `vip_id` - (Required, string) The ID of the VPX load balancer virtual IP address that the service is assigned to.
* `destination_ip_address` - (Required, string) The IP address of the server that traffic is directed to. If NetScaler VPX 10.1 is used, you need to indicate a public IP address in a Bluemix Infrastructure (SoftLayer) account. If NetScaler VPX 10.5 is used, you can use any IP address.
* `destination_port` - (Required, integer) The destination port of the server that traffic is be directed to.
* `weight` - (Required, integer) The percentage of the total connection limit allocated to the load balancer between all your services. See the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service) for details. 
    
    **NOTE**: In VPX 10.5, the weight value is ignored. 
* `connection_limit` - (Required, integer) The connection limit for this service. Acceptable values are `0` ~ `4294967294`. See the [Citrix NetScaler docs](https://docs.citrix.com/en-us/netscaler/11/reference/netscaler-command-reference/basic/service.html) for details.
* `health_check` - (Required, string) Set the health check type. See the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_Service) for details.
* `usip` - (Optional, String) Configures the service to report the source ip of the client to the service being load balanced. Acceptable values are "YES" or "NO". Default is "NO". See the [Citrix NetScaler docs](https://developer-docs.citrix.com/projects/netscaler-nitro-api/en/11.0/configuration/basic/service/service) for details.

    **NOTE**: Only available for VPX 10.5

* `tags` - (Optional, array of strings) Set tags on the VPX load balancer service instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the VPX load balancer service.
