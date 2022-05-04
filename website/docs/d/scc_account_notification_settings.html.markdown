---
layout: "ibm"
page_title: "IBM : ibm_scc_account_notification_settings"
description: |-
  Get information about scc_account_notification_settings
subcategory: "Security and Compliance Center"
---

# ibm_scc_account_notification_settings

Provides a read-only data source for scc_account_notification_settings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

~> **NOTE**: exporting out the environmental variable `IBM_CLOUD_SCC_ADMIN_API_ENDPOINT` will help out if the account fails to resolve.

## Example Usage

```hcl
data "ibm_scc_account_notification_settings" "scc_account_notification_settings" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `instance_crn` - (Optional, String) The Cloud Resource Name (CRN) of the Event Notifications instance that you want to connect.
