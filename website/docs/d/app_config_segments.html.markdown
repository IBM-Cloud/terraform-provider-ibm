---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration segments'
description: |-
  List all the segments.
---

# ibm_app_config_segments

Retrieve information about an existing IBM Cloud App Configuration segments. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration segments, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_segments" "app_config_segments" {
  guid = "guid"
  tags = "tags"
  expand = "expand"
  limit = "limit"
  offset = "limit"
  include = "include"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `tags` - (optional, String) Filter the resources to be returned based on the associated tags. Specify the parameter as a list of comma separated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, String) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `include` - (optional, String) Segment details to include the associated rules in the response.
- `sort` - (optional, String) Sort the segment details based on the specified attribute.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `segments` - (List) Array of Segments.
    
     Nested scheme for `segments`:
  - `name` - (String) Segment name.
  - `segment_id` - (String) Segment id.
  - `description` - (String) Segment description.
  - `tags` - (String) Tags associated with the segments.
  - `created_time` - (String) Creation time of the segment.
  - `updated_time` - (String) Last modified time of the segment data.
  - `href` - (String) Segment URL.
  - `rules` - (List) List of rules that determine if the entity belongs to the segment during feature / property evaluation. An entity is identified by an unique identifier and the attributes that it defines.

     Nested scheme for `rules`:
      - `attribute_name` - (String) Attribute name.
      - `operator` - (String) Operator to be used for the evaluation if the entity belongs to the segment.
      - `values` - (String) List of values. Entities matching any of the given values will be considered to belong to the segment.
  - `features` - (List) List of Features associated with the segment.
     
     Nested scheme for `features`:
      - `feature_id` - Feature id.
      - `name` - Feature name.
        
  - `properties` - (List) List of properties associated with the segment.
    
     Nested scheme for `properties`:
      - `property_id` - Property id.
      - `name` - Property name.
      

- `limit` - (Integer) Number of records returned
- `offset` - (Integer) Skipped number of records
- `total_count` - (Integer) Total number of records
- `first` - (List) The URL to navigate to the first page of records.
   
    Nested scheme for `first`:
    - `href` - (String) The first `href` URL.

- `last` - (List) The URL to navigate to the last page of records.

   Nested scheme for `last`:
    - `href` - (String) The last `href` URL.

- `previous` - (List) The URL to navigate to the previous list of records.

   Nested scheme for `previous`:
    - `href` - (String) The previous `href` URL.

- `next` - (List) The URL to navigate to the next list of records

  Nested scheme for `next`:
  - `href` - (String) The next `href` URL.
