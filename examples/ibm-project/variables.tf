variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for project_instance
variable "project_instance_resource_group" {
  description = "The resource group where the project's data and tools are created."
  type        = string
  default     = "resource_group"
}
variable "project_instance_location" {
  description = "The location where the project's data and tools are created."
  type        = string
  default     = "location"
}
variable "project_instance_name" {
  description = "The project name."
  type        = string
  default     = "acme-microservice"
}
variable "project_instance_description" {
  description = "A project's descriptive text."
  type        = string
  default     = "A microservice to deploy on top of ACME infrastructure."
}

// Data source arguments for project_event_notification
variable "project_event_notification_id" {
  description = "The unique identifier."
  type        = string
  default     = "id"
}
variable "project_event_notification_exclude_configs" {
  description = "When set to true, exclude_configs returns only active configurations. Draft configurations are not returned."
  type        = bool
  default     = false
}
