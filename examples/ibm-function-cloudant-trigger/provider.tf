variable "ibmcloud_api_key" {
  description = "Your IBM Cloud platform API key"
}
variable "function_namespace" {
  description = "Your Cloud Functions namespace"
}

provider "ibm" {
  ibmcloud_api_key = "${var.ibmcloud_api_key}"
  function_namespace = "${var.function_namespace}"
}
