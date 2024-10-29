---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration properties'
description: |-
  List all the properties.
---

# ibm_app_config_properties

Retrieve information about an existing IBM Cloud App Configuration properties. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration properties, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example Usage

```terraform
data "ibm_app_config_properties" "app_config_properties" {
	guid = "guid"
  region = "region"
	environment_id = "environment_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `region` - (Required, String)The region of the App Configuration Instance
- `environment_id` - (Required, String) Environment Id.
- `tags` - (optional, String) Flter the resources to be returned based on the associated tags. Returns resources associated with any of the specified tags.
- `expand` - (optional, bool) If set to `true`, returns expanded view of the resource details.
- `limit` - (optional, int) The number of records to retrieve. By default, the list operation return the first 10 records. To retrieve different set of records, use `limit` with `offset` to page through the available records.
- `offset` - (optional, int) The number of records to skip. By specifying `offset`, you retrieve a subset of items that starts with the `offset` value. Use `offset` with `limit` to page through the available records.
- `sort` - (optiona, String) Sort the feature details based on the specified attribute.
- `collections` - (optiona, array of String) Filter properties by a list of comma separated collections.
- `segments` - (optiona, array of String) Filter properties by a list of comma separated segments.
- `includes` - (optiona, array of String) Include the associated collections or targeting rules details in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the PropertiesList.
  - `properties` -  (List) Array of properties.

     Nested scheme for `properties`
      - `name` - (String) Property name.
      - `property_id` - (String) Property id.
      - `description` - (String) Property description.
      - `type` - (String) Type of the Property (BOOLEAN, STRING, NUMERIC).
      - `value` - Value of the Property. The value can be Boolean, String or a Numeric value as per the `type` attribute.
      - `tags` - (String) Tags associated with the property.
      - `format` -  (String) Format of the property (TEXT, JSON, YAML) and this is a required attribute when TYPE STRING is used, not required for BOOLEAN and NUMERIC types.
    
      - `segment_rules` - (List) Specify the targeting rules that is used to set different property values for different segments. 
  
        Nested scheme for `segment_rules`
        - `rules` - (List) The rules array. 
        
           Nested scheme for `rules`
          - `segments` - (String) List of segment ids that are used for targeting using the rule.
        - `value` - (String) Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
        - `order` - (String) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.
- `segment_exists` - (String) Denotes if the targeting rules are specified for the property.
- `collections` - List of collection id representing the collections that are associated with the specified property.
  
    Nested `collections` blocks have the following structure:
  - `collection_id` - (String) Collection id.
  - `name` - (String) Name of the collection.
  
- `created_time` - Creation time of the property.
- `updated_time` - Last modified time of the property data.
- `evaluation_time` - The last occurrence of the property value evaluation.
- `href` - Property URL.

- `total_count` - Number of records returned in the current response.
- `limit` - Number of records returned
- `offset` - Skipped number of records

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