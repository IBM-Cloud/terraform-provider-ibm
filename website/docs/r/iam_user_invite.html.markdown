---
layout: "ibm"
page_title: "IBM : iam_user_invite"
sidebar_current: "docs-ibm-resource-iam-user-invite"
description: |-
  Manages IBM IAM User Invite.
---

# ibm\_iam_user_invite

Provides a resource for IAM User Invite. This allows batch of users or single user to be invited, updated and deleted. User to be invited can be added to one or more access groups

## Example Usage

### Inviting batch of Users

```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
}

```

### Inviting batch of Users with access groups
```hcl
resource "ibm_iam_user_invite" "invite_user" {
    users = ["test@in.ibm.com"]
    access_groups = ["accessgroup-id-9876543210"]
}

```

## Argument Reference

The following arguments are supported:

* `users` - (Required, list) comma separated list of users email-id. 
* `access_groups` - (Optional, list) comma seperated list of access group ids.

## Import

Import functionality not supported for this resource.