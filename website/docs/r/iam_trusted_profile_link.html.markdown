---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_link"
description: |-
  Manages iam_trusted_profile_link.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_link

Create, update, and delete an IAM trusted profile link. For more information, about IAM trusted profile, see https://cloud.ibm.com/apidocs/iam-identity-token-api#create-link 

## Example usage

```terraform
resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
  name = "test"
}
resource "ibm_iam_trusted_profile_link" "iam_trusted_profile_link" {
  profile_id = ibm_iam_trusted_profile.iam_trusted_profile.id
  cr_type    = "IKS_SA"
  link {
    crn       = "crn:v1:bluemix:public:containers-kubernetes:us-south:a/acc_id:cluster_id::"
    namespace = "namespace"
    name      = "name"
  }
  name = "name"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `cr_type` - (Required, Forces new resource, String) The compute resource type. Supported values are VSI, IKS_SA, ROKS_SA.
* `link` - (Required, Forces new resource, List) Link details.
  Nested scheme for **link**:
	* `crn` - (Required, String) The CRN of the compute resource.
	* `namespace` - (Optional, String) The compute resource namespace, only required if `cr_type` is IKS_SA or ROKS_SA.
	* `name` - (Optional, String) Name of the compute resource, only required if `cr_type` is IKS_SA or ROKS_SA.
* `name` - (Optional, Forces new resource, String) Optional name of the Link.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` -  Id is combination of `profile_id`/ `link_id`.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `entity_tag` - (String) The version of the claim rule.
* `link_id` - The unique identifier of the `iam_trusted_profile_link`.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.

## Import

The  `ibm_iam_trusted_profile_link` resource can be imported by using profile ID and trusted profile link ID 
**Syntax**

```
$ terraform import ibm_iam_trusted_profile_link.example <profile_id>/<link_id>
```
