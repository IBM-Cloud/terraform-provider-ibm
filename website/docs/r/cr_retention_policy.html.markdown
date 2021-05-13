---
layout: "ibm"
page_title: "IBM : cr_retention_policy"
description: |-
  Manages cr_retention_policy.
subcategory: "Container Registry"
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
* `images_per_repo` - (Required, int) Determines how many images will be retained for each repository when the retention policy is executed. The value -1 denotes 'Unlimited' (all images are retained).
* `retain_untagged` - (Optional, bool) Determines if untagged images are retained when executing the retention policy. This is false by default meaning untagged images will be deleted when the policy is executed.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cr_retention_policy.

## Import

You can import the `cr_retention_policy` resource by using `namespace`. The namespace to which the retention policy is attached.

```
$ terraform import cr_retention_policy.cr_retention_policy <namespace>
```
