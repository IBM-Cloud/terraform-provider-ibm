---

subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_action"
description: |-
  Manages IBM Cloud Functions actions.
---

# ibm\_function_action

Create, update, or delete [IBM Cloud Functions actions](https://cloud.ibm.com/docs/openwhisk/openwhisk_actions.html#openwhisk_actions). Actions are stateless code snippets that run on the IBM Cloud Functions platform. An action can be written as a JavaScript, Swift, or Python function, a Java method, or a custom executable program packaged in a Docker container. To bundle and share related actions, use the `function_package` resource.


## Example Usage

###  Simple JavaScript action

```hcl
resource "ibm_function_action" "nodehello" {
  name      = "action-name"
  namespace = "function-namespace-name"

  exec {
    kind = "nodejs:10"
    code = file("hellonode.js")
  }
}

```
### Passing parameters to an action

```hcl
resource "ibm_function_action" "nodehellowithparameter" {
  name      = "hellonodeparam"
  namespace = "function-namespace-name"
  
  exec {
    kind = "nodejs:10"
    code = file("hellonodewithparameter.js")
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
resource "ibm_function_action" "nodezip" {
  name      = "nodezip"
  namespace = "function-namespace-name"

  exec {
    kind      = "nodejs:10"
    code_path = "nodeaction.zip"
  }
}

```

### Creating action sequences

``` hcl
resource "ibm_function_action" "swifthello" {
  name      = "actionsequence"
  namespace = "function-namespace-name"

  exec {
    kind       = "sequence"
    components = ["/whisk.system/utils/split", "/whisk.system/utils/sort"]
  }
}

```

## Creating Docker actions

``` hcl
resource "ibm_function_action" "swifthello" {
  name      = "dockeraction"
  namespace = "function-namespace-name"

  exec {
    kind   = "blackbox"	
    image  = "janesmith/blackboxdemo"
    code   = file("helloSwift.swift")
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The name of the action.
* `namespace` - (Required, string) The name of the function namespace.
* `limits` - (Optional, list) A nested block to describe assigned limits. Nested `limits` blocks have the following structure:
    * `timeout` - The timeout limit to terminate the action, specified in milliseconds. Default value: `60000`.
    * `memory` - The maximum memory for the action, specified in MBs. Default value: `256`.
    * `log_size` - The maximum log size for the action, specified in MBs. Default value: `10`.
* `exec` - (Required, list) A nested block to describe executable binaries. Nested `exec` blocks have the following structure:
    * `image` - (Optional, string) When using the `blackbox` executable, the name of the container image name.  
     **NOTE**: Conflicts with `exec.components`, `exec.code`,`exec.code_path`.
    * `init` - (Optional, string) When using `nodejs`, the optional zipfile reference.  
     **NOTE**: Conflicts with `exec.components`, `exec.image`.
    * `code` - (Optional, string) When not using the `blackbox` executable, the code to execute.  
    **NOTE**: Conflicts with `exec.components`, `exec.image`, `exec.code_path`.
    * `code_path` - (Optional, string) When not using the `blackbox` executable, the file path of code to execute and it supports only .zip extension to create the action.
    **NOTE**: Conflicts with `exec.components`, `exec.image`,`exec.code`.
    * `kind` - (Required, string) The type of action. You can find supported kinds in the [IBM Cloud Functions docs](https://cloud.ibm.com/docs/openwhisk?topic=cloud-functions-runtimes).
    * `main` - (Optional, string) The name of the action entry point (function or fully-qualified method name, when applicable).  
    **NOTE**: Conflicts with `exec.components`, `exec.image`.
    * `components` - (Optional, string) The list of fully qualified actions.  
    **NOTE**: Conflicts with `exec.code`, `exec.image`,`exec.code_path`.
* `publish` - (Optional, boolean) Action visibility.
* `user_defined_annotations` - (Optional, string) Annotations defined in key value format.
* `user_defined_parameters` - (Optional, string) Parameters defined in key value format. Parameter bindings included in the context passed to the action. Cloud Function backend/API.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the action.The id is combination of namespace and actionID delimited by `:` .
* `namespace` - The name of the function namespace.
* `version` - Semantic version of the item.
* `annotations` - All annotations to describe the action, including those set by you or by IBM Cloud Functions.
* `parameters` - All parameters passed to the action when the action is invoked, including those set by you or by IBM Cloud Functions.
* `action_id` - Action ID	
* `target_endpoint_url` - Target endpoint URL of the action.

## Import

`ibm_function_action` can be imported using the namespace and actionID.

Example:

```
$ terraform import ibm_function_action.nodeAction <namespace>:<action_id>

$ terraform import ibm_function_action.nodeAction Namespace-01:nodezip

```
