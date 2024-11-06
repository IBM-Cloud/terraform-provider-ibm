# Examples for sdsaas

These examples illustrate how to use the resources and data sources associated with sdsaas.

The following resources are supported:
* ibm_sds_volume
* ibm_sds_host

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## sdsaas resources

### Resource: ibm_sds_volume

```hcl
resource "ibm_sds_volume" "sds_volume_instance" {
  hostnqnstring = var.sds_volume_hostnqnstring
  capacity = var.sds_volume_capacity
  name = var.sds_volume_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| hostnqnstring | The host nqn. | `string` | false |
| capacity | The capacity of the volume (in gigabytes). | `number` | true |
| name | The name of the volume. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| host_mappings | List of host details that volume is mapped to. |
| created_at | The date and time that the volume was created. |
| resource_type | The resource type of the volume. |
| status | The current status of the volume. |
| status_reasons | Reasons for the current status of the volume. |

### Resource: ibm_sds_host

```hcl
resource "ibm_sds_host" "sds_host_instance" {
  name = var.sds_host_name
  nqn = var.sds_host_nqn
  volume_mappings = var.sds_host_volume_mappings
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | The name for this host. The name must not be used by another host.  If unspecified, the name will be a hyphenated list of randomly-selected words. | `string` | false |
| nqn | The NQN of the host configured in customer's environment. | `string` | true |
| volume_mappings | The host-to-volume map. | `list()` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The date and time that the host was created. |
| service_instance_id | The service instance ID this host should be created in. |


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
