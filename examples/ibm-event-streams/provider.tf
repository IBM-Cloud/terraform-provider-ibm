variable "ibmcloud_api_key" {
    description = "Enter your IBM Cloud API Key, you can get your IBM Cloud API key using: https://cloud.ibm.com/iam#/apikeys"
}

variable "es_reader_api_key" {
    description = "If user wish to use fine-grained access control for Event Streams instances and topics, create a service ID and add access policies for different access roles and scope to different resources. refer to: https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-security"
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  generation       = 1
}