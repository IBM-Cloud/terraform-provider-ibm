---
layout: "ibm"
page_title: "IBM : iam_user_invite"
sidebar_current: "docs-ibm-resource-iam-user-invite"
description: |-
  Manages IBM IAM User Invite.
---

# ibm\_iam_user_invite

Provides a resource for IAM User Invite. This allows batch of users or single user to be invited, updated and deleted.

## Example Usage

### Inviting batch of Users

```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
}

```

## Argument Reference

The following arguments are supported:

* `users` - (Required, list) comma separated list of users email-id. 

## Import

Import functionality not supported for this resource.