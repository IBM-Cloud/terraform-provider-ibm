---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_authorization_policy"
description: |-
  Get information about an IBM IAM service authorizations.
---

# ibm_iam_authorization_policies

Retrieve information about an IAM service authorization policy. For more information, about IAM service authorizations, see [using authorizations to grant access between services](https://cloud.ibm.com/docs/account?topic=account-serviceauth).

## Example usage

```terraform
data "ibm_iam_authorization_policies" "testacc_ds_authorization_policy" {
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `account_id` - (Optional, String) An alpha-numeric value identifying the account ID.
- `transaction_id`- (Optional, String) The TransactionID can be passed to your request for the tracking calls.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `policies` - (List) A nested block describes IAM Authorization Policies in an account.

  Nested scheme for `policies`:
  - `description`  (String) The description of the IAM User Policy.
  - `id` - (String) The unique identifier of the IAM user policy. The ID is composed of `<account_id>/<authorization_policy_id>`.
  - `roles`-  (String) The roles that are assigned to the policy.
  - `resources`- (List of objects) A nested block describes the resources in the policy.

    Nested scheme for `resources`:
    - `source_service_account` - (Optional, Forces new resource, string) The account GUID of source service.
    - `source_service_name` - (Required, Forces new resource, string) The source service name.
    - `target_service_name` - (Required, Forces new resource, string) The target service name.
    - `source_resource_instance_id` - (Optional, Forces new resource, string) The source resource instance id.
    - `target_resource_instance_id` - (Optional, Forces new resource, string) The target resource instance id.
    - `source_resource_type` - (Optional, Forces new resource, string) The resource type of source service.
    - `target_resource_type` - (Optional, Forces new resource, string) The resource type of target service.
    - `source_resource_group_id` - (Optional, Forces new resource, string) The source resource group id.
    - `target_resource_group_id` - (Optional, Forces new resource, string) The target resource group id.
