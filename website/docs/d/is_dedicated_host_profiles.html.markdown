---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_profiles"
description: |-
  Get information about DedicatedHostProfileCollection
---

# ibm\_is_dedicated_host_profiles

Provides a read-only data source for DedicatedHostProfileCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_dedicated_host_profiles" "is_dedicated_host_profiles" {
}
```

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHostProfileCollection.
* `profiles` - Collection of dedicated host profiles. Nested `profiles` blocks have the following structure:
	* `class` - The product class this dedicated host profile belongs to.
	* `family` - The product family this dedicated host profile belongs toThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	* `href` - The URL for this dedicated host.
	* `memory`  Nested `memory` blocks have the following structure:
		* `type` - The type for this profile field.
		* `value` - The value for this profile field.
		* `default` - The default value for this profile field.
		* `max` - The maximum value for this profile field.
		* `min` - The minimum value for this profile field.
		* `step` - The increment step value for this profile field.
		* `values` - The permitted values for this profile field.
	* `name` - The globally unique name for this dedicated host profile.
	* `socket_count`  Nested `socket_count` blocks have the following structure:
		* `type` - The type for this profile field.
		* `value` - The value for this profile field.
		* `default` - The default value for this profile field.
		* `max` - The maximum value for this profile field.
		* `min` - The minimum value for this profile field.
		* `step` - The increment step value for this profile field.
		* `values` - The permitted values for this profile field.
	* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on dedicated hosts with this profile. Nested `supported_instance_profiles` blocks have the following structure:
		* `href` - The URL for this virtual server instance profile.
		* `name` - The globally unique name for this virtual server instance profile.
	* `vcpu_architecture`  Nested `vcpu_architecture` blocks have the following structure:
		* `type` - The type for this profile field.
		* `value` - The VCPU architecture for a dedicated host with this profile.
	* `vcpu_count`  Nested `vcpu_count` blocks have the following structure:
		* `type` - The type for this profile field.
		* `value` - The value for this profile field.
		* `default` - The default value for this profile field.
		* `max` - The maximum value for this profile field.
		* `min` - The minimum value for this profile field.
		* `step` - The increment step value for this profile field.
		* `values` - The permitted values for this profile field.

* `total_count` - The total number of resources across all pages.

