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
  type = var.cd_tekton_pipeline_property_type
  path = var.cd_tekton_pipeline_property_path
}
```
cd_tekton_pipeline_trigger resource:

```hcl
resource "cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  type = var.cd_tekton_pipeline_trigger_type
  name = var.cd_tekton_pipeline_trigger_name
  event_listener = var.cd_tekton_pipeline_trigger_event_listener
  tags = var.cd_tekton_pipeline_trigger_tags
  worker = var.cd_tekton_pipeline_trigger_worker
  max_concurrent_runs = var.cd_tekton_pipeline_trigger_max_concurrent_runs
  disabled = var.cd_tekton_pipeline_trigger_disabled
  secret = var.cd_tekton_pipeline_trigger_secret
  cron = var.cd_tekton_pipeline_trigger_cron
  timezone = var.cd_tekton_pipeline_trigger_timezone
  scm_source = var.cd_tekton_pipeline_trigger_scm_source
  events = var.cd_tekton_pipeline_trigger_events
}
```
cd_tekton_pipeline resource:

```hcl
resource "cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  enable_slack_notifications = var.cd_tekton_pipeline_enable_slack_notifications
  enable_partial_cloning = var.cd_tekton_pipeline_enable_partial_cloning
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
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| trigger_id | The trigger ID. | `string` | true |
| name | Property name. | `string` | false |
| value | Property value. | `string` | false |
| enum | Options for `single_select` property type. Only needed for `single_select` property type. | `list(string)` | false |
| type | Property type. | `string` | false |
| path | A dot notation path for `integration` type properties to select a value from the tool integration. If left blank the full tool integration JSON will be selected. | `string` | false |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| name | Property name. | `string` | false |
| value | Property value. | `string` | false |
| enum | Options for `single_select` property type. Only needed when using `single_select` property type. | `list(string)` | false |
| type | Property type. | `string` | false |
| path | A dot notation path for `integration` type properties to select a value from the tool integration. | `string` | false |
| pipeline_id | The Tekton pipeline ID. | `string` | true |
| type | Trigger type. | `string` | false |
| name | Trigger name. | `string` | false |
| event_listener | Event listener name. The name of the event listener to which the trigger is associated. The event listeners are defined in the definition repositories of the Tekton pipeline. | `string` | false |
| tags | Trigger tags array. | `list(string)` | false |
| worker | Worker used to run the trigger. If not specified the trigger will use the default pipeline worker. | `` | false |
| max_concurrent_runs | Defines the maximum number of concurrent runs for this trigger. Omit this property to disable the concurrency limit. | `number` | false |
| disabled | Flag whether the trigger is disabled. If omitted the trigger is enabled by default. | `bool` | false |
| secret | Only needed for generic webhook trigger type. Secret used to start generic webhook trigger. | `` | false |
| cron | Only needed for timer triggers. Cron expression for timer trigger. | `string` | false |
| timezone | Only needed for timer triggers. Timezone for timer trigger. | `string` | false |
| scm_source | SCM source repository for a Git trigger. Only needed for Git triggers. | `` | false |
| events | Only needed for Git triggers. Events object defines the events to which this Git trigger listens. | `` | false |
| enable_slack_notifications | Flag whether to enable slack notifications for this pipeline. When enabled, pipeline run events will be published on all slack integration specified channels in the enclosing toolchain. | `bool` | false |
| enable_partial_cloning | Flag whether to enable partial cloning for this pipeline. When partial clone is enabled, only the files contained within the paths specified in definition repositories will be read and cloned. This means symbolic links may not work. | `bool` | false |
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
