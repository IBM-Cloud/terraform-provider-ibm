---
layout: "ibm"
page_title: "IBM : ibm_scc_template_attachment"
description: |-
  Manages scc_template_attachment.
subcategory: "Security and Compliance Center"
---

# ibm_scc_template_attachment

Provides a resource for scc_template_attachment. This allows scc_template_attachment to be created, updated and deleted. For more information about Security and Compliance Center template attachments, see [Applying Templates](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-templates-apply&interface=ui).

~> **NOTE**: This resource depends on a `ibm_scc_template` to be created before creating this resource. The object `ibm_scc_template_attachment` must attach to an exiting template.

## Example Usage

```hcl
resource "ibm_scc_template" scc_template_instance {
    // example of a ibm_scc_template that needs to be used in conjunction with ibm_scc_template_attachment
}

resource "ibm_scc_template_attachment" "scc_template_attachment_instance" {
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
	template_id = ibm_scc_template.scc_template_instance.id // from the resource ibm_scc_template
	depends_on  = [
		ibm_scc_template.scc_template_instance          // ensures that the template is created first
	]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Required, String) Your IBM Cloud account ID.
* `included_scope` - (Required, List) The extent at which the template can be attached across your accounts.
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
* `template_id` - (Required, Forces new resource, String) The UUID that uniquely identifies the template.
* `excluded_scopes` - (Optional, List) The extent at which the template can be excluded from the included scope.
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
* `version` - Version of the ibm_scc_template_attachment.

## Import

You can import the `ibm_scc_template_attachment` resource by using `attachment_id`.
The `attachment_id` property can be formed from `template_id`, and `attachment_id` in the following format:

```
<template_id>/<attachment_id>
```
* `template_id`: A string. The UUID that uniquely identifies the template.
* `attachment_id`: A string. The UUID that uniquely identifies the attachment.

# Syntax
```
$ terraform import ibm_scc_template_attachment.scc_template_attachment <template_id>/<attachment_id>
```
