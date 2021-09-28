variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "provider_id" {
  default = "scc"
  description = "Part of parent. This field contains the provider_id for example: providers/{provider_id}"
}

