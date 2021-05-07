
terraform {
  required_providers {
    ibm = {
      source = "ibm-cloud/ibm"
    }
  }
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}