---
layout: "ibm"
page_title: "IBM : function_rule"
sidebar_current: "docs-ibm-datasource-functions-rule"
description: |-
  Get information about an IBM Cloud Functions Rule.
---

# ibm\_function_rule

Import the details of an existing [IBM Cloud Functions rule](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers) as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_function_rule" "rule" {
	name      = "rule-name"
	namespace = "function-namespace-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the rule.
* `namespace` -  The name of the function namespace.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the rule.
* `namespace` -  The name of the function namespace.
* `publish` - Trigger visibility.
* `version` - Semantic version of the trigger.
* `status` - The status of the rule.
* `trigger_name` - The name of the trigger.
* `action_name` - The name of the action.
* `rule_id` - rule ID.
