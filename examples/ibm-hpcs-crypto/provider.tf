provider "ibm" {

}
terraform {
  required_providers {
    ibm = {
      source  = "IBM-Cloud/ibm"
      version = ">=1.29.0"
    }
  }
}