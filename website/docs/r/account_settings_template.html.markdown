---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : ibm_iam_account_settings_template"
description: |-
  Manages IAM account settings templates.
---

# ibm_iam_account_settings_template

Create, update, commit, and delete account_settings_templates with this resource.

## Example Usage

```hcl
resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
	name = "My template name"
	description = "My template description"
	account_settings {
		restrict_create_service_id = "RESTRICTED"
		restrict_create_platform_apikey = "RESTRICTED"
		allowed_ip_addresses = "127.0.0.1"
		mfa = "LEVEL3"
		user_mfa {
			iam_id = "IBMid-123456879"
			mfa = "TOTP"
		}
		session_expiration_in_seconds = "1800"
		session_invalidation_in_seconds = "900"
		max_sessions_per_identity = "3"
		system_access_token_expiration_in_seconds = "NOT_SET"
		system_refresh_token_expiration_in_seconds = "NOT_SET"
	}
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, String) The name of the account settings template.
* `description` - (Optional, String) The description of the account settings template. Describe the template for enterprise account users.
* `account_settings` - (Required, Object) Nested schema for **account_settings**:
	* `allowed_ip_addresses` - (Optional, String) Defines the IP addresses and subnets from which IAM tokens can be created for the account.
	* `max_sessions_per_identity` - (Optional, String) Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.
	* `mfa` - (Optional, String) Defines the MFA trait for the account. 
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
	* `restrict_create_platform_apikey` - (Optional, String) Defines whether or not creating platform API keys is access controlled.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `restrict_create_service_id` - (Optional, String) Defines whether or not creating a service ID is access controlled.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
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

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The identifier of the account_settings_template.
* `version` - The version of the account_settings_template
* `committed` - (Boolean) Committed flag determines if the template is ready for assignment. Committed templates can not be updated.
* `created_at` - (String) Template Created At.
* `created_by_id` - (String) IAMid of the creator.
* `crn` - (String) Cloud resource name.
* `entity_tag` - (String) Entity tag for this templateId-version combination.
* `account_settings` - account settings definition to be applied on assignment
  * `allowed_ip_addresses` - (String) Defines the IP addresses and subnets from which IAM tokens can be created for the account. **Note** value should be a comma separated string.
  * `mfa` - (String) Defines the session expiration in seconds for the account.
  * `user_mfa` - (String) List of users that are exempted from the MFA requirement of the account.
  * `max_sessions_per_identity` - (String) Defines the maximum allowed sessions per identity required by the account.
  * `restrict_create_service_id` - (String) Defines whether or not creating a service ID is access controlled.
  * `restrict_create_platform_apikey` - (String) Defines whether or not creating platform API keys is access controlled.
  * `session_expiration_in_seconds` - (String) Defines the session expiration in seconds for the account.
  * `session_invalidation_in_seconds` - (String) Defines the period of time in seconds in which a session is invalid due to inactivity.
  * `system_access_token_expiration_in_seconds` - (String) Defines the access token expiration in seconds.
  * `system_refresh_token_expiration_in_seconds` - (String) Defines the refresh token expiration in seconds.
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



## Import

You can import the `ibm_iam_account_settings_template` resource by using `version`. Version of the the template.

# Syntax
```
$ terraform import ibm_iam_account_settings_template.account_settings_template <version>
```
