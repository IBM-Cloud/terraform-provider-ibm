---
layout: "ibm"
page_title: "IBM : function_action"
sidebar_current: "docs-ibm-datasource-cloud-function-action"
description: |- 
    Get information on a IBM Cloud Functions action.
---

# ibm\_function_action

Import the details of an existing [IBM Cloud Functions action](https://cloud.ibm.com/docs/openwhisk/openwhisk_actions.html#openwhisk_actions) as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_function_action" "nodehello" {
    name      = "action-name"		  
    namespace = "function-namespace-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the action.
* `namespace` - (Required, string) The name of the function namespace.

## Attributes Reference

The following attributes are exported:

* `id`- The ID of the action.
* `namespace` -  The name of the function namespace.
* `version` - Semantic version of the item.
* `annotations` - Annotations to describe the action, including those set by you or by IBM Cloud Functions.
* `parameters` - Parameters passed to the action when the action is invoked, including those set by you or by IBM Cloud Functions.
* `limits` - A nested block to describe assigned limits. Nested `limits` blocks have the following structure:
    * `timeout` - The timeout limit to terminate the action, specified in milliseconds. Default value: `60000`.
    * `memory` - The maximum memory for the action, specified in MBs. Default value: `256`.
    * `log_size` - The maximum log size for the action, specified in MBs. Default value: `10`.
* `exec` - A nested block to describe executable binaries. Nested `exec` blocks have the following structure:
    * `image` - When using the `blackbox` executable, the name of the container image name.
    * `init` - When using `nodejs`, the optional zipfile reference.
    * `code` - When not using the `blackbox` executable, the code to execute. 
    * `kind` - The type of action. Accepted values: `nodejs`, `blackbox`, `swift`, `sequence`.
    * `main` - The name of the action entry point (function or fully-qualified method name, when applicable).
    * `components` - The list of fully qualified actions.
* `publish` - Action visibility.
* `action_id` - action ID.
