---
layout: "ibm"
page_title: "IBM : ibm_cloud_shell_account_settings"
description: |-
  Manages cloud_shell_account_settings.
subcategory: "IBM Cloud Shell"
---

# ibm_cloud_shell_account_settings

Create, update, and delete cloud_shell_account_settingss with this resource.

## Example Usage

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
  	key = "us-south"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) The account ID in which the account settings belong to.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `default_enable_new_features` - (Optional, Boolean) You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations.
* `default_enable_new_regions` - (Optional, Boolean) Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location.
* `enabled` - (Optional, Boolean) When enabled, Cloud Shell is available to all users in the account.
* `features` - (Optional, List) List of Cloud Shell features.
  * Constraints: The maximum length is `2` items. The minimum length is `0` items.
Nested schema for **features**:
	* `enabled` - (Optional, Boolean) State of the feature.
	* `key` - (Optional, String) Name of the feature.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.]*$/`.
* `regions` - (Optional, List) List of Cloud Shell region settings.
  * Constraints: The maximum length is `3` items. The minimum length is `0` items.
Nested schema for **regions**:
	* `enabled` - (Optional, Boolean) State of the region.
	* `key` - (Optional, String) Name of the region.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z-]*$/`.
* `rev` - (Optional, String) Unique revision number for the settings object.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cloud_shell_account_settings.
* `created_at` - (Integer) Creation timestamp in Unix epoch time.
* `created_by` - (String) IAM ID of creator.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `id` - (String) Unique id of the settings object.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `type` - (String) Type of api response object.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z_]*$/`.
* `updated_at` - (Integer) Timestamp of last update in Unix epoch time.
* `updated_by` - (String) IAM ID of last updater.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.


## Import

You can import the `ibm_cloud_shell_account_settings` resource by using `account_id`.
The `account_id` property can be formed from and `account_id` in the following format:

<pre>
&lt;account_id&gt;
</pre>
* `account_id`: A string. The account ID in which the account settings belong to.

# Syntax
<pre>
$ terraform import ibm_cloud_shell_account_settings.cloud_shell_account_settings &lt;account_id&gt;
</pre>
