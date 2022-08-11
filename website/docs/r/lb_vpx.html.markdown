---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: lb_vpx"
description: |-
  Manages IBM VPX load balancer.
---

# ibm_lb_vpx
Create, delete, and update a [VPX load balancers](https://cloud.ibm.com/docs/citrix-netscaler-vpx?topic=citrix-netscaler-vpx-setting-up-citrix-netscaler-vpx-for-high-availability-ha-).

**Note**

IBM VPX load balancers consist of `Citrix NetscalerVPX` devices (virtual), which are currently priced on a per-month basis. Use caution when creating the resource because the cost for an entire month is incurred immediately upon creation. For more information, about pricing, see the [estimating your cost](https://cloud.ibm.com/docs/billing-usage?topic=billing-usage-cost). 

You can also use the following REST URL to get a listing of VPX choices along with version numbers, speed, and plan type:

```
https://<userName>:<apiKey>@api.softlayer.com/rest/v3/SoftLayer_Product_Package/192/getItems.json?objectMask=id;capacity;description;units;keyName;prices.id;prices.categories.id;prices.categories.name
```

## Example usage
Review the [IBM Cloud Classic Infrastructure documentation](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Application_Delivery_Controller) for more information.

```terraform
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

## Argument reference 
Review the argument references that you can specify for your resource. 

- `datacenter`- (Required, Forces new resource, String) The data center in which you want to provision the VPX load balancer. You can find accepted values in the [data center documentation](https://www.ibm.com/cloud/load-balancer).
- `ip_count` - (Required, Forces new resource, Integer) The number of static public IP addresses assigned to the VPX load balancer. Accepted values are `1`,`2`, `4`, `8`, and `16`.
- `plan`- (Required, Forces new resource, String) The VPX load balancer plan. Accepted values are `Standard` and `Platinum`.
- `public_vlan_id` - (Optional, Forces new resource,  Integer) The public VLAN ID that is used for the public network interface of the VPX load balancer. You can find accepted values in the [VLAN network](https://cloud.ibm.com/classic/network/vlans) by clicking the VLAN that you want to use and noting the ID in the resulting URL. You can also refer to a VLAN name by using a data source.
- `private_vlan_id` - (Optional, Forces new resource,  Integer) The private VLAN ID that is used for the private network interface of the VPX load balancer. You can find accepted values in the [VLAN network](https://cloud.ibm.com/classic/network/vlans) by clicking the VLAN that you want to use and noting the ID in the resulting URL. You can also refer to a VLAN name by using a data source.
- `public_subnet` - (Optional, Forces new resource, String)The public subnet that is used for the public network interface of the VPX load balancer. Accepted values are primary public networks. You can find accepted values in the [subnet documentation](https://cloud.ibm.com/classic/network/subnets).
- `private_subnet`- (Optiona, Forces new resource, String) Public subnet that is used for the private network interface of the VPX load balancer. Accepted values are primary private networks. You can find accepted values in the [subnet documentation](https://cloud.ibm.com/classic/network/subnets).
- `speed` - (Required, Forces new resource,  Integer) The speed, expressed in Mbps. Accepted values are `10`, `200`, and `1000`.
- `tags`- (Optional, Array of Strings) Tags associated with the VPX load balancer instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `version` - (Required, Forces new resource, String)The VPX load balancer version. Accepted values are `10.1`, `10.5`, `11.0`, `11.1` and `12.1`.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The internal identifier of a VPX load balancer.
- `management_ip_address`- (String) The private address of the VPX console.
- `name`- (String) The internal name of a VPX load balancer.
- `vip_pool`- (String) A list of virtual IP addresses for the VPX load balancer.
