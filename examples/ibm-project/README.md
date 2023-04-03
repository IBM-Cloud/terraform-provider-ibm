# Example for ProjectV1

This example illustrates how to use the ProjectV1

These types of resources are supported:

* project

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
  name = var.project_name
  description = var.project_description
  configs = var.project_configs
  resource_group = var.project_resource_group
  location = var.project_location
}
```

## ProjectV1 Data sources

project data source:

```hcl
data "project" "project_instance" {
  id = var.project_id
  exclude_configs = var.project_exclude_configs
  complete = var.project_complete
}
```
event_notification data source:

```hcl
data "event_notification" "event_notification_instance" {
  id = var.event_notification_id
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
| name | The project name. | `string` | true |
| description | A project's descriptive text. | `string` | false |
| configs | The project configurations. | `list()` | false |
| resource_group | Group name of the customized collection of resources. | `string` | false |
| location | Data center locations for resource deployment. | `string` | false |
| id | The ID of the project, which uniquely identifies it. | `string` | true |
| exclude_configs | Only return with the active configuration, no drafts. | `bool` | false |
| complete | The flag to determine if full metadata should be returned. | `bool` | false |
| id | The ID of the project, which uniquely identifies it. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| project | project object |
| project | project object |
| event_notification | event_notification object |
