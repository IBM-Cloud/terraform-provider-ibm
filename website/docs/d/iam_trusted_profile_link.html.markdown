---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_link"
description: |-
  Get information about iam_trusted_profile_link
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_link

Retrieve information about IAM trusted profile link as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about trusted profile link, see [Create a link to a trusted profile](https://cloud.ibm.com/apidocs/iam-identity-token-api#create-link)

## Example usage

```terraform
data "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
	link_id = "link_id"
	profile_id = "profile_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `link_id` - (Required, Forces new resource, String) The ID of the link.
* `profile_id` - (Required, Forces new resource, String) The ID of the trusted profile.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the `iam_trusted_profile_link`.
* `cr_type` - (String) The compute resource type. Supported values are **VSI, IKS_SA, ROKS_SA**.

* `created_at` - (String) If set contains a date time string of the creation date in ISO format.

* `entity_tag` - (String) The version of the claim rule.

* `id` - (String) The unique identifier of the claim rule.

* `link` - (List) 
    Nested scheme for **link**:
	* `crn` - (String) The CRN of the compute resource.
	* `namespace` - (String) The compute resource namespace, only required if `cr_type` is **IKS_SA** or **ROKS_SA**.
	* `name` - (String) Name of the compute resource, only required if cr_type is **IKS_SA** or **ROKS_SA**.

* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.

* `name` - (String) The optional name of the Link.

