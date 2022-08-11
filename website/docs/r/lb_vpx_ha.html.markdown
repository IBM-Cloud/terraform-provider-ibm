---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : lb_vpx_ha"
description: |-
  Configure a high availability pair with two NetScaler VPX devices.
---

# ibm_lb_vpx_ha
Create, update, and delete a VPX load balancer virtual IP addresses. For more information, about VPC load balancer virtual IP address, see [about Citrix Netscaler VPX](https://cloud.ibm.com/docs/citrix-netscaler-vpx?topic=citrix-netscaler-vpx-about-citrix-netscaler-vpx).

Configure a high availability (HA) pair with two `NetscalerVPX` devices. The two `NetscalerVPXs` must be version 10.5 and located in the same subnet. A primary `NetscalerVPX` provides load-balancing services in active mode, and a secondary `NetscalerVPX` provides load-balancing services when the primary `NetscalerVPX` fails. For more information, refer to the  [Citrix support documentation](https://support.citrix.com/article/CTX116748) and the [knowledge layer Netscaler documentation](https://cloud.ibm.com/docs/citrix-netscaler-vpx?topic=citrix-netscaler-vpx-setting-up-citrix-netscaler-vpx-for-high-availability-ha-).

**Note** 

This resource only supports Netscaler VPX 10.5. The [NITRO API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html) is used to configure HA.  Terraform can only access the NITRO API in the IBM Cloud Classic Infrastructure (SoftLayer) private network, so connect to the private network when running  Terraform. You can also use the [SSL VPN](https://www.ibm.com/cloud/vpn-access) to access a private network connection.

The two `NetscalerVPXs` use the same password in HA mode. When you create this resource, Terraform changes the password of the secondary Netscaler VPX to the password of the primary Netscaler VPX. When you destroy this resource, Terraform restores the original password of the secondary Netscaler VPX.

## Example usage

```terraform
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
    public_vlan_id = ibm_lb_vpx.test_pri.public_vlan_id
    private_vlan_id = ibm_lb_vpx.test_pri.private_vlan_id
    public_subnet = ibm_lb_vpx.test_pri.public_subnet
    private_subnet = ibm_lb_vpx.test_pri.private_subnet
}

# Configure high availability with the primary and secondary NetScaler VPXs
resource "ibm_lb_vpx_ha" "test_ha" {
    primary_id = ibm_lb_vpx.test_pri.id
    secondary_id = ibm_lb_vpx.test_sec.id
    stay_secondary = false
}
```

## Argument reference 
Review the argument references that you can specify for your resource. 

- `primary_id` - (Required, Forces new resource, String)The ID of the primary Netscaler VPX.
- `secondary_id` - (Required, Forces new resource, String)The ID of the secondary Netscaler VPX.
- `stay_secondary` - (Optional, Bool) Specifies whether the secondary Netscaler VPX will  take over the service. Set this argument to **true** to prevent the secondary Netscaler VPX from taking over the service even if the primary Netscaler VPX fails. For more information, see the [Citrix Netscaler documentation](https://docs.citrix.com/en-us/netscaler/10-5/ns-system-wrapper-10-con/ns-nw-ha-intro-wrppr-con/ns-nw-ha-frcng-scndry-nd-sty-scndry-tsk.html) and the [Citrix support documentation](https://support.citrix.com/article/CTX116748). The default value is **false**.
- `tags`- (Optional, Array of Strings) Tags associated with the high availability Netscaler VPX pair instance.  *Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the high availability Netscaler VPX pair.
