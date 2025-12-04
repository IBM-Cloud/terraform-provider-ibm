# Examples for Activity Tracker API Version 2

These examples illustrate how to use the resources and data sources associated with Activity Tracker API Version 2.

The following resources are supported:
* ibm_atracker_target
* ibm_atracker_route
* ibm_atracker_settings

The following data sources are supported:
* ibm_atracker_targets
* ibm_atracker_routes

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Activity Tracker API Version 2 resources

### Resource: ibm_atracker_target

```hcl
resource "ibm_atracker_target" "atracker_target_instance" {
  name = var.atracker_target_name
  target_type = var.atracker_target_target_type
  region = var.atracker_target_region
  cos_endpoint = var.atracker_target_cos_endpoint
  eventstreams_endpoint = var.atracker_target_eventstreams_endpoint
  cloudlogs_endpoint = var.atracker_target_cloudlogs_endpoint
  managed_by = var.atracker_target_managed_by
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name of the target resource. | `string` | true |
| target_type | The type of the target. | `string` | true |
| region | Included this optional field if you used it to create a target in a different region other than the one you are connected. | `string` | false |
| cos_endpoint | Property values for a Cloud Object Storage Endpoint in responses. | `` | false |
| eventstreams_endpoint | Property values for the Event Streams Endpoint in responses. | `` | false |
| cloudlogs_endpoint | Property values for the IBM Cloud Logs endpoint in responses. | `` | false |
| managed_by | Identifies who manages the target. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| crn | The crn of the target resource. |
| write_status | The status of the write attempt to the target with the provided endpoint parameters. |
| created_at | The timestamp of the target creation time. |
| updated_at | The timestamp of the target last updated time. |
| message | An optional message containing information about the target. |
| api_version | The API version of the target. |

### Resource: ibm_atracker_route

```hcl
resource "ibm_atracker_route" "atracker_route_instance" {
  name = var.atracker_route_name
  rules = var.atracker_route_rules
  managed_by = var.atracker_route_managed_by
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
| version | The version of the route. |
| created_at | The timestamp of the route creation time. |
| updated_at | The timestamp of the route last updated time. |
| api_version | The API version of the route. |
| message | An optional message containing information about the route. |

### Resource: ibm_atracker_settings

```hcl
resource "ibm_atracker_settings" "atracker_settings_instance" {
  default_targets = var.atracker_settings_default_targets
  permitted_target_regions = var.atracker_settings_permitted_target_regions
  metadata_region_primary = var.atracker_settings_metadata_region_primary
  metadata_region_backup = var.atracker_settings_metadata_region_backup
  private_api_endpoint_only = var.atracker_settings_private_api_endpoint_only
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| default_targets | The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event. Enterprise-managed targets are not supported. | `list(string)` | false |
| permitted_target_regions | If present then only these regions may be used to define a target. | `list(string)` | false |
| metadata_region_primary | To store all your meta data in a single region. | `string` | true |
| metadata_region_backup | To store all your meta data in a backup region. | `string` | false |
| private_api_endpoint_only | If you set this true then you cannot access api through public network. | `bool` | true |

#### Outputs

| Name | Description |
|------|-------------|
| api_version | API version used for configuring IBM Cloud Activity Tracker Event Routing resources in the account. |
| message | An optional message containing information about the audit log locations. |

## Activity Tracker API Version 2 data sources

### Data source: ibm_atracker_targets

```hcl
data "ibm_atracker_targets" "atracker_targets_instance" {
  region = var.atracker_targets_region
  name = var.atracker_targets_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| region | Limit the query to the specified region. | `string` | false |
| name | The name of the target resource. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| targets | A list of target resources. |

### Data source: ibm_atracker_routes

```hcl
data "ibm_atracker_routes" "atracker_routes_instance" {
  name = var.atracker_routes_name
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
