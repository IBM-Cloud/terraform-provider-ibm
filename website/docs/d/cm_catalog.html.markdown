---
layout: "ibm"
page_title: "IBM : ibm_cm_catalog"
description: |-
  Get information about ibm_cm_catalog
subcategory: "Catalog Management"
---

# ibm_cm_catalog

Provides a read-only data source for ibm_cm_catalog. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_catalog" "cm_catalog" {
	catalog_identifier = ibm_cm_catalog.cm_catalog.catalog_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `catalog_identifier` - (Optional, String) Catalog identifier.
* `label` - (Optional, String) Catalog label.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the ibm_cm_catalog.
* `catalog_filters` - (List) Filters for account and catalog filters.
Nested schema for **catalog_filters**:
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

* `catalog_icon_url` - (String) URL for an icon associated with this catalog.

* `catalog_banner_url` - (String) URL for a banner image for this catalog.

* `created` - (String) The date-time this catalog was created.

* `crn` - (String) CRN associated with the catalog.

* `disabled` - (Boolean) Denotes whether a catalog is disabled.

* `features` - (List) List of features associated with this catalog.
Nested scheme for **features**:
	* `description` - (String) Feature description.
	* `description_i18n` - (Map) A map of translated strings, by language code.
	* `title` - (String) Heading.
	* `title_i18n` - (Map) A map of translated strings, by language code.

* `id` - (String) Unique ID.

* `kind` - (String) Kind of catalog. Supported kinds are offering and vpe.

* `label` - (String) Display Name in the requested language.

* `metadata` - (Map) Catalog specific metadata.

* `offerings_url` - (String) URL path to offerings.

* `owning_account` - (String) Account that owns catalog.

* `resource_group_id` - (String) Resource group id the catalog is owned by.

* `rev` - (String) Cloudant revision.

* `short_description` - (String) Description in the requested language.

* `tags` - (List) List of tags associated with this catalog.

* `target_account_contexts` - (List) List of target account contexts for this catalog.
Nested scheme for **target_account_contexts**:
	* `api_key` - (String) API key of the target account.
	* `name` - (String) Unique name/identifier for this target account context.
	* `label` - (String) Label for this target account context.
	* `project_id` - (String) Project ID.
	* `trusted_profile` - (List) Trusted profile information.
	Nested scheme for **trusted_profile**:
		* `trusted_profile_id` - (String) Trusted profile ID.
		* `catalog_crn` - (String) CRN of this catalog.
		* `catalog_name` - (String) Name of this catalog.
		* `target_service_id` - (String) Target service ID.

* `updated` - (String) The date-time this catalog was last updated.

* `url` - (String) The url for this specific catalog.

