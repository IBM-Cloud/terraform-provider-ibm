---
layout: "ibm"
page_title: "IBM : cloud-functions-package"
sidebar_current: "docs-ibm-resource-cloud-functions-package"
description: |-
  Manages IBM Cloud Functions package.
---

# ibm\_openwhisk_package

Create, update, or delete [IBM Cloud Functions packages](https://console.bluemix.net/docs/openwhisk/openwhisk_packages.html#openwhisk_packages). You can use packages to bundle together a set of related actions, and share them with others. To create actions, use the `cloud_functions_action` resource.

## Example Usage

### Create a package

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

### Create a package using a binding

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

* `name` - (Required, string) The name of the package.
* `publish` - (Optional, boolean) Package visibility.
* `user_defined_annotations` - (Optional, string) Annotations defined in key value format.
* `user_defined_parameters` - (Optional, string) Parameters defined in key value format. Parameter bindings included in the context passed to the package.
* `bind_package_name` - (Optional, string) Name of package to be binded.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new package.
* `version` - Semantic version of the item.
* `annotations` - All annotations to describe the package, including those set by you or by IBM Cloud Functions.
* `parameters` - All parameters passed to the package, including those set by you or by IBM Cloud Functions.

## Import

`ibm_cloud_functions_package` can be imported using the ID, 

Example:
```
$ terraform import ibm_cloud_functions_package.sample hello

```
