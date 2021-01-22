---
layout: "ibm"
page_title: "IBM: cr_namespace"
sidebar_current: "docs-ibm-resource-cr-namespace"
description: |-
  Manages IBM Container Registry Namespace.
---

# ibm\_cr_namespace

Creates deletes a Container Registry Namespace. 

## Example Usage

In the following example, you can configure a namespace:

```hcl
resource "ibm_cr_namespace" "test" {
  name              = "test123"
  resource_group_id = "c34128405d5742549538128656d1db57"
}

```
```hcl
data "ibm_resource_group" "rg" {
  name = "default"
}
resource "ibm_cr_namespace" "rg_namespace" {
  name              = "testaasd2312"
  resource_group_id = data.ibm_resource_group.rg.id
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The Name of the Namespace that has to be created.
* `resource_group_id` - (Optional, Forces new resource,string)  Id of the resource group to which the namespace has to be assigned.If not provided, default resource group Id will be assigned

## Attribute Reference

The following attributes are exported:

* `id` - Name of the Namespace.
* `crn` - Crn of the namespace.
* `created_on` - The Created Time of the Namespace.
* `updated_on` - The Updated Time  of the Namespace.

## Import

The `ibm_cr_namespace` resource can be imported using the `id`. The ID is formed from the `Name` (Namespace Name)

id = `name`
```
$ terraform import ibm_cr_namespace.test <name of the namespace>

$ terraform import ibm_cr_namespace.test namespace-name
```