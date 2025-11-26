# Examples for VMware Cloud Foundation as a Service API

These examples illustrate how to use the resources and data sources associated with VMware Cloud Foundation as a Service API.

The following resources are supported:
* ibm_vmaas_vdc

The following data sources are supported:
* ibm_vmaas_vdc

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## VMware Cloud Foundation as a Service API resources

### Resource: ibm_vmaas_vdc

```hcl
resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
  accept_language = var.vmaas_vdc_accept_language
  cpu = var.vmaas_vdc_cpu
  name = var.vmaas_vdc_name
  ram = var.vmaas_vdc_ram
  fast_provisioning_enabled = var.vmaas_vdc_fast_provisioning_enabled
  rhel_byol = var.vmaas_vdc_rhel_byol
  windows_byol = var.vmaas_vdc_windows_byol
  director_site = var.vmaas_vdc_director_site
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| accept_language | Language. | `string` | false |
| cpu | The vCPU usage limit on the virtual data center (VDC). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved. | `number` | false |
| name | A human readable ID for the virtual data center (VDC). | `string` | true |
| ram | The RAM usage limit on the virtual data center (VDC) in GB (1024^3 bytes). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved. | `number` | false |
| fast_provisioning_enabled | Determines whether this virtual data center has fast provisioning enabled or not. | `bool` | false |
| rhel_byol | Indicates if the RHEL VMs will be using the license from IBM or the customer will use their own license (BYOL). | `bool` | false |
| windows_byol | Indicates if the Microsoft Windows VMs will be using the license from IBM or the customer will use their own license (BYOL). | `bool` | false |
| director_site | The Cloud Director site in which to deploy the virtual data center (VDC). | `` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | The URL of this virtual data center (VDC). |
| provisioned_at | The time that the virtual data center (VDC) is provisioned and available to use. |
| crn | A unique ID for the virtual data center (VDC) in IBM Cloud. |
| deleted_at | The time that the virtual data center (VDC) is deleted. |
| ha | Indicates if the VDC is HA-enabled for compute only, compute and network, or network only. If not present, the VDC is not HA-enabled. |
| edges | The VMware NSX-T networking edges deployed on the virtual data center (VDC). NSX-T edges are used for bridging virtualization networking to the physical public-internet and IBM private networking. |
| status_reasons | Information about why the request to create the virtual data center (VDC) cannot be completed. |
| ordered_at | The time that the virtual data center (VDC) is ordered. |
| org_href | The URL of the organization that owns the VDC. |
| org_name | The name of the VMware Cloud Director organization that contains this virtual data center (VDC). VMware Cloud Director organizations are used to create strong boundaries between VDCs. There is a complete isolation of user administration, networking, workloads, and VMware Cloud Director catalogs between different Director organizations. |
| status | Determines the state of the virtual data center. |
| type | Determines whether this virtual data center is in a single-tenant or multitenant Cloud Director site. |

### Resource: ibm_vmaas_transit_gateway_connection

```hcl
resource "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
  vdc_id = var.vmaas_transit_gateway_connection_vdc_id
  edge_id = var.vmaas_transit_gateway_connection_edge_id
  region = var.vmaas_transit_gateway_connection_region
  vmaas_transit_gateway_connection_id = var.vmaas_transit_gateway_connection_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| vdc_id | A unique ID for a virtual data center. | `string` | true |
| edge_id | A unique ID for an edge. | `string` | true |
| region | The region where the IBM Transit Gateway is deployed. | `string` | false |
| vmaas_transit_gateway_connection_id | A unique ID for an IBM Transit Gateway. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| connections | IBM Transit Gateway connections. |
| status | Determines the state of the IBM Transit Gateway based on its connections. |
| vmaas_transit_gateway_connection_id | A unique ID for an IBM Transit Gateway. |


## VMware Cloud Foundation as a Service API data sources

### Data source: ibm_vmaas_vdc

```hcl
data "ibm_vmaas_vdc" "vmaas_vdc_instance" {
  vmaas_vdc_id = var.data_vmaas_vdc_vmaas_vdc_id
  accept_language = var.data_vmaas_vdc_accept_language
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| vmaas_vdc_id | A unique ID for a specified virtual data center. | `string` | true |
| accept_language | Language. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| href | The URL of this virtual data center (VDC). |
| provisioned_at | The time that the virtual data center (VDC) is provisioned and available to use. |
| cpu | The vCPU usage limit on the virtual data center (VDC). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved. |
| crn | A unique ID for the virtual data center (VDC) in IBM Cloud. |
| deleted_at | The time that the virtual data center (VDC) is deleted. |
| ha | Indicates if the VDC is HA-enabled for compute only, compute and network, or network only. If not present, the VDC is not HA-enabled. |
| edges | The VMware NSX-T networking edges deployed on the virtual data center (VDC). NSX-T edges are used for bridging virtualization networking to the physical public-internet and IBM private networking. |
| status_reasons | Information about why the request to create the virtual data center (VDC) cannot be completed. |
| name | A human readable ID for the virtual data center (VDC). |
| ordered_at | The time that the virtual data center (VDC) is ordered. |
| org_href | The URL of the organization that owns the VDC. |
| org_name | The name of the VMware Cloud Director organization that contains this virtual data center (VDC). VMware Cloud Director organizations are used to create strong boundaries between VDCs. There is a complete isolation of user administration, networking, workloads, and VMware Cloud Director catalogs between different Director organizations. |
| ram | The RAM usage limit on the virtual data center (VDC) in GB (1024^3 bytes). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved. |
| status | Determines the state of the virtual data center. |
| type | Determines whether this virtual data center is in a single-tenant or multitenant Cloud Director site. |
| fast_provisioning_enabled | Determines whether this virtual data center has fast provisioning enabled or not. |
| rhel_byol | Indicates if the RHEL VMs will be using the license from IBM or the customer will use their own license (BYOL). |
| windows_byol | Indicates if the Microsoft Windows VMs will be using the license from IBM or the customer will use their own license (BYOL). |
| director_site | The Cloud Director site in which to deploy the virtual data center (VDC). |

### Data source: ibm_vmaas_transit_gateway_connection

```hcl
data "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
  vmaas_transit_gateway_connection_id = var.data_vmaas_transit_gateway_connection_vmaas_transit_gateway_connection_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| vmaas_transit_gateway_connection_id | A unique ID for a specified virtual data center. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| connections | IBM Transit Gateway connections. |
| status | Determines the state of the IBM Transit Gateway based on its connections. |
| region | The region where the IBM Transit Gateway is deployed. |

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
