---
layout: "ibm"
page_title: "IBM : ibm_iam_access_management_account_settings"
description: |-
  Get information about access_management_account_settings
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_management_account_settings

Retrieve information about an existing `iam_access_management_account_settings` data sources. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_access_management_account_settings" "settings" {
    account_id = "accountId-01"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - The account ID that the Access Management Account Settings belong to.
* `accept_language` - (Optional) Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `external_account_identity_interaction` - (Set) Specifies how external accounts can interact in relation to the requested account.
Nested schema for **external_account_identity_interaction**:
  * `identity_types` - (Set) The settings for each identity type.
  Nested schema for **identity_types**:
      * `user` - (Set) The core set of properties associated with a user identity type.
      Nested schema for **user**:
          * `state` - (String) The state of the user identity type.
          * `external_allowed_accounts` - (List) List of accounts that the state applies to for the user identity type.
      * `service_id` - (Set) The core set of properties associated with a serviceId identity type.
      Nested schema for **user**:
          * `state` - (String) The state of the serviceId identity type.
          * `external_allowed_accounts` - (List) List of accounts that the state applies to for the serviceId identity type.
      * `service` - (Set) The core set of properties associated with a service identity type.
      Nested schema for **user**:
          * `state` - (String) The state of the service identity type.
          * `external_allowed_accounts` - (List) List of accounts that the state applies to for the service identity type.
