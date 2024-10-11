---
layout: "ibm"
page_title: "IBM : ibm_config_aggregator_settings"
description: |-
  Get information about config_aggregator_settings
subcategory: "Configuration Aggregator"
---

# ibm_config_aggregator_settings

Provides a read-only data source to retrieve information about config_aggregator_settings. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_config_aggregator_settings" "config_aggregator_settings" {
	instance_id=var.instance_id
	region=var.region	
}
```


## Attribute Reference

After your data source is created, you can read values from the following attributes.
* `instance_id` - (Required, Forces new resource, String) The GUID of the Configuration Aggregator instance.
* `region` - (Optional, Forces new resource, String) The region of the Configuration Aggregator instance. If not provided defaults to the region defined in the IBM provider configuration.
* `id` - The unique identifier of the config_aggregator_settings.
* `additional_scope` - (List) The additional scope that enables resource collection for Enterprise acccounts.
  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
Nested schema for **additional_scope**:
	* `enterprise_id` - (String) The Enterprise ID.
	  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/[a-zA-Z0-9]/`.
	* `profile_template` - (List) The Profile Template details applied on the enterprise account.
	Nested schema for **profile_template**:
		* `id` - (String) The Profile Template ID created in the enterprise account that provides access to App Configuration instance for resource collection.
		  * Constraints: The maximum length is `52` characters. The minimum length is `52` characters. The value must match regular expression `/[a-zA-Z0-9-]/`.
		* `trusted_profile_id` - (String) The trusted profile ID that provides access to App Configuration instance to retrieve template information.
		  * Constraints: The maximum length is `44` characters. The minimum length is `44` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
	* `type` - (String) The type of scope. Currently allowed value is Enterprise.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[a-zA-Z0-9]/`.
* `last_updated` - (String) The last time the settings was last updated.
* `resource_collection_regions` - (List) Regions for which the resource collection is enabled.
  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9-]*$/`. The maximum length is `10` items. The minimum length is `0` items.
* `resource_collection_enabled` - (Boolean) The field to check if the resource collection is enabled.
* `trusted_profile_id` - (String) The trusted profile ID that provides access to App Configuration instance to retrieve resource metadata.
  * Constraints: The maximum length is `44` characters. The minimum length is `44` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.

