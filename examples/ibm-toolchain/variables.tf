variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for toolchain_tool_secretsmanager
variable "toolchain_tool_secretsmanager_toolchain_id" {
  description = "ID of the toolchain to bind integration to."
  type        = string
  default     = "toolchain_id"
}
variable "toolchain_tool_secretsmanager_name" {
  description = "Name of tool integration."
  type        = string
  default     = "name"
}
variable "toolchain_tool_secretsmanager_parameters_references" {
  description = "Decoded values used on provision in the broker that reference fields in the parameters."
  type        = map()
  default     = { "key": null }
}
