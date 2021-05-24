---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag'
description: |-
  Get information about feature flag
---

# ibm_app_config_feature

Provides a read-only data source for `feature flag`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_feature" "app_config_feature" {
  guid = "guid"
  feature_id = "feature_id"
  includes = "includes"
  environment_id = "environment_id"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `feature_id` - (Required, string) Feature Id.
- `environment_id` - (Required, string) Environment Id.
- `includes` - (Optional, string) Include the associated collections in the response.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the feature flag resource.
- `name` - Feature name.
- `description` - Feature description.
- `type` - Type of the feature (BOOLEAN, STRING, NUMERIC).
- `enabled_value` - Value of the feature when it is enabled. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `disabled_value` - Value of the feature when it is disabled. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `enabled` - The state of the feature flag.
- `tags` - Tags associated with the feature.
- `segment_rules` - Specify the targeting rules that is used to set different feature flag values for different segments. Nested `segment_rules` blocks have the following structure:

  - `rules` - The list of targetted segments. Nested `rules` blocks have the following structure:
    - `segments` - List of segment ids that are used for targeting using the rule.
  - `value` - Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
  - `order` - Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.

- `segment_exists` - Denotes if the targeting rules are specified for the feature flag.
- `collections` - List of collection id representing the collections that are associated with the specified feature flag. Nested `collections` blocks have the following structure:

  - `collection_id` - Collection id.

  - `name` - Name of the collection.

- `created_time` - Creation time of the feature flag.
- `updated_time` - Last modified time of the feature flag data.
- `href` - Feature flag URL.
