---
layout: "ibm"
page_title: "IBM : ibm_cloud_shell_account_settings"
description: |-
  Get information about cloud_shell_account_settings
subcategory: "IBM Cloud Shell"
---

# ibm_cloud_shell_account_settings

Provides a read-only data source to retrieve information about cloud_shell_account_settings. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
	account_id = ibm_cloud_shell_account_settings.cloud_shell_account_settings_instance.account_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, Forces new resource, String) The account ID in which the account settings belong to.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cloud_shell_account_settings.
* `created_at` - (Integer) Creation timestamp in Unix epoch time.
* `created_by` - (String) IAM ID of creator.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `default_enable_new_features` - (Boolean) You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.
* `default_enable_new_regions` - (Boolean) Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.
* `enabled` - (Boolean) When enabled, Cloud Shell is available to all users in the account.
* `features` - (List) List of Cloud Shell features.
  * Constraints: The maximum length is `2` items. The minimum length is `0` items.
Nested schema for **features**:
	* `enabled` - (Boolean) State of the feature.
	* `key` - (String) Name of the feature.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.]*$/`.
* `regions` - (List) List of Cloud Shell region settings.
  * Constraints: The maximum length is `3` items. The minimum length is `0` items.
Nested schema for **regions**:
	* `enabled` - (Boolean) State of the region.
	* `key` - (String) Name of the region.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z-]*$/`.
* `rev` - (String) Unique revision number for the settings object.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `type` - (String) Type of api response object.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z_]*$/`.
* `updated_at` - (Integer) Timestamp of last update in Unix epoch time.
* `updated_by` - (String) IAM ID of last updater.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

