---
layout: "ibm"
page_title: "IBM : cloud_shell_account_settings"
description: |-
  Manages cloud_shell_account_settings.
subcategory: "IBM Cloud Shell"
---

# ibm_cloud_shell_account_settings

Provides a resource for cloud_shell_account_settings. This allows cloud_shell_account_settings to be updated.

## Example usage

```terraform
resource "ibm_cloud_shell_account_settings" "cloud_shell_account_settings" {
  account_id = "12345678-abcd-1a2b-a1b2-1234567890ab"
  rev = "130-1bc9ec83d7b9b049890c6d4b74dddb2a"
  default_enable_new_features = true
  default_enable_new_regions = true
  enabled = true
  features {
  	enabled = true
  	key = "server.file_manager"
  }
  features {
  	enabled = true
  	key = "server.web_preview"
  }
  regions {
  	enabled = true
  	key = "eu-de"
  }
  regions {
  	enabled = true
  	key = "jp-tok"
  }
  regions {
  	enabled = true
  	key = "us-south"
  }
}
```

## Argument reference

The following arguments are supported:

* `account_id` - (Required, Forces new resource, string) The account ID in which the account settings belong to.
* `default_enable_new_features` - (Optional, bool) You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.
* `default_enable_new_regions` - (Optional, bool) Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.
* `enabled` - (Optional, bool) When enabled, Cloud Shell is available to all users in the account.
* `features` - (Optional, List) List of Cloud Shell features.
  * `enabled` - (Optional, bool) State of the feature.
  * `key` - (Optional, string) Name of the feature.
* `regions` - (Optional, List) List of Cloud Shell region settings.
  * `enabled` - (Optional, bool) State of the region.
  * `key` - (Optional, string) Name of the region.
* `rev` - (Required, string) Unique revision number for the settings object.  Required it this field is available from the data source.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloud_shell_account_settings.
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
* `rev` - Unique revision number for the settings object.
* `type` - Type of api response object.
* `updated_at` - Timestamp of last update in Unix epoch time.
* `updated_by` - IAM ID of last updater.

## Import

You can import the `cloud_shell_account_settings` resource by using `account_id`.

```
<account_id>
```
* `account_id`: A string. The account ID in which the account settings belong to.

```
$ terraform import cloud_shell_account_settings.cloud_shell_account_settings <account_id>
```
