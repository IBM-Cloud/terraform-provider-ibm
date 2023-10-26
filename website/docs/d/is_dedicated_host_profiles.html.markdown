---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_dedicated_host_profiles"
description: |-
  Get information about dedicated host profiles.
---

# ibm_is_dedicated_host_profiles
Retrieve an information about the dedicated host profiles. For more information, about dedicated host profiles, see [dedicated host profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-dh-profiles).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_dedicated_host_profiles" "example" {
}
```


## Argument reference
Review the argument references that you can specify for your data source. 

- `id` - (String) The unique identifier of the dedicated host profiles.
- `profiles` - (List) Collection of dedicated host profiles. Nested `profiles` blocks have the following structure:

  Nested scheme for `profiles`:
	- `class` - (String) The product class this dedicated host profile belongs to.
	- `disks` - (List) Collection of the dedicated host profile's disks. Nested `disks` blocks have the following structure:

	  Nested scheme for `disks`:
	  - `interface_type` (List) Nested `interface_type` blocks have the following structure:

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
	 - `supported_instance_interface_types`  - (List) Nested `supported_instance_interface_types` blocks have the following structure:
	  
			Nested scheme for `supported_instance_interface_types`:
			- `type` - (String) The type for this profile field.
			- `value` - (String) The instance disk interfaces supported for a dedicated host with this profile.
	- `family` - (String) The product family this dedicated host profile belongs toThe enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	- `href` - (String) The URL for this dedicated host.
	- `memory`  - (List) Nested `memory` blocks have the following structure:

	  Nested scheme for `memory`:
	  - `default` - (String) The default value for this profile field.
	  - `max` - (String) The maximum value for this profile field.
	  - `min` - (String) The minimum value for this profile field.
	  - `step` - (String) The increment step value for this profile field.
	  - `type` - (String) The type for this profile field.
	  - `value` - (String) The value for this profile field.
	  - `values` - (String) The permitted values for this profile field.
	- `name` - (String) The global unique name for this dedicated host profile.
	- `socket_count`  - (List) Nested `socket_count` blocks have the following structure:

	  Nested scheme for `socket_count`:
	  - `default` - (String) The default value for this profile field.
	  - `max` - (String) The maximum value for this profile field.
	  - `min` - (String) The minimum value for this profile field.
	  - `step` - (String) The increment step value for this profile field.
	  - `type` - (String) The type for this profile field.
	  - `value` - (String) The value for this profile field.
	  - `values` - (String) The permitted values for this profile field.
	- `supported_instance_profiles` - (List) Array of instance profiles that can be used by instances placed on dedicated hosts with this profile.

	  Nested scheme for `supported_instance_profiles`:
	  - `href` - (String) The URL for this virtual server instance profile.
	  - `name` - (String) The global unique name for this virtual server instance profile.
	- `status` - (String) The status of the dedicated host profile. Values coule be,  `previous`: This dedicated host profile is an older revision, but remains provisionable and usable. `current`: This profile is the latest revision.
	- `vcpu_architecture` - (List)  Nested `vcpu_architecture` blocks have the following structure:

	  Nested scheme for `vcpu_architecture`:
	  - `type` - (String) The type for this profile field.
	  - `value` - (String) The VCPU architecture for a dedicated host with this profile.
	- `vcpu_count` - (List) Nested `vcpu_count` blocks have the following structure:

	  Nested scheme for `vcpu_count`:
	  - `default` - (String) The default value for this profile field.
	  - `max` - (String) The maximum value for this profile field.
	  - `min` - (String) The minimum value for this profile field.
	  - `step` - (String) The increment step value for this profile field.
	  - `type` - (String) The type for this profile field.
	  - `value` - (String) The value for this profile field.
	  - `values` - (String) The permitted values for this profile field.
	- `vcpu_manufacturer` - (List)  Nested `vcpu_manufacturer` blocks have the following structure:

	  Nested scheme for `vcpu_manufacturer`:
	  - `type` - (String) The type for this profile field.
	  - `value` - (String) The VCPU manufacturer for a dedicated host with this profile.
- `total_count` - (String) The total number of resources across all pages.

