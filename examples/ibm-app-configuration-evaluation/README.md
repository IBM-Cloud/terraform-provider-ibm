# Example for AppConfigurationEvaluation

This example illustrates how to use the AppConfigurationEvaluation for feature flags & properties.

The following types of data sources are supported:

* App Configuration evaluate feature flag.
* App Configuration evaluate property.

## Usage

To run this example, execute the following commands:

```bash
terraform init
terraform plan
terraform apply
```

Run `terraform destroy` when you don't need these data sources.

## AppConfigurationEvaluation data sources

ibm_app_config_evaluate_feature_flag data source:

```hcl
data "ibm_app_config_evaluate_feature_flag" "evaluate_feature_flag" {
  guid              = var.app_config_guid
  environment_id    = var.app_config_environment_id
  collection_id     = var.app_config_collection_id
  feature_id        = var.app_config_feature_id
  entity_id         = var.app_config_entity_id
  entity_attributes = var.app_config_entity_attributes
}
```

ibm_app_config_evaluate_property data source:

```hcl
data "ibm_app_config_evaluate_property" "evaluate_property" {
  guid              = var.app_config_guid
  environment_id    = var.app_config_environment_id
  collection_id     = var.app_config_collection_id
  property_id       = var.app_config_property_id
  entity_id         = var.app_config_entity_id
  entity_attributes = var.app_config_entity_attributes
}
```

## Requirements

| Name      | Version |
| --------- | ------- |
| terraform | ~> 0.12 |

## Providers

| Name | Version |
| ---- | ------- |
| ibm  | 1.60.0  |

## Inputs

| Name               | Description                        | Type     | Required |
| ------------------ | ---------------------------------- | -------- | -------- |
| ibmcloud\_api\_key | IBM Cloud API key                  | `string` | true     |
| region             | IBM Cloud region name.             | `string` | true     |
| guid               | App Configuration instance Id.     | `string` | true     |
| collection_id      | App Configuration Collection Id.   | `string` | true     |
| environment_id     | App Configuration Environment Id.  | `string` | true     |
| feature_id         | App Configuration Feature flag Id. | `string` | true     |
| property_id        | App Configuration Property Id.     | `string` | true     |
| entity_id          | Entity Id.                         | `string` | true     |
| entity_attributes  | Entity attributes.                 | `object` | false    |

## Outputs
| Name           | Description                                                   |
| -------------- | ------------------------------------------------------------- |
| result_boolean | Evaluated value of the BOOLEAN type feature flag or property. |
| result_string  | Evaluated value of the STRING type feature flag or property.  |
| result_numeric | Evaluated value of the NUMERIC type feature flag or property. |