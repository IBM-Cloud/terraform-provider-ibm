# Example for ProjectV1

This example illustrates how to use the ProjectV1

These types of resources are supported:

* Project definition

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## ProjectV1 resources

project_instance resource:

```hcl
resource "project_instance" "project_instance" {
  name = var.project_instance_name
  description = var.project_instance_description
  configs = var.project_instance_configs
  resource_group = var.project_instance_resource_group
  location = var.project_instance_location
}
```

## ProjectV1 Data sources

project_event_notification data source:

```hcl
data "project_event_notification" "project_event_notification_instance" {
  project_id = var.project_event_notification_id
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
| resource_group | The resource group where the project's data and tools are created. | `string` | true |
| location | The location where the project's data and tools are created. | `string` | true |
| project_id | The unique project ID. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| project_instance | project_instance object |
| project_event_notification | project_event_notification object |
