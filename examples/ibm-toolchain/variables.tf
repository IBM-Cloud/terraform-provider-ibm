variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for toolchain_tool_sonarqube
variable "toolchain_tool_sonarqube_toolchain_id" {
  description = "ID of the toolchain to bind integration to."
  type        = string
  default     = "toolchain_id"
}
variable "toolchain_tool_sonarqube_name" {
  description = "Name of tool integration."
  type        = string
  default     = "name"
}
