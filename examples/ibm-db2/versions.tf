terraform {
  required_version = ">= 1.2.0, < 2.0.0"

  required_providers {
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = ">= 1.71.3-beta1"
    }
  }
}