# ========================================
# IBM Cloud Provider Configuration
# ========================================

# The API key for IBM Cloud 
variable "ibmcloud_api_key" {
  description = "IBM Cloud API Key. You can create this at https://cloud.ibm.com/iam#/apikeys"
  sensitive   = true
}

# Default region provider
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.region
}

# Additional region provider for cross-region resources (e.g., snapshot copies)
provider "ibm" {
  alias            = "eu-de"
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "eu-de"
}

# Additional region provider for multi-region scenarios
provider "ibm" {
  alias            = "us-east"
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "us-east"
}