---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_trigger"
description: |-
  Get information on an IBM Cloud Functions trigger.
---


# ibm_function_trigger

Retrieve information about an existing [IBM Cloud Functions trigger](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers) as a read-only data source.


## Example usage
The following example retrieves information about the `mytrigger` trigger. 

```terraform
data "ibm_function_trigger" "trigger" {
	name      = "trigger-name"		  
	namespace = "function-namespace-name"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the trigger.
- `namespace` - (Required, String) The name of the function namespace.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `annotations` - (String) All annotations to describe the trigger, including those set by you or by IBM Cloud Functions.
- `id` - (String) The ID of the trigger.
- `namespace` - (String) The name of the function namespace.
- `parameters` (String) All parameters passed to the trigger, including those set by you or by IBM Cloud Functions.
- `publish`- (Bool) Trigger visibility.
- `trigger_id` - (String) The trigger ID.
- `version` - (String) Semantic version of the trigger.
