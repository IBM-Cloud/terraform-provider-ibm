# Example for ProjectV1

This example illustrates how to use the ProjectV1

The following types of resources are supported:

* project_config
* project

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ProjectV1 resources

project_config resource:

```hcl
resource "project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.project_id
  definition = var.project_config_definition
}
```
project resource:

```hcl
resource "project" "project_instance" {
  location = var.project_location
  resource_group = var.project_resource_group
  definition = var.project_definition
}
```

## ProjectV1 data sources

project_config data source:

```hcl
data "project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  project_config_id = ibm_project_config.project_config_instance.project_config_id
}
```
project data source:

```hcl
data "project" "project_instance" {
  project_id = ibm_project.project_instance.id
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
| project_id | The unique project ID. | `string` | true |
| schematics | A schematics workspace associated to a project configuration, with scripts. | `` | false |
| definition | The type and output of a project configuration. | `` | true |
| location | The IBM Cloud location where a resource is deployed. | `string` | true |
| resource_group | The resource group name where the project's data and tools are created. | `string` | true |
| definition | The definition of the project. | `` | true |
| project_id | The unique project ID. | `string` | true |
| project_config_id | The unique config ID. | `string` | true |
| project_id | The unique project ID. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| project_config | project_config object |
| project | project object |
