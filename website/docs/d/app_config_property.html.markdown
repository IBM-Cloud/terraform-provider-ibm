---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration property'
description: |-
  Get information about property.
---

# ibm_app_config_property

Retrieve information about an existing IBM Cloud App Configuration property. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration property, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example Usage

```terraform
data "ibm_app_config_property" "app_config_property" {
	guid = "guid"
    region = "region"
	environment_id = "environment_id"
	property_id = "property_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `region` - (Required, String)The region of the App Configuration Instance
- `environment_id` - (Required, String) Environment Id.
- `property_id` - (Required, String) Property Id.
- `include` - (Optional, String) Include the associated collections in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `property_id` - The unique identifier of the app_config_property.
- `name` - Property name.
- `description` - Property description.
- `type` - Type of the Property (BOOLEAN, STRING, NUMERIC).
- `value` - Value of the Property. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `tags` - Tags associated with the property.
- `format` - Format of the property (TEXT, JSON, YAML) and this is a required attribute when TYPE STRING is used, not required for BOOLEAN and NUMERIC types.
- `segment_rules` - Specify the targeting rules that is used to set different property values for different segments. Nested `segment_rules` blocks have the following structure:
    - `rules` - Rules array. Nested `rules` blocks have the following structure:
        - `segments` - List of segment ids that are used for targeting using the rule.
    - `value` - Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
    - `order` - Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.

- `segment_exists` - Denotes if the targeting rules are specified for the property.
- `collections` - List of collection id representing the collections that are associated with the specified property. Nested `collections` blocks have the following structure:
    - `collection_id` - Collection id.
    - `name` - Name of the collection.

- `created_time` - Creation time of the property.
- `updated_time` - Last modified time of the property data.
- `evaluation_time` - The last occurrence of the property value evaluation.
- `href` - Property URL.