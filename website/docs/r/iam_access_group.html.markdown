---
layout: "ibm"
page_title: "IBM : iam_access_group"
sidebar_current: "docs-ibm-resource-iam-access-group"
description: |-
  Manages IBM IAM Access Group.
---

# ibm\_iam_access_group

Provides a resource for IAM access group. This allows access group to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_iam_access_group" "accgrp" {
  name        = "test"
  description = "New access group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the access group.
* `description` - (Optional, string) Description of the access group.
* `tags` - (Optional, array of strings) Tags associated with the IAM access group.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the access group.
* `version` - Version of the access group.
