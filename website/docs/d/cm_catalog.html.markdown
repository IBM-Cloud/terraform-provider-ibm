---
layout: "ibm"
page_title: "IBM : ibm_cm_catalog"
description: |-
  Get information about cm_catalog
subcategory: "Catalog Management API"
---

# ibm_cm_catalog

Provides a read-only data source for cm_catalog. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_catalog" "cm_catalog" {
	catalog_identifier = ibm_cm_catalog.cm_catalog.catalog_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `catalog_identifier` - (Required, Forces new resource, String) Catalog identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cm_catalog.
* `catalog_filters` - (List) Filters for account and catalog filters.
Nested scheme for **catalog_filters**:
	* `category_filters` - (Map) Filter against offering properties.
	* `id_filters` - (List) Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
	Nested scheme for **id_filters**:
		* `exclude` - (List) Offering filter terms.
		Nested scheme for **exclude**:
			* `filter_terms` - (List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
		* `include` - (List) Offering filter terms.
		Nested scheme for **include**:
			* `filter_terms` - (List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
	* `include_all` - (Boolean) -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.

* `catalog_icon_url` - (String) URL for an icon associated with this catalog.

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

* `syndication_settings` - (List) Feature information.
Nested scheme for **syndication_settings**:
	* `authorization` - (List) Feature information.
	Nested scheme for **authorization**:
		* `last_run` - (String) Date and time last updated.
		* `token` - (String) Array of syndicated namespaces.
	* `clusters` - (List) Syndication clusters.
	Nested scheme for **clusters**:
		* `all_namespaces` - (Boolean) Syndicated to all namespaces on cluster.
		* `id` - (String) Cluster ID.
		* `name` - (String) Cluster name.
		* `namespaces` - (List) Syndicated namespaces.
		* `region` - (String) Cluster region.
		* `resource_group_name` - (String) Resource group ID.
		* `type` - (String) Syndication type.
	* `history` - (List) Feature information.
	Nested scheme for **history**:
		* `clusters` - (List) Array of syndicated namespaces.
		Nested scheme for **clusters**:
			* `all_namespaces` - (Boolean) Syndicated to all namespaces on cluster.
			* `id` - (String) Cluster ID.
			* `name` - (String) Cluster name.
			* `namespaces` - (List) Syndicated namespaces.
			* `region` - (String) Cluster region.
			* `resource_group_name` - (String) Resource group ID.
			* `type` - (String) Syndication type.
		* `last_run` - (String) Date and time last syndicated.
		* `namespaces` - (List) Array of syndicated namespaces.
	* `remove_related_components` - (Boolean) Remove related components.

* `tags` - (List) List of tags associated with this catalog.

* `updated` - (String) The date-time this catalog was last updated.

* `url` - (String) The url for this specific catalog.

