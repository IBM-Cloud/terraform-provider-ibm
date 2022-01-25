---
layout: "ibm"
page_title: "IBM : cloud_shell_account_settings"
description: |-
  Get information about cloud_shell_account_settings
subcategory: "IBM Cloud Shell"
---

# ibm_cloud_shell_account_settings

Provides a read-only data source for cloud_shell_account_settings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "cloud_shell_account_settings" "cloud_shell_account_settings" {
	account_id = "account_id"
}
```

## Argument reference

The following arguments are supported:

* `account_id` - (Required, string) The account ID in which the account settings belong to.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloud_shell_account_settings.

* `rev` - Unique revision number for the settings object.

* `created_at` - Creation timestamp in Unix epoch time.

* `created_by` - IAM ID of creator.

* `default_enable_new_features` - You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.

* `default_enable_new_regions` - Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.

* `enabled` - When enabled, Cloud Shell is available to all users in the account.

* `features` - List of Cloud Shell features. Nested `features` blocks have the following structure:
	* `enabled` - State of the feature.
	* `key` - Name of the feature.

* `regions` - List of Cloud Shell region settings. Nested `regions` blocks have the following structure:
	* `enabled` - State of the region.
	* `key` - Name of the region.

* `type` - Type of api response object.

* `updated_at` - Timestamp of last update in Unix epoch time.

* `updated_by` - IAM ID of last updater.

