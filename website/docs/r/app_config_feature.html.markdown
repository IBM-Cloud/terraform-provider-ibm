---
subcategory: 'App Configuration'
layout: 'ibm'
page_title: 'IBM : App Configuration feature flag'
description: |-
  Manages feature flag.
---

# ibm_app_config_feature

Provides a resource for `feature flag`. This allows Feature flag to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_app_config_feature" "app_config_feature" {
  guid = "guid"
  name = "name"
  type = "type"
  tags = "tags"
  feature_id = "feature_id"
  enabled_value = "enabled_value"
  environment_id = "environment_id"
  disabled_value = "disabled_value"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) guid of the App Configuration service. Get it from the service instance credentials section of the dashboard.
- `environment_id` - (Required, string) Environment Id.
- `name` - (Required, string) Feature name.
- `feature_id` - (Required, string) Feature id.
- `type` - (Required, string) Type of the feature (BOOLEAN, STRING, NUMERIC).
- `enabled_value` - (Required, string) Value of the feature when it is enabled. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `disabled_value` - (Required, string) Value of the feature when it is disabled. The value can be Boolean, String or a Numeric value as per the `type` attribute.
- `description` - (Optional, string) Feature description.
- `tags` - (Optional, string) Tags associated with the feature.
- `segment_rules` - (Optional, List) Specify the targeting rules that is used to set different feature flag values for different segments.
  - `rules` - (Required, []interface{}) Rules array.
    - `segments` - (Required, Array of Strings)List of segment ids that are used for targeting using the rule.
  - `value` - (Required, string) Value to be used for evaluation for this rule. The value can be Boolean, String or a Numeric value as per the `type` attribute.
  - `order` - (Required, int) Order of the rule, used during evaluation. The evaluation is performed in the order defined and the value associated with the first matching rule is used for evaluation.
- `collections` - (Optional, List) List of collection id representing the collections that are associated with the specified feature flag.
  - `collection_id` - (Required, string) Collection id.

## Attribute Reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the Feature flag resource.
- `enabled` - The state of the feature flag.
- `segment_exists` - Denotes if the targeting rules are specified for the feature flag.
- `created_time` - Creation time of the feature flag.
- `updated_time` - Last modified time of the feature flag data.
- `href` - Feature flag URL.

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
