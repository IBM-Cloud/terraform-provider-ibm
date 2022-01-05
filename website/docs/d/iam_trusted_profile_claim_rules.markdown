---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_claim_rules"
description: |-
  Get information about iam_trusted_profiles_claim_rules
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_trusted_profile_claim_rules

Retrieve list of IAM trusted profile claim rule as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about trusted profile claim rules, see [list claim rule for a trusted profile](https://cloud.ibm.com/apidocs/iam-identity-token-api#list-claim-rule)

## Example usage

```terraform
data "ibm_iam_trusted_profile_claim_rules" "iam_trusted_profiles_claim_rules" {
	profile_id = "profile_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the iam_trusted_profiles_claim_rules.
* `rules` - (List) List of claim rules.
    Nested scheme for **rules**:
	* `cr_type` - (String) The compute resource type. Not required if type is Profile-SAML. Valid values are VSI, IKS_SA, ROKS_SA.
	* `conditions` - (List) Conditions of this claim rule.
	    Nested scheme for **conditions**:
		* `claim` - (String) The claim to evaluate against.
		* `operator` - (String) The operation to perform on the claim. valid values are EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN.
		* `value` - (String) The stringified JSON value that the claim is compared to using the operator.
	* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
	* `entity_tag` - (String) version of the claim rule.
	* `expiration` - (Integer) Session expiration in seconds.
	* `id` - (String) the unique identifier of the claim rule.
	* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
	* `name` - (String) The optional claim rule name.
	* `realm_name` - (String) The realm name of the Idp this claim rule applies to.
	* `type` - (String) Type of the Calim rule, either `Profile-SAML` or `Profile-CR`.

