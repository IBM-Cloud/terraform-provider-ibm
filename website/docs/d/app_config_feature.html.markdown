---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag'
description: |-
  Get information about feature flag
---

# ibm_app_config_feature

Retrieve information about an existing IBM Cloud App Configuration feature. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about App Configuration features, see [App Configuration concepts](https://cloud.ibm.com//docs/app-configuration?topic=app-configuration-ac-overview).

## Example usage

```terraform
data "ibm_app_config_feature" "app_config_feature" {
  guid = "guid"
  feature_id = "feature_id"
  includes = "includes"
  environment_id = "environment_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `guid` - (Required, String) The GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `feature_id` - (Required, String) The feature ID.
- `environment_id` - (Required, String) The environment ID.
- `includes` - (Optional, String) Include the associated collections in the response.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the feature flag resource.
- `name` - (String) Feature name.
- `description` - (String) Feature description.
- `type` - (String) Type of the feature (BOOLEAN, STRING, NUMERIC).
- `enabled_value` - (String) Value of the feature when it is enabled. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `disabled_value` - (String) Value of the feature when it is disabled. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `enabled` - (String) The state of the feature flag.
- `tags` - (String) Tags associated with the feature.
- `rollout_percentage` - (String) Rollout percentage of the feature.
- `segment_rules` - (List) Specify the targeting rules that is used to set different feature flag values for different segments.

  Nested scheme for `segment_rules`:
  - `rules` - (List) The list of targeted segments.

    Nested scheme for `rules`:
    - `segments` - (String) List of segment ids that are used for targeting using the rule.
  - `value` - (String) Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
  - `order` - (String) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.
  - `rollout_percentage` - (String) Rollout percentage for the segment rule.
- `segment_exists` - (String) Denotes if the targeting rules are specified for the feature flag.
- `collections` - (List) List of collection ID representing the collections that are associated with the specified feature flag. 

  Nested scheme for `collections`:

  - `collection_id` - (String) The collection ID.
  - `name` - (String) The name of the collection.
- `created_time` - (Timestamp) The creation time of the feature flag.
- `updated_time` - (Timestamp) The last modified time of the feature flag data.
- `href` - (String) The feature flag URL.
