data "ibm_resource_group" "rg" {
  name = var.resource_group
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

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
resource "ibm_is_vpc" "test_dl_vc_vpc" {
		name = "myVpc"
}  
	
resource "ibm_dl_virtual_connection" "test_dl_gateway_vc"{
		depends_on = [ibm_is_vpc.test_dl_vc_vpc,ibm_dl_gateway.test_dl_gateway]
		gateway = ibm_dl_gateway.test_dl_gateway.id
		name = "myVC"
		type = "vpc"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn
}

data "ibm_dl_gateways" "test_dl_gateways" {
}
 

