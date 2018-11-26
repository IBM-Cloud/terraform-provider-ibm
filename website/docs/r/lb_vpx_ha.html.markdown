---
layout: "ibm"
page_title: "IBM : lb_vpx_ha"
sidebar_current: "docs-ibm-resource-lb-vpx-ha"
description: |-
  Configure a high availability pair with two NetScaler VPX devices
---

# ibm\_lb_vpx_ha

Configure a high availability (HA) pair with two NetScaler VPX devices. The two NetScaler VPXs must be version 10.5 and located in the same subnet. A primary NetScaler VPX provides load balancing services in active mode, and a secondary NetScaler VPX provides load balancing services when the primary NetScaler VPX fails. For additional details, refer to the  [Citrix support docs](https://support.citrix.com/article/CTX116748) and the [KnowledgeLayer NetScaler docs](http://knowledgelayer.softlayer.com/articles/netscaler-vpx-10-high-availability-setup).

**NOTE**: This resource only supports NetScaler VPX 10.5. The [NITRO API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) is used to configure HA. Terraform can only access the NITRO API in the IBM Cloud Infrastructure (SoftLayer) private network, so connect to the private network when running Terraform. You can also use the [SSL VPN](http://www.softlayer.com/VPN-Access) to access a private network connection.

The two NetScaler VPXs use the same password in HA mode. When you create this resource, Terraform changes the password of the secondary NetScaler VPX to the password of the primary NetScaler VPX. When you destroy this resource, Terraform restores the original password of the secondary NetScaler VPX.

## Example Usage

```hcl
# Create a primary NetScaler VPX
resource "ibm_lb_vpx" "test_pri" {
    datacenter = "lon02"
    speed = 10
    version = "10.5"
    plan = "Standard"
    ip_count = 2
}

# Create a secondary NetScaler VPX in the same subnets
resource "ibm_lb_vpx" "test_sec" {
    datacenter = "lon02"
    speed = 10
    version = "10.5"
    plan = "Standard"
    ip_count = 2
    public_vlan_id = "${ibm_lb_vpx.test_pri.public_vlan_id}"
    private_vlan_id = "${ibm_lb_vpx.test_pri.private_vlan_id}"
    public_subnet = "${ibm_lb_vpx.test_pri.public_subnet}"
    private_subnet = "${ibm_lb_vpx.test_pri.private_subnet}"
}

# Configure high availability with the primary and secondary NetScaler VPXs
resource "ibm_lb_vpx_ha" "test_ha" {
    primary_id = "${ibm_lb_vpx.test_pri.id}"
    secondary_id = "${ibm_lb_vpx.test_sec.id}"
    stay_secondary = false
}
```

## Argument Reference

The following arguments are supported:

* `primary_id` - (Required, string) The ID of the primary NetScaler VPX.
* `secondary_id` - (Required, string) The ID of the secondary NetScaler VPX.
* `stay_secondary` - (Optional, boolean) Specifies whether the secondary NetScaler VPX will  take over the service. Set this argument to `true` to prevent the secondary NetScaler VPX from taking over the service even if the primary NetScaler VPX fails. For additional details, see the [Citrix NetScaler docs](https://docs.citrix.com/en-us/netscaler/10-5/ns-system-wrapper-10-con/ns-nw-ha-intro-wrppr-con/ns-nw-ha-frcng-scndry-nd-sty-scndry-tsk.html) and the [Citrix support docs](https://support.citrix.com/article/CTX116748). The default value is `false`.
* `tags` - (Optional, array of strings) Tags associated with the high availability NetScale VPX pair instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the high availability NetScale VPX pair.
