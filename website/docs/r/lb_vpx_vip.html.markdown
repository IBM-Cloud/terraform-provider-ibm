---
layout: "ibm"
page_title: "IBM : lb_vpx_vip"
sidebar_current: "docs-ibm-resource-lb-vpx-vip"
description: |-
  Manages IBM VPX load balancer virtual IP addresses.
---

# ibm\_lb_vpx_vip

Create, update, and delete a VPX load balancer virtual IP address.

**NOTE**: If NetScaler VPX 10.5 is used, Terraform uses NetScaler's [NITRO REST API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html)) to manage the resource. The NITRO API is only accessible in the Bluemix Infrastructure (SoftLayer) private network, so you need to be connected to the private network when running Terraform. You can also use the [SSL VPN](http://www.softlayer.com/VPN-Access) to access a private network connection.

## Example Usage

The following configuration supports NetScaler VPX 10.1 and 10.5

```hcl
resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = 1234567
    load_balancing_method = "lc"
    source_port = 80
    virtual_ip_address = "211.233.12.12"
    type = "HTTP"
}
```

NetScaler VPX 10.5 provides additional options for `load_balancing_method` and `persistence`. A private IP address can be used as a `virtual_ip_address`.

```hcl
resource "ibm_lb_vpx_vip" "testacc_vip" {
    name = "test_load_balancer_vip"
    nad_controller_id = "1234567"
    load_balancing_method = "DESTINATIONIPHASH"
    persistence = "SOURCEIP"
    source_port = 80
    virtual_ip_address = "10.10.2.2"
    type = "HTTP"
}
```

NetScaler VPX 10.5 supports SSL offload. If `type` is set to SSL and the `security_certificate_id` is configured, the `virtual_ip_address` provides `HTTPS` protocol. The following examples describe an SSL-offload configuration.

```hcl
# Create a NetScaler VPX 10.5
resource "ibm_lb_vpx" "test" {
    datacenter = "lon02"
    speed = 10
    version = "10.5"
    plan = "Standard"
    ip_count = 2
}

resource "ibm_lb_vpx_vip" "test_vip1" {
    name = "test_vip1"
    nad_controller_id = "${ibm_lb_vpx.test.id}"
    load_balancing_method = "rr"
    source_port = 443
# SSL type provides SSL offload
    type = "SSL"
    virtual_ip_address = "${ibm_lb_vpx.test.vip_pool[0]}"
# Use a security certificate in the SoftLayer portal
    security_certificate_id = 80347
}

resource "ibm_lb_vpx_service" "testacc_service1" {
  name = "test_load_balancer_service1"
  vip_id = "${ibm_lb_vpx_vip.test_vip1.id}"
# 10.6.218.166 should provides HTTP service with port 80
  destination_ip_address = "10.66.218.166"
  destination_port = 80
  weight = 100
  connection_limit = 4294967294
  health_check = "ICMP"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The unique identifier for the VPX load balancer virtual IP address.
* `nad_controller_id` - (Required, integer) The ID of the VPX load balancer that the virtual IP address is assigned to.
* `load_balancing_method` - (Required, string) See the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_VirtualIpAddress) for available methods. If NetScaler VPX 10.5 is used, see the [Citrix docs](https://docs.citrix.com/en-us/netscaler/10-5/ns-tmg-wrapper-10-con/ns-lb-wrapper-con-10/ns-lb-customizing-lbalgorithms-wrapper-con.html) for additional methods that can be used. 
* `persistence` - (Optional, string) Applies to NetScaler VPX 10.5 only. See the available persistence types in the [Citrix docs](https://docs.citrix.com/en-us/netscaler/10-5/ns-tmg-wrapper-10-con/ns-lb-wrapper-con-10/ns-lb-persistence-wrapper-con/ns-lb-persistence-about-con.html).  
* `virtual_ip_address` - (Required, string) The public facing IP address for the VPX load balancer virtual IP.
* `source_port` - (Required, integer) The source port for the VPX load balancer virtual IP address.
* `type` - (Required, string) The connection type for the VPX load balancer virtual IP address. Accepted values are `HTTP`, `FTP`, `TCP`, `UDP`, `DNS`, and `SSL`. If SSL is configured, `security_certificate_id` is used as certification for SSL offload services.
* `security_certificate_id` - (Optional, integer) Applies to NetScaler VPX 10.5 only. Provides security certification for SSL offload. For additional information, see the  [ibm_compute_ssl_certificate](compute_ssl_certificate.html) resource.
* `tags` - (Optional, array of strings) Set tags on the VPX load balancer virtual IP instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the VPX load balancer virtual IP.
