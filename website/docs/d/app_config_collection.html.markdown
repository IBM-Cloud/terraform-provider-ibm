---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration collection'
description: |-
  Get information about collection
---

# ibm_app_config_collection

Provides a read-only data source for `collection`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.  For more information, about App Configuration features, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example Usage

```hcl
data "ibm_app_config_collection" "app_config_collection" {
	guid = "guid"
    region = "region"
	expand = "expand"
	collection_id = "collection_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.

- `collection_id` - (Required, string) Collection Id of the collection.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `collection_id` - The unique identifier of the collection resouce.

- `name` - Collection name.

- `description` - Collection description.

- `tags` - Tags associated with the collection.

- `created_time` - Creation time of the collection.

- `updated_time` - Last updated time of the collection data.

- `href` - Collection URL.

- `features` - List of Features associated with the collection. Nested `features` blocks have the following structure:

    - `feature_id` - Feature id.

    - `name` - Feature name.

- `properties` - List of properties associated with the collection. Nested `properties` blocks have the following structure:

    - `property_id` - Property id.

    - `name` - Property name.

- `features_count` - Number of features associated with the collection.

- `properties_count` - Number of features associated with the collection.