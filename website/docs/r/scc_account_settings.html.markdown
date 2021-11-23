---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_account_settings"
description: |-
  Manages the account settings scc_account_settings
---

# ibm_scc_account_settings

Provides a resource for the scc_account_settings. This allows the scc_account_settings to be updated.

## Example Usage

```hcl
resource "ibm_scc_account_settings" "ibm_scc_account_settings_instance" {
  location_id = var.ibm_scc_account_settings_location_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `location_id` - (Required, Forces new resource, String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_account_settings.
* `location_id` - (String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.
