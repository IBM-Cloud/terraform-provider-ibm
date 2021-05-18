---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group_dynamic_rule"
description: |-
  Manages Dynamic Rule for IBM IAM Access Group.
---

# ibm\_iam_access_group_dynamic_rule

Provides a resource for Dynamic Rule of an IAM access group. This allows rules to be created, updated and deleted.

## Example Usage

```hcl
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

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the dynamic rule.
* `access_group_id` - (Required, string) ID of the access group.

* `expiration` - (Required, int) The number of hours that the rule lives for (Must be between 1 and 24).
* `identity_provider` - (Required, string) The url of the identity provider.  
* `conditions` - (Required, list) A list of conditions the rule must satisfy:
  * `claim` - (Required, string) The claim to evaluate against. This will be found in the ext claims of a user's login request. 
  * `operator` - (Required, string) The operation to perform on the claim. Valid operators are EQUALS, EQUALS_IGNORE_CASE, IN, NOT_EQUALS_IGNORE_CASE, NOT_EQUALS, and CONTAINS.
  * `value` - (Required, string) The stringified JSON value that the claim is compared to using the operator.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the dynamic rule resource. The id is composed of \<access_group_id\>/\<rule_id\>.<br/>
* `rule_id` - The id of the rule.


## Import

iam_access_group_dynamic_rule can be imported using access group ID and rule id, eg

```
$ terraform import iam_access_group_dynamic_rule.example AccessGroupId-5391772e-1207-45e8-b032-2a21941c11ab/ClaimRule-3c5cd5fd-5b95-45f3-a693-08047eee56b5
```