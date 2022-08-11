---
layout: "ibm"
page_title: "IBM : ibm_cr_namespace"
description: |-
  Manages IBM Cloud Container Registry namespaces
subcategory: "Container Registry"
---

# ibm_cr_namespace

Create or delete a Container Registry namespace. For more information about Container Registry, see [About IBM Cloud Container Registry](https://cloud.ibm.com/docs/Registry?topic=Registry-registry_overview).

## Example usage

The following example shows how you can configure a `namespace`.

```terraform
resource "ibm_cr_namespace" "cr_namespace" {
  name = "name"
}

data "ibm_resource_group" "rg" {
  name = "default"
}
resource "ibm_cr_namespace" "rg_namespace" {
  name              = "testaasd2312"
  resource_group_id = data.ibm_resource_group.rg.id
}
```

## Argument reference

The following arguments are supported:

- `name` - (Required, Forces new resource, string) The name of the namespace that you want to create.
  - Constraints: The maximum length is `30` characters. The minimum length is `4` characters. The value must match regular expression `/^[a-z0-9]+[a-z0-9_-]+[a-z0-9]+$/`
- `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group to which you want to add the namespace. If you don't set this option, the default resource group for the account is used.
- `tags` - (Optional, array of strings) The tags that are associated with the `ibm_cr_namespace`. **Note*- `Tags` are managed locally and not stored on the IBM Cloud service endpoint.

## Attribute reference

In addition to the attributes in the argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the namespace. This identifier is the same as the name of namespace.
- `account` - (String) The IBM Cloud account that owns the namespace.
- `created_date` - (Timestamp) The creation date of the namespace.
- `crn` - (String) If the namespace is assigned to a resource group, the IBM Cloud CRN representing the namespace.
- `resource_created_date` - (Timestamp) The date that the namespace was assigned to a resource group.
- `updated_date` - (Timestamp) The date that the namespace was last updated.

## Import

You can import the `ibm_cr_namespace` resource by using the `name` of the namespace.

**Syntax**

```
$ terraform import ibm_cr_namespace.test <name of the namespace>
```

**Example**

```
$ terraform import ibm_cr_namespace.test namespace-name
```
