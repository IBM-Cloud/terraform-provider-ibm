# Example for ProjectV1

This example illustrates how to use the ProjectV1

These types of resources are supported:

* project
* project_config

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ProjectV1 resources

project resource:

```hcl
resource "project" "project_instance" {
  resource_group = var.project_resource_group
  location = var.project_location
  name = var.project_name
  description = var.project_description
  destroy_on_delete = var.project_destroy_on_delete
  configs = var.project_configs
}
```
project_config resource:

```hcl
resource "project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  name = var.project_config_name
  locator_id = var.project_config_locator_id
  labels = var.project_config_labels
  description = var.project_config_description
  authorizations = var.project_config_authorizations
  compliance_profile = var.project_config_compliance_profile
  input = var.project_config_input
  setting = var.project_config_setting
}
```

## ProjectV1 Data sources

project data source:

```hcl
data "project" "project_instance" {
  id = ibm_project.project_instance.id
}
```
project_config data source:

```hcl
data "project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  id = ibm_project_config.project_config_instance.project_config_id
  version = var.project_config_version
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
| resource_group | The resource group where the project's data and tools are created. | `string` | true |
| location | The location where the project's data and tools are created. | `string` | true |
| name | The project name. | `string` | true |
| description | A project's descriptive text. | `string` | false |
| destroy_on_delete | The policy that indicates whether the resources are destroyed or not when a project is deleted. | `bool` | false |
| configs | The project configurations. | `list()` | false |
| project_id | The unique project ID. | `string` | true |
| name | The configuration name. | `string` | true |
| locator_id | A dotted value of catalogID.versionID. | `string` | true |
| labels | A collection of configuration labels. | `list(string)` | false |
| description | The project configuration description. | `string` | false |
| authorizations | The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager. | `` | false |
| compliance_profile | The profile required for compliance. | `` | false |
| input | The inputs of a Schematics template property. | `list()` | false |
| setting | Schematics environment variables to use to deploy the configuration. | `list()` | false |
| id | The unique project ID. | `string` | true |
| project_id | The unique project ID. | `string` | true |
| id | The unique config ID. | `string` | true |
| version | The version of the configuration to return. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| project | project object |
| project_config | project_config object |
