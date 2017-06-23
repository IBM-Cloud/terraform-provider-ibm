---
layout: "ibm"
page_title: "IBM : object_storage_account"
sidebar_current: "docs-ibm-resource-object-storage-account"
description: |-
  Manages IBM Object Storage Account.
---

# ibm\_object_storage_account

Retrieve the account name for an existing Object Storage instance within your IBM account. If there is no Object Storage instance, you can use this resources to order one for you and remember the account name. 

This resource is not intended for managing the lifecycle (e.g. update, delete) of an Object Storage instance in IBM. For lifecycle management, see the Swift API or Swift resources. 

## Example Usage

```hcl
resource "ibm_object_storage_account" "foo" {
}
```

## Argument Reference

No additional arguments needed.

## Computed Fields

The following attributes are exported:

* `id` - The Object Storage account name, which you can use with Swift resources.
