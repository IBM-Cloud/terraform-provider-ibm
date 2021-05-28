---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segments'
description: |-
  List all the segments.
---

# ibm_app_config_segments

Provides a read-only data source for `segments`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_segments" "app_config_segments" {
  guid = "guid"
  tags = "tags"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `tags` - (optional, string) Flter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, int) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, int) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `sort` - (optiona, string) Sort the feature details based on the specified attribute.
- `includes` - (optiona, string) Segment details to include the associated rules in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the SegmentsList.

- `segments` - Array of Segments. Nested `segments` blocks have the following structure:

  - `name` - Segment name.

  - `segment_id` - Segment id.

  - `description` - Segment description.

  - `tags` - Tags associated with the segments.

  - `rules` - List of rules that determine if the entity is part of the segment. Nested `rules` blocks have the following structure:

    - `attribute_name` - Attribute name.

    - `operator` - Operator to be used for the evaluation if the entity is part of the segment.

    - `values` - List of values. Entities matching any of the given values will be considered to be part of the segment.

  - `created_time` - Creation time of the segment.

  - `updated_time` - Last modified time of the segment data.

  - `href` - Segment URL.

- `total_count` - Total number of records.

- `first` - URL to navigate to the first page of records. Nested `first` blocks have the following structure:

  - `href` - URL.

- `previous` - URL to navigate to the previous list of records. Nested `previous` blocks have the following structure:

  - `href` - URL.

- `last` - URL to navigate to the last list of records. Nested `last` blocks have the following structure:

  - `href` - URL.

- `next` - URL to navigate to the next list of records.. Nested `next` blocks have the following structure:
  - `href` - URL.
