---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration property'
description: |-
  Manages property.
---

# ibm_app_config_property

Provides a resource for `property`. This allows property to be created, updated and deleted. For more information, about App Configuration segment, see [properties](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-properties).

## Example Usage

```hcl
# Example 1
resource "ibm_app_config_property" "app_config_property" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  property_id = "property_id"
  type = "type"
  value = "value"
  description = "description"
  tags = "tag1,tag2"
}

# Example 2
resource "ibm_app_config_property" "app_config_property" {
  guid = "guid"
  environment_id = "environment_id"
  name = "name"
  property_id = "property_id"
  type = "type"
  value = "value"
  description = "description"
  tags = "tag1,tag2"
  collections {
    collection_id = "collection_id1"
  }
  # only use this deleted attribute during 
  # update of property
  collections {
    collection_id = "collection_id2"
    deleted = true
  }
  segment_rules {
    rules {
      segments = [ "segment_id1","segment_id2" ]
    }
    value = "value1"
    order = 1
  }
  segment_rules {
    rules {
      segments = [ "segment_id3","segment_id4" ]
    }
    value = "value2"
    order = 2
  }
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.
- `name` - (Required, string) Property name.
- `property_id` - (Required, string) Property id.
- `type` - (Required, string) Type of the Property (BOOLEAN, STRING, NUMERIC).
- `value` - (Required, TypeMap) Value of the Property. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `description` - (Optional, string) Property description.
- `tags` - (Optional, string) Tags associated with the property.
- `format` - (Optional, string) Format of the property (TEXT, JSON, YAML) and this is a required attribute when TYPE STRING is used, not required for BOOLEAN and NUMERIC types.
- `segment_rules` - (Optional, List) Specify the targeting rules that is used to set different property values for different segments.
    - `rules` - (Required, []interface{}) Rules array.
    - `value` - (Required, TypeMap) Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
    - `order` - (Required, int) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.
    **Index 1 based numbering is used for order**.
- `collections` - (Optional, List) List of collection id representing the collections that are associated with the specified property.
    - `collection_id` - (Required, string) Collection id.
    - `deleted` - (Optional, Boolean) Delete resource association with collection.
  Note:- Only to be used during update operation

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the app_config_property.
- `segment_exists` - Denotes if the targeting rules are specified for the property.
- `created_time` - Creation time of the property.
- `updated_time` - Last modified time of the property data.
- `evaluation_time` - The last occurrence of the property value evaluation.
- `href` - Property URL.

## Import

The `ibm_app_config_property` resource can be imported by using `guid` of the App Configuration instance, `environmentId` and `propertyId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_property.sample  <guid/environmentId/propertyId>

```

**Example**

```
terraform import ibm_app_config_property.sample 272111153-c118-4116-8116-b811fbc31132/dev/sample_property
```
