# Examples for Configuration Aggregator

These examples illustrate how to use the resources and data sources associated with Configuration Aggregator.

The following resources are supported:
* ibm_config_aggregator_settings

The following data sources are supported:
* ibm_config_aggregator_configurations
* ibm_config_aggregator_settings
* ibm_config_aggregator_resource_collection_status

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Configuration Aggregator resources

### Resource: ibm_config_aggregator_settings

```hcl
resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
  resource_collection_enabled = var.config_aggregator_settings_resource_collection_enabled
  trusted_profile_id = var.config_aggregator_settings_trusted_profile_id
  regions = var.config_aggregator_settings_regions
  additional_scope = var.config_aggregator_settings_additional_scope
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| resource_collection_enabled | The field denoting if the resource collection is enabled. | `bool` | false |
| trusted_profile_id | The trusted profile id that provides Reader access to the App Configuration instance to collect resource metadata. | `string` | false |
| regions | The list of regions across which the resource collection is enabled. | `list(string)` | false |
| additional_scope | The additional scope that enables resource collection for Enterprise acccounts. | `list()` | false |

## Configuration Aggregator data sources

### Data source: ibm_config_aggregator_configurations

```hcl
data "ibm_config_aggregator_configurations" "config_aggregator_configurations_instance" {
  config_type = var.config_aggregator_configurations_config_type
  service_name = var.config_aggregator_configurations_service_name
  resource_group_id = var.config_aggregator_configurations_resource_group_id
  location = var.config_aggregator_configurations_location
  resource_crn = var.config_aggregator_configurations_resource_crn
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| config_type | The type of resource configuration that are to be retrieved. | `string` | false |
| service_name | The name of the IBM Cloud service for which resources are to be retrieved. | `string` | false |
| resource_group_id | The resource group id of the resources. | `string` | false |
| location | The location or region in which the resources are created. | `string` | false |
| resource_crn | The crn of the resource. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| prev | The reference to the previous page of entries. |
| configs | Array of resource configurations. |

### Data source: ibm_config_aggregator_settings

```hcl
data "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| resource_collection_enabled | The field to check if the resource collection is enabled. |
| trusted_profile_id | The trusted profile ID that provides access to App Configuration instance to retrieve resource metadata. |
| last_updated | The last time the settings was last updated. |
| regions | Regions for which the resource collection is enabled. |
| additional_scope | The additional scope that enables resource collection for Enterprise acccounts. |

### Data source: ibm_config_aggregator_resource_collection_status

```hcl
data "ibm_config_aggregator_resource_collection_status" "config_aggregator_resource_collection_status_instance" {
}
```

#### Outputs

| Name | Description |
|------|-------------|
| last_config_refresh_time | The timestamp at which the configuration was last refreshed. |
| status | Status of the resource collection. |

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
