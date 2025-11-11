---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration collection'
description: |-
  Manages Collection.
---

# ibm_app_config_collection

Provides a resource for `collection`. This allows collection to be created, updated and deleted. For more information, about App Configuration feature flag, see [Collection](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-collections).

## Example Usage

```hcl
resource "ibm_app_config_collection" "app_config_collection" {
  guid = "guid"
  name = "name"
  tags = "tag1,tag2"
  description = "description"
  collection_id = "collection_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `name` - (Required, string) Collection name.
- `collection_id` - (Required, string) Collection Id.
- `description` - (Optional, string) Collection description.
- `tags` - (Optional, string) Tags associated with the collection.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the Collection.
- `created_time` - Creation time of the collection.
- `updated_time` - Last updated time of the collection data.
- `href` - Collection URL.

## Import

The `ibm_app_config_collection` resource can be imported by using `guid` of the App Configuration instance and `collectionId`. Get `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_collection.sample  <guid/collectionId>

```

**Example**

```
terraform import ibm_app_config_collection.sample 272111153-c118-4116-8116-b811fbc31132/col
```
