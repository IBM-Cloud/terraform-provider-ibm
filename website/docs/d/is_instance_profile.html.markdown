---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Instance Profile"
description: |-
  Manages IBM Cloud virtual server instance profile.
---

# ibm_is_instance_profile
Retrieve information of an existing IBM Cloud virtual server instance profile. For more information, about virtual server instance profile, see [instance profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-profiles).


## Example usage
The following example retrieves information about the `b-2x8` instance profile. 

```terraform

data "ibm_is_instance_profile" "profile" {
  name = "b-2x8"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name for this virtual server instance profile.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `architecture` - (String) The default Operating System architecture for an instance of the profile.
- `architecture_type` - (String) The type for this OS architecture.
- `architecture_values` - (String) The supported OS architecture(s) for an instance with this profile.
- `bandwidth` - (List) Nested `bandwidth` blocks have the following structure:

  Nested scheme for `bandwidth`:
  - `default` - (String) The default value for this profile field.
  - `max` - (String) The maximum value for this profile field.
  - `min` - (String) The minimum value for this profile field.
  - `step` - (String) The increment step value for this profile field.
  - `type` - (String) The type for this profile field.
  - `value` - (String) The value for this profile field.
  - `values` - (String) The permitted values for this profile field.
- `disks` - (List) Collection of the instance profile's disks. Nested `disks` blocks have the following structure:

  Nested scheme for `disks`:
  - `quantity` - (List) Nested `quantity` blocks have the following structure:
   
    Nested scheme for `quantity`:
    - `default` - (String) The default value for this profile field.
    - `max` - (String) The maximum value for this profile field.
    - `min` - (String) The minimum value for this profile field.
    - `step` - (String) The increment step value for this profile field.
    - `type` - (String) The type for this profile field.
    - `value` - (String) The value for this profile field.
    - `values` - (String) The permitted values for this profile field.
  - `size` - (List) Nested `size` blocks have the following structure:

    Nested scheme for `size`
    - `default` - (String) The default value for this profile field.
    - `max` - (String) The maximum value for this profile field.
    - `min` - (String) The minimum value for this profile field.
    - `step` - (String) The increment step value for this profile field.
    - `type` - (String) The type for this profile field.
    - `value` - (String) The value for this profile field.
    - `values` - (String) The permitted values for this profile field.
  - `supported_interface_types` - (List) Nested `supported_interface_types` blocks have the following structure:

    Nested scheme for `supported_interface_types`:
    - `default` - (String) The disk interface used for attaching the disk.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
    - `type` - (String) The type for this profile field.
    - `values` - (String) The supported disk interfaces used for attaching the disk.
- `family` - (String) The family of the virtual server instance profile.
- `href` - (String) The URL for this virtual server instance profile.
- `memory` - (List) Nested `memory` blocks have the following structure:

  Nested scheme for `memory`:
  - `default` - (String) The default value for this profile field.
  - `max` - (String) The maximum value for this profile field.
  - `min` - (String) The minimum value for this profile field.
  - `step` - (String) The increment step value for this profile field.
  - `type` - (String) The type for this profile field.
  - `value` - (String) The value for this profile field.
  - `values` - (String) The permitted values for this profile field.
- `port_speed` - (List) Nested `port_speed` blocks have the following structure:

  Nested scheme for `port_speed`:
  - `type` - (String) The type for this profile field.
  - `value` - (String) The value for this profile field.
- `vcpu_architecture` - (List) Nested `vcpu_architecture` blocks have the following structure:

  Nested scheme for `vcpu_architecture`:
  - `default` - (String) The default VCPU architecture for an instance with this profile.
  - `type` - (String) The type for this profile field.
  - `value` - (String) The VCPU architecture for an instance with this profile.
- `vcpu_count` - (List) Nested `vcpu_count` blocks have the following structure:

  Nested scheme for `vcpu_count`:
  - `default` - (String) The default value for this profile field.
  - `max` - (String) The maximum value for this profile field.
  - `min` - (String) The minimum value for this profile field.
  - `step` - (String) The increment step value for this profile field.
  - `type` - (String) The type for this profile field.
  - `value` - (String) The value for this profile field.
  - `values` - (String) The permitted values for this profile field.
