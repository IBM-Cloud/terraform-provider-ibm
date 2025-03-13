terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
      version = "1.71.1"
    }
  }
}

provider "ibm" {
  ibmcloud_api_key = "uThF6-_vQg7Bx08Pcs0kMNi5sjYySCehr4H9GaMxUCGo"
}