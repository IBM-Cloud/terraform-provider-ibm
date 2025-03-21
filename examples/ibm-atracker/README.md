# Example for AtrackerV2

This example illustrates how to use the AtrackerV2

The following types of resources are supported:

* Activity Tracker Event Routing Target
* Activity Tracker Event Routing Route
* Activity Tracker Event Routing Settings

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## AtrackerV2 resources

atracker_target resource:

```hcl
# Deprecation: logdna_endpoint is no longer in use and will be removed in the next major version of the provider.
resource "atracker_target" "atracker_target_instance" {
  name = var.atracker_target_name
  target_type = var.atracker_target_target_type
  region = var.atracker_target_region
  cos_endpoint = var.atracker_target_cos_endpoint
  logdna_endpoint = var.atracker_target_logdna_endpoint
  eventstreams_endpoint = var.atracker_target_eventstreams_endpoint
  cloudlogs_endpoint = var.atracker_target_cloudlogs_instance
}
```
atracker_route resource:

```hcl
resource "atracker_route" "atracker_route_instance" {
  name = var.atracker_route_name
  rules = var.atracker_route_rules
}
```

atracker_settings resource:

```hcl
resource "atracker_settings" "atracker_settings_instance" {
  metadata_region_primary = var.atracker_settings_metadata_region_primary
  private_api_endpoint_only = var.atracker_settings_private_api_endpoint_only
  default_targets = var.atracker_settings_default_targets
  permitted_target_regions = var.atracker_settings_permitted_target_regions
}
```

## AtrackerV2 Data sources

atracker_targets data source:

```hcl
data "atracker_targets" "atracker_targets_instance" {
  name = var.atracker_targets_name
}
```
atracker_routes data source:

```hcl
data "atracker_routes" "atracker_routes_instance" {
  name = var.atracker_routes_name
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
| name | The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`. | `string` | true |
| target_type | The type of the target. | `string` | true |
| cos_endpoint | Property values for a Cloud Object Storage Endpoint. | `` | true |
| eventstreams_endpoint | Property values for the Event Streams Endpoint in responses. | `` | false |
| logdna_endpoint | Property values for a LogDNA Endpoint. Remove this attribute as it is no longer in use and it will be removed in the next major version of the provider.| `` | false |
| cloudlogs_endpoint | Property values for the IBM Cloud Logs Endpoint in responses. | `` | false |
| name | The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. | `string` | true |
| rules | The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped. | `list()` | true |
| default_targets | The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event. | `list(string)` | false |
| permitted_target_regions | If present then only these regions may be used to define a target. | `list(string)` | false |
| metadata_region_primary | To store all your meta data in a single region. | `string` | true |
| metadata_region_backup | To store all your meta data in a backup region. | `string` | false |
| private_api_endpoint_only | If you set this true then you cannot access api through public network. | `bool` | true |
| region | Limit the query to the specified region. | `string` | false |
| name | The name of the target resource. | `string` | false |
| name | The name of the route. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| atracker_target | atracker_target object |
| atracker_route | atracker_route object |
| atracker_targets | atracker_targets object |
| atracker_routes | atracker_routes object |
| atracker_settings | atracker_settings object |
