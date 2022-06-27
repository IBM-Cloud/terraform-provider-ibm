provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "us-south"
}

terraform {
  required_providers {
    ibm = {
      source  = "zaas/uko"
      version = "0.0.1"
    }
  }
}