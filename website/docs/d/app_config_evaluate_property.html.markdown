---
layout: "ibm"
page_title: "IBM : ibm_app_config_evaluate_property"
description: |-
  Get information about AppConfigurationPropertyEvaluation
subcategory: "App Configuration Evaluation"
---

# ibm_app_config_evaluate_property

Provides a read-only data source for property evaluation. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_app_config_evaluate_property" "evaluate_property" {
  guid              = "0b5571f7-21e6-42b7-91c5-3f5ac9793a46"
  environment_id    = "dev"
  collection_id     = "car-rentals"
  property_id        = "users-location"
  entity_id         = "john_doe"
  entity_attributes = {
    "city" : "Bangalore",
    "radius" : 60,
  }
}
```

**provider.tf**
Please make sure to target right region in the provider block.

```hcl
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.region
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `region` - (Required, String) The region of the App Configuration instance.
* `guid` - (Required, String) The guid or instance id of the App Configuration instance.
* `environment_id` - (Required, String) Id of the environment created in App Configuration instance under the Environments section.
* `collection_id` - (Required, String) Id of the collection created in App Configuration instance under the Collections section.
* `property_id` - (Required, String) Property id required to be evaluated.
* `entity_id` - (Required, String) Id of the Entity. This will be a string identifier related to the Entity against which the property is evaluated. For example, an entity might be an instance of an app that runs on a mobile device, a microservice that runs on the cloud, or a component of infrastructure that runs that microservice. For any entity to interact with App Configuration, it must provide a unique entity ID."
* `entity_attributes` - (Optional, Map) Key value pair consisting of the attribute name and their values that defines the specified entity. This is an optional parameter if the property is not configured with any targeting definition. If the targeting is configured, then entityAttributes should be provided for the rule evaluation. An attribute is a parameter that is used to define a segment. The SDK uses the attribute values to determine if the specified entity satisfies the targeting rules, and returns the appropriate property value.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the AppConfigurationPropertyEvaluation.
* `result_boolean` - (Boolean) Contains the evaluated value of the BOOLEAN type properties.
* `result_string` - (String) Contains the evaluated value of the STRING type properties.
* `result_numeric` - (Number) Contains the evaluated value of the NUMERIC type properties.
