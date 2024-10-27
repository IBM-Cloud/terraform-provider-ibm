---
layout: "ibm"
page_title: "IBM : ibm_is_private_path_service_gateways"
description: |-
  Get information about PrivatePathServiceGatewayCollection
subcategory: "VPC infrastructure"
---

# ibm_is_private_path_service_gateways

Provides a read-only data source for PrivatePathServiceGatewayCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_private_path_service_gateways" "example" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `private_path_service_gateways` - (List) Collection of private path service gateways.
	Nested scheme for **private_path_service_gateways**:
	- `created_at` - (String) The date and time that the private path service gateway was created.
	- `crn` - (String) The CRN for this private path service gateway.
	- `default_access_policy` - (String) The policy to use for bindings from accounts without an explicit account policy.
	- `endpoint_gateway_count` - (Integer) The number of endpoint gateways using this private path service gateway.
	- `endpoint_gateway_binding_auto_delete` - (Boolean) Indicates whether endpoint gateway bindings will be automatically deleted after      endpoint_gateway_binding_auto_delete_timeout hours have passed. At present, this is always true, but may be modifiable in the future.
	- `endpoint_gateway_binding_auto_delete_timeout` - (Integer) If endpoint_gateway_binding_auto_delete is true, the hours after which endpoint gateway bindings   will be automatically deleted. If the value is 0, abandoned endpoint gateway bindings will be deleted immediately. At present, this is always set to 0. This value may be modifiable in the future.
	- `href` - (String) The URL for this private path service gateway.
	- `id` - (String) The unique identifier for this private path service gateway.
	- `lifecycle_state` - (String) The lifecycle state of the private path service gateway.
	- `load_balancer` - (List) The load balancer for this private path service gateway.
		Nested scheme for **load_balancer**:
		- `crn` - (String) The load balancer's CRN.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The load balancer's canonical URL.
		- `id` - (String) The unique identifier for this load balancer.
		- `name` - (String) The name for this load balancer. The name is unique across all load balancers in the VPC.
		- `resource_type` - (String) The resource type.
	- `name` - (String) The name for this private path service gateway. The name is unique across all private path service gateways in the VPC.
	- `published` - (Boolean) Indicates the availability of this private path service gateway- `true`: Any account can request access to this private path service gateway.- `false`: Access is restricted to the account that created this private path service gateway.
	- `region` - (List) The region served by this private path service gateway.
	Nested scheme for **region**:
		- `href` - (String) The URL for this region.
		- `name` - (String) The globally unique name for this region.
	- `resource_group` - (List) The resource group for this private path service gateway.
		Nested scheme for **resource_group**:
		- `href` - (String) The URL for this resource group.
		- `id` - (String) The unique identifier for this resource group.
		- `name` - (String) The name for this resource group.
	- `resource_type` - (String) The resource type.
	- `service_endpoints` - (List of strings) The fully qualified domain names for this private path service gateway.
	- `vpc` - (List) The VPC this private path service gateway resides in.
		Nested scheme for **vpc**:
		- `crn` - (String) The CRN for this VPC.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this VPC.
		- `id` - (String) The unique identifier for this VPC.
		- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
		- `resource_type` - (String) The resource type.
	- `zonal_affinity` - (Boolean) Indicates whether this private path service gateway has zonal affinity.- `true`:  Traffic to the service from a zone will favor service endpoints in           the same zone.- `false`: Traffic to the service from a zone will be load balanced across all zones           in the region the service resides in.



