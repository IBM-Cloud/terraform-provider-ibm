---
layout: "ibm"
page_title: "IBM : cloud_functions_action"
sidebar_current: "docs-ibm-datasource-cloud-functions-action"
description: |- 
    Get information on a IBM Cloud Functions action.
---

# ibm\_cloud_functions_action

Import the details of an existing [IBM Cloud Functions action](https://console.bluemix.net/docs/openwhisk/openwhisk_actions.html#openwhisk_actions)  as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cloud_functions_action" "nodehello" {
    name = "action-name"		  
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of an action.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the action.
* `version` - Semantic version of the item.
* `annotations` -  All annotations set on action by user and those set by the IBM Cloud Function.
* `parameters` -  All parameters set on action by user and those set by the IBM Cloud Function.
* `limits` - A nested block describing the limits assigned to . Nested `limits` blocks have the following structure:
    * `timeout` -  The timeout LIMIT in milliseconds after which the action is terminated. Default value is 60000.
    * `memory` -  The maximum memory LIMIT in MB for the action. Default is 256.
    * `log_size` - The maximum log size LIMIT in MB for the action. Default value is 10.
* `exec` - A nested block describing the exec assigned to . Nested `exec` blocks have the following structure:
    * `image` -  Container image name when kind is 'blackbox'.
    * `init` -  Optional zipfile reference when code kind is 'nodejs'.
    * `code` - The code to execute when kind is not 'blackbox'
    * `kind` -  The type of action. Possible values: nodejs, blackbox, swift, sequence.
    * `main` -  The name of the action entry point (function or fully-qualified method name when applicable)
    * `components` - The List of fully qualified action.
* `publish` - Action visibilty.
