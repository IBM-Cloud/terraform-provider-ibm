---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_action"
description: |-
  Manages IBM Cloud Functions actions.
---

# ibm_function_action

Create, update, or delete an [IBM Cloud Functions action](https://cloud.ibm.com/docs/openwhisk/openwhisk_actions.html#openwhisk_actions). Actions are stateless code snippets that run on the Cloud Functions platform. An action can be written as a JavaScript, Swift, or Python function, a Java method, or a custom executable program packaged in a Docker container. To bundle and share related actions, use the `function_package` resource.


## Example usage
The sample provides the usage of JavaScript, Node.js, Docker, action sequences, by using `ibm_function_action` resources.

###  Simple JavaScript action
The following example creates a JavaScript action. 


```terraform
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
The following example shows how to pass parameters to an action. 


```terraform
resource "ibm_function_action" "nodehellowithparameter" {
  name      = "hellonodeparam"
  namespace = "function-namespace-name"
  
  exec {
    kind = "nodejs:10"
    code = file("hellonodewithparameter.js")
  }  user_defined_parameters = <<EOF
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
The following example packages a JavaScript action to a module. 


``` 
resource "ibm_function_action" "nodezip" {
  name      = "nodezip"
  namespace = "function-namespace-name"  exec {
    kind      = "nodejs:10"
    code_path = "nodeaction.zip"
  }
}
```

### Creating action sequences
The following example creates an action sequence. 


``` 
resource "ibm_function_action" "swifthello" {
  name      = "actionsequence"
  namespace = "function-namespace-name"  exec {
    kind       = "sequence"
    components = ["/whisk.system/utils/split", "/whisk.system/utils/sort"]
  }
}
```


### Creating Docker actions
The following example creates a Docker action. 


``` 
resource "ibm_function_action" "swifthello" {
  name      = "dockeraction"
  namespace = "function-namespace-name"  exec {
    kind   = "blackbox"	
    image  = "janesmith/blackboxdemo"
    code   = file("helloSwift.swift")
  }
}
```


## Argument reference
Review the argument reference that you can specify for your resource. 
 
- `exec` - (Required, List) A nested block to describe executable binaries.

  Nested scheme for `exec`:
  - `code` - (Optional, String) The code to execute, when not using the `blackbox` executable.
    **Note** Conflicts with `exec.components`, `exec.code_path`.
  - `code_path` - (Optional, String)  When not using the `blackbox` executable, the file path of code to execute and supports only `.zip` extension to create the action.
    **Note** Conflicts with `exec.components`, `exec.code`.
  - `components` - (Optional, String) The list of fully qualified actions.
    **Note** Conflicts with `exec.code`, `exec.image`, `exec.code.path`.
  - `image` - (Optional, String)  When using the `blackbox` executable, the name of the container image name.
    **Note** Conflicts with `exec.components`.
  - `init` - (Optional, String)  When using `nodejs`, the optional archive reference.
    **Note** Conflicts with `exec.components`, `exec.image`.
  - `kind` - (Required, String) The type of action. You can find supported kinds in the [IBM Cloud Functions Docs](https://cloud.ibm.com/docs/openwhisk?topic=openwhisk-runtimes).
  - `main` - (Optional, String) The name of the action entry point (function or fully-qualified method name, when applicable).
    **Note** Conflicts with `exec.components`, `exec.image`.
- `limits` - (Optional, List) A nested block to describe assigned limits.

  Nested scheme for `limits`:
  - `timeout` - (Optional, Integer) The timeout limit to terminate the action, specified in milliseconds. Default value is `60000`.
  - `memory` - (Optional, Integer) The maximum memory for the action, specified in megabyte. Default value is `256`.
  - `log_size` - (Optional, Integer) The maximum log size for the action, specified in megabyte. Default value is `10`.
- `name` - (Required, Forces new resource, String) The name of the action.
- `namespace` - (Required, String) The name of the function namespace.
- `publish` - (Optional, Bool) Action visibility.
- `user_defined_annotations` - (Optional, String) Annotations defined in key value format.
- `user_defined_parameters` - (Optional, String) Parameters defined in key value format. Parameter bindings included in the context passed to the action. Cloud Function backend/API.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `action_id` - (String) The action ID.
- `annotations` (List) All annotations to describe the action, including those set by you or by IBM Cloud Functions.
- `id` - (String) The ID of the new action.
- `namespace` - (String) The name of the function namespace.
- `parameters` - (List) All parameters passed to the action when the action is invoked, including those set by you or by the IBM Cloud Functions.
- `target_endpoint_url` - (String) The target endpoint URL of the action.
- `version` - (String) Semantic version of the item.


## Import
The `ibm_function_action` resource can be imported by using the namespace and action ID.

**Syntax**

```
$ terraform import ibm_function_action.nodeAction <namespace>:<action_id>
```

**Example**

```
$ terraform import ibm_function_action.nodeAction Namespace-01:nodezip
```

