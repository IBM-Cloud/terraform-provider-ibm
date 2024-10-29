---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segment'
description: |-
  Get information about segment
---

# ibm_app_config_segment
Retrieve information about an existing IBM Cloud App Configuration segment. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration segment, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_segment" "app_config_segment" {
  guid = "guid"
  segment_id = "segment_id"
  includes = "includes"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.

- `segment_id` - (Required, String) The segment ID.
- `includes` - (Optional, String) Include feature and property details in the response.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `segment_id` - (String) Segment id.
- `name` - (String) Segment name.
- `description` - (String) Segment description.
- `tags` - (String) Tags associated with the segments.
- `created_time` - (Timestamp) Creation time of the segment.
- `updated_time` - (Timestamp) Last modified time of the segment data.
- `href` - (String) Segment URL.
- `rules` - (String) List of rules that determine if the entity belongs to the segment during feature / property evaluation.

  Nested scheme for `rules`:
  - `attribute_name` - (String) Attribute name.
  - `operator` - (String) Operator to be used for the evaluation if the entity belongs to the segment.
  - `values` - (String) List of values. Entities matching any of the given values will be considered to belong to the segment.

- `features` - (List) List of Features associated with the segment.

  Nested scheme for `features`:
  - `feature_id` - (String) Feature Id.
  - `name` - (String) Feature Name.

- `properties` - (List) List of properties associated with the segment.

  Nested scheme for `properties`:
  - `property_id` - (String) Property Id.
  - `name` - (String) Property Name.
