---
layout: "ibm"
page_title: "IBM : ibm_is_share_accessor_bindings"
description: |-
  Get information about ShareAccessorBindingCollection
subcategory: "VPC infrastructure"
---

# ibm_is_share_accessor_bindings

Provides a read-only data source to retrieve information about a ShareAccessorBindingCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_share_accessor_bindings" "is_share_accessor_bindings" {
	share = "shareId"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `share` - (Required, Forces new resource, String) The file share identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the ShareAccessorBindingCollection.
* `accessor_bindings` - (List) Collection of share accessor bindings.
Nested schema for **accessor_bindings**:
	* `accessor` - (List) The accessor for this share accessor binding.The resources supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	Nested schema for **accessor**:
		* `crn` - (String) The CRN for this file share.
		* `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			* `more_info` - (String) Link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `href` - (String) The URL for this file share.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		* `id` - (String) The unique identifier for this file share.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		* `name` - (String) The name for this share. The name is unique across all shares in the region.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested schema for **remote**:
			* `account` - (List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
			Nested schema for **account**:
				* `id` - (String) The unique identifier for this account.
				  * Constraints: The value must match regular expression `/^[0-9a-f]{32}$/`.
				* `resource_type` - (String) The resource type.
				  * Constraints: Allowable values are: `account`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
			* `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested schema for **region**:
				* `href` - (String) The URL for this region.
				  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
				* `name` - (String) The globally unique name for this region.
				  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		* `resource_type` - (String) The resource type.
		  * Constraints: Allowable values are: `share`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `created_at` - (String) The date and time that the share accessor binding was created.
	* `href` - (String) The URL for this share accessor binding.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The unique identifier for this share accessor binding.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	* `lifecycle_state` - (String) The lifecycle state of the file share accessor binding.
	  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
	* `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `share_accessor_binding`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

