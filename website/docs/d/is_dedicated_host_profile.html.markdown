---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_profile"
description: |-
  Get information about dedicated host profile
---

# ibm_is_dedicated_host_profile
Retrieve an information about the dedicated host profile. For more information, about dedicated host groups in your IBM Cloud VPC, see [dedicated host profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-dh-profiles).


## Example usage

```terraform
data "ibm_is_dedicated_host_profile" "is_dedicated_host_profile" {
	name = "dh2-56x464"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The globally unique user defined name for this `VSI` profile.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `class`-  (String) The product class this dedicated host profile belongs to.
- `disks` - (List) Collection of the dedicated host profile's disks. 

  Nested scheme for `disks`:
  - `interface_type`- (List) The interface type.

    Nested scheme for `interface_type`:
    - `type` - (String) The type for this profile field.
    - `value` - (String) The interface of the disk for a dedicated host with this profileThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
 - `quantity` - (List) The number of disks of this type for a dedicated host with this profile. Nested `quantity` blocks have the following structure:

    Nested scheme for `quantity`:
    - `type` - (String) The type for this profile field.
    - `value` - (String) The value for this profile field.
 - `size` - (List) The size of the disk in GB (gigabytes). Nested `size` blocks have the following structure:

    Nested scheme for `size`:
    - `type` - (String) The type for this profile field.
    - `value` - (String) The size of the disk in GB (gigabytes).
 - `supported_instance_interface_types` - (List) Nested `supported_instance_interface_types` blocks have the following structure:

    Nested scheme for `supported_instance_interface_types`:
    - `type` - (String) The type for this profile field.
    - `value` - (String) The instance disk interfaces supported for a dedicated host with this profile.
- `family`-  (String) The product family this dedicated host profile belongs to.
- `href`-  (String) The URL for this dedicated host.
- `id`-  (String) The unique identifier of the dedicated host profile.
- `memory`-  (List) Nested memory blocks have the following structure.

  Nested scheme for `memory`:
  - `default` -  (String) The default value for this profile field.
  - `max` -  (String) The maximum value for this profile field.
  - `min` -  (String) The minimum value for this profile field.
  - `step` -  (String) The increment step value for this profile field.
  - `type` -  (String) The type for this profile field.
  - `value` -  (String) The value for this profile field.
  - `values` -  (String) The permitted values for this profile field.
- `socket_count` - (List) Nested socket_count blocks have the following structure.

  Nested scheme for `socket_count`:
  - `type` -  (String) The type for this profile field.
  - `value` -  (String) The value for this profile field.
  - `default` -  (String) The default value for this profile field.
  - `max` -  (String) The maximum value for this profile field.
  - `min` -  (String) The minimum value for this profile field.
  - `step` -  (String) The increment step value for this profile field.
  - `values` -  (String) The permitted values for this profile field.
- `supported_instance_profiles`-  (List) Array of instance profiles that can be used by instances placed on dedicated hosts with this profile Nested `supported_instance_profiles` blocks have the following structure.

  Nested scheme for `supported_instance_profiles`:
  - `href`-  (String) The URL for this virtual server instance profile.
  - `name`-  (String) The globally unique name for this virtual server instance profile.
- `vcpu_architecture`-  (List) Nested `vcpu_architecture` blocks have the following structure.

  Nested scheme for `vcpu_architecture`:
  - `type`-  (String) The type for this profile field.
  - `value`-  (String) The `VCPU` architecture for a dedicated host with this profile.
- `vcpu_count` - (List) Nested `vcpu_count` blocks have the following structure.

  Nested scheme for `vcpu_count`:
  - `default`-  (String) The default value for this profile field.
  - `max`-  (String) The maximum value for this profile field.
  - `min`-  (String) The minimum value for this profile field.
  - `step`-  (String) The increment step value for this profile field.
  - `type`-  (String) The type for this profile field.
  - `value`-  (String) The value for this profile field.
  - `values`-  (String) The permitted values for this profile field.
