---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_account_settings"
description: |-
  Get information about scc_account_location_settings
---

# ibm_scc_account_settings

Provides a read-only data source for scc_account_location_settings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

~> **NOTE**: exporting out the environmental variable `IBM_CLOUD_SCC_ADMIN_API_ENDPOINT` will help out if the account fails to resolve.
## Example usage

```terraform
data "ibm_scc_account_settings" "scc_account_location_settings" {
}
```


## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - (String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.
