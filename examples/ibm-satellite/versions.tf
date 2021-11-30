terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
    }
    restapi = {
      source  = "fmontezuma/restapi"
      version = "1.14.1"
    }
  }
}

provider "ibm" {
  region           = var.ibm_region
  ibmcloud_api_key = var.ibmcloud_api_key
}