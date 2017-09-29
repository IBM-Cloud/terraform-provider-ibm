---
layout: "ibm"
page_title: "IBM: subnet"
sidebar_current: "docs-ibm-resource-subnet"
description: |-
  Manages IBM Subnet.
---

# ibm\_subnet

 This resource provides portable and static subnets that consist of either IPv4 and IPv6 addresses. Users are able to create 
public portable subnets, private portable subnets, and public static subnets with an IPv4 option and public portable subnets and public static 
subnets with an IPv6 option. 
 
The portable IPv4 subnet is created as a seconday subnet on a VLAN. IP addresses in the portable subnet can be assigned as secondary IP 
 addresses for IBM resources in the VLAN. Each portable subnet has a default gateway IP address, network IP address, and broadcast 
 IP address. For example, if a portable subnet is `10.0.0.0/30`, `10.0.0.0` is a network IP address, `10.0.0.1` is a default gateway IP address, 
 and `10.0.0.3` is a broadcast IP address. Therefore, only `10.0.0.2` can be assigned to IBM resources as a secondary IP address. 
 Number of usuable IP addresses is `capacity` - 3. If `capacity` is 4, the number of usuable IP addresses is 4 - 3 = 1. If `capacity` is 8, the 
 number of usuable IP addresses is 8 - 3 = 5. For additional details, refer to [Static and Portable IP blocks](https://knowledgelayer.softlayer.com/articles/static-and-portable-ip-blocks).

The static IPv4 subnet provides secondary IP addresses for primary IP addresses. It provides secondary IP addresses for IBM resources such as 
virtual servers, bare metal servers, and netscaler VPXs. Suppose that a virtual server requires secondary IP addresses. Then, users can create 
a static subnet on the public IP address of the virtual server. Unlike the portable subnet, `capacity` is same with a number of usuable IP address. 
For example, if a static subnet is `10.0.0.0/30`, `capacity` is 4 and four IP addresses(10.0.0.0 ~ 10.0.0.3) can be used as secondary IP addresses. 
For additional details, refer to [Subnet](https://knowledgelayer.softlayer.com/topic/subnets).

Both the public portable IPv6 subnet and the public static IP only accept `64` as a value of `capacity` attribute. They provide 2^64 IP addresses. For additional detail, refer to [IPv6 address](http://blog.softlayer.com/tag/ipv6)

The following example will create a private portable subnet which has one available IPv4 address. 
##### Example Usage of portable subnet

```hcl
# Create a new portable subnet
resource "ibm_subnet" "portable_subnet" {
  type = "Portable"
  private = true
  ip_version = 4
  capacity = 4
  vlan_id = 1234567
  notes = "portable_subnet"
}
```

The following example will create a public static subnet which has four available IPv4 address.
##### Example Usage of static subnet

```hcl
# Create a new static subnet
resource "ibm_subnet" "static_subnet" {
  type = "Static"
  private = false
  ip_version = 4
  capacity = 4
  endpoint_ip="151.1.1.1"
  notes = "static_subnet_updated"
}
```

Sometimes, users need to get IP addresses on a subnet. Terraform built-in functions can be used to get IP addresses from `subnet`. 
The following example returns first IP address in the subnet `test`:

```hcl
resource "ibm_subnet" "test" {
  type = "Static"
  private = false
  ip_version = 4
  capacity = 4
  endpoint_ip="159.8.181.82"
}

# Use a built-in function cidrhost with index 0.
output "first_ip_address" {
  value = "${cidrhost(ibm_subnet.test.subnet,0)}"
}

```

##### Argument Reference

The following arguments are supported:

* `private` - (Optional,boolean) Set the network property of the subnet if it is public or private.
* `type` - (Required,string) Set the type of the subnet. Accepted values are Portable and Static.
* `ip_version` - (Optional,integer) Set the IP version of the subnet. Accepted values are 4 and 6. *Default*: true
* `capacity` - (Required,integer)
    * Set the size of the subnet.
    * Accepted values for a public portable IPv4 subnet are 4, 8, 16, and 32.
    * Accepted values for a private portable IPv4 subnet are 4, 8, 16, 32, and 64.
    * Accepted values for a public static IPv4 subnet are 1, 2, 4, 8, 16, and 32.
    * Accepted value for a public portable IPv6 subnet is 64. /64 block is created and 2^64 IP addresses are provided.
    * Accepted value for a public static IPv6 subnet is 64.  /64 block is created and 2^64 IP addresses are provided.
* `vlan_id` - (Optional,integer)
    * VLAN id for portable subnet. It should be configured when the subnet is a portable subnet. Both public VLAN ID and private VLAN ID can 
    be configured. Accepted values can be found [here](https://control.softlayer.com/network/vlans). Click on the desired VLAN and note the 
    ID on the resulting URL. Or, you can also [refer to a VLAN by name using a data source](https://github.com/IBM-Bluemix/terraform-provider-ibm/blob/master/website/docs/d/network_vlan.html.markdown). 
* `endpoint_ip` - (Optional,string)
    * Target primary IP address for static subnet. It should be configured when the subnet is a static subnet. Only public IP address can be 
    configured as a `endpoint_ip`. It can be public IP address of virtual servers, bare metal servers, and netscaler VPXs. `static subnet` will 
    be created on VLAN where `endpoint_ip` is located in.
* `notes` | - (Optional,string)
    * Set comments for the subnet.

##### Attributes Reference

The following attributes are exported:

* `id` - id of the subnet.
* `subnet_cidr` - It provides IP address/cidr format (ex. 10.10.10.10/28). It can be used to get an available IP address in `subnet`. 