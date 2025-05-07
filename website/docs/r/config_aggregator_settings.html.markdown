---
layout: "ibm"
page_title: "IBM : ibm_config_aggregator_settings"
description: |-
  Manages config_aggregator_settings.
subcategory: "Configuration Aggregator"
---

# ibm_config_aggregator_settings

Create, update, and delete config_aggregator_settingss with this resource.

## Example Usage

```hcl
resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
  instance_id = var.instance_id
  region = var.region
  additional_scope {
		type = "Enterprise"
		enterprise_id = "enterprise_id"
		profile_template {
			id = "ProfileTemplate-adb55769-ae22-4c60-aead-bd1f84f93c57"
			trusted_profile_id = "Profile-6bb60124-8fc3-4d18-b63d-0b99560865d3"
		}
  }
  resource_collection_regions = ["all"]
  resource_collection_enabled = true
  trusted_profile_id = "Profile-1260aec2-f2fc-44e2-8697-2cc15a447560"
}
```

## Argument Reference

You can specify the following arguments for this resource.
* `instance_id` - (Required, Forces new resource, String) The GUID of the Configuration Aggregator instance.
* `region` - (Optional, Forces new resource, List) The region of the Configuration Aggregator instance. If not provided defaults to the region defined in the IBM provider configuration.
* `additional_scope` - (Optional, Forces new resource, List) The additional scope that enables resource collection for Enterprise acccounts.
  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
Nested schema for **additional_scope**:
	* `enterprise_id` - (Optional, String) The Enterprise ID.
	  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/[a-zA-Z0-9]/`.
	* `profile_template` - (Optional, List) The Profile Template details applied on the enterprise account.
	Nested schema for **profile_template**:
		* `id` - (Optional, String) The Profile Template ID created in the enterprise account that provides access to App Configuration instance for resource collection.
		  * Constraints: The maximum length is `52` characters. The minimum length is `52` characters. The value must match regular expression `/[a-zA-Z0-9-]/`.
		* `trusted_profile_id` - (Optional, String) The trusted profile ID that provides access to App Configuration instance to retrieve template information.
		  * Constraints: The maximum length is `44` characters. The minimum length is `44` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
	* `type` - (Optional, String) The type of scope. Currently allowed value is Enterprise.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[a-zA-Z0-9]/`.
* `resource_collection_regions` - (Required,List) Regions for which the resource collection is enabled.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9-]*$/`. The maximum length is `10` items. The minimum length is `0` items.
* `resource_collection_enabled` - (Required, Boolean) The field denoting if the resource collection is enabled.
* `trusted_profile_id` - (Required, String) The trusted profile id that provides Reader access to the App Configuration instance to collect resource metadata.
  * Constraints: The maximum length is `44` characters. The minimum length is `44` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the config_aggregator_settings.


## Import

You can import the `ibm_config_aggregator_settings` resource by using `region` and `instance_id`. 
# Syntax
<pre>
$ terraform import ibm_config_aggregator_settings.config_aggregator_settings <region>/<instance_id>
</pre>

# Example
```
$ terraform import ibm_config_aggregator_settings.config_aggregator_settings us-south/23243-3223-2323-333
```
