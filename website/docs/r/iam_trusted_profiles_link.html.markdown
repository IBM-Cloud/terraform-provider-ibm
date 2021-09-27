---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_link"
description: |-
  Manages iam_trusted_profile_link.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_link

Create, update, and delete an IAM trusted profile link. For more information, about IAM trusted profile, see https://cloud.ibm.com/apidocs/iam-identity-token-api#create-link 

## Example usage

```terraform
resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
  cr_type = "cr_type"
  link = { "crn" : "crn", "namespace" : "namespace" }
  profile_id = "profile_id"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `cr_type` - (Required, Forces new resource, String) The compute resource type. Supported values are VSI, IKS_SA, ROKS_SA.
* `link` - (Required, Forces new resource, List) Link details.
  Nested scheme for **link**:
	* `crn` - (Required, String) The CRN of the compute resource.
	* `namespace` - (Required, String) The compute resource namespace, only required if `cr_type` is IKS_SA or ROKS_SA.
	* `name` - (Optional, String) Name of the compute resource, only required if `cr_type` is IKS_SA or ROKS_SA.
* `name` - (Optional, Forces new resource, String) Optional name of the Link.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the `iam_trusted_profile_link`.
* `created_at` - (Required, String) If set contains a date time string of the creation date in ISO format.
* `entity_tag` - (Required, String) The version of the claim rule.
* `modified_at` - (Required, String) If set contains a date time string of the last modification date in ISO format.

## Import

You can import the `ibm_iam_trusted_profile_link` resource by using `id`.
The `id` property can be formed from `profile-id`, and `link-id` in the following format:

```
<profile-id>/<link-id>
```
* `profile-id`: A string. ID of the trusted profile.
* `link-id`: A string. ID of the link.

# Syntax
```
$ terraform import ibm_iam_trusted_profile_link.iam_trusted_profile_link <profile-id>/<link-id>
```