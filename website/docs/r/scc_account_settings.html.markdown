---
layout: "ibm"
page_title: "IBM : ibm_scc_account_settings"
description: |-
  Manages the account settings scc_account_settings
subcategory: "Security and Compliance Center"
---

# ibm_scc_account_settings

Provides a resource for scc_account_settings. This allows scc_account_settings to be created, updated and deleted.

~> **NOTE**: exporting out the environmental variable `IBM_CLOUD_SCC_ADMIN_API_ENDPOINT` will help out if the account fails to resolve(e.g. `export IBMCLOUD_SCC_ADMIN_API_ENDPOINT=https://compliance.cloud.ibm.com`)

## Example Usage

```terraform
resource "ibm_scc_account_settings" "scc_account_settings" {
    location {
        location_id = "uk"
    }
    event_notifications {
        instance_crn        = "<instance_crn for event_notifications>" // Optional field
    }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

> `location_id` will be Deprecated in the near future. Please adjust to using the argument `location`
* `location_id` **Deprecated** - (Optional, Forces new resource, String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.
* `event_notifications` - (Optional, List) The Event Notification settings to register.
Nested scheme for **event_notifications**:
	* `instance_crn` - (Optional, String) The Cloud Resource Name (CRN) of the Event Notifications instance that you want to connect. If this field is left blank, no Event Notifications instance will be used.
* `location` - (Optional, List) Location Settings.
Nested scheme for **location**:
	* `location_id` - (Required, String) The programatic ID of the location that you want to work in.
	  * Constraints: Allowable values are: `us`, `eu`, `uk`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the scc_account_settings.

## Import

You can import the `ibm_scc_account_settings` resource by using `terraform import <string>`, with the `<string>` being anything you want.

# Syntax
```
$ terraform import ibm_scc_account_settings.scc_account_settings <string>
```
