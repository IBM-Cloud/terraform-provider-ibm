---
layout: "ibm"
page_title: "IBM : ibm_cm_account"
description: |-
  Get information about cm_account
subcategory: "Catalog Management API"
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

