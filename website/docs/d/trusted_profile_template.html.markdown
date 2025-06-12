---
layout: "ibm"
page_title: "IBM : ibm_trusted_profile_template"
description: |-
  Get information about an IAM trusted profile template
subcategory: "Identity & Access Management (IAM)"
---

# ibm_trusted_profile_template

Provides a read-only data source to retrieve information about a trusted_profile_template. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_trusted_profile_template" "trusted_profile_template" {
	template_id = "${var.template_id}"
	version = "${var.version}"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `template_id` - (Required, String) ID of the trusted profile template.
* `version` - (Optional, String) Version of the trusted profile template. If the template_id provided comes from a created terraform resource then the version is not required, as the terraform resource id contains the ID and version already. If the template is pre-existing however, then both the template_id and version must be provided
* `include_history` - (Optional, Boolean) Defines if the entity history is included in the response.
	* Constraints: The default value is `false`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the trusted_profile_template.
* `version` (Required, Integer) - The version of the trusted_profile_template
* `name` - (String) The name of the trusted profile template. This is visible only in the enterprise account.
* `description` - (String) The description of the trusted profile template. Describe the template for enterprise account users.
* `account_id` - (String) ID of the account where the template resides.
* `committed` - (Boolean) Committed flag determines if the template is ready for assignment.
* `profile` - (List) Input body parameters for the TemplateProfileComponent.
  Nested schema for **profile**:
	* `name` - (String) Name of the Profile.
	* `description` - (String) Description of the Profile.
	* `rules` - (List) Rules for the Profile.
	  Nested schema for **rules**:
		* `conditions` - (List) Conditions of this claim rule.
		  Nested schema for **conditions**:
			* `claim` - (String) The claim to evaluate against.
			* `operator` - (String) The operation to perform on the claim. valid values are EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN.
			* `value` - (String) The stringified JSON value that the claim is compared to using the operator.
		* `expiration` - (Integer) Session expiration in seconds, only required if type is 'Profile-SAML'.
		* `name` - (String) Name of the claim rule to be created or updated.
		* `realm_name` - (String) The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'.
		* `type` - (String) Type of the claim rule.
			* Constraints: Allowable values are: `Profile-SAML`.
	* `identities` - (List) Identities for the Profile.
	  Nested schema for **identities**:
		* `accounts` - (List) Only valid for the type user. Accounts from which a user can assume the trusted profile.
		* `description` - (String) Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'.
		* `iam_id` - (String) IAM ID of the identity.
		* `identifier` - (String) Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN.
		* `type` - (String) Type of the identity.
			* Constraints: Allowable values are: `user`, `serviceid`, `crn`.
* `policy_template_references` - (List) Existing policy templates that you can reference to assign access in the trusted profile component.
  Nested schema for **policy_template_references**:
	* `id` - (String) ID of Access Policy Template.
	* `version` - (String) Version of Access Policy Template.
* `history` - (List) History of the trusted profile template.
Nested schema for **history**:
	* `action` - (String) Action of the history entry.
	* `iam_id` - (String) IAM ID of the identity which triggered the action.
	* `iam_id_account` - (String) Account of the identity which triggered the action.
	* `message` - (String) Message which summarizes the executed action.
	* `params` - (List) Params of the history entry.
	* `timestamp` - (String) Timestamp when the action was triggered.
* `crn` - (String) Cloud resource name.
* `entity_tag` - (String) Entity tag for this templateId-version combination.
* `created_at` - (String) Timestamp of when the template was created.
* `created_by_id` - (String) IAMid of the creator.
* `last_modified_at` - (String) Timestamp of when the template was last modified.
* `last_modified_by_id` - (String) IAMid of the identity that made the latest modification.
