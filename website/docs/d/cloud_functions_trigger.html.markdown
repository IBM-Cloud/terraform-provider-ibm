---
layout: "ibm"
page_title: "IBM : cloud_functions_trigger"
sidebar_current: "docs-ibm-datasource-cloud-functions-trigger"
description: |-
  Get information on an IBM Cloud Functions Trigger.
---

# ibm\_cloud_functions_trigger

Import the details of an existing [IBM Cloud Functions trigger](https://console.bluemix.net/docs/openwhisk/openwhisk_triggers_rules.html#openwhisk_triggers) as a read-only data source. The fields of the data source can then be referenced by other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_cloud_functions_trigger" "trigger" {
			name = "trigger-name"		  
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of trigger.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the trigger.
* `publish` - Trigger Visibility.
* `version` - Semantic version of the trigger.
* `annotations` - All annotations set on trigger by user and those set by the IBM Cloud Function backend/API.
* `parameters` - All parameters set on trigger by user and those set by the IBM Cloud Function backend/API.
