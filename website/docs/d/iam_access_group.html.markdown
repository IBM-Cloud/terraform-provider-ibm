---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group"
description: |-
  Get information about IBM IAM Access Group and all the members and dynamic rules associated with the group.
---

# ibm_iam_access_group

Retrieve information about an [IAM Access Group](https://cloud.ibm.com/iam/groups). Access groups can be used to define a set of permissions that you want to grant to a group of users.

## Example usage


```terraform
data "ibm_iam_access_group" "accgroup" {
  access_group_name = ibm_iam_access_group.accgroup.name
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `access_group_name` - (Optional, String) The name of the access group that you want to retrieve details for. If no access group is specified, all access groups that exist in the IBM Cloud account are returned. 

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `groups`- (List) A list of IAM access groups that are set up for an IBM Cloud account.

  Nested scheme for `groups`:
  
  - `description` - (String) The description of the IAM access group.
  - `iam_service_ids` - (Array of Strings) A list of service IDs that belong to the access group.
  - `iam_profile_ids` - (Array of Strings) A list of trusted profile IDs that belong to the access group.
  - `ibm_ids` - (Array of Strings) A list of IBM ID that belong to the access group.
  - `id` - (String) The ID of the IAM access group.
  - `name` - (String) The name of the IAM access group.
  - `rules`- (List) A list of dynamic rules that are applied to the IAM access group.

    Nested scheme for `rules`:
	- `conditions`- (List) A list of conditions that the rule must satisfy.
	  
	   Nested scheme for `conditions`:
	   - `claim` - (String) The key value to evaluate the condition against. The key depends on what key-value pairs your identity provider provides. For example, your identity provider might include a key that is named `blueGroups` and that holds all the user groups that have access. To apply a condition for a specific user group within the `blueGroups` key, you specify `blueGroups` as your claim and add the value that you are looking for in `value`.
	   - `operator` - (String) The operation to perform on the claim. Supported values are `EQUALS`, `QUALS_IGNORE_CASE`, `IN`, `NOT_EQUALS_IGNORE_CASE`, `NOT_EQUALS`, and `CONTAINS`.
	   - `value` - (String) The value that the claim is compared to by using the `operator`.
	- `expiration`- (Integer) The number of hours that authenticated users can work in IBM Cloud before they must refresh their access.
	- `identity_provider` - (String) The URI of your identity provider. This is the SAML "entity ID" field, which is sometimes referred to as the issuer ID, for the identity provider as part of the federation configuration for onboarding with IBMID.
	- `name` - (String) The name of the dynamic rule.
	- `rule_id` - (String) The ID of the dynamic rule.
