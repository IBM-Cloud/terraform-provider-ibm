---
layout: "ibm"
page_title: "IBM : ibm_atracker_settings"
description: |-
  Manages atracker_settings.
subcategory: "Activity Tracker API Version 2"
---

# ibm_atracker_settings

Create, update, and delete atracker_settingss with this resource.

## Example Usage

```hcl
resource "ibm_atracker_settings" "atracker_settings_instance" {
  default_targets = [ ibm_atracker_target.atracker_target_instance.id ]
  metadata_region_backup = "eu-de"
  metadata_region_primary = "us-south"
  permitted_target_regions = us-south
  private_api_endpoint_only = false
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `default_targets` - (Optional, List) The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event. Enterprise-managed targets are not supported.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -]/`.
* `metadata_region_backup` - (Optional, String) To store all your meta data in a backup region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -_]/`.
* `metadata_region_primary` - (Required, String) To store all your meta data in a single region.
  * Constraints: The maximum length is `256` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-zA-Z0-9 -_]/`.
* `permitted_target_regions` - (Optional, List) If present then only these regions may be used to define a target.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9 -_]/`.
* `private_api_endpoint_only` - (Required, Boolean) If you set this true then you cannot access api through public network.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the atracker_settings.
* `api_version` - (Integer) API version used for configuring IBM Cloud Activity Tracker Event Routing resources in the account.
  * Constraints: The maximum value is `2`. The minimum value is `2`.
* `message` - (String) An optional message containing information about the audit log locations.


## Import

You can import the `ibm_atracker_settings` resource by using `metadata_region_primary`. To store all your meta data in a single region.

# Syntax
<pre>
$ terraform import ibm_atracker_settings.atracker_settings &lt;metadata_region_primary&gt;
</pre>

# Example
```
$ terraform import ibm_atracker_settings.atracker_settings us-south
```
