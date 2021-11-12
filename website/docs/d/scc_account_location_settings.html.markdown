---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_account_location_settings"
description: |-
  Get information about scc_account_location_settings
---

# ibm_scc_account_location_settings

Provides a read-only data source for scc_account_location_settings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_account_location_settings" "scc_account_location_settings" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_account_location_settings.
* `cleanup_inactive_locations` - (Required, Boolean) Used to determine when to delete data in an inactive location.

* `id` - (Required, String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.

* `modified` - (Required, String) The time that the location was last updated.

