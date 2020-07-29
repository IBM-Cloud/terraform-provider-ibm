---
layout: "ibm"
page_title: "IBM : function_namespace"
sidebar_current: "docs-ibm-datasource-function-namespace"
description: |-
  Get information on an IBM Cloud Functions namespace.
---

# ibm\_function_namespace

Import the details of an existing [IBM Cloud Functions namespace](https://cloud.ibm.com/docs/openwhisk/openwhisk_namespaces.html#openwhisk_namespaces) as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_function_namespace" "test_namespace" {
	name = var.namespace
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the namespace.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the namespace.
* `resource_group_id` - The ID of the resource group.
* `location` - Target location of the namespace.

