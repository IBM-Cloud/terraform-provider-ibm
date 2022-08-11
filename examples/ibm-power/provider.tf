provider "ibm" {
  ibmcloud_api_key = var.ibm_cloud_api_key // export IC_API_KEY = "<api key>"
  region           = var.region            // export IBMCLOUD_REGION = "<region>"
  zone             = var.zone              // export IBMCLOUD_ZONE = "<zone>"
}