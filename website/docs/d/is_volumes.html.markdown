---
layout: "ibm"
page_title: "IBM : ibm_is_volumes"
description: |-
  Get information about VolumeCollection
subcategory: "VPC infrastructure"
---

# ibm_is_volumes

Provides a read-only data source for VolumeCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_volumes" "example" {
}
```
```hcl
data "ibm_is_volumes" "example" {
  volume_name = "my-example-volume"
  zone_name = "us-south-2"
  attachment_state = "unattached"
  encryption = "provider_managed"
  operating_system_family = "Ubuntu Server"
  operating_system_architecture = "amd64"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `volume_name` - (Required, String) Filters the collection to resources with the exact specified name.
- `zone_name` - (Optional, String) Filters the collection to resources in the zone with the exact specified name.
- `attachment_state` - (Optional, String) Filters the collection to volumes with the specified attachment state.
- `encryption` - (Optional, String) Filters the collection to resources with the specified encryption type.
- `operating_system_family` - (Optional, String) Filters the collection to resources with the exact specified operating system family.
- `operating_system_architecture` - (Optional, String) Filters the collection to resources with the exact specified operating system architecture.
## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VolumeCollection.
- `volumes` - (List) Collection of volumes.
	Nested scheme for **volumes**:
	- `access_tags`  - (List) Access management tags associated for the volume.
	- `active` - (Boolean) Indicates whether a running virtual server instance has an attachment to this volume.
	- `adjustable_capacity_states` - (List) The attachment states that support adjustable capacity for this volume. Allowable list items are: `attached`, `unattached`, `unusable`. 
	- `adjustable_iops_states` - (List) The attachment states that support adjustable IOPS for this volume. Allowable list items are: `attached`, `unattached`, `unusable`.
	- `attachment_state` - (Boolean) The attachment state of the volume
	- `allowed_use` - (List) The usage constraints to be matched against the requested instance or bare metal server properties to determine compatibility. Only present for boot volumes. The value of this property will be inherited from the source image or snapshot at volume creation, but can be changed.
    
    Nested schema for `allowed_use`:
    - `api_version` - (String) The API version with which to evaluate the expressions.
	  
    - `bare_metal_server` - (String) The expression that must be satisfied by the properties of a bare metal server provisioned using the image data in this volume. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
    
    ~> **NOTE** </br> In addition, the following property is supported, corresponding to the BareMetalServer property: </br>
      **&#x2022;** `enable_secure_boot` - (boolean) Indicates whether secure boot is enabled.
	  
    - `instance` - (String) The expression that must be satisfied by the properties of a virtual server instance provisioned using this volume. The expression follows [Common Expression Language](https://github.com/google/cel-spec/blob/master/doc/langdef.md), but does not support built-in functions and macros. 
    
    ~> **NOTE** </br> In addition, the following variables are supported, corresponding to `Instance` </br>
       **&#x2022;** `gpu.count` - (integer) The number of GPUs. </br>
       **&#x2022;** `gpu.manufacturer` - (string) The GPU manufacturer. </br>
       **&#x2022;** `gpu.memory` - (integer) The overall amount of GPU memory in GiB (gibibytes). </br>
       **&#x2022;** `gpu.model` - (string) The GPU model. </br>
       **&#x2022;** `enable_secure_boot` - (boolean)Indicates whether secure boot is enabled. </br>
	- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.
	- `busy` - (Boolean) Indicates whether this volume is performing an operation that must be serialized. This must be `false` to perform an operation that is specified to require serialization.
	- `capacity` - (Integer) The capacity to use for the volume (in gigabytes). The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
	  - Constraints: The minimum value is `1`.
	- `catalog_offering` - (List) The catalog offering this volume was created from. If a virtual server instance is provisioned with a boot_volume_attachment specifying this volume, the virtual server instance will use this volume's catalog offering, including its pricing plan.If absent, this volume was not created from a catalog offering.
	  
	  Nested scheme for `catalog_offering`:
	    - `version_crn` - (String) The CRN for this version of a catalog offering
		- `plan_crn` - (String) The CRN for this catalog offering version's billing plan
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
		
		  Nested schema for `deleted`:
		    - `more_info`  - (String) Link to documentation about deleted resources.
	- `created_at` - (String) The date and time that the volume was created.
	- `crn` - (String) The CRN for this volume.
	- `encryption` - (String) The type of encryption used on the volume.
	  - Constraints: The default value is `provider_managed`. Allowable values are: `provider_managed`, `user_managed`.
	- `encryption_key` - (Optional, List) The root key used to wrap the data encryption key for the volume.This property will be present for volumes with an `encryption` type of`user_managed`.
	Nested scheme for **encryption_key**:
		- `crn` - (String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
	- `health_reasons` - (List) The reasons for the current health_state (if any).
	
  	Nested scheme for `health_reasons`:
  		- `code` - (String) A snake case string succinctly identifying the reason for this health state.
  		- `message` - (String) An explanation of the reason for this health state.
  		- `more_info` - (String) Link to documentation about the reason for this health state.
	- `health_state` - (String) The health of this resource.
	- `href` - (String) The URL for this volume.
	- `id` - (String) The unique identifier for this volume.
	- `iops` - (Integer) The maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profile `family` of `custom`.
	- `name` - (String) The unique user-defined name for this volume.
	- `operating_system` - (Optional, List) The operating system associated with this volume. If absent, this volume was notcreated from an image, or the image did not include an operating system.
		Nested scheme for **operating_system**:
		- `architecture` - (String) The operating system architecture
		- `dedicated_host_only` - (Boolean) Images with this operating system can only be used on dedicated hosts or dedicated host groups
		- `display_name` - (String) A unique, display-friendly name for the operating system
		- `family` - (String) The software family for this operating system
		- `href` - (String) The URL for this operating system.
		- `name` - (String) The globally unique name for this operating system.
		- `vendor` - (String) The vendor of the operating system
		- `version` - (String) The major release version of this operating system
	- `profile` - (List) The profile this volume uses.
		Nested scheme for **profile**:
		- `href` - (String) The URL for this volume profile.
		- `name` - (String) The globally unique name for this volume profile.
	- `resource_group` - (List) The resource group object, for this volume.
		Nested scheme for **resource_group**:
		- `href` - (String) The URL for this resource group.
		- `id` - (String) The unique identifier for this resource group.
		- `name` - (String) The user-defined name for this resource group.
	- `source_image` - (Optional, List) The image from which this volume was created (this may be[deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).If absent, this volume was not created from an image.
		Nested scheme for **source_image**:
		- `crn` - (String) The CRN for this image.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this image.
		- `id` - (String) The unique identifier for this image.
		- `name` - (String) The user-defined or system-provided name for this image.
	- `source_snapshot` - (Optional, List) The snapshot from which this volume was cloned.
		Nested scheme for **source_snapshot**:
		- `crn` - (String) The CRN for this snapshot.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this snapshot.
		- `id` - (String) The unique identifier for this snapshot.
		- `name` - (String) The user-defined name for this snapshot.
		- `resource_type` - (String) The resource type.
		  - Constraints: Allowable values are: `snapshot`.
	- `status` - (String) The status of the volume.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the volume on which the unexpected property value was encountered.
	  - Constraints: Allowable values are: `available`, `failed`, `pending`, `pending_deletion`, `unusable`.
	- `status_reasons` - (List) The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
		Nested scheme for **status_reasons**:
		- `code` - (String) A snake case string succinctly identifying the status reason.
		- `message` - (String) An explanation of the status reason.
		- `more_info` - (Optional, String) Link to documentation about this status reason.
	- `storage_generation` - (Int) The storage generation indicates which generation the profile family belongs to. For the custom and tiered profiles, this value is 1. For the sdp profile, this value is 2.
	- `tags` - (String) User Tags associated with the volume. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
	- `volume_attachments` - (List) The volume attachments for this volume.
		Nested scheme for **volume_attachments**:
		- `delete_volume_on_instance_delete` - (Boolean) If set to true, when deleting the instance the volume will also be deleted.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `device` - (Optional, List) Information about how the volume is exposed to the instance operating system.This property may be absent if the volume attachment's `status` is not `attached`.
			Nested scheme for **device**:
			- `id` - (Optional, String) A unique identifier for the device which is exposed to the instance operating system.
		- `href` - (String) The URL for this volume attachment.
		- `id` - (String) The unique identifier for this volume attachment.
		- `instance` - (List) The attached instance.
			Nested scheme for **instance**:
			- `crn` - (String) The CRN for this virtual server instance.
			- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
				Nested scheme for **deleted**:
				- `more_info` - (String) Link to documentation about deleted resources.
			- `href` - (String) The URL for this virtual server instance.
			- `id` - (String) The unique identifier for this virtual server instance.
			- `name` - (String) The user-defined name for this virtual server instance (and default system hostname).
		- `name` - (String) The user-defined name for this volume attachment.
		- `type` - (String) The type of volume attachment. Allowable values are: `boot`, `data`.
	- `zone` - (List) The zone this volume resides in.
		Nested scheme for **zone**:
		- `name` - (String) The globally unique name for this zone.

