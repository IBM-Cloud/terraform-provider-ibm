---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function-package"
description: |-
  Manages IBM Cloud Functions package.
---

# ibm_function_package

Create, update, or delete an [IBM Cloud functions package](https://cloud.ibm.com/docs/openwhisk/openwhisk_packages.html#openwhisk_packages). You can use the packages to bundle together a set of related actions, and share with other resources. To create actions, use the `function_action` resource.

## Example usage
The sample example provides the usage of package and to bind the package by using `ibm_function_package` resource.

### Create a package
The following example creates the `mypackage` package. 

```terraform
resource "ibm_function_package" "package" {
  name = "mypackage"  user_defined_annotations = <<EOF
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


### Create a package by using a binding


```terraform
resource "ibm_function_package" "bindpackage" {
  name              = "bindalaram"
  bind_package_name = "/whisk.system/alarms/alarm"  user_defined_parameters = <<EOF
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

## Argument reference
Review the argument reference that you can specify for your resource. 

- `bind_package_name` - (Optional, Forces new resource,String)  Name of package to be bound.
- `name` - (Required, Forces new resource, String) The name of the package.
- `namespace` - (Required, String) The name of the function namespace.
- `publish` - (Optional, Bool) Package visibility.
- `user_defined_annotations` - (Optional, String) Annotations defined in key value format.
- `user_defined_parameters` - (Optional, String) Parameters defined in key value format. Parameter bindings included in the context passed to the package.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `annotations` - (String) All annotations to describe the package, including those you set or by IBM Cloud Functions.
- `id` - (String) The ID of the new package. The ID is a combination of namespace and package ID delimited by `:`.
- `namespace` - (String) The name of the function namespace.
- `package_id` - (String) The package ID.
- `parameters` - (String) All parameters passed to the package, including those you set or by IBM Cloud Functions.
- `version` - (String) Semantic version of the item.

## Import

The `ibm_function_package` resource can be imported by using the namespace and packageID.

**Syntax**

```
$ terraform import ibm_function_package.sample <namespace>:<package_id>
```

**Example**

```
$ terraform import ibm_function_package.sample Namespace-01:util

```
