variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cd_toolchain_tool_sonarqube
variable "cd_toolchain_tool_sonarqube_toolchain_id" {
  description = "ID of the toolchain to bind integration to."
  type        = string
  default     = "toolchain_id"
}
variable "cd_toolchain_tool_sonarqube_name" {
  description = "Name of tool integration."
  type        = string
  default     = "name"
}

// Data source arguments for cd_toolchain_tool_sonarqube
variable "cd_toolchain_tool_sonarqube_toolchain_id" {
  description = "ID of the toolchain."
  type        = string
  default     = "toolchain_id"
}
variable "cd_toolchain_tool_sonarqube_integration_id" {
  description = "ID of the tool integration bound to the toolchain."
  type        = string
  default     = "integration_id"
}
