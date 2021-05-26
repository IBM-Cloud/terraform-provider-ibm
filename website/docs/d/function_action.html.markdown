---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_action"
description: |- 
    Get information on an IBM Cloud Functions action.
---

# ibm_function_action

Retrieve information about an action. Import the details of an existing [IBM Cloud Functions action](https://cloud.ibm.com/docs/openwhisk/openwhisk_actions.html#openwhisk_actions) as a read-only data source. 


## Example usage
The following example retrieves information about the `myaction` action. 


```terraform
data "ibm_function_action" "nodehello" {
    name      = "action-name"		  
    namespace = "function-namespace-name"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the action.
- `namespace` - (Required, String) The name of the function namespace.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created. 

- `action_id` - (String) Action ID.
- `annotations` (List) Annotations to describe the action, including those set by you or by IBM Cloud Functions.
- `exec` - (List of Objects) A nested block to describe executable binaries.

  Nested scheme for `exec`:
  - `code` - (String) When not using the `blackbox` executable, the code to execute. 
  - `components` (String) The list of fully qualified actions.
  - `image` - (String) When using the `blackbox` executable, the name of the container image name.
  - `init` - (String) When using `nodejs`, the optional reference to the compressed file.
  - `kind` - (String) The type of action. Accepted values: `nodejs`, `blackbox`, `swift`, `sequence`.
  - `main` - (String) The name of the action entry point (function or fully-qualified method name, when applicable).
- `id`-  (String) The ID of the action.
- `limits`- (List) A nested block to describe assigned 
	
   Nested scheme for `limits`:
   - `timeout`- (Integer) The timeout limit to terminate the action, specified in milliseconds. Default value is `60000`.
   - `memory`- (Integer) The maximum memory for the action, specified in megabytes. Default value is `256`.
   - `log_size`- (Integer) The maximum log size for the action, specified in megabytes. Default value is `10`.
- `namespace` - (String) The name of the function namespace.
- `parameters` (List) Parameters passed to the action when the action is invoked, including those set by you or by IBM Cloud Functions.
- `publish`- (Bool) Action visibility.
- `target_endpoint_url` - (String) Target endpoint URL of the action.
- `version` - (String) The version of the action.
