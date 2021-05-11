---
subcategory: "Cloud Foundry"
layout: "ibm"
page_title: "IBM: ibm_account"
description: |-
  Get information about an IBM Cloud account.
---

# `ibm_account`

Retrieve information about an existing IBM Cloud account. For more information, about IBM account, see [How do I create an IBM Cloud account?](https://cloud.ibm.com/docs/account?topic=account-accountfaqs).

**Important**

Import the details of an existing IBM Cloud account as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.


## Example usage
The following example retrieves information about an IBM Cloud account that belongs to the `myorg` Cloud Foundry organization. 

```
data "ibm_org" "orgData" {
  org = "example.com"
}

data "ibm_account" "accountData" {
  org_guid = data.ibm_org.orgData.id
}
```

## Argument reference
Review the input parameters that you can specify for your data source.

- `org_guid` - (Required, String) The GUID of the IBM Cloud organization. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs --guid` command.

## Attribute reference
Review the output parameters that you can access after you retrieved your data source. 

- `account_users` - (List of Objects) The list of account user's in the account. The nested `account_users` has the following structure:
	- `email` - (String) The email address of the account user.
	- `id` - (String) The user ID of the account user.
	- `role` -  (String) The Cloud Foundry account role that is assigned to the account user.
	- `state` - (String) The state of the account user.
- `id` - (String) The unique identifier of the IBM Cloud account.


