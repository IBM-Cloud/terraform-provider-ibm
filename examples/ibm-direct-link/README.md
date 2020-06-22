# Example for Direct Link resources

This example shows how to create Direct Link resources.

Following types of resources are supported:

* [Direct Link Gateways](https://cloud.ibm.com/docs/terraform)
* [Direct Link Virtual Connections](https://cloud.ibm.com/docs/terraform)
* [Direct Link Offering Information](https://cloud.ibm.com/docs/terraform)
* [Direct Link Ports](https://cloud.ibm.com/docs/terraform)


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

Create a direct link gateway:

```hcl
resource ibm_dl_gateway test_dl_gateway {
  bgp_asn =  64999
  bgp_base_cidr =  "169.254.0.0/16"
  bgp_ibm_cidr =  "169.254.0.29/30"
  bgp_cer_cidr =  "169.254.0.30/30"
  global = true 
  metered = false
  name = "terraformtestGateway"
  resource_group = data.ibm_resource_group.rg.id
  speed_mbps = 1000 
  loa_reject_reason = "The port mentioned was incorrect"
  operational_status = "loa_accepted"
  type =  "dedicated" 
  cross_connect_router = "LAB-xcr01.dal09"
  location_name = "dal09"
  customer_name = "Customer1" 
  carrier_name = "Carrier1"

}   
```

## Examples

* [ Direct Link ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-direct-link)

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
| bgp\_base_cidr | The BGP base CIDR of the Gateway to be created. | `string` | yes |
| global | Gateways with global routing (true) can connect to networks outside their associated region. | `boolean` | yes |
| metered | Metered billing option. When true gateway usage is billed per gigabyte. When false there is no per gigabyte usage charge, instead a flat rate is charged for the gateway. | `boolean` | yes |
| name | The unique user-defined name for this gateway. | `string` | yes |
| speed\_mbps | Gateway speed in megabits per second. | `integer` | yes |
| type | Gateway type. Allowable values: [dedicated,connect]. | `string` | yes |
| bgp\_cer_cidr | BGP customer edge router CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is 169.254.0.0/16, this field can be ommitted and a CIDR will be selected automatically. | `string` | no |
| bgp\_ibm_cidr | BGP IBM CIDR. Specify a value within bgp_base_cidr. If bgp_base_cidr is 169.254.0.0/16, this field can be ommitted and a CIDR will be selected automatically. | `string` | no |
| resource\_group | Resource group for this resource. If unspecified, the account's default resource group is used.  | `string` | no |
| carrier\_name | Carrier name. | `string` | yes |
| cross\_connect_router | Cross connect router. | `string` | yes |
| customer\_name | Customer name. | `string` | yes |
| location\_name |  Gateway location. | `string` | yes | 
| loa\_reject_reason | Use this field during LOA rejection to provide the reason for the rejection. | `string` | no |
| operational\_status | ateway operational status. For gateways pending LOA approval, patch operational_status to the appropriate value to approve or reject its LOA. When rejecting an LOA, provide reject reasoning in loa_reject_reason. | `string` | no | 


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
