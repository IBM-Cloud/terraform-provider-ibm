---
layout: "ibm"
page_title: "IBM : ibm_cr_retention_policy"
description: |-
  Manages retention policies in IBM Cloud Container Registry.
subcategory: "Container Registry"
---

# ibm_cr_retention_policy

Create, update, and delete about the IBM Cloud Container Registry retention policy resource. For more information, about IBM Cloud Container Registry retention policy, see [Managing access for Container Registry](https://cloud.ibm.com/docs/Registry?topic=Registry-iam).

## Example usage

```terraform
resource "ibm_cr_retention_policy" "cr_retention_policy" {
  namespace = "birds"
  images_per_repo = 10
  retain_untagged = false
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `namespace` - (Required, String) The namespace to which the retention policy is attached.
- `images_per_repo` - (Required, Integer) Determines how many images are retained in each repository when the retention policy is processed. The value `-1` denotes `Unlimited` (all images are retained).
- `retain_untagged` - (Optional, Bool) Determines whether untagged images are retained when the retention policy is processed. Default value is **false**, means untagged images can be deleted when the policy runs.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - The unique identifier of the cr_retention_policy. This identifier is the same as the name of namespace to which the retention policy is attached.

## Import

You can import the `ibm_cr_retention_policy` resource by adding the `namespace` option, where `namespace` is the namespace to which the retention policy is attached.

```
$ terraform import ibm_cr_retention_policy.cr_retention_policy <namespace>
```
