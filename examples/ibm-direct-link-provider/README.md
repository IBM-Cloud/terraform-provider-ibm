# Example for Direct Link resources

This example shows how to create Direct Link Provider resources.

Following types of resources are supported:

* [Direct Link Provider Gateway](https://cloud.ibm.com/docs/terraform?topic=terraform-dl-provider-gateway-resource#dl-provider-gwy)



## Terraform versions

Terraform 0.12. Pin module version to `~> v1.5.1`. Branch - `master`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Direct Link Provider Resources



## Direct Link Provider Data Sources

List ports :
```hcl
data "ibm_dl_provider_ports" "test_ds_dl_ports" {
 }
```


## Examples

* [ Direct Link ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-direct-link-provider)

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


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
