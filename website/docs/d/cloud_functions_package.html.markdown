---
layout: "ibm"
page_title: "IBM : cloud_functions_package"
sidebar_current: "docs-ibm-datasource-cloud-functions-package"
description: |-
  Get information on an IBM Cloud Functions Package.
---

# ibm\_cloud_functions_package

Import the details of an existing [IBM Cloud Functions package](https://console.bluemix.net/docs/openwhisk/openwhisk_packages.html#openwhisk_packages) as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cloud_functions_package" "package" {
  name = "package_name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of package.


## Attributes Reference

The following attributes are exported:

* `id` - The ID of the package.
* `version` - Semantic version of the package.
* `publish` - Package Visibility.
* `annotations` - All annotations set on package by user and those set by the IBM Cloud Function backend/API.
* `parameters` - All parameters set on package by user and those set by the IBM Cloud Function backend/API.
