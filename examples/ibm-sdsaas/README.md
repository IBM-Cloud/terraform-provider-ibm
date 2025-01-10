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
  sds_endpoint = var.sds_endpoint
  hostnqnstring = var.sds_volume_hostnqnstring
  capacity = var.sds_volume_capacity
  name = var.sds_volume_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| sds_endpoint | IBM Cloud Endpoint | `string` | false |
| hostnqnstring | The host nqn. | `string` | false |
| capacity | The capacity of the volume (in gigabytes). | `number` | true |
| name | The name of the volume. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| bandwidth | The maximum bandwidth (in megabits per second) for the volume. |
| created_at | The date and time that the volume was created. |
| hosts | List of host details that volume is mapped to. |
| iops | Iops The maximum I/O operations per second (IOPS) for this volume. |
| resource_type | The resource type of the volume. |
| status | The current status of the volume. |
| status_reasons | Reasons for the current status of the volume. |

### Resource: ibm_sds_host

```hcl
resource "ibm_sds_host" "sds_host_instance" {
  sds_endpoint = var.sds_endpoint
  name = var.sds_host_name
  nqn = var.sds_host_nqn
  volumes = var.sds_host_volumes
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| sds_endpoint | IBM Cloud Endpoint | `string` | false |
| name | The name for this host. The name must not be used by another host.  If unspecified, the name will be a hyphenated list of randomly-selected words. | `string` | false |
| nqn | The NQN of the host configured in customer's environment. | `string` | true |
| volumes | The host-to-volume map. | `list()` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The date and time that the host was created. |


## Assumptions

The `IBMCLOUD_SDS_ENDPOINT` can optionally be set instead of setting `sds_endpoint` in each of the resources. This is the endpoint provided to customers to perform operations against their service.

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
