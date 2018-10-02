variable "bluemix_api_key" {
  description = "Your IBM Cloud platform API key"
}
variable "function_namespace" {
  description = "Your Cloud Functions namespace"
}

provider "ibm" {
  bluemix_api_key = "${var.bluemix_api_key}"
  function_namespace = "${var.function_namespace}"
}
