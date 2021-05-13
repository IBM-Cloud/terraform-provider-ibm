---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_account_settings"
description: |-
  Get information about iam_account_settings
---

# ibm\_iam_account_settings

Provides a read-only data source for iam_account_settings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_account_settings" "iam_account_settings" {
}
```

## Argument Reference

The following arguments are supported:

* `include_history` - (Optional, boolean) Defines if the entity history is included in the response.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the iam_account_settings.

* `account_id` - Unique ID of the account.

* `restrict_create_service_id` - Defines whether or not creating a Service Id is access controlled.
  
* `restrict_create_platform_apikey` - Defines whether or not creating platform API keys is access controlled. 
  
* `allowed_ip_addresses` - Defines the IP addresses and subnets from which IAM tokens can be created for the account.

* `entity_tag` - Version of the account settings.

* `mfa` - Defines the MFA trait for the account.

* `history` - History of the Account Settings. Nested `history` blocks have the following structure:
	* `timestamp` - Timestamp when the action was triggered.
	* `iam_id` - IAM ID of the identity which triggered the action.
	* `iam_id_account` - Account of the identity which triggered the action.
	* `action` - Action of the history entry.
	* `params` - Params of the history entry.
	* `message` - Message which summarizes the executed action.

* `session_expiration_in_seconds` - Defines the session expiration in seconds for the account.

* `session_invalidation_in_seconds` - Defines the period of time in seconds in which a session will be invalidated due  to inactivity.

* `max_sessions_per_identity` - Defines the max allowed sessions per identity required by the account.