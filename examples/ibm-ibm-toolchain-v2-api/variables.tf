variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for toolchain

// Resource arguments for toolchain_integration
variable "toolchain_integration_toolchain_id" {
  description = "ID of the toolchain to bind integration to."
  type        = string
  default     = "toolchain_id"
}
variable "toolchain_integration_service_id" {
  description = "The unique short name of the service that should be provisioned."
  type        = string
  default     = "todolist"
}
variable "toolchain_integration_name" {
  description = "Name of tool integration."
  type        = string
  default     = "name"
}
variable "toolchain_integration_parameters" {
  description = "Arbitrary JSON data."
  type        = map()
  default     = { "key": null }
}
