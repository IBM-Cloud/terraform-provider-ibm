---
layout: "ibm"
page_title: "IBM : ibm_config_aggregator_configurations"
description: |-
  Get information about config_aggregator_configurations
subcategory: "Configuration Aggregator"
---

# ibm_config_aggregator_configurations

Provides a read-only data source to retrieve information about config_aggregator_configurations. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_config_aggregator_configurations" "config_aggregator_configurations" {
	instance_id=var.instance_id
	region=var.region
}
```

## Argument Reference

You can specify the following arguments for this data source.
* `instance_id` - (Required, Forces new resource, String) The GUID of the Configuration Aggregator instance.
* `region` - (Optional, Forces new resource, String) The region of the Configuration Aggregator instance. If not provided defaults to the region defined in the IBM provider configuration.
* `config_type` - (Optional, String) The type of resource configuration that are to be retrieved.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 ,\\-_]+$/`.
* `location` - (Optional, String) The location or region in which the resources are created.
  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^$|[a-z]-[a-z]/`.
* `resource_crn` - (Optional, String) The crn of the resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
* `resource_group_id` - (Optional, String) The resource group id of the resources.
  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `service_name` - (Optional, String) The name of the IBM Cloud service for which resources are to be retrieved.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 ,\\-_]+$/`.
* `sub_account` - (Optional, String) Filter the resource configurations from the specified sub-account in an enterprise hierarchy. Used for fetching enterprise child accounts configurations.
  * Constraints: The length is `32` characters. The value must match regular expression `[a-zA-Z0-9]`.
* `access_tags` - (Optional, String) Filter the resource configurations attached with the specified access tags.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `[a-zA-Z0-9]`.
* `user_tags` - (Optional, String)Filter the resource configurations attached with the specified user tags.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `[a-zA-Z0-9]`.
* `service_tags` - (Optional, String) Filter the resource configurations attached with the specified service tags.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `[a-zA-Z0-9]`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the config_aggregator_configurations.
* `configs` - (List) Array of resource configurations.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **configs**:
	* `about` - (Map) The basic metadata fetched from the query API.
	Nested schema for **about**:
		* `account_id` - (String) The account ID in which the resource exists.
		  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
		* `config_type` - (String) The type of configuration of the retrieved resource.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
		* `last_config_refresh_time` - (String) Date/time stamp identifying when the information was last collected. Must be in the RFC 3339 format.
		* `location` - (String) Location of the resource specified.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `0` characters. The value must match regular expression `/^$|[a-z]-[a-z]/`.
		* `resource_crn` - (String) The unique CRN of the IBM Cloud resource.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
		* `resource_group_id` - (String) The account ID.
		  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
		* `resource_name` - (String) User defined name of the resource.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
		* `service_name` - (String) The name of the service to which the resources belongs.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.\\:\/-]+$/`.
		* `tags` - (List) Tags associated with the resource.
		Nested schema for **tags**:
			* `tag` - (String) The name of the tag.
			  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
	* `config` - (String) The configuration of the resource.
	Nested schema for **config**:
* `prev` - (List) The reference to the previous page of entries.
Nested schema for **prev**:
	* `href` - (String) The reference to the previous page of entries.
	* `start` - (String) the start string for the query to view the page.

