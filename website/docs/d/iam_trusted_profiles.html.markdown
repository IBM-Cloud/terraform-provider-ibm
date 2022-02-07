---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profiles"
description: |-
  Get information about iam_trusted_profiles
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profiles

List IAM trusted profile resource. For more information, about list trusted profile, see [List a trusted profile](https://cloud.ibm.com/apidocs/iam-identity-token-api#list-profile)

## Example usage

```terraform
data "ibm_iam_trusted_profiles" "iam_trusted_profiles" {
	account_id = "account_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `account_id` - (Optional, String) Account ID to query for trusted profiles.
* `name` - (Optional, String) Name of the trusted profile to query.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the iam_trusted_profiles.

* `profiles` - (List) List of trusted profiles.
  Nested scheme for **profiles**:
    * `account_id` - (String) ID of the account that this trusted profile belong to.
	* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
	* `crn` - (String) Cloud Resource Name of the item. Example Cloud Resource Name: `crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::profile:Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c`.
	* `description` - (String) The optional description of the trusted profile. The 'description' property is only available if a description was provided during a 
	* `entity_tag` - (String) Version of the trusted profile details object. You need to specify this value when updating the trusted profile to avoid stale updates.
	* `history` - (List) History of the trusted profile.
	  Nested scheme for **history**:
		* `action` - (String) Action of the history entry.
		* `iam_id` - (String) IAM ID of the identity which triggered the action.
		* `iam_id_account` - (String) Account of the identity which triggered the action.
		* `message` - (String) Message which summarizes the executed action.
		* `params` - (List) Params of the history entry.
		* `timestamp` - (String) Timestamp when the action was triggered.
	* `iam_id` - (String) The iam_id of this trusted profile.
	* `id` - (String) the unique identifier of the trusted profile. Example:`Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c`.
	* `ims_account_id` - (Integer) IMS acount ID of the trusted profile.
	* `ims_user_id` - (Integer) IMS user ID of the trusted profile.
	* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
	* `name` - (String) Name of the trusted profile. The name is checked for uniqueness. Therefore trusted profiles with the same names can not exist in the same account.

