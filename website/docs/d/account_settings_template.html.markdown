---
layout: "ibm"
page_title: "IBM : ibm_iam_account_settings_template"
description: |-
  Get information about an IAM account settings template
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_account_settings_template

Provides a read-only data source to retrieve information about an account_settings_template. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_account_settings_template" "account_settings_template" {
	template_id = "${var.template_id}"
	version = "${var.version}"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `template_id` - (Required, String) ID of the account settings template.
* `version` - (Optional, String) Version of the account settings template. If the template_id provided comes from a created terraform resource then the version is not required, as the terraform resource id contains the ID and version already. If the template is pre-existing however, then both the template_id and version must be provided
* `include_history` - (Optional, Boolean) Defines if the entity history is included in the response.
  * Constraints: The default value is `false`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - (Required, String) The unique identifier of the account_settings_template.
* `version` (Required, Integer) - The version of the account_settings_template
* `name` - (Required, String) The name of the account settings template.
* `description` - (Optional, String) The description of the account settings template. Describe the template for enterprise account users.
* `account_settings` - (List) Nested schema for **account_settings**:
  * `allowed_ip_addresses` - (Optional, String) Defines the IP addresses and subnets from which IAM tokens can be created for the account.
  * `max_sessions_per_identity` - (Optional, String) Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.
  * `mfa` - (Optional, String) Defines the MFA trait for the account. 
   * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
  * `restrict_create_platform_apikey` - (Optional, String) Defines whether or not creating platform API keys is access controlled.
   * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
  * `restrict_create_service_id` - (Optional, String) Defines whether or not creating a service ID is access controlled.
   * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
  * `restrict_user_domains` - (List)
		Nested schema for **restrict_user_domains**:
			* `account_sufficient` - (Boolean)
			* `restrictions` - (List) Defines if account invitations are restricted to specified domains. To remove an entry for a realm_id, perform an update (PUT) request with only the realm_id set.
			Nested schema for **restrictions**:
				* `invitation_email_allow_patterns` - (List) The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit.
				* `realm_id` - (String) The realm that the restrictions apply to.
				* `restrict_invitation` - (Boolean) When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted.
	* `restrict_user_list_visibility` - (String) Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console  * NOT_SET - to 'unset' a previous set value.
		  * Constraints: Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
  * `session_expiration_in_seconds` - (Optional, String) Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.
   * Constraints: The default value is `86400`.
  * `session_invalidation_in_seconds` - (Optional, String) Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.
   * Constraints: The default value is `7200`.
  * `system_access_token_expiration_in_seconds` - (Optional, String) Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.
   * Constraints: The default value is `3600`.
  * `system_refresh_token_expiration_in_seconds` - (Optional, String) Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.
   * Constraints: The default value is `259200`.
  * `user_mfa` - (Optional, List) List of users that are exempted from the MFA requirement of the account.
  Nested schema for **user_mfa**:
  * `iam_id` - (Required, String) The iam_id of the user.
  * `mfa` - (Required, String) Defines the MFA requirement for the user, overriding the account level MFA.
    * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
* `committed` - (Boolean) Committed flag determines if the template is ready for assignment. Committed templates can not be updated.
* `created_at` - (String) Template Created At.
* `created_by_id` - (String) IAMid of the creator.
* `crn` - (String) Cloud resource name.
* `entity_tag` - (String) Entity tag for this templateId-version combination.

* `history` - (List) History of the Template.
Nested schema for **history**:
	* `action` - (String) Action of the history entry.
	* `iam_id` - (String) IAM ID of the identity which triggered the action.
	* `iam_id_account` - (String) Account of the identity which triggered the action.
	* `message` - (String) Message which summarizes the executed action.
	* `params` - (List) Params of the history entry.
	* `timestamp` - (String) Timestamp when the action was triggered.

* `last_modified_at` - (String) Template last modified at.

* `last_modified_by_id` - (String) IAMid of the identity that made the latest modification.


