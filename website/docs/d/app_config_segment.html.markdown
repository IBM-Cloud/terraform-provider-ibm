---
layout: 'ibm'
page_title: 'IBM : app_config_segment'
sidebar_current: 'docs-ibm-datasource-app-config-segment'
description: |-
  Get information about Segment
subcategory: 'App Configuration'
---

# ibm_app_config_segment

Provides a read-only data source for Segment. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_segment" "app_config_segment" {
  guid = "guid"
  segment_id = "segment_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `segment_id` - (Required, string) Segment Id.
- `includes` - (optiona, array of string) Include feature and property details in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the Segment.

- `name` - Segment name.

- `description` - Segment description.

- `tags` - Tags associated with the segments.

- `rules` - List of rules that determine if the entity is part of the segment. Nested `rules` blocks have the following structure:

  - `attribute_name` - Attribute name.

  - `operator` - Operator to be used for the evaluation if the entity is part of the segment.

  - `values` - List of values. Entities matching any of the given values will be considered to be part of the segment.

- `created_time` - Creation time of the segment.

- `updated_time` - Last modified time of the segment data.

- `href` - Segment URL.
