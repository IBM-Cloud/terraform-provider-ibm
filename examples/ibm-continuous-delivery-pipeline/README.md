# Example for ContinuousDeliveryPipelineV2

This example illustrates how to use the ContinuousDeliveryPipelineV2

These types of resources are supported:

* tekton_pipeline_definition
* tekton_pipeline_trigger_property
* tekton_pipeline_property
* tekton_pipeline_trigger
* tekton_pipeline

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ContinuousDeliveryPipelineV2 resources

tekton_pipeline_definition resource:

```hcl
resource "tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  create_tekton_pipeline_definition_request = var.tekton_pipeline_definition_create_tekton_pipeline_definition_request
}
```
tekton_pipeline_trigger_property resource:

```hcl
resource "tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  create_tekton_pipeline_trigger_properties_request = var.tekton_pipeline_trigger_property_create_tekton_pipeline_trigger_properties_request
}
```
tekton_pipeline_property resource:

```hcl
resource "tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  create_tekton_pipeline_properties_request = var.tekton_pipeline_property_create_tekton_pipeline_properties_request
}
```
tekton_pipeline_trigger resource:

```hcl
resource "tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  create_tekton_pipeline_trigger_request = var.tekton_pipeline_trigger_create_tekton_pipeline_trigger_request
}
```
tekton_pipeline resource:

```hcl
resource "tekton_pipeline" "tekton_pipeline_instance" {
  integration_instance_id = var.tekton_pipeline_integration_instance_id
  worker = var.tekton_pipeline_worker
}
```

## ContinuousDeliveryPipelineV2 Data sources

tekton_pipeline_definition data source:

```hcl
data "tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  definition_id = var.tekton_pipeline_definition_definition_id
}
```
tekton_pipeline_trigger_property data source:

```hcl
data "tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  property_name = var.tekton_pipeline_trigger_property_property_name
}
```
tekton_pipeline_property data source:

```hcl
data "tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  property_name = var.tekton_pipeline_property_property_name
}
```
tekton_pipeline_workers data source:

```hcl
data "tekton_pipeline_workers" "tekton_pipeline_workers_instance" {
  pipeline_id = var.tekton_pipeline_workers_pipeline_id
}
```
tekton_pipeline_trigger data source:

```hcl
data "tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_trigger_id
}
```
tekton_pipeline data source:

```hcl
data "tekton_pipeline" "tekton_pipeline_instance" {
  id = var.tekton_pipeline_id
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
| pipeline_id | The tekton pipeline ID. | `string` | true |
| create_tekton_pipeline_definition_request |  | `` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| create_tekton_pipeline_trigger_properties_request |  | `` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| create_tekton_pipeline_properties_request |  | `` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| create_tekton_pipeline_trigger_request |  | `` | false |
| integration_instance_id | UUID. | `string` | false |
| worker | Worker object with just worker ID. | `` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| definition_id | The definition ID. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| property_name | The property's name. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| property_name | The property's name. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| id | ID of current instance. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| tekton_pipeline_definition | tekton_pipeline_definition object |
| tekton_pipeline_trigger_property | tekton_pipeline_trigger_property object |
| tekton_pipeline_property | tekton_pipeline_property object |
| tekton_pipeline_trigger | tekton_pipeline_trigger object |
| tekton_pipeline | tekton_pipeline object |
| tekton_pipeline_definition | tekton_pipeline_definition object |
| tekton_pipeline_trigger_property | tekton_pipeline_trigger_property object |
| tekton_pipeline_property | tekton_pipeline_property object |
| tekton_pipeline_workers | tekton_pipeline_workers object |
| tekton_pipeline_trigger | tekton_pipeline_trigger object |
| tekton_pipeline | tekton_pipeline object |
