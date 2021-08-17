---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_claim_rule"
description: |-
  Get information about iam_trusted_profile_claim_rule
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_claim_rule

Retrieve information about IAM trusted profile claim rule as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about trusted profile claim rules, see [Create claim rule for a trusted profile](https://cloud.ibm.com/apidocs/iam-identity-token-api#create-claim-rule)

## Example usage

```terraform
data "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
  profile_id = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.profile_id
  rule_id    = ibm_iam_trusted_profile_claim_rule.iam_trusted_profile_claim_rule.rule_id
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, Forces new resource, String) The ID of the trusted profile.
* `rule_id` - (Required, Forces new resource, String) ID of the claim rule to fetch.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `conditions` - (List) Conditions of this claim rule.
    Nested scheme for **conditions**:
	* `claim` - (String) The claim to evaluate against.
	* `operator` - (String) The operation to perform on the claim. Supported values are **EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN**.
	* `value` - (String) The stringified JSON value that the claim is compared to using the operator.

* `cr_type` - (String) The compute resource type. The compute resource type not required if type is set as Profile-SAML. Valid values are **VSI, IKS_SA, ROKS_SA**

* `created_at` - (String) If set contains a date time string of the creation date in ISO format.

* `entity_tag` - (String) The version of the claim rule.

* `expiration` - (Integer) The session expiration in seconds.

* `id` - (String) Id is combination of `profile_id`/ `rule_id`.

* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.

* `name` - (String) The optional claim rule name.

* `realm_name` - (String) The realm name of the Identity Provider(Idp) this claim rule applies to.

* `type` - (String) Type of the Calim rule. Supported values are **Profile-SAML** or **Profile-CR**.

