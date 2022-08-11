---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_package"
description: |-
  Get information on an IBM Cloud Functions package.
---

# ibm_function_package

Retrieve the information about an existing [IBM Cloud Functions package](https://cloud.ibm.com/docs/openwhisk/openwhisk_packages.html#openwhisk_packages).


## Example usage
The following example retrieves information about the `mypackage` package. 


```terraform
data "ibm_function_package" "package" {
  name      = "package_name"
  namespace = "function_namespace_name"

```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the package.
- `namespace` - (Required, String) The name of the function namespace.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `annotations` - (List) All annotations to describe the package, including those set by you or by IBM Cloud Functions.
- `id` - (String) The ID of the package.
- `parameters` (List) All parameters passed to the package, including those set by you or by IBM Cloud Functions.
- `package_id` - (String) The package ID.
- `publish`- (Bool) Package visibility.
- `namespace` - (String) The name of the function namespace.
- `version` - (String) Semantic version of the package.
