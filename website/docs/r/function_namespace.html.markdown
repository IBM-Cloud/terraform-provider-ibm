---

subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function-namespace"
description: |-
  Manages IBM Cloud Functions namespace.
---

# ibm\_function_namespace

Create, update, or delete [IBM Cloud Functions namespace](https://cloud.ibm.com/docs/openwhisk?topic=openwhisk-namespaces). You can create Identity and Access (IAM) managed namespaces to group entities, such as actions or triggers, together.

## Example Usage

In the following example, you can create a IAM based namespace and package at specified location :

```hcl
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

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the namespace.
* `description` - (Optional, string) Description of the namespace.
* `resource_group_id` - (Required, ForceNew, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`.

## Attributes Reference

The following attributes are exported:

* `id` - The unique identifier of the namespace.
* `location` - Target location of the namespace.

## Import

`ibm_function_namespace` can be imported using the namespace ID.

**NOTE:**: Namespace import will not return value for `resource_group_id` attribute.

Example:

```
$ terraform import ibm_function_namespace.namespace <namespaceID>

$ terraform import ibm_function_namespace.namespace 4cf78bb1-2298-413f-8575-2464948a344b

```

