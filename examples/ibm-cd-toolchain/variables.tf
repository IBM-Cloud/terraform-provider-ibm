variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for cd_toolchain_tool_sonarqube
variable "cd_toolchain_tool_sonarqube_toolchain_id" {
  description = "ID of the toolchain to bind tool to."
  type        = string
  default     = "toolchain_id"
}
variable "cd_toolchain_tool_sonarqube_name" {
  description = "Name of tool."
  type        = string
  default     = "name"
}

// Data source arguments for cd_toolchain_tool_sonarqube
variable "cd_toolchain_tool_sonarqube_toolchain_id" {
  description = "ID of the toolchain."
  type        = string
  default     = "toolchain_id"
}
variable "cd_toolchain_tool_sonarqube_tool_id" {
  description = "ID of the tool bound to the toolchain."
  type        = string
  default     = "tool_id"
}
