provider "ibm" {
  region = var.ibm_region
}


terraform {
  required_providers {
    ibm = {
      source = "ibm-cloud/ibm"
    }
  }
}