variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for project_config
variable "project_config_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}

// Resource arguments for project
variable "project_resource_group" {
  description = "The resource group where the project's data and tools are created."
  type        = string
  default     = "Default"
}
variable "project_location" {
  description = "The location where the project's data and tools are created."
  type        = string
  default     = "us-south"
}

// Data source arguments for project_config
variable "project_config_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}
variable "project_config_id" {
  description = "The unique config ID."
  type        = string
  default     = "id"
}

// Data source arguments for project
variable "project_id" {
  description = "The unique project ID."
  type        = string
  default     = "id"
}
