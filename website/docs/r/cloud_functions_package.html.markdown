---
layout: "ibm"
page_title: "IBM : cloud-functions-package"
sidebar_current: "docs-ibm-resource-cloud-functions-package"
description: |-
  Manages IBM Cloud Functions package.
---

# ibm\_openwhisk_package

Create, update, or delete [IBM Cloud Functions package](https://console.bluemix.net/docs/openwhisk/openwhisk_packages.html#openwhisk_packages).

## Example Usage

### Create package

```hcl
resource "ibm_cloud_functions_package" "package" {
  name = "package-name"

  user_defined_annotations = <<EOF
        [
    {
        "key":"description",
        "value":"Count words in a string"
    },
    {
        "key":"sampleOutput",
        "value": {
                        "count": 3
                }
    },
    {
        "key":"final",
        "value": [
                        {
                                "description": "A string",
                                "name": "payload",
                                "required": true
                        }
                ]
    }
]
EOF
}
```

### Create package using binding

``` hcl
resource "ibm_cloud_functions_package" "bindpackage" {
  name              = "bindalaram"
  bind_package_name = "/whisk.system/alarms/alarm"

  user_defined_parameters = <<EOF
        [
    {
        "key":"cron",
        "value":"0 0 1 0 *"
    },
    {
        "key":"trigger_payload ",
        "value":"{'message':'bye old Year!'}"
    },
    {
        "key":"maxTriggers",
        "value":1
    },
    {
        "key":"userdefined",
        "value":"test"
    }
]
EOF
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of package.
* `publish` - (Optional, boolean) Package visibilty.
* `user_defined_annotations` - (Optional, string) Annotation values in KEY VALUE format.
* `user_defined_parameters` - (Optional, string) Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the package.
* `bind_package_name` - (Optional, string) Name of package to be binded.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new package.
* `version` - Semantic version of the item.
* `annotations` -  All annotations set on package by user and those set by the IBM Cloud Function backend/API.
* `parameters` - All parameters set on package by user and those set by the IBM Cloud Function backend/API.

## Import

ibm_cloud_functions_package can be imported using their id, e.g.

```
$ terraform import ibm_cloud_functions_package.sample hello

```
