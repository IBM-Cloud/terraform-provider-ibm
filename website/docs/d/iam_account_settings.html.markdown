---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_account_settings"
description: |-
  Get information about an IAM account settings.
---

# ibm_iam_account_settings

Retrieve information about an existing `iam_account_settings` data sources. For more information, about IAM account settings, refer to [setting up your IBM Cloud](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).

## Example usage

```terraform
data "ibm_iam_account_settings" "iam_account_settings" {
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `include_history` - (Optional, Bool) Defines if the entity history is included in the response.

## Attribute reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `account_id` - (String) The unique ID of an account.
- `allowed_ip_addresses` - (String) Defines the IP addresses and subnets from which IAM tokens is created for an account.
- `entity_tag` - (String) The version of an account settings.
- `history` - (String) The history of an account settings. Nested history blocks have the following structure.
  Nested scheme for `history`:
  - `action` - (String) The action of the history entry.
  - `iam_id` - (String) The IAM ID of the identity that triggered an action.
  - `iam_id_account` - (String) The account of an identity that trigger an action.- `params` - (String) The parameters of the history entry.
  - `message` - (String) The message that summarizes the executed action.
  - `params` - (String) Params of the history entry.
  - `timestamp` - (String) The timestamp when an action is triggered.
- `id` - (String) The unique identifier of an iam_account_settings.
- `max_sessions_per_identity` - (Integer) Defines the maximum allowed sessions per identity required by an account.
- `mfa` - (String) Defines the MFA trait for an account. Valid values are **NONE** No MFA trait set. **TOTP** For all non-federated IBMID users **TOTP4ALL** For all users. **LEVEL1** The Email based MFA for all users. **LEVEL2** TOTP based MFA for all users. **LEVEL3** U2F MFA for all users.
- `user_mfa` - (List) List of users that are exempted from the MFA requirement of the account.
 Nested scheme for `user_mfa`:
 - `iam_id` - (String) The iam_id of the user.
 - `mfa` - (String) Defines the MFA requirement for the user. Valid values are **NONE** No MFA trait set. **TOTP** For all non-federated IBMID users **TOTP4ALL** For all users. **LEVEL1** The Email based MFA for all users. **LEVEL2** TOTP based MFA for all users. **LEVEL3** U2F MFA for all users.
- `restrict_create_service_id` - (String) Defines whether creating a service ID is access controlled. Valid values are  **RESTRICTED** to apply access control. **NOT_RESTRICTED** to remove access control. **NOT_SET** to `unset` a previous set value.
- `restrict_create_platform_apikey` - (String) Defines whether creating platform API keys is access controlled. Valid values are **RESTRICTED** to apply access control. **NOT_RESTRICTED** to remove access control. **NOT_SET** to `unset` a previous set value.
- `session_expiration_in_seconds` - (String) Defines the session expiration in seconds for the account. Valid values are Any whole number between between `900` and `86400`, and **NOT_SET** to unset account setting and use the service default.
- `session_invalidation_in_seconds` - (String) Defines the period of time in seconds in which a session is invalid due to inactivity. Valid values are Any whole number between `900` and `7200`, and **NOT_SET** to unset account setting and use the service default.
- `system_access_token_expiration_in_seconds` - (String) Defines the access token expiration in seconds. Valid values:  * Any whole number between '900' and '3600'  * NOT_SET - To unset account setting and use service default.
  - Constraints: The default value is `3600`.
- `system_refresh_token_expiration_in_seconds` - (String) Defines the refresh token expiration in seconds. Valid values:  * Any whole number between '900' and '2592000'  * NOT_SET - To unset account setting and use service default.
  - Constraints: The default value is `2592000`.
