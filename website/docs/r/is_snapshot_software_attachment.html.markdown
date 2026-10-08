---
layout: "ibm"
page_title: "IBM : ibm_is_snapshot_software_attachment"
description: |-
  Manages SnapshotSoftwareAttachment.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_snapshot_software_attachment

Manage a snapshot software attachment with this resource.

A software attachment cannot be created or deleted through the API. It is created automatically when a snapshot is taken from a volume that has software attachments, for example the boot volume of a virtual server instance provisioned from a software-licensed [catalog offering](https://cloud.ibm.com/docs/vpc?topic=vpc-about-images). This resource adopts an existing attachment, identified by `snapshot_id` and `snapshot_software_attachment_id`, and manages its mutable properties (currently only the `name`).

~> **Note:** Running `terraform destroy` on this resource only removes it from Terraform state. It does not delete the software attachment from the snapshot, and any `name` you set remains in place. The attachment is removed only when its snapshot is deleted.

## Example Usage

```hcl
resource "ibm_is_snapshot" "example" {
  name          = "my-boot-snapshot"
  source_volume = ibm_is_instance.example.boot_volume.0.volume_id
}

data "ibm_is_snapshot_software_attachments" "example" {
  snapshot_id = ibm_is_snapshot.example.id
}

resource "ibm_is_snapshot_software_attachment" "is_snapshot_software_attachment_instance" {
  snapshot_id                     = ibm_is_snapshot.example.id
  snapshot_software_attachment_id = data.ibm_is_snapshot_software_attachments.example.software_attachments.0.id
  name                            = "my-software-attachment"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Optional, String) The name for this snapshot software attachment. The name is unique across all software attachments for the snapshot. If you do not set it, the current name of the attachment is read from the API. Removing `name` from the configuration does not reset it on the attachment.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `snapshot_id` - (Required, Forces new resource, String) The snapshot identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `snapshot_software_attachment_id` - (Required, Forces new resource, String) The unique identifier of the existing snapshot software attachment to manage. You can get it from the `software_attachments` attribute of the `ibm_is_snapshot` data source or from the `ibm_is_snapshot_software_attachments` data source.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the SnapshotSoftwareAttachment, in the format `<snapshot_id>/<snapshot_software_attachment_id>`.
* `catalog_offering` - (List) The [catalog](https://cloud.ibm.com/docs/account?topic=account-restrict-by-user)offering for this snapshot software attachment. May be absent if`software_attachment.lifecycle_state` is not `stable`.
Nested schema for **catalog_offering**:
	* `plan` - (List) The billing plan for the catalog offering version associated with this snapshot software attachment. If absent, no billing plan is associated with the catalog offering version (free).
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
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `snapshot_software_attachment`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.


## Import

You can import the `ibm_is_snapshot_software_attachment` resource by using `id`.
The `id` property can be formed from `snapshot_id`, and `snapshot_software_attachment_id` in the following format:

<pre>
&lt;snapshot_id&gt;/&lt;snapshot_software_attachment_id&gt;
</pre>
* `snapshot_id`: A string. The snapshot identifier.
* `snapshot_software_attachment_id`: A string in the format `0717-7ec86020-1c6e-4889-b3f0-a15f2e50f87e`. The unique identifier for this snapshot software attachment.

# Syntax
<pre>
$ terraform import ibm_is_snapshot_software_attachment.is_snapshot_software_attachment &lt;snapshot_id&gt;/&lt;snapshot_software_attachment_id&gt;
</pre>