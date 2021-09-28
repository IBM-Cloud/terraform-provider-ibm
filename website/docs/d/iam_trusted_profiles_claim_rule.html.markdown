---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_claim_rule"
description: |-
  Get information about iam_trusted_profile_claim_rule
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_claim_rule

Retrieve information about IAM trusted profile claim rule as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about trusted profile claim rules, see [Create claim rule for a trusted profile](https://cloud.ibm.com/apidocs/iam-identity-token-api#create-claim-rule)

## Example usage

```terraform
data "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
	profile_id = "profile_id"
	rule_id = "rule_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, Forces new resource, String) The ID of the trusted profile.
* `rule_id` - (Required, Forces new resource, String) ID of the claim rule to fetch.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the `iam_trusted_profile_claim_rule`.
* `conditions` - (Required, List) Conditions of this claim rule.
    Nested scheme for **conditions**:
	* `claim` - (Required, String) The claim to evaluate against.
	* `operator` - (Required, String) The operation to perform on the claim. Supported values are **EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN**.
	* `value` - (Required, String) The stringified JSON value that the claim is compared to using the operator.

* `cr_type` - (Optional, String) The compute resource type. The compute resource type not required if type is set as Profile-SAML. Valid values are **VSI, IKS_SA, ROKS_SA**

* `created_at` - (Required, String) If set contains a date time string of the creation date in ISO format.

* `entity_tag` - (Required, String) The version of the claim rule.

* `expiration` - (Required, Integer) The session expiration in seconds.

* `id` - (Required, String) the unique identifier of The claim rule.

* `modified_at` - (Optional, String) If set contains a date time string of the last modification date in ISO format.

* `name` - (Optional, String) The optional claim rule name.

* `realm_name` - (Optional, String) The realm name of the Identity Provider(Idp) this claim rule applies to.

* `type` - (Required, String) Type of the Calim rule. Supported values are **Profile-SAML** or **Profile-CR**.

