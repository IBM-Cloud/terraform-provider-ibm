---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function-namespace"
description: |-
  Manages IBM Cloud Functions namespace.
---

# ibm_function_namespace

Create, update, or delete an IBM Cloud Functions namespace. For more information, about managing namespace, see [managing namespace](https://cloud.ibm.com/docs/openwhisk?topic=openwhisk-namespaces). Then, you can create IAM managed namespaces to group entities such as actions, triggers or both.

## Example usage
The following example creates an IAM based namespace and package at a specific location.

```terraform
provider "ibm" {
  ibmcloud_api_key   = var.ibmcloud_api_key
  region = var.region
}

data "ibm_resource_group" "resource-group" {
   name = var.resource_group
}

resource "ibm_function_namespace" "namespace" {
   name                = var.namespace
   resource_group_id   = data.ibm_resource_group.resource-group.id
}

resource "ibm_function_package" "package" {
  name      = var.packagename
  namespace = ibm_function_namespace.namespace.name
}
```


## Argument reference
Review the argument reference that you can specify for your resource. 

- `description` - (Optional, String) The description of the namespace.
- `name` - (Required, String) The name of the namespace.
- `resource_group_id` - (Required, Forces new resource, String) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique ID of the new namespace.
- `location` - (String) Target locations of the namespace.

## Import
The `ibm_function_namespace` resource can be imported by using the namespace ID.

**Note** 
Namespace import will not return the value for `resource_group_id` attribute.


**Syntax**

```
$ terraform import ibm_function_namespace.namespace <namespaceID>

```

**Example**

```
$ terraform import ibm_function_namespace.namespace 4cf78bb1-2298-413f-8575-2464948a344b

```


