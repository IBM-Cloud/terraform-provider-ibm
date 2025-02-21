---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile"
description: |-
  Manages iam_trusted_profile.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile

Create, update, and delete iam_trusted_profiles with this resource.

## Example Usage

```hcl
resource "ibm_iam_trusted_profile" "iam_trusted_profile_instance" {
  name = "name"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) The optional description of the trusted profile. The 'description' property is only available if a description was provided during a create of a trusted profile.
* `name` - (Required, String) Name of the trusted profile. The name is checked for uniqueness. Therefore trusted profiles with the same names can not exist in the same account.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_trusted_profile.
* `assignment_id` - (String) ID of the assignment that was used to create an enterprise-managed trusted profile in your account. When returned, this indicates that the trusted profile is created from and managed by a template in the root enterprise account.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `crn` - (String) Cloud Resource Name of the item. Example Cloud Resource Name: 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::profile:Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c'.
* `entity_tag` - (String) Version of the trusted profile details object. You need to specify this value when updating the trusted profile to avoid stale updates.
* `history` - (List) History of the trusted profile.
Nested schema for **history**:
	* `action` - (String) Action of the history entry.
	* `iam_id` - (String) IAM ID of the identity which triggered the action.
	* `iam_id_account` - (String) Account of the identity which triggered the action.
	* `message` - (String) Message which summarizes the executed action.
	* `params` - (List) Params of the history entry.
	* `timestamp` - (String) Timestamp when the action was triggered.
* `iam_id` - (String) The iam_id of this trusted profile.
* `id` - (String) the unique identifier of the trusted profile. Example:'Profile-94497d0d-2ac3-41bf-a993-a49d1b14627c'.
* `ims_account_id` - (Integer) IMS acount ID of the trusted profile.
* `ims_user_id` - (Integer) IMS user ID of the trusted profile.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
* `template_id` - (String) ID of the IAM template that was used to create an enterprise-managed trusted profile in your account. When returned, this indicates that the trusted profile is created from and managed by a template in the root enterprise account.


## Import

You can import the `ibm_iam_trusted_profile` resource by using `profile_id`. ID of the account that this trusted profile belong to.

# Syntax
<pre>
$ terraform import ibm_iam_trusted_profile.iam_trusted_profile &lt;account_id&gt;
</pre>
