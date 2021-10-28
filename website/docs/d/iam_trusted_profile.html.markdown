---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile"
description: |-
  Get information about iam_trusted_profile
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile

Create, update and delete an IAM trusted profile resource. For more information, about trusted profile, see [Create a trusted profile](https://cloud.ibm.com/apidocs/iam-identity-token-api#create-profile)

## Example usage

```terraform
data "ibm_iam_trusted_profile" "iam_trusted_profile" {
	profile_id = ibm_iam_trusted_profile_link.iam_trusted_profile_link.profile_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile to get.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `account_id` - (String) ID of the account that this trusted profile belong to.

* `created_at` - (String) If set contains a date time string of the creation date in ISO format.

* `crn` - (String) Cloud Resource Name of the item. For example: `crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::profile:Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c`.

* `description` - (String) The optional description of the trusted profile. The 'description' property is only available if a description was provided during a create of a trusted profile.

* `entity_tag` - (String) The Version of the trusted profile details object. You need to specify this value when updating the trusted profile to avoid stale updates.

* `history` - (List) History of the trusted profile.
    Nested scheme for **history**:
	* `timestamp` - (String) The timestamp when the action was triggered.
	* `iam_id` - (String) The IAM ID of an identity which triggered an action.
	* `iam_id_account` - (String) The account of the identity which triggered the action.
	* `action` - (String) The action of the history entry.
	* `params` - (List) The params of the history entry.
	* `message` - (String) The message which summarizes the executed action.

* `iam_id` - (String) The `iam_id` of this trusted profile.

* `id` - (String) The unique identifier of the trusted profile. For example:`Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c`.

* `ims_account_id` - (Integer) IMS acount ID of the trusted profile.

* `ims_user_id` - (Integer) IMS user ID of the trusted profile.

* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.

* `name` - (String) The unique name of the trusted profile. Therefore trusted profiles with the same names can not exist in the same account.