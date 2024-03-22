---
layout: "ibm"
page_title: "IBM : ibm_is_vpc_dns_resolution_binding"
description: |-
  Get information about VPCDNSResolutionBinding
subcategory: "VPC infrastructure"
---

# ibm_is_vpc_dns_resolution_binding

Provides a read-only data source for VPCDNSResolutionBinding. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding" {
	identifier = "id"
	vpc_id = "vpc_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `identifier` - (Required, Forces new resource, String) The DNS resolution binding identifier.
- `vpc_id` - (Required, Forces new resource, String) The VPC identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPCDNSResolutionBinding.
- `created_at` - (String) The date and time that the DNS resolution binding was created.

- `endpoint_gateways` - (List) The endpoint gateways in the bound to VPC that are allowed to participate in this DNS resolution binding.The endpoint gateways may be remote and therefore may not be directly retrievable.
  - Constraints: The minimum length is `0` items.
	Nested scheme for **endpoint_gateways**:
	- `crn` - (String) The CRN for this endpoint gateway.
	  - Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	- `href` - (String) The URL for this endpoint gateway.
	- `id` - (String) The unique identifier for this endpoint gateway.
	- `name` - (String) The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.
	- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested scheme for **remote**:
		- `account` - (List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
			Nested scheme for **account**:
			- `id` - (String) The unique identifier for this account.
			  - Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
			- `resource_type` - (String) The resource type.
		- `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested scheme for **region**:
			- `href` - (String) The URL for this region.
			- `name` - (String) The globally unique name for this region.
	- `resource_type` - (String) The resource type.

- `health_reasons` - (List) The reasons for the current `health_state` (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
	Nested schema for **health_reasons**:
	- `code` - (String) A snake case string succinctly identifying the reason for this health state.
	- `message` - (String) An explanation of the reason for this health state.
	- `more_info` - (String) Link to documentation about the reason for this health state.
	
- `health_state` - (String) The health of this resource.- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state. Constraints: Allowable values are: `degraded`, `faulted`, `inapplicable`, `ok`.

- `href` - (String) The URL for this DNS resolution binding.

- `lifecycle_state` - (String) The lifecycle state of the DNS resolution binding.
  - Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.

- `name` - (String) The name for this DNS resolution binding. The name is unique across all DNS resolution bindings for the VPC.

- `resource_type` - (String) The resource type.

- `vpc` - (List) The VPC bound to for DNS resolution.The VPC may be remote and therefore may not be directly retrievable.
	Nested scheme for **vpc**:
	- `crn` - (String) The CRN for this VPC.
	  - Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
	- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested scheme for **remote**:
		- `account` - (List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
			Nested scheme for **account**:
			- `id` - (String) The unique identifier for this account.
			- `resource_type` - (String) The resource type.
		- `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested scheme for **region**:
			- `href` - (String) The URL for this region.
			- `name` - (String) The globally unique name for this region.
	- `resource_type` - (String) The resource type.

