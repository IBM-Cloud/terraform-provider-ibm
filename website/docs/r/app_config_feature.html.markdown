---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag'
description: |-
  Manages feature flag.
---

# ibm_app_config_feature

Create, update, or delete an environment by using IBM Cloud™ App Configuration feature flag. For more information, about App Configuration feature flag, see [Feature flags](https://cloud.ibm.com/docs/app-configuration?topic=app-configuration-ac-feature-flags).

## Example usage

```terraform
resource "ibm_app_config_feature" "app_config_feature" {
  guid = "guid"
  name = "name"
  type = "type"
  tags = "tags"
  feature_id = "feature_id"
  enabled_value = "enabled_value"
  environment_id = "environment_id"
  disabled_value = "disabled_value"
  rollout_percentage = "rollout_percentage"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `guid` - (Required, String) The GUID of the App Configuration service. Fetch GUID from the service instance credentials section of the dashboard.
- `environment_id` - (Required, String) The environment ID.
- `name` - (Required, String) The feature name.
- `feature_id` - (Required, String) The feature ID.
- `type` - (Required, String) The feature type. Supported values are **BOOLEAN**, **STRING**, or **NUMERIC**.
- `enabled_value` - (Required, String) The value of the feature when it is enabled. The value can be **BOOLEAN**, **STRING**, or **NUMERIC** value as per the `type` attribute.
- `disabled_value` - (Required, String) The value of the feature when it is disabled. The value can be **BOOLEAN**, **STRING**, or **NUMERIC** value as per the `type` attribute.
- `description` - (Optional, String) The feature description.
- `tags` - (Optional, String) Tags associated with the feature.
- `rollout_percentage` - (String) Rollout percentage of the feature.
- `segment_rules` - (Optional, List) Specify the targeting rules that is used to set different feature flag values for different segments.
  - `rules` - (Required, []interface{}) The rules array.
    - `segments` - (Required, Array of Strings) The list of segment IDs that are used for targeting using the rule.
  - `value` - (Required, String) The value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
  - `order` - (Required, Integer) The order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.
  - `rollout_percentage` - (String) Rollout percentage for the segment rule.
- `collections` - (Optional, List) The list of collection ID representing the collections that are associated with the specified feature flag.
  - `collection_id` - (Required, String) Collection ID.

## Attribute reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the Feature flag resource.
- `enabled` - (String) The state of the feature flag.
- `segment_exists` - (String) Denotes if the targeting rules are specified for the feature flag.
- `created_time` - (Timestamp) The creation time of the feature flag.
- `updated_time` - (Timestamp) The last modified time of the feature flag data.
- `href` - (String) The feature flag URL.

## Import

The `ibm_app_config_feature` resource can be imported by using `guid` of the App Configuration instance, `environmentId` and `featureId`. Get the `guid` from the service instance credentials section of the dashboard.

**Syntax**

```
terraform import ibm_app_config_feature.sample  <guid/environmentId/featureId>

```

**Example**

```
terraform import ibm_app_config_feature.sample 272111153-c118-4116-8116-b811fbc31132/dev/sample_feature
```
