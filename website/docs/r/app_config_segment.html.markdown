---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segments'
description: |-
  Manages segments.
---

# ibm_app_config_segment

Provides a resource for `segment`. This allows `segment` to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_app_config_segment" "app_config_segment" {
  guid = "guid"
  name = "name"
  segment_id = "segment_id"
  description = "description"
  tags = "tags"
  rules = "rules"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `name` - (Required, string) Segment name.
- `segment_id` - (Required, string) Segment id.
- `description` - (Optional, string) Segment description.
- `tags` - (Optional, string) Tags associated with the segments.
- `rules` - (Required, List) List of rules that determine if the entity is part of the segment.

  - `attribute_name` - (Required, string) Attribute name.

  - `operator` - (Required, string) Operator to be used for the evaluation if the entity is part of the segment.

  - `values` - (Required, []interface{}) List of values. Entities matching any of the given values will be considered to be part of the segment.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the resource segment.
- `created_time` - Creation time of the segment.
- `updated_time` - Last modified time of the segment data.
- `href` - Segment URL.

## Import

The `ibm_app_config_segment` resource can be imported by using `guid` of the App Configuration instance and `segmentId`. Get `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_segment.sample  <guid/segmentId>

```

**Example**

```
terraform import ibm_app_config_segment.sample 272111153-c118-4116-8116-b811fbc31132/sample_segment
```
