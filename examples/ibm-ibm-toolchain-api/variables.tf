variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for toolchain_tool_git
variable "toolchain_tool_git_git_provider" {
  description = ""
  type        = string
  default     = "git_provider"
}
variable "toolchain_tool_git_toolchain_id" {
  description = ""
  type        = string
  default     = "toolchain_id"
}
