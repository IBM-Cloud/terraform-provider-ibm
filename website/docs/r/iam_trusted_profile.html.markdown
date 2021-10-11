---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile"
description: |-
  Manages iam_trusted_profile.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile

Provides a resource for iam_trusted_profile. This allows iam_trusted_profile to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_iam_trusted_profile" "iam_trusted_profile" {
  name = "name"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `description` - (Optional, String) The optional description of the trusted profile. The 'description' property is only available if a description was provided during creation of trusted profile.
* `name` - (Required, String) Name of the trusted profile. The name is checked for uniqueness. Therefore trusted profiles with the same names can not exist in the same account.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `account_id` - (String) The account ID of the trusted profile.
* `id` - The unique identifier of the iam_trusted_profile.
* `created_at` - (Optional, String) If set contains a date time string of the creation date in ISO format.
* `crn` - (Required, String) The Cloud Resource Name of the item. For example: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::profile:Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c'.
* `entity_tag` - (Required, String) Version of the trusted profile details object. You need to specify this value when updating the trusted profile to avoid stale updates.
* `history` - (Optional, List) History of the trusted profile.
Nested scheme for **history**:
	* `timestamp` - (Required, String) Timestamp when the action was triggered.
	* `iam_id` - (Required, String) IAM ID of the identity which triggered the action.
	* `iam_id_account` - (Required, String) Account of the identity which triggered the action.
	* `action` - (Required, String) Action of the history entry.
	* `params` - (Required, List) Params of the history entry.
	* `message` - (Required, String) Message which summarizes the executed action.
* `iam_id` - (Required, String) The iam_id of this trusted profile.
* `id` - (Required, String) the unique identifier of the trusted profile. Example:'Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c'.
* `ims_account_id` - (Optional, Integer) IMS acount ID of the trusted profile.
* `ims_user_id` - (Optional, Integer) IMS user ID of the trusted profile.
* `modified_at` - (Optional, String) If set contains a date time string of the last modification date in ISO format.

## Import

You can import the `ibm_iam_trusted_profile` resource by using `profile_id`. ID of the account that this trusted profile belong to.

# Syntax
```
$ terraform import ibm_iam_trusted_profile.iam_trusted_profiles <profile_id>
```
