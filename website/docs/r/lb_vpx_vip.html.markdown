---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lb_vpx_vip"
description: |-
  Manages IBM VPX load balancer virtual IP addresses.
---

# ibm_lb_vpx_vip
Create, update, and delete a VPX load balancer virtual IP addresses. For more information, about VPC load balancer virtual IP address, see [about Citrix Netscaler VPX](https://cloud.ibm.com/docs/citrix-netscaler-vpx?topic=citrix-netscaler-vpx-about-citrix-netscaler-vpx).

**Note** 

If you use Netscaler VPX 10.5, Terraform uses Netscaler's [NITRO REST API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) to manage the resource. Terraform can only access the NITRO API in the IBM Cloud Classic Infrastructure (SoftLayer) private network, so connect to the private network when running  Terraform. You can also use the [SSL VPN](https://www.ibm.com/cloud/vpn-access) to access a private network connection.

## Example usage
The following example configuration supports NetScaler VPX 10.1 and 10.5:

```terraform
    name = "test_load_balancer_vip"
    nad_controller_id = 1234567
    load_balancing_method = "lc"
    source_port = 80
    virtual_ip_address = "211.233.12.12"
    type = "HTTP"
}
```

The following example configuration supports only Netscaler VPX 10.5. More options for the `load_balancing_method` and `persistence` arguments are shown. A private IP address can be used for the `virtual_ip_address` argument.

```terraform
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

Netscaler VPX 10.5 also supports SSL offload. If you set the `type` argument to `SSL` and configure the `security_certificate_id` argument, then the `virtual_ip_address` argument provides the `HTTPS` protocol. The following example shows an SSL-offload configuration:

```terraform
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
    nad_controller_id = ibm_lb_vpx.test.id
    load_balancing_method = "rr"
    source_port = 443
# SSL type provides SSL offload
    type = "SSL"
    virtual_ip_address = ibm_lb_vpx.test.vip_pool[0]
# Use a security certificate in the SoftLayer portal
    security_certificate_id = 80347
}

resource "ibm_lb_vpx_service" "testacc_service1" {
  name = "test_load_balancer_service1"
  vip_id = ibm_lb_vpx_vip.test_vip1.id
# 10.6.218.166 should provides HTTP service with port 80
  destination_ip_address = "10.66.218.166"
  destination_port = 80
  weight = 100
  connection_limit = 4294967294
  health_check = "ICMP"
}
```


## Argument reference 
Review the argument references that you can specify for your resource. 

- `load_balancing_method` - (Required, String) See [IBM Cloud Classic Infrastructure (SoftLayer) API documentation](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_LoadBalancer_VirtualIpAddress) for available methods. If you use Netscaler VPX 10.5, see [Citrix documentation](https://docs.citrix.com/en-us/netscaler/10-5/ns-tmg-wrapper-10-con/ns-lb-wrapper-con-10/ns-lb-customizing-lbalgorithms-wrapper-con.html) for more methods that you can use.
- `name`- (Required, Forces new resource, String) The ID of the VPX load balancer virtual IP address.
- `nad_controller_id` - (Required, Integer) The ID of the VPX load balancer that the virtual IP address is assigned to.
- `persistence` -  (Optional, String) Applies to Netscaler VPX 10.5 only. See the available persistence types in the [Citrix documentation](https://docs.citrix.com/en-us/netscaler/10-5/ns-tmg-wrapper-10-con/ns-lb-wrapper-con-10/ns-lb-persistence-wrapper-con/ns-lb-persistence-about-con.html).
- `security_certificate_id` - (Optional, Forces new resource, Integer)Applies to Netscaler VPX 10.5 only. The ID of a security certificate that you want to use. This argument provides security certification for SSL offload services. For more information, see [`ibm_compute_ssl_certificate resource`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/compute_ssl_certificate).
- `source_port` - (Required, Integer)  - The source port for the VPX load balancer virtual IP address.
- `tags`- (Optional, Array of string)  Tags associated with the VPX load balancer virtual IP instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `type` - (Required, Forces new resource, String)The connection type for the VPX load balancer virtual IP address. Accepted values are `HTTP`, `FTP`, `TCP`, `UDP`, `DNS`, and `SSL`. If you set the type to `SSL`, then `security_certificate_id` provides certification for SSL offload services.
- `virtual_ip_address`- (Required, String) The public IP address for the VPX load balancer virtual IP.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the VPX load balancer virtual IP.
