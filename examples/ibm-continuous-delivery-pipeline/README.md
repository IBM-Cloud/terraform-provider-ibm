# Example for ContinuousDeliveryPipelineV2

This example illustrates how to use the ContinuousDeliveryPipelineV2

These types of resources are supported:

* tekton_pipeline_property

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ContinuousDeliveryPipelineV2 resources

tekton_pipeline_property resource:

```hcl
resource "tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  name = var.tekton_pipeline_property_name
  value = var.tekton_pipeline_property_value
  options = var.tekton_pipeline_property_options
  type = var.tekton_pipeline_property_type
  path = var.tekton_pipeline_property_path
}
```

## ContinuousDeliveryPipelineV2 Data sources

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
| name | Property name. | `string` | false |
| value | String format property value. | `string` | false |
| options | Options for SINGLE_SELECT property type. | `` | false |
| type | Property type. | `string` | false |
| path | property path for INTEGRATION type properties. | `string` | false |
| pipeline_id | The tekton pipeline ID. | `string` | true |
| property_name | The property's name. | `string` | true |
| pipeline_id | The tekton pipeline ID. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| tekton_pipeline_property | tekton_pipeline_property object |
| tekton_pipeline_property | tekton_pipeline_property object |
| tekton_pipeline_workers | tekton_pipeline_workers object |
