---
layout: "ibm"
page_title: "IBM: lb_vpx"
sidebar_current: "docs-ibm-resource-lb-vpx"
description: |-
  Manages IBM VPX Load Balancer.
---

# ibm\_lb_vpx

Provides a resource for VPX load balancers. This allows VPX load balancers to be created, updated, and deleted.

**NOTE**: IBM VPX load balancers consist of Citrix NetScaler VPX devices (virtual), which are currently priced on a per-month basis. Use caution when creating the resource because the cost for an entire month is incurred immediately upon creation. For more information about pricing, see the [network appliance docs](http://www.softlayer.com/network-appliances). Under the Citrix log, click **See more pricing** for a current price matrix.

You can also use the following REST URL to get a listing of VPX choices along with version numbers, speed, and plan type:

```
https://<userName>:<apiKey>@api.softlayer.com/rest/v3/SoftLayer_Product_Package/192/getItems.json?objectMask=id;capacity;description;units;keyName;prices.id;prices.categories.id;prices.categories.name
```

## Example Usage

Review the [IBM Cloud Infrastructure (SoftLayer) docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller) for more information.

```hcl
resource "ibm_lb_vpx" "test_vpx" {
    datacenter = "dal06"
    speed = 10
    version = "10.1"
    plan = "Standard"
    ip_count = 2
    public_vlan_id = 1251234
    private_vlan_id = 1540786
    public_subnet = "23.246.226.248/29"
    private_subnet = "10.107.180.0/26"
}
```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Required, string) The data center in which you want to provision the VPX load balancer. You can find accepted values in the [data center docs](http://www.softlayer.com/data-centers).
* `speed` - (Required, integer) The speed, expressed in Mbps. Accepted values are `10`, `200`, and `1000`.
* `version` - (Required, string) The VPX load balancer version. Accepted values are `10.1`, `10.5`, `11.0` and `11.1`.
* `plan` - (Required, string) The VPX load balancer plan. Accepted values are `Standard` and `Platinum`.
* `ip_count` - (Required, integer) The number of static public IP addresses assigned to the VPX load balancer. Accepted values are `1`,`2`, `4`, `8`, and `16`.
* `public_vlan_id` - (Optional, integer) The public VLAN ID that is used for the public network interface of the VPX load balancer. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans) by clicking the desired VLAN and noting the ID in the resulting URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).
* `private_vlan_id` - (Optional, integer) The private VLAN ID that is used for the private network interface of the VPX load balancer. You can find accepted values in the [VLAN docs](https://control.softlayer.com/network/vlans) by clicking the desired VLAN and noting the ID in the resulting URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).
* `public_subnet` - (Optional, string) The public subnet that is used for the public network interface of the VPX load balancer. Accepted values are primary public networks. You can find accepted values in the [subnet docs](https://control.softlayer.com/network/subnets).
* `private_subnet` - (Optional, string) Public subnet that is used for the private network interface of the VPX load balancer. Accepted values are primary private networks. You can find accepted values in the [subnet docs](https://control.softlayer.com/network/subnets).
* `tags` - (Optional, array of strings) Tags associated with the VPX load balancer instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The internal identifier of a VPX load balancer
* `name` - The internal name of a VPX load balancer.
* `vip_pool` - A list of virtual IP addresses for the VPX load balancer.
* `management_ip_address` - The private address of the VPX UI.
