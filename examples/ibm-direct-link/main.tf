data "ibm_resource_group" "rg" {
  name = var.resource_group
}

provider "ibm" {
 }
data "ibm_dl_routers" "test_dl_routers" {
  offering_type = var.type
  location_name = var.location_name
}

resource ibm_dl_gateway test_dl_gateway {
  bgp_asn              = var.bgp_asn
  bgp_base_cidr        = var.bgp_base_cidr
  bgp_ibm_cidr         = var.bgp_ibm_cidr 
  bgp_cer_cidr         = var.bgp_cer_cidr 
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
data "ibm_dl_ports" "test_ds_dl_ports" {
 
 }

# # datasource to list all dl gateways
# data "ibm_dl_gateways" "test_dl_gateways" {
# }
# # datasource to list all dl speeds for directlink dedicated
# data "ibm_dl_offering_speeds" "test_dl_speeds" {
#   offering_type = "dedicated"
# }
# # datasource to read a directlink gateway by name
# data "ibm_dl_gateway" "test_dl_gateway_vc" {
#   name =  ibm_dl_gateway.test_dl_gateway.name
# }
# # datasource to list all directlink ports
# data "ibm_dl_ports" "test_ds_dl_ports" {
# }
# # datasource to read a port for directlink
# data "ibm_dl_port" "test_ds_dl_port" {
#   port_id = "2f41cf65-e72a-4522-9526-e156e4ca02b5"
# }
# # datasource to list all locations for directlink dedicated
# data "ibm_dl_locations" "test_dl_locations"{
# 		offering_type = "dedicated"
# }



