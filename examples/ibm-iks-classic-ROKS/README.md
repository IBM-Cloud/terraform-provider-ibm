# Example for IKS Classic ROKS resources

This example shows how to create a Kubernetes Cluster with openshift.

Following types of resources are supported:

* [ Container Cluster Resource](https://cloud.ibm.com/docs/terraform?topic=terraform-container-resources)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.5.1`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.27.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create a container cluster:

```hcl
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "b3c.4x16"
  hardware        = "shared"
  private_vlan_id = "2709721"
  private_service_endpoint = true
}
```

```hcl

data "ibm_container_cluster" "cluster_foo" {
  cluster_name_id = "FOO"
}

```

## Examples

* [ Container cluster ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-iks-openshift)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | Name of the cluster. | `string` | yes |
| datacenter | The datacenter of the worker nodes. Default: `wdc04`| `string` | yes |
| kube\_version | The desired Kubernetes version of the created cluster. Default: `3.11_openshift`. | `string` | no |
| machine\_type | The machine type of the worker nodes. Default: `b3c.4x16`| `string` | no |
| hardware | The level of hardware isolation for your worker node. Default: `shared` | `string` | no |
| public\_vlan_id | The public VLAN ID for the worker node. | `string` | no |
| private\_vlan_id | The private VLAN of the worker node. | `string` | no |


## Outputs

| Name | Description |
|------|-------------|
| cluster\_config\_file\_path | The path to the cluster configuration file. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

