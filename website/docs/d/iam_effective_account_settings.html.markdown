---
layout: "ibm"
page_title: "IBM : ibm_iam_effective_account_settings"
description: |-
  Get information about iam_effective_account_settings
subcategory: "IAM Identity Services"
---

# ibm_iam_effective_account_settings

Provides a read-only data source to retrieve information about iam_effective_account_settings. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_effective_account_settings" "iam_effective_account_settings" {
	account_id = ibm_iam_effective_account_settings.iam_effective_account_settings_instance.account_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, Forces new resource, String) Unique ID of the account.
* `include_history` - (Optional, Boolean) Defines if the entity history is included in the response.
  * Constraints: The default value is `false`.
* `resolve_user_mfa` - (Optional, Boolean) Enrich MFA exemptions with user information.
  * Constraints: The default value is `false`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_effective_account_settings.
* `account` - (List) Input body parameters for the Account Settings REST request.
Nested schema for **account**:
	* `account_id` - (String) Unique ID of the account.
	* `allowed_ip_addresses` - (String) Defines the IP addresses and subnets from which IAM tokens can be created for the account.
	* `entity_tag` - (String) Version of the account settings.
	* `history` - (List) History of the Account Settings.
	Nested schema for **history**:
		* `action` - (String) Action of the history entry.
		* `iam_id` - (String) IAM ID of the identity which triggered the action.
		* `iam_id_account` - (String) Account of the identity which triggered the action.
		* `message` - (String) Message which summarizes the executed action.
		* `params` - (List) Params of the history entry.
		* `timestamp` - (String) Timestamp when the action was triggered.
	* `max_sessions_per_identity` - (String) Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.
	* `mfa` - (String) MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
	* `restrict_create_platform_apikey` - (String) Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `restrict_create_service_id` - (String) Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `restrict_user_domains` - (List) Defines if account invitations are restricted to specified domains. To remove an entry for a realm_id, perform an update (PUT) request with only the realm_id set.
	Nested schema for **restrict_user_domains**:
		* `invitation_email_allow_patterns` - (List) The list of allowed email patterns. Wildcard syntax is supported, '*' represents any sequence of zero or more characters in the string, except for '.' and '@'. The sequence ends if a '.' or '@' was found. '**' represents any sequence of zero or more characters in the string - without limit.
		* `realm_id` - (String) The realm that the restrictions apply to.
		* `restrict_invitation` - (Boolean) When true invites will only be possible to the domain patterns provided, otherwise invites are unrestricted.
	* `restrict_user_list_visibility` - (String) Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.
	  * Constraints: The default value is `NOT_RESTRICTED`. Allowable values are: `NOT_RESTRICTED`, `RESTRICTED`.
	* `session_expiration_in_seconds` - (String) Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `86400`.
	* `session_invalidation_in_seconds` - (String) Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `7200`.
	* `system_access_token_expiration_in_seconds` - (String) Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `3600`.
	* `system_refresh_token_expiration_in_seconds` - (String) Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `259200`.
	* `user_mfa` - (List) List of users that are exempted from the MFA requirement of the account.
	Nested schema for **user_mfa**:
		* `description` - (String) optional description.
		* `email` - (String) email of the user.
		* `iam_id` - (String) The iam_id of the user.
		* `mfa` - (String) MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
		  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
		* `name` - (String) name of the user account.
		* `user_name` - (String) userName of the user.
* `assigned_templates` - (List) assigned template section.
Nested schema for **assigned_templates**:
	* `allowed_ip_addresses` - (String) Defines the IP addresses and subnets from which IAM tokens can be created for the account.
	* `max_sessions_per_identity` - (String) Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.
	* `mfa` - (String) MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
	* `restrict_create_platform_apikey` - (String) Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `restrict_create_service_id` - (String) Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `session_expiration_in_seconds` - (String) Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `86400`.
	* `session_invalidation_in_seconds` - (String) Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `7200`.
	* `system_access_token_expiration_in_seconds` - (String) Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `3600`.
	* `system_refresh_token_expiration_in_seconds` - (String) Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `259200`.
	* `template_id` - (String) Template Id.
	* `template_name` - (String) Template name.
	* `template_version` - (Integer) Template version.
	* `user_mfa` - (List) List of users that are exempted from the MFA requirement of the account.
	Nested schema for **user_mfa**:
		* `description` - (String) optional description.
		* `email` - (String) email of the user.
		* `iam_id` - (String) The iam_id of the user.
		* `mfa` - (String) MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
		  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
		* `name` - (String) name of the user account.
		* `user_name` - (String) userName of the user.
* `context` - (List) Context with key properties for problem determination.
Nested schema for **context**:
	* `cluster_name` - (String) The cluster name.
	* `elapsed_time` - (String) The elapsed time in msec.
	* `end_time` - (String) The finish time of the request.
	* `host` - (String) The host of the server instance processing the request.
	* `instance_id` - (String) The instance ID of the server instance processing the request.
	* `operation` - (String) The operation of the inbound REST request.
	* `start_time` - (String) The start time of the request.
	* `thread_id` - (String) The thread ID of the server instance processing the request.
	* `transaction_id` - (String) The transaction ID of the inbound REST request.
	* `url` - (String) The URL of that cluster.
	* `user_agent` - (String) The user agent of the inbound REST request.
* `effective` - (List) 
Nested schema for **effective**:
	* `allowed_ip_addresses` - (String) Defines the IP addresses and subnets from which IAM tokens can be created for the account.
	* `max_sessions_per_identity` - (String) Defines the max allowed sessions per identity required by the account. Valid values:  * Any whole number greater than 0  * NOT_SET - To unset account setting and use service default.
	* `mfa` - (String) MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
	  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
	* `restrict_create_platform_apikey` - (String) Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `restrict_create_service_id` - (String) Defines whether or not creating the resource is access controlled. Valid values:  * RESTRICTED - only users assigned the 'Service ID creator' role on the IAM Identity Service can create service IDs, including the account owner  * NOT_RESTRICTED - all members of an account can create service IDs  * NOT_SET - to 'unset' a previous set value.
	  * Constraints: The default value is `NOT_SET`. Allowable values are: `RESTRICTED`, `NOT_RESTRICTED`, `NOT_SET`.
	* `restrict_user_list_visibility` - (String) Defines whether or not user visibility is access controlled. Valid values:  * RESTRICTED - users can view only specific types of users in the account, such as those the user has invited to the account, or descendants of those users based on the classic infrastructure hierarchy  * NOT_RESTRICTED - any user in the account can view other users from the Users page in IBM Cloud console.
	  * Constraints: The default value is `NOT_RESTRICTED`. Allowable values are: `NOT_RESTRICTED`, `RESTRICTED`.
	* `session_expiration_in_seconds` - (String) Defines the session expiration in seconds for the account. Valid values:  * Any whole number between between '900' and '86400'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `86400`.
	* `session_invalidation_in_seconds` - (String) Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:  * Any whole number between '900' and '7200'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `7200`.
	* `system_access_token_expiration_in_seconds` - (String) Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `3600`.
	* `system_refresh_token_expiration_in_seconds` - (String) Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '259200'  * NOT_SET - To unset account setting and use service default.
	  * Constraints: The default value is `259200`.
	* `user_mfa` - (List) List of users that are exempted from the MFA requirement of the account.
	Nested schema for **user_mfa**:
		* `description` - (String) optional description.
		* `email` - (String) email of the user.
		* `iam_id` - (String) The iam_id of the user.
		* `mfa` - (String) MFA trait definitions as follows:  * NONE - No MFA trait set  * NONE_NO_ROPC- No MFA, disable CLI logins with only a password  * TOTP - For all non-federated IBMId users  * TOTP4ALL - For all users  * LEVEL1 - Email-based MFA for all users  * LEVEL2 - TOTP-based MFA for all users  * LEVEL3 - U2F MFA for all users.
		  * Constraints: Allowable values are: `NONE`, `NONE_NO_ROPC`, `TOTP`, `TOTP4ALL`, `LEVEL1`, `LEVEL2`, `LEVEL3`.
		* `name` - (String) name of the user account.
		* `user_name` - (String) userName of the user.

