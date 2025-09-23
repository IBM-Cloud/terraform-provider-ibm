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

Direct Link Connect gateway resource :

resource ibm_dl_provider_gateway test_dl_gateway {
  bgp_asn              = var.bgp_asn
  bgp_ibm_cidr         = var.bgp_ibm_cidr
  bgp_cer_cidr         = var.bgp_cer_cidr
  name                 = var.name
  speed_mbps           = var.speed_mbps
  port                 = data.ibm_dl_provider_ports.test_ds_dl_ports.ports[0].port_id
  customer_account_id  = var.customerAccID
}


## Direct Link Provider Data Sources

List ports :
```hcl
data "ibm_dl_provider_ports" "test_ds_dl_ports" {
 }
```
List all Direct Link Connect gateways created by this provider :
```hcl
data "ibm_dl_provider_gateways" "test_ibm_dl_provider_gws"{
}
```

## Examples

* [ Direct Link ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-direct-link-provider)

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
| bgp\_asn | The BGP ASN of the Gateway to be created. | `string` | yes |
| name | The unique user-defined name for this gateway. | `string` | yes |
| speed\_mbps | Gateway speed in megabits per second. | `integer` | yes |
| bgp\_cer_cidr | BGP customer edge router CIDR. | `string` | no |
| bgp\_ibm_cidr | BGP IBM CIDR. | `string` | no |
| customerAccID | Customer IBM Cloud account ID for the new gateway. A gateway object containing the pending create request will become available in the specified account. | `string` | yes |
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
