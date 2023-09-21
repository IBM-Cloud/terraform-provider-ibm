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
variable "project_location" {
  description = "The IBM Cloud location where a resource is deployed."
  type        = string
  default     = "us-south"
}
variable "project_resource_group" {
  description = "The resource group where the project's data and tools are created."
  type        = string
  default     = "Default"
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
