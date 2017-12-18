---
layout: "ibm"
page_title: "IBM : cloud_functions_trigger"
sidebar_current: "docs-ibm-resource-cloud-functions-trigger"
description: |-
  Manages IBM Cloud Functions trigger.
---

# ibm\_cloud_functions_trigger

Create, update, or delete for [IBM Cloud Functions trigger](https://console.bluemix.net/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers).

## Example Usage

### Creating triggers

```hcl
resource "ibm_cloud_functions_trigger" "trigger" {
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

### Creating trigger feed
```hcl
resource "ibm_cloud_functions_trigger" "feedtrigger" {
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

* `name` - (Required, string) Name of trigger.
* `feed` - (Optional, set) A nested block describing the feed assigned to . Nested `feed` blocks have   the following structure:
    * `name` - (Required, string) Trigger feed ACTION_NAME
    * `parameters` - (Optinal, string) Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the action invoke.
* `user_defined_annotations` - (Optional, string) Annotation values in KEY VALUE format.
* `user_defined_parameters` - (Optional, string) Parameters values in KEY VALUE format. Parameter bindings included in the context passed to the trigger.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new trigger.
* `publish` - Trigger visbility.
* `version` - Semantic version of the item.
* `annotations` -  All annotations set on trigger by user and those set by the IBM Cloud Function backend/API.
* `parameters` - All parameters set on trigger by user and those set by the IBM Cloud Function backend/API.

## Import

ibm_cloud_functions_trigger can be imported using their id, e.g.

```
$ terraform import ibm_cloud_functions_trigger.alaramtrigger alaram

```
