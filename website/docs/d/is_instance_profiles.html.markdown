---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance profiles"
description: |-
  Manages IBM Cloud virtual server instance profiles.
---

# ibm\_is_instance_profiles

Import the details of an existing IBM Cloud virtual server instance profiles as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_instance_profiles" "ds_instance_profiles" {
}

```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `profiles` - List of all server instance profiles in the region.
  * `name` - The name for this virtual server instance profile.
  * `family` - The family of the virtual server instance profile.
  * `architecture` - The default OS architecture for an instance with this profile.
  * `architecture_type` - The type for this OS architecture.
  * `architecture_values` - The supported OS architecture(s) for an instance with this profile.
  * `bandwidth`  Nested `bandwidth` blocks have the following structure:
	* `type` - The type for this profile field.
	* `value` - The value for this profile field.
	* `default` - The default value for this profile field.
	* `max` - The maximum value for this profile field.
	* `min` - The minimum value for this profile field.
	* `step` - The increment step value for this profile field.
	* `values` - The permitted values for this profile field.
  * `disks` - Collection of the instance profile's disks. Nested `disks` blocks have the following structure:
    * `quantity`  Nested `quantity` blocks have the following structure:
      * `type` - The type for this profile field.
      * `value` - The value for this profile field.
      * `default` - The default value for this profile field.
      * `max` - The maximum value for this profile field.
      * `min` - The minimum value for this profile field.
      * `step` - The increment step value for this profile field.
      * `values` - The permitted values for this profile field.
    * `size`  Nested `size` blocks have the following structure:
      * `type` - The type for this profile field.
      * `value` - The value for this profile field.
      * `default` - The default value for this profile field.
      * `max` - The maximum value for this profile field.
      * `min` - The minimum value for this profile field.
      * `step` - The increment step value for this profile field.
      * `values` - The permitted values for this profile field.
    * `supported_interface_types`  Nested `supported_interface_types` blocks have the following structure:
      * `default` - The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
      * `type` - The type for this profile field.
      * `values` - The supported disk interfaces used for attaching the disk.
  * `href` - The URL for this virtual server instance profile.
  * `memory`  Nested `memory` blocks have the following structure:
    * `type` - The type for this profile field.
    * `value` - The value for this profile field.
    * `default` - The default value for this profile field.
    * `max` - The maximum value for this profile field.
    * `min` - The minimum value for this profile field.
    * `step` - The increment step value for this profile field.
    * `values` - The permitted values for this profile field.
  * `port_speed`  Nested `port_speed` blocks have the following structure:
	* `type` - The type for this profile field.
	* `value` - The value for this profile field.
  * `vcpu_architecture`  Nested `vcpu_architecture` blocks have the following structure:
	* `default` - The default VCPU architecture for an instance with this profile.
	* `type` - The type for this profile field.
	* `value` - The VCPU architecture for an instance with this profile.
  * `vcpu_count`  Nested `vcpu_count` blocks have the following structure:
	* `type` - The type for this profile field.
	* `value` - The value for this profile field.
	* `default` - The default value for this profile field.
	* `max` - The maximum value for this profile field.
	* `min` - The minimum value for this profile field.
	* `step` - The increment step value for this profile field.
	* `values` - The permitted values for this profile field.
