# Example for Direct Link resources

This example shows how to create Direct Link resources.

Following types of resources are supported:

* [Direct Link Gateway](https://cloud.ibm.com/docs/terraform?topic=terraform-dl-gateway-resource#dl-gwy)
* [Direct Link Virtual Connection](https://cloud.ibm.com/docs/terraform?topic=terraform-dl-gateway-resource#dl-vc)


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

## Direct Link Resources

Direct Link Dedicated gateway resource :

```hcl

resource ibm_dl_gateway test_dl_gateway {
  bgp_asn              = var.bgp_asn
  bgp_base_cidr        = var.bgp_base_cidr
  bgp_ibm_cidr         = var.bgp_ibm_cidr
  bgp_cer_cidr         =  var.bgp_cer_cidr
  global               = true
  metered              = false
  name                 = var.name
  resource_group       = data.ibm_resource_group.rg.id
  speed_mbps           = var.speed_mbps
  type                 = var.type
  cross_connect_router = data.ibm_dl_routers.test_dl_routers.cross_connect_routers[0].router_name
  location_name = data.ibm_dl_routers.test_dl_routers.location_name
  customer_name        = var.customer_name
  carrier_name         = var.carrier_name

}  
```

Create a virtual connection to the specified network :

```hcl
resource "ibm_is_vpc" "test_dl_vc_vpc" {
  name = var.vpc_name
}
resource "ibm_dl_virtual_connection" "test_dl_gateway_vc" {
  depends_on = [ibm_is_vpc.test_dl_vc_vpc, ibm_dl_gateway.test_dl_gateway]
  gateway    = ibm_dl_gateway.test_dl_gateway.id
  name       = var.vc_name
  type       = var.vc_type
  network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn
}   
```
Direct Link Connect gateway resource :

```hcl
resource "ibm_dl_gateway" "test_dl_connect" {
  bgp_asn =  var.bgp_asn
  bgp_base_cidr =  var.bgp_base_cidr
  global = true
  metered = false
  name = var.dl_connect_gw_name
  speed_mbps = 1000
  type =  "connect"
  port =  data.ibm_dl_ports.test_ds_dl_ports.ports[0].port_id
}  
```

## Direct Link Data Sources

Retrieve location specific cross connect router information. Only valid for offering_type=dedicated locations :

```hcl
data "ibm_dl_routers" "test_dl_routers" {
  offering_type = "dedicated"
  location_name = "dal10"
}
```

List available locations :

```hcl
data "ibm_dl_locations" "test_dl_locations"{
 		offering_type = "dedicated"
}
```

List speed options :

```hcl
data "ibm_dl_offering_speeds" "test_dl_speeds" {
  offering_type = "dedicated"
 }
```
List ports :
```hcl
data "ibm_dl_ports" "test_ds_dl_ports" {
 }
```

Get port :
```hcl
data "ibm_dl_port" "test_ds_dl_port" {
 port_id = data.ibm_dl_ports.test_ds_dl_ports.ports[0].port_id
 }
```

## Examples

* [ Direct Link ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-direct-link)

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
| vc\_name | Virtual Connection name. | `string` | yes |
| vc\_type | The type of virtual connection.Allowable values: [classic,vpc]. | `string` | yes |
| dl_connect_gw_name | The unique user-defined name for the direct link connect gateway. | `string` | yes |



<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
