---
layout: "ibm"
page_title: "IBM : is_dedicated_host_profile"
sidebar_current: "docs-ibm-datasource-is-dedicated-host-profile"
description: |-
  Get information about DedicatedHostProfile
---

# ibm\_is_dedicated_host_profile

Provides a read-only data source for DedicatedHostProfile. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_dedicated_host_profile" "is_dedicated_host_profile" {
	name = "dh2-56x464"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The globally unique name for this virtual server instance profile.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the DedicatedHostProfile.
* `class` - The product class this dedicated host profile belongs to.

* `family` - The product family this dedicated host profile belongs to.

* `href` - The URL for this dedicated host.

* `memory`  Nested `memory` blocks have the following structure:
	* `type` - The type for this profile field.
	* `value` - The value for this profile field.
	* `default` - The default value for this profile field.
	* `max` - The maximum value for this profile field.
	* `min` - The minimum value for this profile field.
	* `step` - The increment step value for this profile field.
	* `values` - The permitted values for this profile field.

* `socket_count`  Nested `socket_count` blocks have the following structure:
	* `type` - The type for this profile field.
	* `value` - The value for this profile field.
	* `default` - The default value for this profile field.
	* `max` - The maximum value for this profile field.
	* `min` - The minimum value for this profile field.
	* `step` - The increment step value for this profile field.
	* `values` - The permitted values for this profile field.

* `supported_instance_profiles` - Array of instance profiles that can be used by instances placed on dedicated hosts with this profile Nested`supported_instance_profiles` blocks have the following structure:
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

