---
subcategory: "Functions"
layout: "ibm"
page_title: "IBM : function_trigger"
description: |-
  Manages an IBM Cloud Functions trigger.
---

# ibm_function_trigger

Create, update, or delete an [IBM Cloud Functions trigger](https://cloud.ibm.com/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers). Events from external and internal event sources are channeled through a trigger, and rules allow your actions to react to these events. To set rules, use the `function_rule` resource. 

## Example usage

### Creating triggers
The following example creates the `mytrigger` trigger. 


```terraform
resource "ibm_function_trigger" "trigger" {
  name = "mytrigger"  user_defined_parameters = <<EOF
                        [
                                {
                                        "key":"place",
                                        "value":"India"
                           }
                   ]
           EOF  user_defined_annotations = <<EOF
           [
                   {
                          "key":"Description",
                           "value":"Example usage to display hello"
                  }
          ]
  EOF
}
```


### Creating a trigger feed
The following example creates a feed for the `alarmFeed` trigger.


```terraform
resource "ibm_function_trigger" "feedtrigger" {
  name = "alarmFeed"  feed = [
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
  ]  user_defined_annotations = <<EOF
                 [
         {
                 "key":"sample trigger",
                 "value":"Trigger for hello action"
         }
                 ]
                 EOF
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `feed` (List, Forces new resource, Optional)  A nested block to describe the feed.
  
  Nested scheme for `feed`:
  - `name` - (Required, Forces new resource, String) Trigger feed `ACTION_NAME`. 
  - `parameters` - (Optional, String) Parameters definitions in key value format. Parameter bindings are included in the context and passed when the action is invoked.
- `name` - (Required, Forces new resource, String) The name of the trigger.
- `namespace` - (Required, String) The name of the function namespace.
- `user_defined_annotations` - (Optional, String)  Annotation definitions in key value format.
- `user_defined_parameters` - (Optional, String) Parameters definitions in key value format. Parameter bindings are included in the context and passed to the trigger.


## Attribute reference
In addition to all arguments listed, you can access the following attribute references after your resource is created. 

- `annotations` - (String) All annotations to describe the trigger, including those set by you or by IBM Cloud Functions.
- `id` - (String) The ID of the new trigger.
- `namespace` - (String) The name of the function namespace.
- `parameters` - (String) All parameters passed to the trigger, including those set by you or by IBM Cloud Functions.
- `publish`- (Bool) Trigger visibility.
- `trigger_id` - (String) The trigger ID.
- `version` - (String) Semantic version of the item.


## Import
The `ibm_function_trigger` resource can be imported by using the `namespace` and `trigger ID`.


**Syntax**

```
$ terraform import ibm_function_trigger.trigger <namespace>:<trigger_id>

```
**Example**

```
$ terraform import ibm_function_trigger.trigger Namespace01:alarmtrigger

```
