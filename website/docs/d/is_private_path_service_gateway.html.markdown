---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateway"
description: |-
  Get information about PrivatePathServiceGateway
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateway

Provides a read-only data source for PrivatePathServiceGateway. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_private_path_service_gateway" "is_private_path_service_gateway" {
	id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The private path service gateway identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the PrivatePathServiceGateway.
* `created_at` - (String) The date and time that the private path service gateway was created.

* `crn` - (String) The CRN for this private path service gateway.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.

* `default_access_policy` - (String) The policy to use for bindings from accounts without an explicit account policy.
  * Constraints: Allowable values are: `deny`, `permit`, `review`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

* `endpoint_gateways_count` - (Integer) The number of endpoint gateways using this private path service gateway.
  * Constraints: The minimum value is `0`.

* `href` - (String) The URL for this private path service gateway.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `lifecycle_state` - (String) The lifecycle state of the private path service gateway.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.

* `load_balancer` - (List) The load balancer for this private path service gateway.
Nested scheme for **load_balancer**:
	* `crn` - (String) The load balancer's CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested scheme for **deleted**:
		* `more_info` - (String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (String) The load balancer's canonical URL.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this load balancer.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this load balancer. The name is unique across all load balancers in the VPC.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `load_balancer`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

* `name` - (String) The name for this private path service gateway. The name is unique across all private path service gateways in the VPC.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.

* `published` - (Boolean) Indicates the availability of this private path service gateway- `true`: Any account can request access to this private path service gateway.- `false`: Access is restricted to the account that created this private path service gateway.

* `region` - (List) The region served by this private path service gateway.
Nested scheme for **region**:
	* `href` - (String) The URL for this region.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `name` - (String) The globally unique name for this region.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.

* `resource_group` - (List) The resource group for this private path service gateway.
Nested scheme for **resource_group**:
	* `href` - (String) The URL for this resource group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this resource group.
	  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.
	* `name` - (String) The name for this resource group.
	  * Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`.

* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `private_path_service_gateway`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

* `service_endpoints` - (List) The fully qualified domain names for this private path service gateway.
  * Constraints: The list items must match regular expression `/^((?=[A-Za-z0-9-]{1,63}\\.)[A-Za-z0-9-]*\\.)+[A-Za-z]{2,63}\\.?$/`. The maximum length is `10` items. The minimum length is `1` item.

* `vpc` - (List) The VPC this private path service gateway resides in.
Nested scheme for **vpc**:
	* `crn` - (String) The CRN for this VPC.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested scheme for **deleted**:
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

* `zonal_affinity` - (Boolean) Indicates whether this private path service gateway has zonal affinity.- `true`:  Traffic to the service from a zone will favor service endpoints in           the same zone.- `false`: Traffic to the service from a zone will be load balanced across all zones           in the region the service resides in.

