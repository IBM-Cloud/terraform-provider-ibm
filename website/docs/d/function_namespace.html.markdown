---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_namespace"
description: |-
  Get information on an IBM Cloud Functions namespace.
---

# ibm_function_namespace

Import the details of an existing IBM Cloud Functions namespace. For more information, about managing namespace, see [managing namespace](https://cloud.ibm.com/docs/openwhisk?topic=openwhisk-namespaces). 

## Example usage
The following example creates the namespace and package at a specific location.

```terraform
data "ibm_function_namespace" "test_namespace" {
	name = var.namespace
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `name` - (Required, String) The name of the namespace.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created.

- `id` - (String) The ID of the namespace.
- `location` - (String) The target location of the namespace.
- `resource_group_id` - (String) The ID of the resource group.


