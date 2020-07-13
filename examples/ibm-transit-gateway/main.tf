data "ibm_resource_group" "rg" {
  name = var.resource_group
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_tg_gateway" "new_tg_gw"{
name="tg-gateway-1"
location="us-south"
global=true
resource_group = data.ibm_resource_group.rg.id
} 

data "ibm_tg_gateway" "tg_gateway" {
name="tg-gateway-1"
}

data "ibm_tg_gateways" "all_tg_gws"{

}

data "ibm_tg_locations" "tg_locations" {
}
 
