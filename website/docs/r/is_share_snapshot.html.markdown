---
layout: "ibm"
page_title: "IBM : ibm_is_share_snapshot"
description: |-
  Manages ShareSnapshot.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_share_snapshot

Create, update, and delete ShareSnapshots with this resource.

## Example Usage

```hcl
resource "ibm_is_share_snapshot" "is_share_snapshot_instance" {
  name = "my-share-snapshot"
  share_id = "share_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Optional, String) The name for this share snapshot. The name is unique across all snapshots for the file share.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
* `share_id` - (Required, Forces new resource, String) The file share identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `user_tags` - (Optional, List) The [user tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags) associated with this share snapshot.
  * Constraints: The list items must match regular expression `/^[A-Za-z0-9:_ .-]+$/`. The maximum length is `1000` items. The minimum length is `0` items.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the ShareSnapshot.
* `backup_policy_plan` - (List) If present, the backup policy plan which created this share snapshot.
Nested schema for **backup_policy_plan**:
	* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		* `more_info` - (String) Link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `href` - (String) The URL for this backup policy plan.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this backup policy plan.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `name` - (String) The name for this backup policy plan. The name is unique across all plans in the backup policy.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
	Nested schema for **remote**:
		* `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
		Nested schema for **region**:
			* `href` - (String) The URL for this region.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
			* `name` - (String) The globally unique name for this region.
			  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `backup_policy_plan`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `captured_at` - (String) The date and time the data capture for this share snapshot was completed.If absent, this snapshot's data has not yet been captured.
* `created_at` - (String) The date and time that the share snapshot was created.
* `crn` - (String) The CRN for this share snapshot.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$|^crn:\\[\\.\\.\\.\\]$/`.
* `fingerprint` - (String) The fingerprint for this snapshot.
* `href` - (String) The URL for this share snapshot.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `is_share_snapshot_id` - (String) The unique identifier for this share snapshot.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `lifecycle_state` - (String) The lifecycle state of this share snapshot.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `minimum_size` - (Integer) The minimum size of a share created from this snapshot. When a snapshot is created, this will be set to the size of the `source_share`.
  * Constraints: The minimum value is `10`.
* `resource_group` - (List) The resource group for this file share.
Nested schema for **resource_group**:
	* `href` - (String) The URL for this resource group.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this resource group.
	  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-f]{32}$/`.
	* `name` - (String) The name for this resource group.
	  * Constraints: The maximum length is `40` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-_ ]+$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `share_snapshot`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status` - (String) The status of the share snapshot:- `available`: The share snapshot is available for use.- `failed`: The share snapshot is irrecoverably unusable.- `pending`: The share snapshot is being provisioned and is not yet usable.- `unusable`: The share snapshot is not currently usable (see `status_reasons`)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `available`, `failed`, `pending`, `unusable`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status_reasons` - (List) The reasons for the current status (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) A reason code for the status:- `encryption_key_deleted`: File share snapshot is unusable  because its `encryption_key` was deleted- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `encryption_key_deleted`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the status reason.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
	* `more_info` - (String) Link to documentation about this status reason.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `zone` - (List) The zone this share snapshot resides in.
Nested schema for **zone**:
	* `href` - (String) The URL for this zone.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `name` - (String) The globally unique name for this zone.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.

* `etag` - ETag identifier for ShareSnapshot.

## Import

You can import the `ibm_is_share_snapshot` resource by using `id`.
The `id` property can be formed from `share_id`, and `is_share_snapshot_id` in the following format:

<pre>
&lt;share_id&gt;/&lt;is_share_snapshot_id&gt;
</pre>
* `share_id`: A string. The file share identifier.
* `is_share_snapshot_id`: A string in the format `r006-e13ee54f-baa4-40d3-b35c-b9ec163972b4`. The unique identifier for this share snapshot.

# Syntax
<pre>
$ terraform import ibm_is_share_snapshot.is_share_snapshot &lt;share_id&gt;/&lt;is_share_snapshot_id&gt;
</pre>
