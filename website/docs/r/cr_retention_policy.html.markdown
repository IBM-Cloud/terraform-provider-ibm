---
layout: "ibm"
page_title: "IBM : cr_retention_policy"
sidebar_current: "docs-ibm-resource-cr-retention-policy"
description: |-
  Manages cr_retention_policy.
---

# ibm\_cr_retention_policy

Provides a resource for cr_retention_policy. This allows cr_retention_policy to be created, updated and deleted.

## Example Usage

```hcl
resource "cr_retention_policy" "cr_retention_policy" {
  namespace = "birds"
  images_per_repo = 10
  retain_untagged = false
}
```

## Argument Reference

The following arguments are supported:

* `namespace` - (Required, string) The namespace to which the retention policy is attached.
* `images_per_repo` - (Optional, int) Determines how many images will be retained for each repository when the retention policy is executed. The value -1 denotes 'Unlimited' (all images are retained).
* `retain_untagged` - (Optional, bool) Determines if untagged images are retained when executing the retention policy. This is false by default meaning untagged images will be deleted when the policy is executed.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cr_retention_policy.
