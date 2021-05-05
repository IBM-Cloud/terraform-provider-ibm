---
layout: "ibm"
page_title: "IBM : cr_namespace"
sidebar_current: "docs-ibm-resource-cr-namespace"
description: |-
  Manages cr_namespace.
---

# ibm\_cr_namespace

Provides a resource for cr_namespace. This allows cr_namespace to be created, updated and deleted.

## Example Usage

```hcl
resource "cr_namespace" "cr_namespace" {
  name = "name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The name of the namespace.
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group that the namespace will be created within.
* `tags` - (Optional, array of strings) Tags associated with the cr_namespace.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cr_namespace.
* `account` - The IBM Cloud account that owns the namespace.
* `created_date` - When the namespace was created.
* `crn` - If the namespace has been assigned to a resource group, this is the IBM Cloud CRN representing the namespace.
* `resource_created_date` - When the namespace was assigned to a resource group.
* `updated_date` - When the namespace was last updated.
