# Example for MetricsRouterV3

This example illustrates how to use the MetricsRouterV3

These types of resources are supported:

* Metrics Router Target
* Metrics Router Route
* Metrics Router Settings

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
  metadata_region_primary = var.metrics_router_settings_metadata_region_primary
  private_api_endpoint_only = var.metrics_router_settings_private_api_endpoint_only
  default_targets = var.metrics_router_settings_default_targets
  permitted_target_regions = var.metrics_router_settings_permitted_target_regions
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

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names. | `string` | true |
| destination_crn | The CRN of a destination service instance or resource. | `string` | true |
| region | Include this optional field if you want to create a target in a different region other than the one you are connected. | `string` | false |
| name | The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names. | `string` | true |
| rules | Routing rules that will be evaluated in their order of the array. | `list()` | true |
| metadata_region_primary | To store all your meta data in a single region. | `string` | true |
| private_api_endpoint_only | If you set this true then you cannot access api through public network. | `bool` | true |
| default_targets | The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event. | `list(string)` | false |
| permitted_target_regions | If present then only these regions may be used to define a target. | `list(string)` | false |
| name | The name of the target resource. | `string` | false |
| name | The name of the route. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| metrics_router_target | metrics_router_target object |
| metrics_router_route | metrics_router_route object |
| metrics_router_settings | metrics_router_settings object |
| metrics_router_targets | metrics_router_targets object |
| metrics_router_routes | metrics_router_routes object |
