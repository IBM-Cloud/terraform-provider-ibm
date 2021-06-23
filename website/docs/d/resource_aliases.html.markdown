---
layout: "ibm"
page_title: "IBM : resource_aliases"
description: |-
  Get information about resource_aliases
subcategory: "Resource management"
---

# ibm\_resource_aliases

Provides a read-only data source for resource_aliases. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_resource_aliases" "resource_aliases" {
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The human-readable name of the alias.
* `guid` - (Optional, string) Short ID of the alias.
* `resource_instance_id` - (Optional, string) Resource instance CRN.
* `region_instance_id` - (Optional, string) Short ID of the instance in a specific targeted environment.
* `resource_id` - (Optional, string) The unique ID of the offering (service name). This value is provided by and stored in the global catalog.
* `resource_group_id` - (Optional, string) Short ID of Resource group.
## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the resource_aliases.

* `aliases` - A list of resource aliases. Nested `aliases` blocks have the following structure:
	* `id` - The ID associated with the alias.
	* `guid` - When you create a new alias, a globally unique identifier (GUID) is assigned. This GUID is a unique internal indentifier managed by the resource controller that corresponds to the alias.
	* `url` - When you created a new alias, a relative URL path is created identifying the location of the alias.
	* `created_at` - The date when the alias was created.
	* `updated_at` - The date when the alias was last updated.
	* `deleted_at` - The date when the alias was deleted.
	* `created_by` - The subject who created the alias.
	* `updated_by` - The subject who updated the alias.
	* `deleted_by` - The subject who deleted the alias.
	* `name` - The human-readable name of the alias.
	* `resource_instance_id` - The ID of the resource instance that is being aliased.
	* `target_crn` - The CRN of the target namespace in the specific environment.
	* `account_id` - An alpha-numeric value identifying the account ID.
	* `resource_id` - The unique ID of the offering. This value is provided by and stored in the global catalog.
	* `resource_group_id` - The ID of the resource group.
	* `crn` - The CRN of the alias. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
	* `region_instance_id` - The ID of the instance in the specific target environment, for example, `service_instance_id` in a given IBM Cloud environment.
	* `region_instance_crn` - The CRN of the instance in the specific target environment.
	* `state` - The state of the alias.
	* `migrated` - A boolean that dictates if the alias was migrated from a previous CF instance.
	* `resource_instance_url` - The relative path to the resource instance.
	* `resource_bindings_url` - The relative path to the resource bindings for the alias.
	* `resource_keys_url` - The relative path to the resource keys for the alias.

