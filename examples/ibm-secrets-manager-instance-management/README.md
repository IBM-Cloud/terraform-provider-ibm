# Examples for IBM Cloud Secrets Manager Instance Management API

These examples illustrate how to use the resources and data sources associated with IBM Cloud Secrets Manager Instance Management API.

The following resources are supported:
* ibm_sm_destination

The following data sources are supported:
* ibm_sm_instance
* ibm_sm_destinations
* ibm_sm_destination

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Cloud Secrets Manager Instance Management API resources

### Resource: ibm_sm_destination

```hcl
resource "ibm_sm_destination" "sm_destination_instance" {
  instance_id = var.sm_destination_instance_id
  name = var.sm_destination_name
  type = var.sm_destination_type
  description = var.sm_destination_description
  crn = var.sm_destination_crn
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | Secrets Manager instance ID. | `string` | true |
| name | Destination name. | `string` | true |
| type | Destination type. | `string` | true |
| description | Optional description. | `string` | false |
| crn | IBM Cloud Database service instance CRN. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | The URL of the destination resource. |
| state | Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress. |
| created_at | Timestamp when the destination was created. |
| updated_at | Timestamp when the destination was last updated. |
| created_by | IAM identity that created the destination. |
| message | Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed. |
| destination_id | Destination ID. |

## IBM Cloud Secrets Manager Instance Management API data sources

### Data source: ibm_sm_instance

```hcl
data "ibm_sm_instance" "sm_instance_instance" {
  sm_instance_id = var.sm_instance_sm_instance_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| sm_instance_id | Secrets Manager instance ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| name | The instance name. |
| instance_crn | The instance CRN identifier. |
| plan | Instance plan name. |
| vault_cluster | Vault cluster information for Vault Dedicated instances. |
| endpoints | Instance endpoints for Vault Dedicated instances. |
| encryption | Vault encryption configuration for Vault Dedicated instances. |
| href | The URL of the instance resource. |

### Data source: ibm_sm_destinations

```hcl
data "ibm_sm_destinations" "sm_destinations_instance" {
  instance_id = var.sm_destinations_instance_id
  state = var.sm_destinations_state
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| state | Filter by destination state. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| destinations | List of destinations. |
| total | Total number of destinations. Maximum 20 per instance. |

### Data source: ibm_sm_destination

```hcl
data "ibm_sm_destination" "sm_destination_instance" {
  instance_id = var.data_sm_destination_instance_id
  destination_id = var.data_sm_destination_destination_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | Secrets Manager instance ID. | `string` | true |
| destination_id | Destination ID. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | The URL of the destination resource. |
| name | Destination name. |
| type | Destination type. |
| description | Optional description. |
| state | Destination state:- `not_started`: Initial state before the first provisioning attempt begins- `provisioning`: Provisioning in progress — poll until `succeeded` or `failed`- `succeeded`: Destination ready and usable- `failed`: Terminal state reached when provisioning or deletion fails. A `failed` destination still counts against the per-instance quota until deleted.- `deleting`: Deletion in progress. |
| created_at | Timestamp when the destination was created. |
| updated_at | Timestamp when the destination was last updated. |
| created_by | IAM identity that created the destination. |
| message | Human-readable message providing additional context about the current state. Present only when non-empty — set when `state` is `failed`, describing why provisioning or deletion failed. |
| crn | IBM Cloud Database service instance CRN. |

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
