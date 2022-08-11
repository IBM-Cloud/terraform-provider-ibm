data "ibm_resource_group" "rg" {
  name = var.resource_group
}

provider "ibm" {
 }

data "ibm_dl_provider_ports" "test_ds_dl_ports" {
 
 }

resource ibm_dl_provider_gateway test_dl_gateway {
  bgp_asn              = var.bgp_asn
  bgp_ibm_cidr         = var.bgp_ibm_cidr
  bgp_cer_cidr         = var.bgp_cer_cidr
  name                 = var.name
  speed_mbps           = var.speed_mbps
  port                 = data.ibm_dl_provider_ports.test_ds_dl_ports.ports[0].port_id
  customer_account_id  = var.customerAccID
}
data "ibm_dl_provider_gateways" "test_ibm_dl_provider_gws" {
  
}