---
layout: "ibm"
page_title: "IBM : ibm_cr_retention_policy"
description: |-
  Manages retention policies in IBM Cloud Container Registry.
subcategory: "Container Registry"
---

# ibm_cr_retention_policy

Provides a resource for ibm_cr_retention_policy. You can use this resource to create, update, and delete a retention policy.

## Example Usage

```terraform
resource "ibm_cr_retention_policy" "cr_retention_policy" {
  namespace = "birds"
  images_per_repo = 10
  retain_untagged = false
}
```

## Argument Reference

The following arguments are supported:

- `namespace` - (Required, string) The namespace to which the retention policy is attached.
- `images_per_repo` - (Required, int) Determines how many images are retained in each repository when the retention policy is processed. The value -1 denotes 'Unlimited' (all images are retained).
- `retain_untagged` - (Optional, bool) Determines whether untagged images are retained when the retention policy is processed. The value is false by default, which means that untagged images can be deleted when the policy runs.

## Attribute Reference

In addition to the arguments in the Argument Reference section, the following attributes are exported:

- `id` - The unique identifier of the cr_retention_policy. This identifier is the same as the name of namespace to which the retention policy is attached.

## Import

You can import the `ibm_cr_retention_policy` resource by adding the `namespace` option, where `namespace` is the namespace to which the retention policy is attached.

```
$ terraform import ibm_cr_retention_policy.cr_retention_policy <namespace>
```
