---
layout: "ibm"
page_title: "IBM : ibm_atracker_settings"
description: |-
  Manages atracker_settings.
subcategory: "Activity Tracker"
---

# ibm_atracker_settings

Provides a resource for atracker_settings. This allows atracker_settings to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_atracker_settings" "atracker_settings" {
  default_targets = [ ibm_atracker_target.atracker_target.id ]
  metadata_region_primary = "us-south"
  permitted_target_regions = us-south
  private_api_endpoint_only = false
  # Optional but recommended lifecycle flag to ensure target delete order is correct
  lifecycle {
    create_before_destroy = true
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `default_targets` - (Optional, List) The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -]/`.
* `metadata_region_primary` - (Required, String) To store all your meta data in a single region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -_]/`.
* `permitted_target_regions` - (Optional, List) If present then only these regions may be used to define a target.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -_]/`.
* `private_api_endpoint_only` - (Required, Boolean) If you set this true then you cannot access api through public network.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the atracker_settings (only one).
* `api_version` - (Required, Integer) The lowest API version of targets or routes that customer might have under his or her account.

## Import

You can import the `ibm_atracker_settings` resource by using `metadata_region_primary`. To store all your meta data in a single region.

# Syntax
```
$ terraform import ibm_atracker_settings.atracker_settings <metadata_region_primary>
```

# Example
```
$ terraform import ibm_atracker_settings.atracker_settings us-south
```
