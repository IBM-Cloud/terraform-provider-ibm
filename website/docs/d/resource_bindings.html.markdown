---
layout: "ibm"
page_title: "IBM : resource_bindings"
description: |-
  Get information about resource_bindings
subcategory: "Resource management"
---

# ibm\_resource_bindings

Provides a read-only data source for resource_bindings. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_resource_bindings" "resource_bindings" {
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The human-readable name of the binding.
* `guid` - (Optional, string) Short ID of the binding.
* `resource_group_id` - (Optional, string) Short ID of resource group.
* `resource_id` - (Optional, string) The unique ID of the offering (service name). This value is provided by and stored in the global catalog.
* `region_binding_id` - (Optional, string) Short ID of the instance in a specific targeted environment.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the resource_bindings.

* `bindings` - A list of resource bindings. Nested `bindings` blocks have the following structure:
	* `id` - The ID associated with the binding.
	* `guid` - When you create a new binding, a globally unique identifier (GUID) is assigned. This GUID is a unique internal identifier managed by the resource controller that corresponds to the binding.
	* `url` - When you provision a new binding, a relative URL path is created identifying the location of the binding.
	* `created_at` - The date when the binding was created.
	* `updated_at` - The date when the binding was last updated.
	* `deleted_at` - The date when the binding was deleted.
	* `created_by` - The subject who created the binding.
	* `updated_by` - The subject who updated the binding.
	* `deleted_by` - The subject who deleted the binding.
	* `source_crn` - The CRN of resource alias associated to the binding.
	* `target_crn` - The CRN of target resource, for example, application, in a specific environment.
	* `crn` - The full Cloud Resource Name (CRN) associated with the binding. For more information about this format, see [Cloud Resource Names](https://cloud.ibm.com/docs/overview?topic=overview-crn).
	* `region_binding_id` - The ID of the binding in the specific target environment, for example, `service_binding_id` in a given IBM Cloud environment.
	* `region_binding_crn` - The CRN of the binding in the specific target environment.
	* `name` - The human-readable name of the binding.
	* `account_id` - An alpha-numeric value identifying the account ID.
	* `resource_group_id` - The ID of the resource group.
	* `state` - The state of the binding.
	* `credentials` - The credentials for the binding. Additional key-value pairs are passed through from the resource brokers.  For additional details, see the service’s documentation. Nested `credentials` blocks have the following structure:
		* `apikey` - The API key for the credentials.
		* `iam_apikey_description` - The optional description of the API key.
		* `iam_apikey_name` - The name of the API key.
		* `iam_role_crn` - The Cloud Resource Name for the role of the credentials.
		* `iam_serviceid_crn` - The Cloud Resource Name for the service ID of the credentials.
	* `iam_compatible` - Specifies whether the binding’s credentials support IAM.
	* `resource_id` - The unique ID of the offering. This value is provided by and stored in the global catalog.
	* `migrated` - A boolean that dictates if the alias was migrated from a previous CF instance.
	* `resource_alias_url` - The relative path to the resource alias that this binding is associated with.

