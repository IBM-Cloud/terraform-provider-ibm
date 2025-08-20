# Examples for DrAutomation Service

These examples illustrate how to use the resources and data sources associated with DrAutomation Service.

The following resources are supported:
* ibm_pdr_managedr
* ibm_pdr_validate_apikey

The following data sources are supported:
* ibm_pdr_get_deployment_status
* ibm_pdr_get_event
* ibm_pdr_get_events
* ibm_pdr_get_machine_types
* ibm_pdr_get_managed_vm_list
* ibm_pdr_last_operation
* ibm_pdr_validate_clustertype
* ibm_pdr_validate_proxyip
* ibm_pdr_validate_workspace
* ibm_pdr_get_dr_summary_response

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## DrAutomation Service resources

### Resource: ibm_pdr_managedr

```hcl
resource "ibm_pdr_managedr" "pdr_managedr_instance" {
  instance_id = var.pdr_managedr_instance_id
  stand_by_redeploy = var.pdr_managedr_stand_by_redeploy
  accept_language = var.pdr_managedr_accept_language
  if_none_match = var.pdr_managedr_if_none_match
  accepts_incomplete = var.pdr_managedr_accepts_incomplete
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | instance id of instance to provision. | `string` | true |
| stand_by_redeploy | Flag to indicate if standby should be redeployed (must be "true" or "false"). | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |
| accepts_incomplete | A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| dashboard_url |  |
| instance_id |  |

### Resource: ibm_pdr_validate_apikey

```hcl
resource "ibm_pdr_validate_apikey" "pdr_validate_apikey_instance" {
  instance_id = var.pdr_validate_apikey_instance_id
  accept_language = var.pdr_validate_apikey_accept_language
  if_none_match = var.pdr_validate_apikey_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | instance id of instance to provision. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| description | Validation result message. |
| status | Status of the API key. |
| instance_id |  |

## DrAutomation Service data sources

### Data source: ibm_pdr_get_deployment_status

```hcl
data "ibm_pdr_get_deployment_status" "pdr_get_deployment_status_instance" {
  instance_id = var.pdr_get_deployment_status_instance_id
  if_none_match = var.pdr_get_deployment_status_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| orch_ext_connectivity_status |  |
| orch_standby_node_addition_status |  |
| orchestrator_cluster_message |  |
| orchestrator_cluster_type |  |
| orchestrator_config_status |  |
| orchestrator_group_leader |  |
| orchestrator_name |  |
| orchestrator_status |  |
| schematic_workspace_name |  |
| schematic_workspace_status |  |
| ssh_key_name |  |
| standby_orchestrator_name |  |
| standby_orchestrator_status |  |

### Data source: ibm_pdr_get_event

```hcl
data "ibm_pdr_get_event" "pdr_get_event_instance" {
  provision_id = var.pdr_get_event_provision_id
  event_id = var.pdr_get_event_event_id
  accept_language = var.pdr_get_event_accept_language
  if_none_match = var.pdr_get_event_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| provision_id | provision id. | `string` | true |
| event_id | Event ID. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| action | Type of action for this event. |
| api_source | Source of API when it being executed. |
| level | Level of the event (notice, info, warning, error). |
| message | The (translated) message of the event. |
| message_data |  |
| metadata | Any metadata associated with the event. |
| resource | Type of resource for this event. |
| time | Time of activity in ISO 8601 - RFC3339. |
| timestamp | Time of activity in unix epoch. |
| user |  |

### Data source: ibm_pdr_get_events

```hcl
data "ibm_pdr_get_events" "pdr_get_events_instance" {
  provision_id = var.pdr_get_events_provision_id
  time = var.pdr_get_events_time
  from_time = var.pdr_get_events_from_time
  to_time = var.pdr_get_events_to_time
  accept_language = var.pdr_get_events_accept_language
  if_none_match = var.pdr_get_events_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| provision_id | provision id. | `string` | true |
| time | (deprecated - use from_time) A time in either ISO 8601 or unix epoch format. | `string` | false |
| from_time | A from query time in either ISO 8601 or unix epoch format. | `string` | false |
| to_time | A to query time in either ISO 8601 or unix epoch format. | `string` | false |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| events | Events. |

### Data source: ibm_pdr_get_machine_types

```hcl
data "ibm_pdr_get_machine_types" "pdr_get_machine_types_instance" {
  instance_id = var.pdr_get_machine_types_instance_id
  primary_workspace_name = var.pdr_get_machine_types_primary_workspace_name
  accept_language = var.pdr_get_machine_types_accept_language
  if_none_match = var.pdr_get_machine_types_if_none_match
  standby_workspace_name = var.pdr_get_machine_types_standby_workspace_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| primary_workspace_name | Primary Workspace Name. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |
| standby_workspace_name | Standby Workspace Name. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| workspaces | Map of workspace IDs to lists of machine types. |

### Data source: ibm_pdr_get_managed_vm_list

```hcl
data "ibm_pdr_get_managed_vm_list" "pdr_get_managed_vm_list_instance" {
  instance_id = var.pdr_get_managed_vm_list_instance_id
  accept_language = var.pdr_get_managed_vm_list_accept_language
  if_none_match = var.pdr_get_managed_vm_list_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| managed_vms |  |

### Data source: ibm_pdr_last_operation

```hcl
data "ibm_pdr_last_operation" "pdr_last_operation_instance" {
  instance_id = var.pdr_last_operation_instance_id
  accept_language = var.pdr_last_operation_accept_language
  if_none_match = var.pdr_last_operation_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| crn |  |
| deployment_name |  |
| is_ksys_ha |  |
| orch_ext_connectivity_status |  |
| orch_standby_node_addtion_status |  |
| orchestrator_cluster_message |  |
| orchestrator_config_status |  |
| primary_description |  |
| primary_ip_address |  |
| primary_orchestrator_status |  |
| recovery_location |  |
| resource_group |  |
| standby_description |  |
| standby_ip_address |  |
| standby_status |  |
| status |  |

### Data source: ibm_pdr_validate_clustertype

```hcl
data "ibm_pdr_validate_clustertype" "pdr_validate_clustertype_instance" {
  instance_id = var.pdr_validate_clustertype_instance_id
  orchestrator_cluster_type = var.pdr_validate_clustertype_orchestrator_cluster_type
  accept_language = var.pdr_validate_clustertype_accept_language
  if_none_match = var.pdr_validate_clustertype_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| orchestrator_cluster_type | orchestrator cluster type value. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| description |  |
| status |  |

### Data source: ibm_pdr_validate_proxyip

```hcl
data "ibm_pdr_validate_proxyip" "pdr_validate_proxyip_instance" {
  instance_id = var.pdr_validate_proxyip_instance_id
  proxyip = var.pdr_validate_proxyip_proxyip
  vpc_location = var.pdr_validate_proxyip_vpc_location
  vpc_id = var.pdr_validate_proxyip_vpc_id
  if_none_match = var.pdr_validate_proxyip_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| proxyip | proxyip value. | `string` | true |
| vpc_location | vpc location value. | `string` | true |
| vpc_id | vpc id value. | `string` | true |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| description |  |
| status |  |
| warning | Indicates whether the proxy IP is valid but has an advisory (e.g., not in reserved IPs). |

### Data source: ibm_pdr_validate_workspace

```hcl
data "ibm_pdr_validate_workspace" "pdr_validate_workspace_instance" {
  instance_id = var.pdr_validate_workspace_instance_id
  workspace_id = var.pdr_validate_workspace_workspace_id
  crn = var.pdr_validate_workspace_crn
  location_url = var.pdr_validate_workspace_location_url
  if_none_match = var.pdr_validate_workspace_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| workspace_id | standBy workspaceID value. | `string` | true |
| crn | crn value. | `string` | true |
| location_url | schematic_workspace_id value. | `string` | true |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| description |  |
| status |  |

### Data source: ibm_pdr_get_dr_summary_response

```hcl
data "ibm_pdr_get_dr_summary_response" "pdr_get_dr_summary_response_instance" {
  instance_id = var.pdr_get_dr_summary_response_instance_id
  accept_language = var.pdr_get_dr_summary_response_accept_language
  if_none_match = var.pdr_get_dr_summary_response_if_none_match
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | instance id of instance to provision. | `string` | true |
| accept_language | The language requested for the return document. | `string` | false |
| if_none_match | ETag for conditional requests (optional). | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| managed_vm_list |  |
| orchestrator_details | Contains details about the orchestrator configuration. |
| service_details | Contains details about the DR automation service. |
| dashboard_url |  |

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
