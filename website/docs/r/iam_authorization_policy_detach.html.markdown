---

subcategory: "Identity & Access (IAM)"
layout: "ibm"
page_title: "IBM : iam_authorization_policy_detach"
description: |-
  Manages IBM IAM Service Authorizations detach.
---

# ibm\_authorization_policy

Provides a resource for IAM Service Authorizations policy to be detached. This allows authorization policy to deleted.

## Example Usage

### Authorization policy detach

```hcl
resource "ibm_iam_authorization_policy_detach" "policy" {
  authorization_policy_id = "971164c3-add8-4ac3-bcb4-7376fd2a505e"
}

```

## Argument Reference

The following arguments are supported:

* `authorization_policy_id` - (Required, Forces new resource, string) The valid authorization policy ID.