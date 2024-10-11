terraform {
  required_version = ">= 1.0"
  required_providers {
    ibm = {
      source  = "terraform.local/ibm-cloud/ibm"
    }
  }
}
