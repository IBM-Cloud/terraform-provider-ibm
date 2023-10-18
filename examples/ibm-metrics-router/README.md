# Example for IBM Cloud Metrics Routing

This example illustrates how to use the MetricsRouterV3

These types of resources are supported:

* metrics_router_target
* metrics_router_route
* metrics_router_settings

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## MetricsRouterV3 resources

metrics_router_target resource:

```hcl
resource "metrics_router_target" "metrics_router_target_instance" {
  name = var.metrics_router_target_name
  destination_crn = var.metrics_router_target_destination_crn
  region = var.metrics_router_target_region
}
```
metrics_router_route resource:

```hcl
resource "metrics_router_route" "metrics_router_route_instance" {
  name = var.metrics_router_route_name
  rules = var.metrics_router_route_rules
}
```
metrics_router_settings resource:

```hcl
resource "metrics_router_settings" "metrics_router_settings_instance" {
  default_targets = var.metrics_router_settings_default_targets
  permitted_target_regions = var.metrics_router_settings_permitted_target_regions
  primary_metadata_region = var.metrics_router_settings_primary_metadata_region
  backup_metadata_region = var.metrics_router_settings_backup_metadata_region
  private_api_endpoint_only = var.metrics_router_settings_private_api_endpoint_only
}
```

## MetricsRouterV3 Data sources

metrics_router_targets data source:

```hcl
data "metrics_router_targets" "metrics_router_targets_instance" {
  name = var.metrics_router_targets_name
}
```
metrics_router_routes data source:

```hcl
data "metrics_router_routes" "metrics_router_routes_instance" {
  name = var.metrics_router_routes_name
}
```

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

## Inputs

### Target

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names. | `string` | true |
| destination_crn | The CRN of the destination resource. Ensure you have a service authorization between IBM Cloud Metrics Routing and your Cloud resource. Read [S2S authorization](https://cloud.ibm.com/docs/metrics-router?topic=metrics-router-target-monitoring&interface=ui#target-monitoring-ui) for details.| `string` | true |
| region | Include this optional field if you want to create a target in a different region other than the one you are connected. | `string` | false |

### Route

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names. | `string` | true |
| rules | Routing rules that will be evaluated in their order of the array. | `list()` | true |

### Settings

| Name | Description | Type | Required |
|------|-------------|------|---------|
| default_targets | A list of default target references. | `list()` | false |
| permitted_target_regions | If present then only these regions may be used to define a target. | `list(string)` | false |
| primary_metadata_region | To store all your meta data in a single region. For new accounts, all target / route creation will fail until primary_metadata_region is set. | `string` | false |
| backup_metadata_region | To backup all your meta data in a different region. | `string` | false |
| private_api_endpoint_only | If you set this true then you cannot access api through public network. | `bool` | false |

### Data Source For Target

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of the target resource. | `string` | false |

### Data Source For Route

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of the route. | `string` | false |

## Outputs


| Name | Description |
|------|-------------|
| metrics_router_target | metrics_router_target object |
| metrics_router_route | metrics_router_route object |
| metrics_router_settings | metrics_router_settings object |
| metrics_router_targets | metrics_router_targets object |
| metrics_router_routes | metrics_router_routes object |
