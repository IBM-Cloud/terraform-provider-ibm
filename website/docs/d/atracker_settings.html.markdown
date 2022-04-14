---
layout: "ibm"
page_title: "IBM : ibm_atracker_settings"
description: |-
  Get information about atracker_settings
subcategory: "Activity Tracker"
---

# ibm_atracker_settings

Provides a read-only data source for atracker_settings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_atracker_settings" "atracker_settings" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the atracker_settings.
* `api_version` - (Required, Integer) The lowest API version of targets or routes that customer might have under his or her account.

* `default_targets` - (Required, List) The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -]/`.

* `metadata_region_primary` - (Optional, String) To store all your meta data in a single region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -_]/`.

* `permitted_target_regions` - (Required, List) If present then only these regions may be used to define a target.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -_]/`.

* `private_api_endpoint_only` - (Required, Boolean) If you set this true then you cannot access api through public network.

