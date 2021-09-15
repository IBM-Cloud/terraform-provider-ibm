---
layout: "ibm"
page_title: "IBM : ibm_iam_trusted_profiles_claim_rule"
description: |-
  Manages iam_trusted_profiles_claim_rule.
subcategory: "IAM Identity Services"
---

# ibm_iam_trusted_profiles_claim_rule

Provides a resource for iam_trusted_profiles_claim_rule. This allows iam_trusted_profiles_claim_rule to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_iam_trusted_profiles_claim_rule" "iam_trusted_profiles_claim_rule" {
  conditions = { "claim" : "claim", "operator" : "operator", "value" : "value" }
  profile_id = "profile_id"
  type = "type"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `conditions` - (Required, List) Conditions of this claim rule.
Nested scheme for **conditions**:
	* `claim` - (Required, String) The claim to evaluate against.
	* `operator` - (Required, String) The operation to perform on the claim. valid values are EQUALS, NOT_EQUALS, EQUALS_IGNORE_CASE, NOT_EQUALS_IGNORE_CASE, CONTAINS, IN.
	* `value` - (Required, String) The stringified JSON value that the claim is compared to using the operator.
* `context` - (Optional, List) Context with key properties for problem determination.
Nested scheme for **context**:
	* `transaction_id` - (Optional, String) The transaction ID of the inbound REST request.
	* `operation` - (Optional, String) The operation of the inbound REST request.
	* `user_agent` - (Optional, String) The user agent of the inbound REST request.
	* `url` - (Optional, String) The URL of that cluster.
	* `instance_id` - (Optional, String) The instance ID of the server instance processing the request.
	* `thread_id` - (Optional, String) The thread ID of the server instance processing the request.
	* `host` - (Optional, String) The host of the server instance processing the request.
	* `start_time` - (Optional, String) The start time of the request.
	* `end_time` - (Optional, String) The finish time of the request.
	* `elapsed_time` - (Optional, String) The elapsed time in msec.
	* `cluster_name` - (Optional, String) The cluster name.
* `cr_type` - (Optional, String) The compute resource type the rule applies to, required only if type is specified as 'Profile-CR'. Valid values are VSI, IKS_SA, ROKS_SA.
* `expiration` - (Optional, Integer) Session expiration in seconds, only required if type is 'Profile-SAML'.
* `name` - (Optional, String) Name of the claim rule to be created or updated.
* `profile_id` - (Required, Forces new resource, String) ID of the trusted profile to create a claim rule.
* `realm_name` - (Optional, String) The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'.
* `type` - (Required, String) Type of the calim rule, either 'Profile-SAML' or 'Profile-CR'.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the iam_trusted_profiles_claim_rule.
* `created_at` - (Required, String) If set contains a date time string of the creation date in ISO format.
* `entity_tag` - (Required, String) version of the claim rule.
* `modified_at` - (Optional, String) If set contains a date time string of the last modification date in ISO format.

## Import

You can import the `ibm_iam_trusted_profiles_claim_rule` resource by using `id`.
The `id` property can be formed from `profile-id`, and `rule-id` in the following format:

```
<profile-id>/<rule-id>
```
* `profile-id`: A string. ID of the trusted profile to create a claim rule.
* `rule-id`: A string. ID of the claim rule to delete.

# Syntax
```
$ terraform import ibm_iam_trusted_profiles_claim_rule.iam_trusted_profiles_claim_rule <profile-id>/<rule-id>
```
