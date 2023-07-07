---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface"
description: |-
  Manages VirtualNetworkInterface.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_virtual_network_interface

Create, update, and delete VirtualNetworkInterfaces with this resource.

## Example Usage

```terraform
resource "ibm_is_virtual_network_interface" "is_virtual_network_interface_instance" {
  allow_ip_spoofing = true
  auto_delete = false
  enable_infrastructure_nat = true
  ips {
		address = "192.168.3.4"
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		name = "my-reserved-ip"
		resource_type = "subnet_reserved_ip"
  }
  name = "my-virtual-network-interface"
  primary_ip {
		address = "192.168.3.4"
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e/reserved_ips/6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		id = "6d353a0f-aeb1-4ae1-832e-1110d10981bb"
		name = "my-reserved-ip"
		resource_type = "subnet_reserved_ip"
  }
  resource_group {
		href = "https://resource-controller.cloud.ibm.com/v2/resource_groups/fee82deba12e4c0fb69c3b09d1f12345"
		id = "fee82deba12e4c0fb69c3b09d1f12345"
		name = "my-resource-group"
  }
  security_groups {
		crn = "crn:v1:bluemix:public:is:us-south:a/123456::security-group:be5df5ca-12a0-494b-907e-aa6ec2bfa271"
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/security_groups/be5df5ca-12a0-494b-907e-aa6ec2bfa271"
		id = "be5df5ca-12a0-494b-907e-aa6ec2bfa271"
		name = "my-security-group"
  }
  subnet {
		crn = "crn:v1:bluemix:public:is:us-south-1:a/123456::subnet:7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		id = "7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
		name = "my-subnet"
		resource_type = "subnet"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.
* `auto_delete` - (Optional, Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.
* `enable_infrastructure_nat` - (Optional, Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
* `ips` - (Optional, List) The reserved IPs bound to this virtual network interface.May be empty when `lifecycle_state` is `pending`.
  * Constraints: The minimum length is `0` items.
Nested schema for **ips**:
	* `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Required, String) The URL for this reserved IP.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The unique identifier for this reserved IP.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (Computed, String) The resource type.
	  * Constraints: Allowable values are: `subnet_reserved_ip`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `name` - (Optional, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `primary_ip` - (Optional, List) The reserved IP for this virtual network interface.
Nested schema for **primary_ip**:
	* `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Required, String) The URL for this reserved IP.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The unique identifier for this reserved IP.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (Computed, String) The resource type.
	  * Constraints: Allowable values are: `subnet_reserved_ip`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `resource_group` - (Optional, List) The resource group for this virtual network interface.
Nested schema for **resource_group**:
	* `href` - (Computed, String) The URL for this resource group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The unique identifier for this resource group.
	  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.
	* `name` - (Computed, String) The name for this resource group.
	  * Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`.
* `security_groups` - (Optional, List) The security groups for this virtual network interface.
  * Constraints: The minimum length is `1` item.
Nested schema for **security_groups**:
	* `crn` - (Required, String) The security group's CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Required, String) The security group's canonical URL.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The unique identifier for this security group.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Required, String) The name for this security group. The name is unique across all security groups for the VPC.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `subnet` - (Optional, List) The associated subnet.
Nested schema for **subnet**:
	* `crn` - (Required, String) The CRN for this subnet.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Required, String) The URL for this subnet.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (Required, String) The unique identifier for this subnet.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Computed, String) The name for this subnet. The name is unique across all subnets in the VPC.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (Computed, String) The resource type.
	  * Constraints: Allowable values are: `subnet`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the VirtualNetworkInterface.
* `created_at` - (String) The date and time that the virtual network interface was created.
* `crn` - (String) The CRN for this virtual network interface.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
* `href` - (String) The URL for this virtual network interface.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the virtual network interface.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
* `mac_address` - (String) The MAC address of the interface. Absent when the interface is not attached to a target.
  * Constraints: The maximum length is `17` characters. The minimum length is `17` characters. The value must match regular expression `/^([0-9A-F]{2}:){5}[0-9A-F]{2}$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `virtual_network_interface`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `target` - (List) The target of this virtual network interface.If absent, this virtual network interface is not attached to a target.
Nested schema for **target**:
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (String) The URL for this share mount target.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this share mount target.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this share mount target. The name is unique across all mount targets for the file share.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `share_mount_target`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `vpc` - (List) The VPC this virtual network interface resides in.
Nested schema for **vpc**:
	* `crn` - (String) The CRN for this VPC.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (String) The URL for this VPC.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this VPC.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `vpc`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `zone` - (List) The zone this virtual network interface resides in.
Nested schema for **zone**:
	* `href` - (String) The URL for this zone.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `name` - (String) The globally unique name for this zone.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.


## Import

You can import the `ibm_is_virtual_network_interface` resource by using `id`. The unique identifier for this virtual network interface.

# Syntax
```
$ terraform import ibm_is_virtual_network_interface.is_virtual_network_interface <id>
```

# Example
```
$ terraform import ibm_is_virtual_network_interface.is_virtual_network_interface 0767-fa41aecb-4f21-423d-8082-630bfba1e1d9
```
