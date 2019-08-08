---
layout: "ibm"
page_title: "IBM : function_rule"
sidebar_current: "docs-ibm-resource-function-rule"
description: |-
  Manages IBM Cloud Functions rule.
---

# ibm\_function_rule

Create, update, or delete [IBM Cloud Functions triggers](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers). Events from external and internal event sources are channeled through a trigger, and rules allow your actions to react to these events. To set triggers, use the `function_trigger` resource.

## Example Usage

```hcl
resource "ibm_function_action" "action" {
  name = "hello"

  exec = {
    kind = "nodejs:6"
    code = "${file("test-fixtures/hellonode.js")}"
  }
}

resource "ibm_function_trigger" "trigger" {
  name = "alarmtrigger"

  feed = [
    {
      name = "/whisk.system/alarms/alarm"

      parameters = <<EOF
					[
						{
							"key":"cron",
							"value":"0 */2 * * *"
						}
					]
                EOF
    },
  ]
}

resource "ibm_function_rule" "rule" {
  name         = "alarmrule"
  trigger_name = "${ibm_function_trigger.trigger.name}"
  action_name  = "${ibm_function_action.action.name}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the rule.
* `trigger_name` - (Required, string) The name of the trigger.
* `action_name` - (Required, string) The name of the action.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new rule.
* `publish` - Rule visibility.
* `version` - Semantic version of the item.
* `status` - The status of the rule.

## Import

`ibm_function_rule` can be imported using the ID.

Example: 

```
$ terraform import ibm_function_rule.sampleRule alarmrule

```
