---
layout: "ibm"
page_title: "IBM : cloud_functions_action"
sidebar_current: "docs-ibm-resource-cloud-functions-action"
description: |-
  Manages IBM Cloud Functions action.
---

# ibm\_cloud_functions_action

Create, update, or delete  [IBM Cloud Functions action](https://console.bluemix.net/docs/openwhisk/openwhisk_actions.html#openwhisk_actions).


## Example Usage

###  Simple JavaScript action

```hcl
resource "ibm_cloud_functions_action" "nodehello" {
  name = "action-name"

  exec = {
    kind = "nodejs:6"
    code = "${file("hellonode.js")}"
  }
}

```
### Passing Parameters to action

```hcl
resource "ibm_cloud_functions_action" "nodehellowithparameter" {
  name = "hellonodeparam"

  exec = {
    kind = "nodejs:6"
    code = "${file("hellonodewithparameter.js")}"
  }

  user_defined_parameters = <<EOF
        [
    {
        "key":"place",
        "value":"India"
    }
        ]
        EOF
}

```

### Packaging an action as a Node.js module

``` hcl
resource "ibm_cloud_functions_action" "nodezip" {
  name = "nodezip"

  exec = {
    kind = "nodejs:6"
    code = "${base64encode("${file("nodeaction.zip")}")}"
  }
}

```

### Creating action sequences

``` hcl
resource "ibm_cloud_functions_action" "swifthello" {
  name = "actionsequence"

  exec = {
    kind = "sequence"
    components = ["/whisk.system/utils/split","/whisk.system/utils/sort"]
  }
}

```

## Creating docker actions

``` hcl
resource "ibm_cloud_functions_action" "swifthello" {
  name = "dockeraction"

  exec = {
    kind = "janesmith/blackboxdemo"
    image = "${file("helloSwift.swift")}"
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of action.
* `limits` - (Optional, set) A nested block describing the limits assigned to . Nested `limits` blocks have the following structure:
    * `timeout` - (Optional, integer) The timeout LIMIT in milliseconds after which the action is terminated. Default value is 60000.
    * `memory` - (Optional, integer) The maximum memory LIMIT in MB for the action. Default is 256.
    * `log_size` - (Optional, integer) The maximum log size LIMIT in MB for the action. Default value is 10.
* `exec` - (Required, set) A nested block describing the exec assigned to . Nested `exec` blocks have the following structure:
    * `image` - (Optional, string) Container image name when kind is 'blackbox'. **NOTE**: Conflicts with `exec.components`, `exec.code`.
    * `init` - (Optional, string) Optional zipfile reference. **NOTE**: Conflicts with `exec.components`, `exec.image`.
    * `code` - (Optional, string) The code to execute when kind is not 'blackbox'. **NOTE**: Conflicts with `exec.components`, `exec.image`.
    * `kind` - (Required, string) The type of action. Possible values: php:7.1, nodejs:8, swift:3, nodejs, blackbox, java, sequence, nodejs:6, python:3, python, python:2, swift, swift:3.1.1.
    * `main` - (Optional, string) The name of the action entry point (function or fully-qualified method name when applicable). **NOTE**: Conflicts with `exec.components`, `exec.image`.
    * `components` - (Optional, string) The List of fully qualified action. **NOTE**: Conflicts with `exec.code`, `exec.image`.
* `publish` - (Optional, boolean) Action visibilty.
* `user_defined_annotations` - (Optional, string) Annotation values in KEY VALUE format.
* `user_defined_parameters` - (Optional, string) Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the action.Cloud Function backend/API.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new action.
* `version` - Semantic version of the item.
* `annotations` -  All annotations set on action by user and those set by the IBM Cloud Function backend/API.
* `parameters` - All parameters set on action by user and those set by the IBM Cloud Function backend/API.


## Import

ibm_cloud_functions_action can be imported using their id, e.g.

```
$ terraform import ibm_cloud_functions_action.nodeAction hello

```
