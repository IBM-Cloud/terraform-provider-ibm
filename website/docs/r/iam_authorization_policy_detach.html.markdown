---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_authorization_policy_detach"
description: |-
  Manages IBM IAM service authorizations detach.
---

# ibm_authorization_policy

Provides a resource for IAM Service Authorizations policy to be detached. This allows authorization policy to delete. For more information, about IAM service authorizations detach, see [using authorizations to grant access between services](https://cloud.ibm.com/docs/account?topic=account-serviceauth).

## Example usage

### Authorization policy detach

```terraform
resource "ibm_iam_authorization_policy_detach" "policy" {
  authorization_policy_id = "971164c3-add8-4ac3-bcb4-7376fd2a505e"
}

```

## Argument reference

Review the argument references that you can specify for your resource. 

- `authorization_policy_id` - (Required, Forces new resource, String) The authorization policy ID.

## Attribute reference
This resource does not provide attribute reference.