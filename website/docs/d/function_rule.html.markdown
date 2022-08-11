---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_rule"
description: |-
  Get information about an IBM Cloud Functions rule.
---

# ibm_function_rule

Retrieve the information about an existing [IBM Cloud Functions rule](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers) as a read only data source.


## Example usage
The following example retrieves information about the `myrule` rule. 

```terraform
data "ibm_function_rule" "rule" {
	name      = "rule-name"
	namespace = "function-namespace-name"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Required, String) The name of the rule.
- `namespace` - (Required, String) The name of the function namespace.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `action_name` - (String) The name of the action that the rule belongs to.
- `id` - (String) The ID of the rule.
- `namespace` - (String) The name of the function namespace.
- `publish`- (Bool) Rule visibility.
- `rule_id` - (String) The rule ID.
- `status` - (String) The status of the rule.
- `trigger_name` - (String) The name of the trigger that the rule belongs to.
- `version` - (String) Semantic version of the rule.
