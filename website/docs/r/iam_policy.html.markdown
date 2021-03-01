---
layout: "ibm"
page_title: "IBM : iam_policy"
sidebar_current: "docs-ibm-resource-iam-policy"
description: |-
  Manages iam_policy.
---

# ibm\_iam_policy

Provides a resource for iam_policy. This allows iam_policy to be created, updated and deleted.

## Example Usage

```hcl
resource "iam_policy" "iam_policy" {
  type = "type"
  subjects = { example: "object" }
  roles = { example: "object" }
  resources = { example: "object" }
}
```

## Argument Reference

The following arguments are supported:

* `type` - (Required, string) The policy type; either 'access' or 'authorization'.
* `subjects` - (Required, List) The subjects associated with a policy.
  * `attributes` - (Optional, []interface{}) List of subject attributes.
* `roles` - (Required, List) A set of role cloud resource names (CRNs) granted by the policy.
  * `role_id` - (Required, string) The role cloud resource name granted by the policy.
  * `display_name` - (Optional, string) The display name of the role.
  * `description` - (Optional, string) The description of the role.
* `resources` - (Required, List) The resources associated with a policy.
  * `attributes` - (Optional, []interface{}) List of resource attributes.
* `description` - (Optional, string) Customer-defined description.
* `accept_language` - (Optional, string) Translation language code.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the iam_policy.
* `href` - The href link back to the policy.
* `created_at` - The UTC timestamp when the policy was created.
* `created_by_id` - The iam ID of the entity that created the policy.
* `last_modified_at` - The UTC timestamp when the policy was last modified.
* `last_modified_by_id` - The iam ID of the entity that last modified the policy.
