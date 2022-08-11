---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_role_actions"
description: |-
  Manages IBM IAM role actions.
---

# ibm_iam_role_actions

Retrieve a list of actions for an IBM Cloud service that are included in an IAM service access role.  For more information, about IAM role action, see [actions and roles for account management services](https://cloud.ibm.com/docs/account?topic=account-account-services#account-management-actions-roles).

## Example usage

```terraform
data "ibm_iam_role_actions" "test" {
  service = "kms"
}

```

## Argument reference

Review the argument references that you can specify for your data source.

- `service` - (Required, String) The name of the IBM Cloud service for which you want to list supported actions. For account management services, you can find supported values in the [documentation](https://cloud.ibm.com/docs/account?topic=account-account-services#api-acct-mgmt). For other services, run the `ibmcloud catalog service-marketplace` command and retrieve the value from the **Name** column of your command line output.

## Attribute reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the service.
- `actions`- (Map of (string, string)) A map containing all roles and actions in key value format. The key contains a string equal to the role name and value contains a string of all the actions separated by a comma (",").
- `manager`- (List of strings) A list of supported actions that require the **Manager** service access role.
- `reader`- (List of strings) A list of supported actions that require the **Reader** service access role.
- `reader_plus`- (List of strings) A list of supported actions that require the **Reader plus** service access role.
- `writer`- (List of strings) A list of supported actions that require the **Writer** service access role.



  
