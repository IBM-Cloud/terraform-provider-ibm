variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
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
variable "project_name" {
  description = "The name of the project."
  type        = string
  default     = "acme-microservice"
}
variable "project_description" {
  description = "A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description."
  type        = string
  default     = "A microservice to deploy on top of ACME infrastructure."
}
variable "project_destroy_on_delete" {
  description = "The policy that indicates whether the resources are destroyed or not when a project is deleted."
  type        = bool
  default     = true
}

// Resource arguments for project_config
variable "project_config_project_id" {
  description = "The unique project ID."
  type        = string
  default     = "project_id"
}
variable "project_config_name" {
  description = "The name of the configuration."
  type        = string
  default     = "name"
}
variable "project_config_labels" {
  description = "A collection of configuration labels."
  type        = list(string)
  default     = [ "labels" ]
}
variable "project_config_description" {
  description = "The description of the project configuration."
  type        = string
  default     = "description"
}
variable "project_config_locator_id" {
  description = "A dotted value of catalogID.versionID."
  type        = string
  default     = "locator_id"
}

// Data source arguments for project
variable "project_id" {
  description = "The unique project ID."
  type        = string
  default     = "id"
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
