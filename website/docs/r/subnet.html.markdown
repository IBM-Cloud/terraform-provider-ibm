---
layout: "ibm"
page_title: "IBM: subnet"
sidebar_current: "docs-ibm-resource-subnet"
description: |-
  Manages IBM Subnet.
---

# ibm\_subnet

This resource provides portable and static subnets that consist of either IPv4 and IPv6 addresses. Users are able to create 
public portable subnets, private portable subnets, and public static subnets with an IPv4 option, and public portable subnets and public static subnets with an IPv6 option. 
 
The portable IPv4 subnet is created as a secondary subnet on a VLAN. IP addresses in the portable subnet can be assigned as secondary IP 
addresses for IBM resources in the VLAN. Because each portable subnet has a default gateway IP address, network IP address, and broadcast IP address, the number of usable IP addresses is `capacity` - 3. A `capacity` of 4 means that the number of usable IP addresses is 1; a `capacity` of 8 means that the number of usable IP addresses is 5. For example, consider a portable subnet of `10.0.0.0/30` that has `10.0.0.1` as a default gateway IP address, `10.0.0.0` as a network IP address, and `10.0.0.3` as a broadcast IP address. Only `10.0.0.2` can be assigned to IBM resources as a secondary IP address. For additional details, refer to [Static and Portable IP blocks](https://knowledgelayer.softlayer.com/articles/static-and-portable-ip-blocks).

The static IPv4 subnet provides secondary IP addresses for primary IP addresses. It provides secondary IP addresses for IBM resources such as virtual servers, bare metal servers, and netscaler VPXs. Consider a virtual server that requires secondary IP addresses. Users can create a static subnet on the public IP address of the virtual server. Unlike the portable subnet, the number of usable IP addresses for the stactic subnet is the same as the value of `capacity`. For example, when a static subnet of `10.0.0.0/30` has a `capacity` of 4, then four IP addresses (10.0.0.0 - 10.0.0.3) can be used as secondary IP addresses. For additional details, refer to [Subnet](https://knowledgelayer.softlayer.com/topic/subnets).

Both the public portable IPv6 subnet and the public static IP only accept `64` as a value for the `capacity` attribute. They provide 2^64 IP addresses. For additional detail, refer to [IPv6 address](http://blog.softlayer.com/tag/ipv6)

##### Example Usage of portable subnet
The following example creates a private portable subnet which has one available IPv4 address:

```hcl
resource "ibm_subnet" "portable_subnet" {
  type = "Portable"
  private = true
  ip_version = 4
  capacity = 4
  vlan_id = 1234567
  notes = "portable_subnet"
}
```

##### Example Usage of static subnet
The following example creates a public static subnet which has four available IPv4 address:

```hcl
resource "ibm_subnet" "static_subnet" {
  type = "Static"
  private = false
  ip_version = 4
  capacity = 4
  endpoint_ip="151.1.1.1"
  notes = "static_subnet_updated"
}
```

Users can use Terraform built-in functions to get IP addresses from `subnet`. The following example returns the first IP address in the subnet `test`:

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

* `private` - (Optional, boolean) Specifies whether the network is public or private.
* `type` - (Required, string) The type of the subnet. Accepted values are `portable` and `static`.
* `ip_version` - (Optional, integer) The IP version of the subnet. Accepted values are 4 and 6.
* `capacity` - (Required, integer) The size of the subnet.
    * Accepted values for a public portable IPv4 subnet are 4, 8, 16, and 32.
    * Accepted values for a private portable IPv4 subnet are 4, 8, 16, 32, and 64.
    * Accepted values for a public static IPv4 subnet are 1, 2, 4, 8, 16, and 32.
    * Accepted value for a public portable IPv6 subnet is 64. A /64 block is created and 2^64 IP addresses are provided.
    * Accepted value for a public static IPv6 subnet is 64. A /64 block is created and 2^64 IP addresses are provided.
* `vlan_id` - (Optional, integer) The VLAN ID for portable subnet. You can configure both public and private VLAN ID. You can find accepted values in the [Softlayer VLAN documentation](https://control.softlayer.com/network/vlans) by clicking on the desired VLAN and noting the ID in the resulting URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html.markdown).
* `endpoint_ip` - (Optional, string) The target primary IP address for a static subnet. Only public IP addresses of virtual servers, bare metal servers, and netscaler VPXs can be configured as an `endpoint_ip`. The `static subnet` will be created on the VLAN where the `endpoint_ip` is located.
* `notes` - (Optional, string) Descriptive text or comments about the subnet.
* `tags` - (Optional, array of strings) Tags associated with the subnet instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

##### Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the subnet.
* `subnet_cidr` - The IP address/cidr format (ex. 10.10.10.10/28), which you can use to get an available IP address in `subnet`.
