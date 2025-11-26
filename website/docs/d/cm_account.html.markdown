---
layout: "ibm"
page_title: "IBM : ibm_cm_account"
description: |-
  Get information about cm_account
subcategory: "Catalog Management"
---

# ibm_cm_account

Provides a read-only data source to retrieve information about a cm_account. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_account" "cm_account" {
}
```


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cm_account.
* `account_filters` - (List) Filters for account and catalog filters.
Nested schema for **account_filters**:
	* `category_filters` - (List) Filter against offering properties.
	Nested schema for **category_filters**:
    	* `category_name` - (String) Name of the category.
    	* `include` -  (Boolean) Whether to include the category in the catalog filter.
    	* `filter` - (List) Filter terms related to the category.
		Nested schema for **filter**:
			* `filter_terms` - (List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
	* `id_filters` - (List) Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
	Nested schema for **id_filters**:
		* `exclude` - (List) Offering filter terms.
		Nested schema for **exclude**:
			* `filter_terms` - (List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
		* `include` - (List) Offering filter terms.
		Nested schema for **include**:
			* `filter_terms` - (List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
	* `include_all` - (Boolean) -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.
* `hide_ibm_cloud_catalog` - (Boolean) Hide the public catalog in this account.
* `region_filter` - (String) Region filter string.
* `rev` - (String) Cloudant revision.
* `terraform_engines` - (List) List of terraform engines configured for this account.
Nested schema for **terraform_engines**:
	* `api_token` - (String) The api key used to access the engine instance.
	* `da_creation` - (List) The settings that determines how deployable architectures are auto-created from workspaces in the terraform engine.
	Nested schema for **da_creation**:
		* `default_private_catalog_id` - (String) Default private catalog to create the deployable architectures in.
		* `enabled` - (Boolean) Determines whether deployable architectures are auto-created from workspaces in the engine.
		* `polling_info` - (List) Determines which workspace scope to query to auto-create deployable architectures from.
		Nested schema for **polling_info**:
			* `last_polling_status` - (List) Last polling status of the engine scope.
			Nested schema for **last_polling_status**:
				* `code` - (Integer) Status code of the last polling attempt.
				* `message` - (String) Status message from the last polling attempt.
			* `scopes` - (List) List of scopes to auto-create deployable architectures from workspaces in the engine.
			Nested schema for **scopes**:
				* `name` - (String) Identifier for the specified type in the scope.
				* `type` - (String) Scope to auto-create deployable architectures from. The supported scopes today are workspace, org, and project.
	* `name` - (String) User provided name for the specified engine.
	* `private_endpoint` - (String) The private endpoint for the engine instance.
	* `public_endpoint` - (String) The public endpoint for the engine instance.
	* `type` - (String) The terraform engine type. The only one supported at the moment is terraform-enterprise.

