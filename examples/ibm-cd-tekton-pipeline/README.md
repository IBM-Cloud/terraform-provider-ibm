# Example for CdTektonPipelineV2

This example illustrates how to use the CdTektonPipelineV2

These types of resources are supported:

* cd_tekton_pipeline_definition
* cd_tekton_pipeline_trigger_property
* cd_tekton_pipeline_property
* cd_tekton_pipeline_trigger
* cd_tekton_pipeline

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CdTektonPipelineV2 resources

cd_tekton_pipeline_definition resource:

```hcl
resource "cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = var.cd_tekton_pipeline_definition_pipeline_id
  scm_source = var.cd_tekton_pipeline_definition_scm_source
  service_instance_id = var.cd_tekton_pipeline_definition_service_instance_id
}
```
cd_tekton_pipeline_trigger_property resource:

```hcl
resource "cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  name = var.cd_tekton_pipeline_trigger_property_name
  value = var.cd_tekton_pipeline_trigger_property_value
  enum = var.cd_tekton_pipeline_trigger_property_enum
  default = var.cd_tekton_pipeline_trigger_property_default
  type = var.cd_tekton_pipeline_trigger_property_type
  path = var.cd_tekton_pipeline_trigger_property_path
}
```
cd_tekton_pipeline_property resource:

```hcl
resource "cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  name = var.cd_tekton_pipeline_property_name
  value = var.cd_tekton_pipeline_property_value
  enum = var.cd_tekton_pipeline_property_enum
  default = var.cd_tekton_pipeline_property_default
  type = var.cd_tekton_pipeline_property_type
  path = var.cd_tekton_pipeline_property_path
}
```
cd_tekton_pipeline_trigger resource:

```hcl
resource "cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  trigger = var.cd_tekton_pipeline_trigger_trigger
}
```
cd_tekton_pipeline resource:

```hcl
resource "cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  worker = var.cd_tekton_pipeline_worker
}
```

## CdTektonPipelineV2 Data sources

cd_tekton_pipeline_definition data source:

```hcl
data "cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = var.cd_tekton_pipeline_definition_pipeline_id
  definition_id = var.cd_tekton_pipeline_definition_definition_id
}
```
cd_tekton_pipeline_trigger_property data source:

```hcl
data "cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  property_name = var.cd_tekton_pipeline_trigger_property_property_name
}
```
cd_tekton_pipeline_property data source:

```hcl
data "cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  property_name = var.cd_tekton_pipeline_property_property_name
}
```
cd_tekton_pipeline_trigger data source:

```hcl
data "cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_trigger_id
}
```
cd_tekton_pipeline data source:

```hcl
data "cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  id = var.cd_tekton_pipeline_id
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
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| scm_source | SCM source for Tekton pipeline definition. | `` | false |
| service_instance_id | ID of the SCM repository service instance. | `string` | false |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| name | Property name. | `string` | false |
| value | Property value. Can be empty and should be omitted for SINGLE_SELECT property type. | `string` | false |
| enum | Options for SINGLE_SELECT property type. Only needed for SINGLE_SELECT property type. | `list(string)` | false |
| default | Default option for SINGLE_SELECT property type. Only needed for SINGLE_SELECT property type. | `string` | false |
| type | Property type. | `string` | false |
| path | A dot notation path for INTEGRATION type properties to select a value from the tool integration. If left blank the full tool integration JSON will be selected. | `string` | false |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| name | Property name. | `string` | false |
| value | Property value. | `string` | false |
| enum | Options for SINGLE_SELECT property type. Only needed when using SINGLE_SELECT property type. | `list(string)` | false |
| default | Default option for SINGLE_SELECT property type. Only needed when using SINGLE_SELECT property type. | `string` | false |
| type | Property type. | `string` | false |
| path | A dot notation path for INTEGRATION type properties to select a value from the tool integration. | `string` | false |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| trigger | Tekton pipeline trigger. | `` | false |
| worker | Worker object containing worker ID only. If omitted the IBM Managed shared workers are used by default. | `` | false |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| definition_id | The definition ID. | `string` | true |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| property_name | The property name. | `string` | true |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| property_name | The property name. | `string` | true |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| id | ID of current instance. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| cd_tekton_pipeline_definition | cd_tekton_pipeline_definition object |
| cd_tekton_pipeline_trigger_property | cd_tekton_pipeline_trigger_property object |
| cd_tekton_pipeline_property | cd_tekton_pipeline_property object |
| cd_tekton_pipeline_trigger | cd_tekton_pipeline_trigger object |
| cd_tekton_pipeline | cd_tekton_pipeline object |
| cd_tekton_pipeline_definition | cd_tekton_pipeline_definition object |
| cd_tekton_pipeline_trigger_property | cd_tekton_pipeline_trigger_property object |
| cd_tekton_pipeline_property | cd_tekton_pipeline_property object |
| cd_tekton_pipeline_trigger | cd_tekton_pipeline_trigger object |
| cd_tekton_pipeline | cd_tekton_pipeline object |
