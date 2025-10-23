---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : ibm_iam_account_settings"
description: |-
  Manages iam_account_settings.
---

# ibm_iam_account_settings

Create or update iam_account_settingss with this resource.

## Example Usage

```hcl
resource "ibm_iam_account_settings" "iam_account_settings_instance" {
  mfa = "LEVEL3"
  session_expiration_in_seconds = "40000"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) Unique ID of the account.
* `include_history` - (Optional, Boolean) Defines if the entity history is included in the response.
  * Constraints: The default value is `false`.
* `resolve_user_mfa` - (Optional, Boolean) Enrich MFA exemptions with user PI.
  * Constraints: The default value is `false`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_account_settings.
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


## Import

You can import the `ibm_iam_account_settings` resource by using `account_id`.
The `account_id` property can be formed from and `account_id` in the following format:

<pre>
&lt;account_id&gt;
</pre>
* `account_id`: A string. Unique ID of the account.

# Syntax
<pre>
$ terraform import ibm_iam_account_settings.iam_account_settings &lt;account_id&gt;
</pre>
