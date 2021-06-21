---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: subnet"
description: |-
  Manages IBM Subnet.
---

# ibm_subnet
This resource provides portable and static subnets that consist of either IPv4 and IPv6 addresses. Users are able to create 
public portable subnets, private portable subnets, and public static subnets with an IPv4 option, and public portable subnets and public static subnets with an IPv6 option. 
 
The portable IPv4 subnet is created as a secondary subnet on a VLAN. IP addresses in the portable subnet can be assigned as secondary IP addresses for IBM resources in the VLAN. Because each portable subnet has a default gateway IP address, network IP address, and broadcast IP address, the number of usable IP addresses is `capacity` - 3. A `capacity` of 4 means that the number of usable IP addresses is 1; a `capacity` of 8 means that the number of usable IP addresses is 5. For example, consider a portable subnet of `10.0.0.0/30` that has `10.0.0.1` as a default gateway IP address, `10.0.0.0` as a network IP address, and `10.0.0.3` as a broadcast IP address. Only `10.0.0.2` can be assigned to IBM resources as a secondary IP address. 

The static IPv4 subnet provides secondary IP addresses for primary IP addresses. It provides secondary IP addresses for IBM resources such as virtual servers, Bare Metal servers, and `NetscalerVPXs`. Consider a virtual server that requires secondary IP addresses. Users can create a static subnet on the public IP address of the virtual server. Unlike the portable subnet, the number of usable IP addresses for the static subnet is the same as the value of `capacity`. For example, when a static subnet of `10.0.0.0/30` has a `capacity` of 4, then four IP addresses (10.0.0.0 - 10.0.0.3) can be used as secondary IP addresses. 

Both the public portable IPv6 subnet and the public static IP only accept `64` as a value for the `capacity` attribute. They provide 2^64 IP addresses. For more information, see [about subnets and IPs](https://cloud.ibm.com/docs/subnets?topic=subnets-about-subnets-and-ips).

## Example usage
The following example creates a private portable subnet which has one available IPv4 address:

```terraform
resource "ibm_subnet" "portable_subnet" {
  type       = "Portable"
  private    = true
  ip_version = 4
  capacity   = 4
  vlan_id    = 1234567
  notes      = "portable_subnet"

  //User can increase timeouts
  timeouts {
    create = "45m"
  }
}
```

Users can use Terraform built-in functions to get IP addresses from `portable subnet`. The following example returns the first usable IP address of the portable subnet `test`.:

```terraform
resource "ibm_subnet" "test" {
  type = "Portable"
  private = true
  ip_version = 4
  capacity = 4
  vlan_id = 1234567
}

# Use a built-in function cidrhost with index 1.

output "first_ip_address" {
  value = cidrhost(ibm_subnet.test.subnet_cidr,1)
}

```

### Example usage of static subnet
The following example creates a public static subnet which has four available IPv4 address:

```terraform
resource "ibm_subnet" "static_subnet" {
  type = "Static"
  private = false
  ip_version = 4
  capacity = 4
  endpoint_ip="151.1.1.1"
  notes = "static_subnet_updated"
}
```

Users can use Terraform built-in functions to get IP addresses from `subnet`. The following example returns the first usable IP address in the static subnet `test`:

```terraform
resource "ibm_subnet" "test" {
  type        = "Static"
  private     = false
  ip_version  = 4
  capacity    = 4
  endpoint_ip = "159.8.181.82"
}

# Use a built-in function cidrhost with index 0.
output "first_ip_address" {
  value = cidrhost(ibm_subnet.test.subnet_cidr, 0)
}

```

## Timeouts

The `ibm_subnet` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating instance.


## Argument reference
Review the argument references that you can specify for your resource.

- `capacity` - (Required, Forces new resource,Integer) The size of the subnet. <ul><li>Accepted values for a public portable IPv4 subnet are 4, 8, 16, and 32. </li><li> Accepted values for a private portable IPv4 subnet are 4, 8, 16, 32, and 64. </li><li>Accepted values for a public static IPv4 subnet are 1, 2, 4, 8, 16, and 32. </li><li>Accepted value for a public portable IPv6 subnet is 64. A /64 block is created and 2^64 IP addresses are provided. </li><li>Accepted value for a public static IPv6 subnet is 64. A /64 block is created and 2^64 IP addresses are provided.</li></ul>.
- `endpoint_ip` - (Optional, Forces new resource, String) The target primary IP address for a static subnet. Only public IP addresses of virtual servers, Bare Metal servers, and `NetscalerVPXs` can be configured as an `endpoint_ip`. The `static subnet` will be created on the VLAN where the `endpoint_ip` is located.
- `ip_version` - (Optional, Forces new resource, Integer)The IP version of the subnet. Accepted values are 4 and 6.
- `notes`- (Optional, String) Descriptive text or comments about the subnet.
- `private` -  (Optional, Forces new resource, Bool) Specifies whether the network is public or private.
- `tags`- (Optional, Array of Strings) Tags associated with the subnet instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
- `type` - (Required, Forces new resource, String) The type of the subnet. Accepted values are `portable` and `static`.
- `vlan_id` - (Optional, Forces new resource, Integer) The VLAN ID for portable subnet. You can configure both public and private VLAN ID. You can find accepted values in the [SoftLayer VLAN documentation](https://cloud.ibm.com/classic/network/vlans) by clicking the VLAN that you want and noting the ID in the resulting URL. You can also refer to a [`ibm_network_vlan`](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/network_vlan) data source.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id`- (String) The unique identifier of the subnet.
- `subnet_cidr`- (String) The IP address / CIDR format (For example, 10.10.10.10/28), which you can use to get an available IP address in `subnet`.
