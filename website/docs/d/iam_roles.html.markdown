---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_roles"
description: |-
  Manages IBM IAM roles.
---

# ibm_iam_roles

Retrieve information about supported IAM roles for an IBM Cloud service. For more information, about IAM role action, see [actions and roles for account management services](https://cloud.ibm.com/docs/account?topic=account-account-services#account-management-actions-roles).

## Example usage

```terraform
data "ibm_iam_roles" "test" {
  service = "kms"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `service` - (Optional, String) The name of the IBM Cloud service for which you want to list supported IAM  For account management services, you can find supported values in the [documentation](https://cloud.ibm.com/docs/account?topic=account-account-services#api-acct-mgmt). For other services, run the `ibmcloud catalog service-marketplace` command and retrieve the value from the **Name** column of your command line output.

## Attribute reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The ID of your IBM Cloud account.
- `roles`- (List) A list of supported IAM service access, platform, and custom roles for an IBM Cloud service.
	- `description` - (String) The description of the role.
	- `name` - (String) The name of the role.
	- `type` - (String) The type of the role. Supported values are `service`, `platform`, and `custom`.



  
