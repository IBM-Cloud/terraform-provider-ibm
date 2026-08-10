---
layout: "ibm"
page_title: "IBM : ibm_is_dynamic_route_server_member"
description: |-
  Manages DynamicRouteServerMember.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_dynamic_route_server_member

Create, update, and delete DynamicRouteServerMembers with this resource.

## Example Usage

```hcl
resource "ibm_is_dynamic_route_server_member" "is_dynamic_route_server_member_instance" {
  dynamic_route_server_id = "dynamic_route_server_id"
  name = "my-dynamic-route-server-member1"
  virtual_network_interfaces {
		crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::virtual-network-interface:0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		deleted {
			more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
		}
		href = "https://us-south.iaas.cloud.ibm.com/v1/virtual_network_interfaces/0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		id = "0717-54eb57ee-86f2-4796-90bb-d7874e0831ef"
		name = "my-virtual-network-interface"
		primary_ip {
			address = "192.168.3.4"
			deleted {
				more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
			}
			href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-bea6a632-5e13-42a4-b4b8-31dc877abfe4/reserved_ips/0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
			id = "0717-6d353a0f-aeb1-4ae1-832e-1110d10981bb"
			name = "my-reserved-ip"
			resource_type = "subnet_reserved_ip"
		}
		resource_type = "virtual_network_interface"
		subnet {
			crn = "crn:v1:bluemix:public:is:us-south-1:a/aa2432b1fa4d4ace891e9b80fc104e34::subnet:0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
			deleted {
				more_info = "https://cloud.ibm.com/apidocs/vpc#deleted-resources"
			}
			href = "https://us-south.iaas.cloud.ibm.com/v1/subnets/0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
			id = "0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
			name = "my-subnet"
			resource_type = "subnet"
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `dynamic_route_server_id` - (Required, Forces new resource, String) The dynamic route server identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `name` - (Required, Forces new resource, String) The name for this dynamic route server member. The name is unique across all members in the dynamic route server.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `virtual_network_interfaces` - (Required, Forces new resource, List) The virtual network interfaces for this dynamic route server member.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested schema for **virtual_network_interfaces**:
	* `crn` - (Required, String) The CRN for this virtual network interface.
	  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
	* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (Required, String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (Required, String) The URL for this virtual network interface.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^https:\/\/([^\/?#]*)([^?#]*)\/virtual_network_interfaces\/[-0-9a-z_]+$/`.
	* `id` - (Required, String) The unique identifier for this virtual network interface.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (Required, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `primary_ip` - (Required, List) The primary IP for this virtual network interface.
	Nested schema for **primary_ip**:
		* `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (Required, String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (Required, String) The URL for this reserved IP.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (Required, String) The unique identifier for this reserved IP.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (Required, String) The resource type.
		  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `resource_type` - (Required, String) The resource type.
	  * Constraints: Allowable values are: `virtual_network_interface`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `subnet` - (Required, List) The associated subnet.
	Nested schema for **subnet**:
		* `crn` - (Required, String) The CRN for this subnet.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
		* `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (Required, String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (Required, String) The URL for this subnet.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (Required, String) The unique identifier for this subnet.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (Required, String) The name for this subnet. The name is unique across all subnets in the VPC.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (Required, String) The resource type.
		  * Constraints: Allowable values are: `subnet`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the DynamicRouteServerMember.
* `health_reasons` - (List) The reasons for the current `health_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **health_reasons**:
	* `code` - (String) A reason code for this health state.
	  * Constraints: Allowable values are: `cannot_start_capacity`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this health state.
	* `more_info` - (String) A link to documentation about the reason for this health state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `health_state` - (String) The health of this dynamic route server member:- `ok`: No abnormal behavior detected- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A member with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` member may also have this state.
  * Constraints: Allowable values are: `faulted`, `inapplicable`, `ok`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `href` - (String) The URL for this dynamic route server member.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_dynamic_route_server_member_id` - (String) The unique identifier for this dynamic route server member.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	* `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the reason for this lifecycle state.
	* `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `lifecycle_state` - (String) The lifecycle state of the dynamic route server member.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `dynamic_route_server_member`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.


## Import

You can import the `ibm_is_dynamic_route_server_member` resource by using `id`.
The `id` property can be formed from `dynamic_route_server_id`, and `is_dynamic_route_server_member_id` in the following format:

<pre>
&lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_member_id&gt;
</pre>
* `dynamic_route_server_id`: A string. The dynamic route server identifier.
* `is_dynamic_route_server_member_id`: A string in the format `0717-9dce0ab3-a719-4582-ad71-00abfaae43ed`. The unique identifier for this dynamic route server member.

# Syntax
<pre>
$ terraform import ibm_is_dynamic_route_server_member.is_dynamic_route_server_member &lt;dynamic_route_server_id&gt;/&lt;is_dynamic_route_server_member_id&gt;
</pre>
