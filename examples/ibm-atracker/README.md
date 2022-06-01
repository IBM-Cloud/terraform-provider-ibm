# Example for AtrackerV1

This example illustrates how to use the AtrackerV1

These types of resources are supported:

* Activity Tracker Target
* Activity Tracker Route
* Activity Tracker Settings

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## AtrackerV1 resources

atracker_target resource:

```hcl
resource "atracker_target" "atracker_target_instance" {
  name = var.atracker_target_name
  target_type = var.atracker_target_target_type
  cos_endpoint = var.atracker_target_cos_endpoint
}
```
atracker_route resource:

```hcl
resource "atracker_route" "atracker_route_instance" {
  name = var.atracker_route_name
  receive_global_events = var.atracker_route_receive_global_events
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

## AtrackerV1 Data sources

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
atracker_endpoints data source:

```hcl
data "atracker_endpoints" "atracker_endpoints_instance" {
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
| name | The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. | `string` | true |
| receive_global_events | Indicates whether or not all global events should be forwarded to this region. | `bool` | true |
| rules | Routing rules that will be evaluated in their order of the array. | `list()` | true |
| name | The name of the target resource. | `string` | false |
| name | The name of the route. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| atracker_target | atracker_target object |
| atracker_route | atracker_route object |
| atracker_targets | atracker_targets object |
| atracker_routes | atracker_routes object |
| atracker_endpoints | atracker_endpoints object |
| atracker_settings | atracker_settings object |
