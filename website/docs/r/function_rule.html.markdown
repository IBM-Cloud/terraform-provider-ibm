---

subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_rule"
description: |-
  Manages IBM Cloud Functions rule.
---

# ibm\_function_rule

Create, update, or delete [IBM Cloud Functions triggers](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers). Events from external and internal event sources are channeled through a trigger, and rules allow your actions to react to these events. To set triggers, use the `function_trigger` resource.

## Example Usage

```hcl
resource "ibm_function_action" "action" {
  name      = "hello"
  namespace = "function-namespace-name

  exec {
    kind = "nodejs:10"
    code = file("test-fixtures/hellonode.js")
  }
}

resource "ibm_function_trigger" "trigger" {
  name      = "alarmtrigger"
  namespace = "function-namespace-name

  feed {
    name      = "/whisk.system/alarms/alarm"
    namespace = "function-namespace-name
    parameters = <<EOF
                                        [
                                                {
                                                        "key":"cron",
                                                        "value":"0 */2 * * *"
                                                }
                                        ]

EOF

  }
}

resource "ibm_function_rule" "rule" {
  name         = "alarmrule"
  namespace    = "function-namespace-name
  trigger_name = ibm_function_trigger.trigger.name
  action_name  = ibm_function_action.action.name
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The name of the rule.
* `namespace` - (Required, string) The name of the function namespace.
* `trigger_name` - (Required, string) The name of the trigger.
* `action_name` - (Required, string) The name of the action.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the rule.The id is combination of namespace and ruleID delimited by `:`.
* `namespace` - The name of the function namespace.
* `publish` - Rule visibility.
* `version` - Semantic version of the item.
* `status` - The status of the rule.
* `rule_id` - Rule ID	

## Import

`ibm_function_rule` can be imported using the namespace and ruleID.

Example: 

```
$ terraform import ibm_function_rule.sampleRule <namespace>:<rule_id>

$ terraform import ibm_function_rule.sampleRule Namespace-01:alaramrule

```
