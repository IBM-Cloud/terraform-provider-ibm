---
layout: "ibm"
page_title: "IBM : ibm_is_vpc_dns_resolution_binding"
description: |-
  create VPCDNSResolutionBinding
subcategory: "VPC infrastructure"
---

# ibm_is_vpc_dns_resolution_binding

Provides a resource for VPCDNSResolutionBinding. You can then reference the fields of the resource in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
resource "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding_by_id" {
	name = "example-dns"
	vpc_id = "vpc_id"
	vpc {
		id = "<target_vpc_id">
	}
}
resource "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding_by_crn" {
	name = "example-dns"
	vpc_id = "vpc_id"
	vpc {
		crn = "<target_vpc_crn">
	}
}
resource "ibm_is_vpc_dns_resolution_binding" "is_vpc_dns_resolution_binding_href" {
	name = "example-dns"
	vpc_id = "vpc_id"
	vpc {
		href = "<target_vpc_href">
	}
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The DNS resolution binding name.
- `vpc_id` - (Required, Forces new resource, String) The VPC identifier of the source vpc.
- `vpc` - (Required, Forces new resource, String) The VPC identifier/href/crn of the target.
	Nested scheme for **vpc**:
	- `crn` - (Optional, String) The CRN for this target vpc.
	- `href` - (Optional, String) The href for this target vpc.
	- `id` - (Optional, String) The unique identifier for this target vpc.

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
	  - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `id` - (String) The unique identifier for this endpoint gateway.
	  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	- `name` - (String) The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.
	  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested scheme for **remote**:
		- `account` - (List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
			Nested scheme for **account**:
			- `id` - (String) The unique identifier for this account.
			  - Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
			- `resource_type` - (String) The resource type.
			  - Constraints: Allowable values are: `account`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		- `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested scheme for **region**:
			- `href` - (String) The URL for this region.
			  - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			- `name` - (String) The globally unique name for this region.
			  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `resource_type` - (String) The resource type.
	  - Constraints: Allowable values are: `endpoint_gateway`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

- `health_reasons` - (List) The reasons for the current `health_state` (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
	Nested schema for **health_reasons**:
	- `code` - (String) A snake case string succinctly identifying the reason for this health state.
	- `message` - (String) An explanation of the reason for this health state.
	- `more_info` - (String) Link to documentation about the reason for this health state.
	
- `health_state` - (String) The health of this resource.- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle state. A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state. Constraints: Allowable values are: `degraded`, `faulted`, `inapplicable`, `ok`.

- `href` - (String) The URL for this DNS resolution binding.
  - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

- `lifecycle_state` - (String) The lifecycle state of the DNS resolution binding.
  - Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.

- `name` - (String) The name for this DNS resolution binding. The name is unique across all DNS resolution bindings for the VPC.
  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.

- `resource_type` - (String) The resource type.
  - Constraints: Allowable values are: `vpc_dns_resolution_binding`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

- `vpc` - (List) The VPC bound to for DNS resolution.The VPC may be remote and therefore may not be directly retrievable.
	Nested scheme for **vpc**:
	- `crn` - (String) The CRN for this VPC.
	  - Constraints: The maximum length is `512` characters. The minimum length is `9` characters.
	- `href` - (String) The URL for this VPC.
	  - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `id` - (String) The unique identifier for this VPC.
	  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
	  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested scheme for **remote**:
		- `account` - (List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
			Nested scheme for **account**:
			- `id` - (String) The unique identifier for this account.
			  - Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
			- `resource_type` - (String) The resource type.
			  - Constraints: Allowable values are: `account`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		- `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested scheme for **region**:
			- `href` - (String) The URL for this region.
			  - Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			- `name` - (String) The globally unique name for this region.
			  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `resource_type` - (String) The resource type.
	  - Constraints: Allowable values are: `vpc`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

