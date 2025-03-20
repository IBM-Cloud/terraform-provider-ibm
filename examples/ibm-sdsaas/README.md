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
  capacity = var.sds_volume_capacity
  name = var.sds_volume_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| sds_endpoint | IBM Cloud Endpoint | `string` | false |
| capacity | The capacity of the volume (in gigabytes). | `number` | true |
| name | Unique name of the host. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| bandwidth | The maximum bandwidth (in megabits per second) for the volume. |
| created_at | The date and time that the volume was created. |
| href | The URL for this resource. |
| iops | Iops The maximum I/O operations per second (IOPS) for this volume. |
| resource_type | The resource type of the volume. |
| status | The status of the volume resource. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered. |
| status_reasons | The reasons for the current status (if any). |
| volume_mappings | List of volume mappings for this volume. |

### Resource: ibm_sds_host

```hcl
resource "ibm_sds_host" "sds_host_instance" {
  sds_endpoint = var.sds_endpoint
  name = var.sds_host_name
  nqn = var.sds_host_nqn
  volume_mappings = var.sds_host_volume_mappings
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| sds_endpoint | IBM Cloud Endpoint | `string` | false |
| name | Unique name of the host. | `string` | false |
| nqn | The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage. | `string` | true |
| volume_mappings | The host-to-volume map. | `list()` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The date and time when the resource was created. |
| href | The URL for this resource. |


### Resource: ibm_sds_volume_mapping

```hcl
resource "ibm_sds_volume_mapping" "sds_volume_mapping_instance" {
  host_id = ibm_sds_host.sds_host_instance.id
  volume = var.sds_volume_mapping_volume
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| host_id | A unique host ID. | `string` | true |
| volume | The volume reference. | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| status | The status of the volume mapping. The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered. |
| storage_identifier | Storage network and ID information associated with a volume/host mapping. |
| href | The URL for this resource. |
| host | Host mapping schema. |
| subsystem_nqn | The NVMe target subsystem NQN (NVMe Qualified Name) that can be used for doing NVMe connect by the initiator. |
| namespace | The NVMe namespace properties for a given volume mapping. |
| gateways | List of NVMe gateways. |
| volume_mapping_id | Unique identifier of the mapping. |

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
