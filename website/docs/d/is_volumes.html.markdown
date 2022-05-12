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


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VolumeCollection.
- `volumes` - (List) Collection of volumes.
Nested scheme for **volumes**:
	- `active` - (Boolean) Indicates whether a running virtual server instance has an attachment to this volume.
	- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.
	- `busy` - (Boolean) Indicates whether this volume is performing an operation that must be serialized. If an operation specifies that it requires serialization, the operation will fail unless this property is `false`.
	- `capacity` - (Integer) The capacity to use for the volume (in gigabytes). The specified minimum and maximum capacity values for creating or updating volumes may expand in the future.
	  - Constraints: The minimum value is `1`.
	- `created_at` - (String) The date and time that the volume was created.
	- `crn` - (String) The CRN for this volume.
	- `encryption` - (String) The type of encryption used on the volume.
	  - Constraints: The default value is `provider_managed`. Allowable values are: `provider_managed`, `user_managed`.
	- `encryption_key` - (Optional, List) The root key used to wrap the data encryption key for the volume.This property will be present for volumes with an `encryption` type of`user_managed`.
	Nested scheme for **encryption_key**:
		- `crn` - (String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
	- `href` - (String) The URL for this volume.
	  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `id` - (String) The unique identifier for this volume.
	  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	- `iops` - (Integer) The maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profile `family` of `custom`.
	- `name` - (String) The unique user-defined name for this volume.
	  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `operating_system` - (Optional, List) The operating system associated with this volume. If absent, this volume was notcreated from an image, or the image did not include an operating system.
	Nested scheme for **operating_system**:
		- `href` - (String) The URL for this operating system.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `name` - (String) The globally unique name for this operating system.
		  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `profile` - (List) The profile this volume uses.
	Nested scheme for **profile**:
		- `href` - (String) The URL for this volume profile.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `name` - (String) The globally unique name for this volume profile.
		  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `resource_group` - (List) The resource group for this volume.
	Nested scheme for **resource_group**:
		- `href` - (String) The URL for this resource group.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this resource group.
		  - Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
		- `name` - (String) The user-defined name for this resource group.
		  - Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`.
	- `source_image` - (Optional, List) The image from which this volume was created (this may be[deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).If absent, this volume was not created from an image.
	Nested scheme for **source_image**:
		- `crn` - (String) The CRN for this image.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
			  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `href` - (String) The URL for this image.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this image.
		  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		- `name` - (String) The user-defined or system-provided name for this image.
		  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `source_snapshot` - (Optional, List) The snapshot from which this volume was cloned.
	Nested scheme for **source_snapshot**:
		- `crn` - (String) The CRN for this snapshot.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
			  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `href` - (String) The URL for this snapshot.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this snapshot.
		  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		- `name` - (String) The user-defined name for this snapshot.
		  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		- `resource_type` - (String) The resource type.
		  - Constraints: Allowable values are: `snapshot`.
	- `status` - (String) The status of the volume.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the volume on which the unexpected property value was encountered.
	  - Constraints: Allowable values are: `available`, `failed`, `pending`, `pending_deletion`, `unusable`.
	- `status_reasons` - (List) The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
	Nested scheme for **status_reasons**:
		- `code` - (String) A snake case string succinctly identifying the status reason.
		  - Constraints: Allowable values are: `encryption_key_deleted`. The value must match regular expression `/^[a-z]+(_[a-z]+)*$/`.
		- `message` - (String) An explanation of the status reason.
		- `more_info` - (Optional, String) Link to documentation about this status reason.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `volume_attachments` - (List) The volume attachments for this volume.
	Nested scheme for **volume_attachments**:
		- `delete_volume_on_instance_delete` - (Boolean) If set to true, when deleting the instance the volume will also be deleted.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
			  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `device` - (Optional, List) Information about how the volume is exposed to the instance operating system.This property may be absent if the volume attachment's `status` is not `attached`.
		Nested scheme for **device**:
			- `id` - (Optional, String) A unique identifier for the device which is exposed to the instance operating system.
		- `href` - (String) The URL for this volume attachment.
		  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this volume attachment.
		  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		- `instance` - (List) The attached instance.
		Nested scheme for **instance**:
			- `crn` - (String) The CRN for this virtual server instance.
			- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for **deleted**:
				- `more_info` - (String) Link to documentation about deleted resources.
				  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			- `href` - (String) The URL for this virtual server instance.
			  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			- `id` - (String) The unique identifier for this virtual server instance.
			  - Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
			- `name` - (String) The user-defined name for this virtual server instance (and default system hostname).
			  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		- `name` - (String) The user-defined name for this volume attachment.
		  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		- `type` - (String) The type of volume attachment.
		  - Constraints: Allowable values are: `boot`, `data`.
	- `zone` - (List) The zone this volume resides in.
	Nested scheme for **zone**:
		- `name` - (String) The globally unique name for this zone.
		  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.

