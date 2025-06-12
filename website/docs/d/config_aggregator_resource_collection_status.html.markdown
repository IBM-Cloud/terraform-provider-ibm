---
layout: "ibm"
page_title: "IBM : ibm_config_aggregator_resource_collection_status"
description: |-
  Get information about config_aggregator_resource_collection_status
subcategory: "Configuration Aggregator"
---

# ibm_config_aggregator_resource_collection_status

Provides a read-only data source to retrieve information about config_aggregator_resource_collection_status. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_config_aggregator_resource_collection_status" "config_aggregator_resource_collection_status" {
	instance_id=var.instance_id
	region=var.region
}
```


## Attribute Reference

After your data source is created, you can read values from the following attributes.
* `instance_id` - (Required, Forces new resource, String) The GUID of the Configuration Aggregator instance.
* `region` - (Optional, Forces new resource, String) The region of the Configuration Aggregator instance. If not provided defaults to the region defined in the IBM provider configuration.
* `id` - The unique identifier of the config_aggregator_resource_collection_status.
* `last_config_refresh_time` - (String) The timestamp at which the configuration was last refreshed.
* `status` - (String) Status of the resource collection.
  * Constraints: Allowable values are: `initiated`, `inprogress`, `complete`.

