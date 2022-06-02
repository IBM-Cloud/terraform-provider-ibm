# Example for CdTektonPipelineV2

This example illustrates how to use the CdTektonPipelineV2

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


## CdTektonPipelineV2 resources

tekton_pipeline_definition resource:

```hcl
resource "tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  scm_source = var.tekton_pipeline_definition_scm_source
}
```
tekton_pipeline_trigger_property resource:

```hcl
resource "tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  name = var.tekton_pipeline_trigger_property_name
  value = var.tekton_pipeline_trigger_property_value
  enum = var.tekton_pipeline_trigger_property_enum
  default = var.tekton_pipeline_trigger_property_default
  type = var.tekton_pipeline_trigger_property_type
  path = var.tekton_pipeline_trigger_property_path
}
```
tekton_pipeline_property resource:

```hcl
resource "tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  name = var.tekton_pipeline_property_name
  value = var.tekton_pipeline_property_value
  enum = var.tekton_pipeline_property_enum
  default = var.tekton_pipeline_property_default
  type = var.tekton_pipeline_property_type
  path = var.tekton_pipeline_property_path
}
```
tekton_pipeline_trigger resource:

```hcl
resource "tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  trigger = var.tekton_pipeline_trigger_trigger
}
```
tekton_pipeline resource:

```hcl
resource "tekton_pipeline" "tekton_pipeline_instance" {
  worker = var.tekton_pipeline_worker
}
```

## CdTektonPipelineV2 Data sources

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
| scm_source | Scm source for tekton pipeline defintion. | `` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| name | Property name. | `string` | false |
| value | String format property value. | `string` | false |
| enum | Options for SINGLE_SELECT property type. | `list(string)` | false |
| default | Default option for SINGLE_SELECT property type. | `string` | false |
| type | Property type. | `string` | false |
| path | property path for INTEGRATION type properties. | `string` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| name | Property name. | `string` | false |
| value | String format property value. | `string` | false |
| enum | Options for SINGLE_SELECT property type. | `list(string)` | false |
| default | Default option for SINGLE_SELECT property type. | `string` | false |
| type | Property type. | `string` | false |
| path | property path for INTEGRATION type properties. | `string` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| trigger | Tekton pipeline trigger object. | `` | false |
| worker | Worker object with worker ID only. | `` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| definition_id | The definition ID. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| property_name | The property's name. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| property_name | The property's name. | `string` | true |
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
| tekton_pipeline_trigger | tekton_pipeline_trigger object |
| tekton_pipeline | tekton_pipeline object |
