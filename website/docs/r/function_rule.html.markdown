---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_rule"
description: |-
  Manages IBM Cloud Functions rule.
---

# ibm_function_rule

Create, update, or delete an IBM Cloud Functions rule. Events from external and internal event sources are channeled through a trigger, and rules allow your actions to react to these events. To set triggers, use the `function_trigger` resource. For more information, see [getting started with IBM Cloud Functions](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers).


## Example usage
The following example creates a rule for an action. 

```terraform
resource "ibm_function_action" "action" {
  name = "hello"  exec {
    kind = "nodejs:10"
    code = file("test-fixtures/hellonode.js")
  }
}

resource "ibm_function_trigger" "trigger" {
  name = "alarmtrigger"  feed = [
    {
      name = "/whisk.system/alarms/alarm"      parameters = <<EOF
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
  trigger_name = ibm_function_trigger.trigger.name
  action_name  = ibm_function_action.action.name
}

```


## Argument reference
Review the argument reference that you can specify for your resource. 

- `action_name` - (Required, String) The name of the action.
- `name` - (Required, Forces new resource, String) The name of the rule.
- `namespace` - (Required, String) The name of the function namespace.
- `trigger_name` - (Required, String) The name of the trigger.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `id` - (String) The ID of the new rule.
- `namespace` - (String) The name of the function namespace.
- `publish`- (Bool) Rule visibility.
- `rule_id` - (String) The rule ID.
- `status` - (String) The status of the rule.
- `version` - (String) Semantic version of the item.

## Import
The `ibm_function_rule` resource can be imported by using the `namespace` and `rule_id`.

**Syntax**

```
$ terraform import ibm_function_rule.sampleRule <namespace>:<rule_id>

```
**Example**

```
$ terraform import ibm_function_rule.sampleRule alarmrule
```
