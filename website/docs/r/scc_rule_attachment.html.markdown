---
layout: "ibm"
page_title: "IBM : ibm_scc_rule_attachment"
description: |-
  Manages scc_rule_attachment.
subcategory: "Security and Compliance Center"
---

# ibm_scc_rule_attachment

Provides a resource for ibm_scc_rule_attachment. This allows ibm_scc_rule_attachment to be created, updated and deleted. For more information about Security and Compliance Center rule attachments, see [Applying Rules](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-rules-apply&interface=ui).

~> **NOTE**: This resource depends on a `ibm_scc_rule` to be created before creating this resource. The object `ibm_scc_rule_attachment` must attach to an exiting rule.

## Example Usage

```hcl
resource "ibm_scc_rule" scc_rule_instance {
	// example of a ibm_scc_rule that needs to be used in conjunction with ibm_scc_rule_attachment
}

resource "ibm_scc_rule_attachment" "scc_rule_attachment_instance" {
	account_id = "thisIsAFake32CharacterAccountID"
	included_scope {
		note       = "This is a note to reference my account"
		scope_id   = "thisIsAFake32CharacterAccountID" // value determined by scope type
		scope_type = "account"
	}
	excluded_scopes {
		note       = "This is a note to exclude a specific resource group"
		scope_id   = "<resource_group_id>"     // value determined by scope type
		scope_type = "account.resource_group"
	}
	rule_id    = ibm_scc_rule.scc_rule_instance.id // from the resource ibm_scc_rule
	depends_on = [
		ibm_scc_rule.scc_rule_instance          // ensures that the rule is created first
	]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Required, String) Your IBM Cloud account ID.
* `included_scope` - (Required, List) The extent at which the rule can be attached across your accounts.
Nested scheme for **included_scope**:
	* `note` - (Optional, String) A short description or alias to assign to the scope.
	* `scope_id` - (Required, String) The ID of the scope, such as an enterprise, account, or account group, that you want to evaluate.
	* `scope_type` - (Required, String) The type of scope that you want to evaluate.
		* Constraints: Allowable values are:
			* `enterprise`,
			* `enterprise.account_group`,
			* `enterprise.account`,
			* `account`,
			* `account.resource_group`.
    * `Constraints`: Only one `included_scope` item is allowed
* `rule_id` - (Required, Forces new resource, String) The UUID that uniquely identifies the rule.
* `excluded_scopes` - (Optional, List) The extent at which the rule can be excluded from the included scope.
Nested scheme for **excluded_scopes**:
	* `note` - (Optional, String) A short description or alias to assign to the scope.
	* `scope_id` - (Required, String) The ID of the scope, such as an enterprise, account, or account group, that you want to evaluate.
	* `scope_type` - (Required, String) The type of scope that you want to evaluate.
		* Constraints: Allowable values are:
    		* `enterprise`,
    		* `enterprise.account_group`,
    		* `enterprise.account`,
    		* `account`,
    		* `account.resource_group`.
## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `attachment_id` - (Required, String) The UUID that uniquely identifies the attachment
* `version` - Version of the ibm_scc_rule_attachment.

## Import

You can import the `ibm_scc_rule_attachment` resource by using `attachment_id`.
The `attachment_id` property can be formed from `rule_id`, and `attachment_id` in the following format:

```
<rule_id>/<attachment_id>
```
* `rule_id`: A string. The UUID that uniquely identifies the rule.
* `attachment_id`: A string. The UUID that uniquely identifies the attachment.

# Syntax
```
$ terraform import ibm_scc_rule_attachment.scc_rule_attachment <rule_id>/<attachment_id>
```
