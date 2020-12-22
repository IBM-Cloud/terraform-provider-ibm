# Example for Container Registry resources

This example illustrates how to use the Container Registry resources to create a namespace on to an account

These types of resources are supported:

* [ Namespace ](https://cloud.ibm.com/docs/Registry?topic=container-registry-cli-plugin-containerregcli#bx_cr_namespace_add)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.18.0`. Branch - `master`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Container Registry Resources

`Creating Namespace`:

```hcl
resource "ibm_cr_namespace" "namespace" {
  name              = var.name
  resource_group_id = data.ibm_resource_group.rg.id
}
```
##  Container Registry Data Source
`List all Namespaces:`

```hcl

data "ibm_cr_namespaces" "namespaces"{}

```

## Examples

* [ Container Registry ](./main.tf)


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

| Name | Description | Type | Required | Default |
|------|-------------|------|---------|
| name | Name Space Name. Force New Attribute| `string` | yes | N/A |
| resource_group_id | The Id of resource group to which the namespace has to be imported. Force New attribute | `string` | No | Default RG|

## Outputs

| Name | Description |
|------|-------------|
| namespace_crn | CRN of namespace resource. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
