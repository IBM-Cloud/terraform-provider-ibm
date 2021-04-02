---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_package"
description: |-
  Get information on an IBM Cloud Functions Package.
---

# ibm\_function_package

Import the details of an existing [IBM Cloud Functions package](https://cloud.ibm.com/docs/openwhisk/openwhisk_packages.html#openwhisk_packages) as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_function_package" "package" {
  name      = "package_name"
  namespace = "function_namespace_name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the package.
* `namespace` - (Required, string) The name of the function namespace.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the package.
* `namespace` -  The name of the function namespace.
* `version` - Semantic version of the package.
* `publish` - Package visibility.
* `annotations` - All annotations to describe the package, including those set by you or by IBM Cloud Functions.
* `parameters` - All parameters passed to the package, including those set by you or by IBM Cloud Functions.
* `package_id` - package ID.
