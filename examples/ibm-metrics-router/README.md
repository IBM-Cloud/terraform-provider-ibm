# Examples for Metrics Routing API Version 3

These examples illustrate how to use the resources and data sources associated with Metrics Routing API Version 3.

The following resources are supported:
* ibm_metrics_router_target
* ibm_metrics_router_route
* ibm_metrics_router_settings

The following data sources are supported:
* ibm_metrics_router_targets
* ibm_metrics_router_routes

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Metrics Routing API Version 3 resources

### Resource: ibm_metrics_router_target

```hcl
resource "ibm_metrics_router_target" "metrics_router_target_instance" {
  name = var.metrics_router_target_name
  destination_crn = var.metrics_router_target_destination_crn
  region = var.metrics_router_target_region
  managed_by = var.metrics_router_target_managed_by
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the target resource. | `string` | true |
| destination_crn | Cloud Resource Name (CRN) of the destination resource. Ensure you have a service authorization between IBM Cloud Metrics Routing and your Cloud resource. See [service-to-service authorization](https://cloud.ibm.com/docs/metrics-router?topic=metrics-router-target-monitoring&interface=ui#target-monitoring-ui) for details. | `string` | true |
| region | Include this optional field if you used it to create a target in a different region other than the one you are connected. | `string` | false |
| managed_by | Present when the target is enterprise-managed (`managed_by: enterprise`). For account-managed targets this field is omitted. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| crn | The crn of the target resource. |
| target_type | The type of the target. |
| created_at | The timestamp of the target creation time. |
| updated_at | The timestamp of the target last updated time. |

### Resource: ibm_metrics_router_route

```hcl
resource "ibm_metrics_router_route" "metrics_router_route_instance" {
  name = var.metrics_router_route_name
  rules = var.metrics_router_route_rules
  managed_by = var.metrics_router_route_managed_by
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the route. | `string` | true |
| rules | The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped. | `list()` | true |
| managed_by | Present when the route is enterprise-managed (`managed_by: enterprise`). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| crn | The crn of the route resource. |
| created_at | The timestamp of the route creation time. |
| updated_at | The timestamp of the route last updated time. |

### Resource: ibm_metrics_router_settings

```hcl
resource "ibm_metrics_router_settings" "metrics_router_settings_instance" {
  default_targets = var.metrics_router_settings_default_targets
  permitted_target_regions = var.metrics_router_settings_permitted_target_regions
  primary_metadata_region = var.metrics_router_settings_primary_metadata_region
  backup_metadata_region = var.metrics_router_settings_backup_metadata_region
  private_api_endpoint_only = var.metrics_router_settings_private_api_endpoint_only
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| default_targets | A list of default target references. Enterprise-managed targets are not supported. | `list()` | false |
| permitted_target_regions | If present then only these regions may be used to define a target. | `list(string)` | false |
| primary_metadata_region | To store all your meta data in a single region. | `string` | false |
| backup_metadata_region | To backup all your meta data in a different region. | `string` | false |
| private_api_endpoint_only | If you set this true then you cannot access api through public network. | `bool` | false |

## Metrics Routing API Version 3 data sources

### Data source: ibm_metrics_router_targets

```hcl
data "ibm_metrics_router_targets" "metrics_router_targets_instance" {
  name = var.metrics_router_targets_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of the target resource. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| targets | A list of target resources. |

### Data source: ibm_metrics_router_routes

```hcl
data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
  name = var.metrics_router_routes_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The name of the route. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| routes | A list of route resources. |

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
