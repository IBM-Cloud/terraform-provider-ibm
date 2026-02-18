---
layout: "ibm"
page_title: "IBM : ibm_is_snapshot_software_attachment"
description: |-
  Get information about SnapshotSoftwareAttachment
subcategory: "Virtual Private Cloud API"
---

# ibm_is_snapshot_software_attachment

Provides a read-only data source to retrieve information about a SnapshotSoftwareAttachment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_snapshot_software_attachment" "is_snapshot_software_attachment" {
	is_snapshot_software_attachment_id = ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance.is_snapshot_software_attachment_id
	snapshot_id = ibm_is_snapshot_software_attachment.is_snapshot_software_attachment_instance.snapshot_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `is_snapshot_software_attachment_id` - (Required, Forces new resource, String) The snapshot software attachment identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `snapshot_id` - (Required, Forces new resource, String) The snapshot identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the SnapshotSoftwareAttachment.
* `catalog_offering` - (List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this snapshot software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.
Nested schema for **catalog_offering**:
	* `plan` - (List) The billing plan for the catalog offering version associated with this snapshot softwareattachment.If absent, no billing plan is associated with the catalog offering version (free).
	Nested schema for **plan**:
		* `crn` - (String) The CRN for this[catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user) offering version's billing plan.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
		* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `version` - (List) The catalog offering version associated with this snapshot software attachment.
	Nested schema for **version**:
		* `crn` - (String) The CRN for this version of a[catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user) offering.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
* `created_at` - (String) The date and time that the snapshot software attachment was created.
* `entitlement` - (List) The entitlement for the snapshot software attachment's licensable software.
Nested schema for **entitlement**:
	* `licensable_software` - (List) The licensable software for this snapshot software attachment entitlement. The software will be licensed when an instance is provisioned from this snapshot.
	  * Constraints: The minimum length is `0` items.
	Nested schema for **licensable_software**:
		* `sku` - (String) The SKU for this licensable software.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~]+$/`.
* `href` - (String) The URL for this snapshot software attachment.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `name` - (String) The name for this snapshot software attachment. The name is unique across all software attachments for the snapshot.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `snapshot_software_attachment`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

