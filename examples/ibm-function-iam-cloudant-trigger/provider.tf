variable "ibmcloud_api_key" {
  description = "Your IBM Cloud platform API key"
}

variable "function_namespace" {
  description = "Your Cloud Functions namespace"
}

variable "region" {
  description = "region"
}


provider "ibm" {
  function_namespace = var.function_namespace
  ibmcloud_api_key   = var.ibmcloud_api_key
  region = var.region
}



