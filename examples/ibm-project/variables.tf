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
  description = "The resource group name where the project's data and tools are created."
  type        = string
  default     = "Default"
}

// Resource arguments for project_environment
variable "project_environment_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}

// Data source arguments for project_config
variable "data_project_config_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}
variable "data_project_config_project_config_id" {
  description = "The unique configuration ID."
  type        = string
  default     = "project_config_id"
}

// Data source arguments for project
variable "data_project_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}

// Data source arguments for project_environment
variable "data_project_environment_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}
variable "data_project_environment_project_environment_id" {
  description = "The environment ID."
  type        = string
  default     = "project_environment_id"
}
