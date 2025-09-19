# Examples for Projects API

These examples illustrate how to use the resources and data sources associated with Projects API.

The following resources are supported:
* ibm_project_config
* ibm_project
* ibm_project_environment

The following data sources are supported:
* ibm_project_config
* ibm_project
* ibm_project_environment

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Projects API resources

### Resource: ibm_project_config

```hcl
resource "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  definition = var.project_config_definition
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| project_id | The unique project ID. | `string` | true |
| schematics | A Schematics workspace that is associated to a project configuration, with scripts. | `` | false |
| definition |  | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| version | The version of the configuration. |
| needs_attention_state | The needs attention state of a configuration. |
| created_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| modified_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| outputs | The outputs of a Schematics template property. |
| references | The resolved references that are used by the configuration. |
| state | The state of the configuration. |
| state_code | Computed state code clarifying the prerequisites for validation for the configuration. |
| config_error | The error from config actions. |
| href | A Url. |
| is_draft | The flag that indicates whether the version of the configuration is draft, or active. |
| last_saved_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| project | The project that is referenced by this resource. |
| update_available | The flag that indicates whether a configuration update is available. |
| template_id | The stack definition identifier. |
| member_of | The stack config parent of which this configuration is a member of. |
| deployment_model | The configuration type. |
| approved_version | A summary of a project configuration version. |
| deployed_version | A summary of a project configuration version. |
| project_config_id | The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration. |

### Resource: ibm_project

```hcl
resource "ibm_project" "project_instance" {
  location = var.project_location
  resource_group = var.project_resource_group
  definition = var.project_definition
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| location | The IBM Cloud location where a resource is deployed. | `string` | true |
| resource_group | The resource group name where the project's data and tools are created. | `string` | true |
| definition | The definition of the project. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| crn | An IBM Cloud resource name that uniquely identifies a resource. |
| created_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| cumulative_needs_attention_view | The cumulative list of needs attention items for a project. If the view is successfully retrieved, an empty or nonempty array is returned. |
| cumulative_needs_attention_view_error | A value of `true` indicates that the fetch of the needs attention items failed. This property only exists if there was an error while retrieving the cumulative needs attention view. |
| resource_group_id | The resource group ID where the project's data and tools are created. |
| state | The project status value. |
| href | A Url. |
| event_notifications_crn | The CRN of the Event Notifications instance if one is connected to this project. |
| configs | The project configurations. These configurations are only included in the response of creating a project if a configuration array is specified in the request payload. |
| environments | The project environment. These environments are only included in the response if project environments were created on the project. |

### Resource: ibm_project_environment

```hcl
resource "ibm_project_environment" "project_environment_instance" {
  project_id = ibm_project.project_instance.id
  definition = var.project_environment_definition
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| project_id | The unique project ID. | `string` | true |
| definition | The environment definition. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| project | The project that is referenced by this resource. |
| created_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| target_account | The target account ID derived from the authentication block values. The target account exists only if the environment currently has an authorization block. |
| modified_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| href | A Url. |
| project_environment_id | The environment ID as a friendly name. |

## Projects API data sources

### Data source: ibm_project_config

```hcl
data "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  project_config_id = ibm_project_config.project_config_instance.project_config_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The unique project ID. | `string` | true |
| project_config_id | The unique configuration ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| version | The version of the configuration. |
| needs_attention_state | The needs attention state of a configuration. |
| created_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| modified_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| outputs | The outputs of a Schematics template property. |
| references | The resolved references that are used by the configuration. |
| state | The state of the configuration. |
| state_code | Computed state code clarifying the prerequisites for validation for the configuration. |
| config_error | The error from config actions. |
| href | A Url. |
| is_draft | The flag that indicates whether the version of the configuration is draft, or active. |
| last_saved_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| project | The project that is referenced by this resource. |
| schematics | A Schematics workspace that is associated to a project configuration, with scripts. |
| update_available | The flag that indicates whether a configuration update is available. |
| template_id | The stack definition identifier. |
| member_of | The stack config parent of which this configuration is a member of. |
| deployment_model | The configuration type. |
| definition |  |
| approved_version | A summary of a project configuration version. |
| deployed_version | A summary of a project configuration version. |

### Data source: ibm_project

```hcl
data "ibm_project" "project_instance" {
  project_id = ibm_project.project_instance.id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The unique project ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| crn | An IBM Cloud resource name that uniquely identifies a resource. |
| created_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| cumulative_needs_attention_view | The cumulative list of needs attention items for a project. If the view is successfully retrieved, an empty or nonempty array is returned. |
| cumulative_needs_attention_view_error | A value of `true` indicates that the fetch of the needs attention items failed. This property only exists if there was an error while retrieving the cumulative needs attention view. |
| location | The IBM Cloud location where a resource is deployed. |
| resource_group_id | The resource group ID where the project's data and tools are created. |
| state | The project status value. |
| href | A Url. |
| resource_group | The resource group name where the project's data and tools are created. |
| event_notifications_crn | The CRN of the Event Notifications instance if one is connected to this project. |
| configs | The project configurations. These configurations are only included in the response of creating a project if a configuration array is specified in the request payload. |
| environments | The project environment. These environments are only included in the response if project environments were created on the project. |
| definition | The definition of the project. |

### Data source: ibm_project_environment

```hcl
data "ibm_project_environment" "project_environment_instance" {
  project_id = ibm_project.project_instance.id
  project_environment_id = ibm_project_environment.project_environment_instance.project_environment_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| project_id | The unique project ID. | `string` | true |
| project_environment_id | The environment ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| project | The project that is referenced by this resource. |
| created_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| target_account | The target account ID derived from the authentication block values. The target account exists only if the environment currently has an authorization block. |
| modified_at | A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339. |
| href | A Url. |
| definition | The environment definition. |

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
