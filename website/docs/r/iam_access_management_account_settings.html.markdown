---
layout: "ibm"
page_title: "IBM : ibm_iam_access_management_account_settings"
description: |-
  Manages access_management_account_settings
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_management_account_settings

Update, and reset an `iam_access_management_account_settings` with this resource. 

**Note**: The resource is already initialized with default values. Therefore, create operation does not actually create the resource. 
As a result, the `terraform apply` command would apply the values supplied 
in the plan and this would override any existing settings values previously set. 
Also note that the delete operation (`terraform destroy`) resets the resource with default values.


## Example Usage

```hcl
resource "ibm_iam_access_management_account_settings" "settings" {
    account_id = "accountId-01"

    external_account_identity_interaction {
      identity_types {
        user {
          state                     = "monitor"
          external_allowed_accounts = ["accountId-02", "accountId-03"]
        }

        service_id {
          state                     = "enabled"
          external_allowed_accounts = ["accountId-02", "accountId-04"]
        }

        service {
          state                     = "limited"
          external_allowed_accounts = ["accountId-03"]
        }
      }
    }
}
```

## Argument Reference

Following arguments can be specified for this resource.

* `account_id` - The account ID that the Access Management Account Settings belong to.
* `accept_language` - (Optional) Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).
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

## Attribute reference
All argument reference list can be accessed after the resource is fetched/modified.
