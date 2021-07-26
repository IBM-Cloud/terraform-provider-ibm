---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_dynamic_rule"
description: |-
  Manages Dynamic Rule for the IBM IAM access group.
---

# ibm_iam_access_group_dynamic_rule

Provides a resource for Dynamic Rule of an IAM access group. This allows rules to be created, updated and deleted.

Create, update, or delete a dynamic rule for an IAM access group. With dynamic rules, you can automatically add federated users to access groups based on specific identity attributes. When your users log in with a federated ID, the data from the identity provider dynamically maps your users to an access group based on the rules that you set. For more information, about IAM access group dynamic rule, see [reating dynamic rules for access groups](https://cloud.ibm.com/docs/account?topic=account-rules).


## Example usage

```terraform
resource "ibm_iam_access_group_dynamic_rule" "rule1" {
  name              = "newrule"
  access_group_id   = "AccessGroupId-dsnd4bvsaf"
  expiration        = 4
  identity_provider = "test-idp.com"
  conditions {
    claim    = "blueGroups"
    operator = "CONTAINS"
    value    = "\"test-bluegroup-saml\""
  }
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `access_group_id` - (Required, String) The ID of the access group.
- `conditions`- (Required, List) A list of conditions that the rule must satisfy.

  Nested scheme for `conditions`:
  - `claim` - (Required, String) The key value to evaluate the condition against. The key that you enter depends on what key-value pairs your identity provider provides. For example, your identity provider might include a key that is named `blueGroups` and that holds all the user groups that have access. To apply a condition for a specific user group within the `blueGroups` key, you specify `blueGroups` as your claim and add the value that you are looking for in `conditions.value`.
  - `operator` - (Required, String) The operation to perform on the claim. Supported values are `EQUALS`, `EQUALS_IGNORE_CASE`, `IN`, `NOT_EQUALS_IGNORE_CASE`, `NOT_EQUALS`, and `CONTAINS`.
  - `value` - (Required, String) The value that the claim is compared by using the `conditions.operator`.
- `expiration`- (Required, Integer) The number of hours that authenticated users can work in IBM Cloud before they must refresh their access. This value must be between 1 and 24.
- `identity_provider` - (Required, String) Enter the URI for your identity provider. This is the SAML `entity ID` field, which is sometimes referred to as the issuer ID, for the identity provider as part of the federation configuration for onboarding with IBMID. For example, `https://idp.example.org/SAML2`.
- `name` - (Required, String) The name of the dynamic rule for the IAM access group.


## Attribute reference

In addition to all arguments listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the dynamic rule. The ID is composed of `<access_group_ID>/<rule_ID>`.
- `rule_id` - (String) The ID of the rule.


## Import

The `iam_access_group_dynamic_rule` resource can be imported by using access group ID and rule ID.

**Syntax**

```
$ terraform import ibm_iam_access_group_dynamic_rule.example <access_group_ID>/<rule_ID>
```

**Example**

```
$ terraform import iam_access_group_dynamic_rule.example AccessGroupId-5391772e-1207-45e8-b032-2a21941c11ab/ClaimRule-3c5cd5fd-5b95-45f3-a693-08047eee56b5
```
