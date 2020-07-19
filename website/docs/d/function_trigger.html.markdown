---
layout: "ibm"
page_title: "IBM : function_trigger"
sidebar_current: "docs-ibm-datasource-functions-trigger"
description: |-
  Get information on an IBM Cloud Functions Trigger.
---

# ibm\_function_trigger

Import the details of an existing [IBM Cloud Functions trigger](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers) as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_function_trigger" "trigger" {
	name      = "trigger-name"		  
	namespace = "function-namespace-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the trigger.
* `namespace` - (Required, string) The name of the function namespace.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the trigger.
* `namespace` - The name of the function namespace.
* `publish` - Trigger visibility.
* `version` - Semantic version of the trigger.
* `annotations` - All annotations to describe the trigger, including those set by you or by IBM Cloud Functions.
* `parameters` - All parameters passed to the trigger, including those set by you or by IBM Cloud Functions.
* `trigger_id` - trigger ID.
