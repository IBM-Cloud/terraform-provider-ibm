---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profile_claim_rule"
description: |-
  Manages iam_trusted_profile_claim_rule.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profile_claim_rule

Create, update, or delete an IAM trusted profiles claim rule resource. For more information, about IAM trusted profiles claim rule, see https://cloud.ibm.com/apidocs/iam-identity-token-api#create-claim-rule

## Example usage

```terraform
resource "ibm_iam_trusted_profile_claim_rule" "iam_trusted_profile_claim_rule" {
  conditions = { "claim" : "claim", "operator" : "operator", "value" : "value" }
  profile_id = "profile_id"
  type = "type"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `conditions` - (Required, List) The conditions of this claim rule.
Nested scheme for **conditions**:
	* `claim` - (Required, String) The claim to evaluate against.
	* `operator` - (Required, String) The operation to perform on the claim. Supported values are EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN.
	* `value` - (Required, String) The stringified JSON value that the claim is compared to using the operator.
* `cr_type` - (Optional, String) The compute resource type the rule applies to, required only if type is specified as 'Profile-CR'. Supported values are VSI, IKS_SA, ROKS_SA.
* `expiration` - (Optional, Integer) Session expiration in seconds, only required if type is 'Profile-SAML'.
* `name` - (Optional, String) Name of the claim rule to be created or updated.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile to create a claim rule.
* `realm_name` - (Optional, String) The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'.
* `type` - (Required, String) The type of the calim rule, either 'Profile-SAML', or 'Profile-CR'.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the `iam_trusted_profiles_claim_rule`.
* `created_at` - (String) If set contains a date time string of the creation date in ISO format.
* `entity_tag` - (String) The version of the claim rule.
* `modified_at` - (String) If set contains a date time string of the last modification date in ISO format.
