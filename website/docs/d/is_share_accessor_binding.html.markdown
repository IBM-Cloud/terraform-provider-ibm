---
layout: "ibm"
page_title: "IBM : ibm_is_share_accessor_binding"
description: |-
  Get information about ShareAccessorBinding
subcategory: "VPC infrastructure"
---

# ibm_is_share_accessor_binding

Provides a read-only data source to retrieve information about a ShareAccessorBinding. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_share_accessor_bindings" "is_share_accessor_bindings" {
	share = "shareId"
}
data "ibm_is_share_accessor_binding" "is_share_accessor_binding" {
	accessor_binding = "share_accessor_binding_id"
	share = "share_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `accessor_binding` - (Required, Forces new resource, String) The file share accessor binding identifier.
- `share` - (Required, Forces new resource, String) The file share identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the ShareAccessorBinding.
- `accessor` - (List) The accessor for this share accessor binding.The resources supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	Nested schema for **accessor**:
	- `crn` - (String) The CRN for this file share.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this file share.
	- `id` - (String) The unique identifier for this file share.
	- `name` - (String) The name for this share. The name is unique across all shares in the region.
	- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested schema for **remote**:
		- `account` - (List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
			Nested schema for **account**:
			- `id` - (String) The unique identifier for this account.
			- `resource_type` - (String) The resource type.
		- `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested schema for **region**:
			- `href` - (String) The URL for this region.
			- `name` - (String) The globally unique name for this region.
	- `resource_type` - (String) The resource type.
- `created_at` - (String) The date and time that the share accessor binding was created.
- `href` - (String) The URL for this share accessor binding.
- `lifecycle_state` - (String) The lifecycle state of the file share accessor binding.
- `resource_type` - (String) The resource type.

