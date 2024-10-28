---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration collections'
description: |-
  List all the collections.
---

# ibm_app_config_collections

Provides a read-only data source for `collection`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration features flag, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example Usage

```hcl
data "app_config_collections" "app_config_collections" {
	guid = "guid"
  region="region"
  tags = "tags"
  expand = "expand"
  limit = "limit"
  offset = "limit"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `region` - (Required, String)The region of the App Configuration Instance
- `sort` - (optional, string) Sort the collection details based on the specified attribute.
- `tags` - (optional, string) Flter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, int) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, int) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `properties` - (optional, array of string) Filter collections by a list of comma separated properties.
- `features` - (optional, array of string) Include the associated collections or targeting rules details in the response.
- `includes` - (optional, array of string) Include feature and property details in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the collections source.
- `collections` - Array of collections. Nested `collections` blocks have the following structure:

- `name` - Collection name.

- `collection_id` - Collection Id.

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

- `total_count` - Number of records returned in the current response.

- `first` - (List) The URL to navigate to the first page of records.

  Nested scheme for `first`:
  - `href` - (String) The first `href` URL.
  
- `previous` - (List) The URL to navigate to the previous list of records.

  Nested scheme for `previous`:
  - `href` - (String) The previous `href` URL.

- `last` - (List) The URL to navigate to the last list of records.

  Nested scheme for `last`:
  - `href` - (String) The last `href` URL.

- `next` - (List) The URL to navigate to the next list of records.

  Nested scheme for `next`:
  - `href` - (String) The next `href` URL.
