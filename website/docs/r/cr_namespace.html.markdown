---
subcategory: "Container Registry"
layout: "ibm"
page_title: "IBM : ibm_cr_namespace"
description: |-
  Manages Container Registry namespace.

---

# ibm_cr_namespace
Create, update, or delete a container registry namespace. For more information, about container registry, see [About IBM Cloud Container Registry](https://cloud.ibm.com/docs/Registry?topic=Registry-registry_overview).

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

Review the argument references that you can specify for your resource. 

- `name` - (Required, Forces new resource, String) The name of the namespaces to create.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group to which the namespace is assigned. If not provided, default resource group ID will be assigned.
- `tags` - (Optional, Array of Strings) Tags associated with the `ibm_cr_namespace`. **Note*- `Tags` are managed locally and not stored on the IBM Cloud service endpoint.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the `ibm_cr_namespace`.
- `account` - (String) The IBM Cloud account that owns the namespace.
- `created_date` - (Timestamp) The creation date of the namespace.
- `crn` - (String) If the namespace has been assigned to a resource group, the IBM Cloud CRN represent the namespace.
- `resource_created_date` - (Timestamp) The namespace assigned date to a resource group.
- `updated_date` - (Timestamp) The last updated of the namespace.

## Import
The `ibm_cr_namespace` resource can be imported by using the `name` of the namespace.

**Syntax**

```
$ terraform import ibm_cr_namespace.test <name of the namespace>
```

**Example**

```
$ terraform import ibm_cr_namespace.test namespace-name
```
