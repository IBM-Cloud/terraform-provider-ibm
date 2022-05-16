---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration features'
description: |-
  List all the feature flags.
---

# ibm_app_config_features

Retrieve information about an existing IBM Cloud App Configuration features flag. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration features flag, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_features" "app_config_features" {
  guid = "guid"
  tags = "tags"
  expand = "expand"
  limit = "limit"
  offset = "limit"
  environment_id = "environment_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String)  Environment ID.
- `tags` - (optional, String) Flter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, Bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, Integer) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, Integer) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `sort` - (optiona, String) Sort the feature details based on the specified attribute.
- `collections` - (optiona, Array of String) Filter features by a list of comma separated collections.
- `segments` - (optiona, Array of String) Filter features by a list of comma separated segments.
- `includes` - (optiona, Array of String) Include the associated collections or targeting rules details in the response.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the Features datasource.
- `features` - (List) Array of Features. 

   Nested scheme for `features`:

  - `name` - (String) The feature name.
  - `feature_id` - (String) The feature ID.
  - `description` - (String) The feature description.
  - `type` -  (String) The type of the feature. Supported values are `BOOLEAN`, `STRING`, and `NUMERIC`).
  - `enabled_value` - (String) The value of the feature when it is enabled.
  - `disabled_value` - (String) The value of the feature when it is disabled.
  - `tags` - (String) The tags associated with the feature.
  - `rollout_percentage` - (String) Rollout percentage of the feature.
  - `enabled` - (String) The state of the feature flag.
  - `segment_exists` - (String) Denotes if the targeting rules are specified for the feature flag.
  - `segment_rules` - (List) The segment rules array. 
  
    Nested scheme for `segment_rules`:
    - `rules` - (List) The rules array. 
    
      Nested scheme for `rules`:
      - `segments` - (String) The Segments array.
    - `value` - (String) Value of the segment.
    - `order` - (String) Order of the segment, used during evaluation.
    - `rollout_percentage` - (String) Rollout percentage for the segment rule.
  - `collections` - (List) The collection array. 
  
    Nested scheme for `collections`:
    - `collection_id` - (String) The collection ID.
    - `name` - (String) Name of the collection.
  - `created_time` - (String) Creation time of the feature flag.
  - `updated_time` - (String) Last modified time of the feature flag data.
  - `href` - (String) Feature flag URL.
- `total_count` - (String) Number of records returned in the current response.
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
