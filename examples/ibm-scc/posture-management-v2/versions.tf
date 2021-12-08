terraform {
  required_providers {
    ibm = {
      source = "github.ibm.com/compliance-terraform/ibm"
      version = "0.0.2"
    }
    google = {
      source = "hashicorp/google"
      version = "3.66.0"
    }
  }
}
