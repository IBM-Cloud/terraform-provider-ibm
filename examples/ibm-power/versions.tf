terraform {
  required_version = ">= 0.13"
}

terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
      version = "<desired_provider_version>"
    }
  }
}
