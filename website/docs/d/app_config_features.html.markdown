---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration features'
description: |-
  List all the feature flags.
---

# ibm_app_config_features

Provides a read-only data source for all `feature flags`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_features" "app_config_features" {
  guid = "guid"
  tags = "tags"
  expand = "expand"
  limit = "limit"
  offset = "limit"
  environment_id = "environment_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.
- `tags` - (optional, string) Flter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, int) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, int) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `sort` - (optiona, string) Sort the feature details based on the specified attribute.
- `collections` - (optiona, array of string) Filter features by a list of comma separated collections.
- `segments` - (optiona, array of string) Filter features by a list of comma separated segments.
- `includes` - (optiona, array of string) Include the associated collections or targeting rules details in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the Features datasource.
- `features` - Array of Features. Nested `features` blocks have the following structure:

  - `name` - Feature name.

  - `feature_id` - Feature id.

  - `description` - Feature description.

  - `type` - Type of the feature (BOOLEAN, STRING, NUMERIC).

  - `enabled_value` - Value of the feature when it is enabled.

  - `disabled_value` - Value of the feature when it is disabled.

  - `tags` - Tags associated with the feature.

  - `enabled` - The state of the feature flag.

  - `segment_exists` - Denotes if the targeting rules are specified for the feature flag.

  - `segment_rules` - Segment Rules array. Nested `segment_rules` blocks have the following structure:

    - `rules` - Rules array. Nested `rules` blocks have the following structure:

      - `segments` - Segments array.

    - `value` - Value of the segment.

    - `order` - Order of the segment, used during evaluation.

  - `collections` - Collection array. Nested `collections` blocks have the following structure:

    - `collection_id` - Collection id.

    - `name` - Name of the collection.

  - `created_time` - Creation time of the feature flag.

  - `updated_time` - Last modified time of the feature flag data.

  - `href` - Feature flag URL.

- `total_count` - Number of records returned in the current response.

- `first` - URL to navigate to the first page of records. Nested `first` blocks have the following structure:

  - `href` - URL.

- `previous` - URL to navigate to the previous list of records. Nested `previous` blocks have the following structure:

  - `href` - URL.

- `last` - URL to navigate to the last list of records. Nested `last` blocks have the following structure:

  - `href` - URL.

- `next` - URL to navigate to the next list of records.. Nested `next` blocks have the following structure:
  - `href` - URL.
