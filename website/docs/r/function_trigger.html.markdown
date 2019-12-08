---
layout: "ibm"
page_title: "IBM : function_trigger"
sidebar_current: "docs-ibm-resource-function-trigger"
description: |-
  Manages IBM Cloud Functions trigger.
---

# ibm\_function_trigger

Create, update, or delete [IBM Cloud Functions triggers](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers). Events from external and internal event sources are channeled through a trigger, and rules allow your actions to react to these events. To set rules, use the `function_rule` resource.

## Example Usage

### Creating triggers

```hcl
resource "ibm_function_trigger" "trigger" {
  name = "trigger-name"

  user_defined_parameters = <<EOF
                        [
                                {
                                        "key":"place",
                                        "value":"India"
                           }
                   ]
           EOF

  user_defined_annotations = <<EOF
           [
                   {
                          "key":"Description",
                           "value":"Sample code to display hello"
                  }
          ]
  EOF
}
```

### Creating a trigger feed
```hcl
resource "ibm_function_trigger" "feedtrigger" {
  name = "alarmFeed"

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

  user_defined_annotations = <<EOF
                 [
         {
                 "key":"sample trigger",
                 "value":"Trigger for hello action"
         }
                 ]
                 EOF
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The name of the trigger.
* `feed` - (Optional, Forces new resource, list) A nested block to describe the feed. Nested `feed` blocks have the following structure:
    * `name` - (Required, Forces new resource, string) Trigger feed `ACTION_NAME`.
    * `parameters` - (Optional, string) Parameters definitions in key value format. Parameter bindings are included in the context and passed when the action is invoked.
* `user_defined_annotations` - (Optional, string) Annotation definitions in key value format.
* `user_defined_parameters` - (Optional, string) Parameters definitions in key value format. Parameter bindings are included in the context and passed to the trigger.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new trigger.
* `publish` - Trigger visibility.
* `version` - Semantic version of the item.
* `annotations` - All annotations to describe the trigger, including those set by you or by IBM Cloud Functions.
* `parameters` - All parameters passed to the trigger, including those set by you or by IBM Cloud Functions.

## Import

`ibm_function_trigger` can be imported using the ID.

Example:

```
$ terraform import ibm_function_trigger.trigger alaram

```
