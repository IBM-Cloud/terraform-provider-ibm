---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profiles_link"
description: |-
  Get information about iam_trusted_profiles_link
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profiles_link

Provides a read-only data source for iam_trusted_profiles_link. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_trusted_profiles_link" "iam_trusted_profiles_link" {
	link_id = "link_id"
	profile_id = "profile_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `link_id` - (Required, Forces new resource, String) ID of the link.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the iam_trusted_profiles_link.
* `cr_type` - (Required, String) The compute resource type. Valid values are VSI, IKS_SA, ROKS_SA.

* `created_at` - (Required, String) If set contains a date time string of the creation date in ISO format.

* `entity_tag` - (Required, String) version of the claim rule.

* `id` - (Required, String) the unique identifier of the claim rule.

* `link` - (Required, List) 
Nested scheme for **link**:
	* `crn` - (Optional, String) The CRN of the compute resource.
	* `namespace` - (Optional, String) The compute resource namespace, only required if cr_type is IKS_SA or ROKS_SA.
	* `name` - (Optional, String) Name of the compute resource, only required if cr_type is IKS_SA or ROKS_SA.

* `modified_at` - (Required, String) If set contains a date time string of the last modification date in ISO format.

* `name` - (Optional, String) Optional name of the Link.

