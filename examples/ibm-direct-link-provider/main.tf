data "ibm_resource_group" "rg" {
  name = var.resource_group
}

provider "ibm" {
 }

data "ibm_dl_provider_ports" "test_ds_dl_ports" {
 
 }
