---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segment'
description: |-
  Manages segment.
---

# ibm_app_config_segment

Create, update, or delete an segment by using IBM Cloud™ App Configuration. For more information, about App Configuration segment, see [segments](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-segments).

## Example usage

```terraform
resource "ibm_app_config_segment" "app_config_segment" {
  guid = "guid"
  name = "name"
  description = "description"
  tags = "tag1,tag2"
  segment_id = "segment_id"
  rules {
    attribute_name = "attribute_name"
    operator = "operator"
    values = "values"
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) The GUID of the App Configuration service. Fetch GUID from the service instance credentials section of the dashboard.
- `name` - (Required, String) The Segment name.
- `rules` - (Required, List) List of rules that determine if the entity belongs to the segment during feature / property evaluation.
  
  Nested scheme for `rules`:
    - `attribute_name` - (Required, String) The Attribute name.
    - `operator` - (Required, String) The Operator to be used for the evaluation if the entity belongs to the segment.
    - `values` - (Required, Array of Strings) List of values. Entities matching any of the given values will be considered to belong to the segment.
  
- `description` - (Optional, String) The Segment description.
- `tags` - (Optional, String) Tags associated with the segments.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `segment_id` - (String) Segment id.
- `created_time` - (Timestamp) Creation time of the segment.
- `updated_time` - (Timestamp) Last modified time of the segment data.
- `href` - (String) Segment URL.
- `features` - (List) List of Features associated with the segment.
   
  Nested scheme for `features`:
    - `feature_id` - (String) Feature id.
    - `name` - (String) Feature name.

- `properties` - (List) List of properties associated with the segment.

  Nested scheme for `properties`:
    - `property_id` - (String) Property id.
    - `name` - (String) Property name.

## Import

The `ibm_app_config_segment` resource can be imported by using `guid` of the App Configuration instance and `segmentId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_segment.sample  <guid/segmentId>

```

**Example**

```
terraform import ibm_app_config_segment.sample 272111153-c118-4116-8116-b811fbc31132/sample_segment
```
