provider "ibm" {
  iaas_classic_username = var.iaas_classic_username
  iaas_classic_api_key  = var.iaas_classic_api_key
  ibmcloud_api_key= var.apikey
  // softlayer_timeout  = "300"
}
