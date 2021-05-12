---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_account_settings"
description: |-
  Manages iam_account_settings.
---

# ibm\_iam_account_settings

Provides a resource for iam_account_settings. This allows iam_account_settings to be created and updated.

## Example Usage

```hcl
resource "ibm_iam_account_settings" "iam_account_settings_instance" {
  mfa = "LEVEL3"
  session_expiration_in_seconds = "40000"
}
```

## Argument Reference

The following arguments are supported:

* `include_history` - (Optional, boolean) Defines if the entity history is included in the response.
* `if_match` - (Optional, string) Version of the account settings to be updated, if no value is supplied then the default value `*` will be used to indicate to update any version available. This might result in stale updates.
* `restrict_create_service_id` - (Optional, string) Defines whether or not creating a Service Id is access controlled. Valid values:  
  * RESTRICTED - to apply access control  
  * NOT_RESTRICTED - to remove access control  
  * NOT_SET - to 'unset' a previous set value.
* `restrict_create_platform_apikey` - (Optional, string) Defines whether or not creating platform API keys is access controlled. Valid values:  
  * RESTRICTED - to apply access control  
  * NOT_RESTRICTED - to remove access control  
  * NOT_SET - to 'unset' a previous set value.
* `allowed_ip_addresses` - (Optional, string) Defines the IP addresses and subnets from which IAM tokens can be created for the account. Value should be a comma separated string. 
* `mfa` - (Optional, string) Defines the session expiration in seconds for the account.  Valid values:  
  * NONE - No MFA trait set  
  * TOTP - For all non-federated IBMId users
  * TOTP4ALL - For all users
  * LEVEL1 - Email-based MFA for all users
  * LEVEL2 - TOTP-based MFA for all users
  * LEVEL3 - U2F MFA for all users.
* `session_expiration_in_seconds` - (Optional, string) Defines the session expiration in seconds for the account. Valid values:  
  * Any whole number between between '900' and '86400'  
  * NOT_SET - To unset account setting and use service default.
* `session_invalidation_in_seconds` - (Optional, string) Defines the period of time in seconds in which a session will be invalidated due to inactivity. Valid values:
  * Any whole number between '900' and '7200'
  * NOT_SET - To unset account setting and use service default.
* `max_sessions_per_identity` - (Optional, string) Defines the max allowed sessions per identity required by the account. Valid values:
  * Any whole number greater than '0' 
  * NOT_SET - To unset account setting and use service default.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `history` - The update history of the settings instance
* `entity_tag` - The version of the account settings object. You need to specify this value when updating the account settings to avoid stale updates.
* `restrict_create_service_id` - Defines whether or not creating a Service Id is access controlled.
* `restrict_create_platform_apikey` - Defines whether or not creating platform API keys is access controlled.
* `allowed_ip_addresses` - Defines the IP addresses and subnets from which IAM tokens can be created for the account. Value should be a comma separated string.* 
* `mfa` - Defines the session expiration in seconds for the account.
* `session_expiration_in_seconds` - Defines the session expiration in seconds for the account.
* `session_invalidation_in_seconds` - Defines the period of time in seconds in which a session will be invalidated due to inactivity.
* `max_sessions_per_identity` - Defines the max allowed sessions per identity required by the account.
* `account_id` - Unique Id of the account.
* `id` - Unique Id of the account settings instance