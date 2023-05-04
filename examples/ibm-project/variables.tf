variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for project_instance
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
variable "project_instance_resource_group" {
  description = "Group name of the customized collection of resources."
  type        = string
  default     = "resource_group"
}
variable "project_instance_location" {
  description = "Data center locations for resource deployment."
  type        = string
  default     = "location"
}

// Data source arguments for project_event_notification
variable "project_event_notification_id" {
  description = "The ID of the project, which uniquely identifies it."
  type        = string
  default     = "id"
}
variable "project_event_notification_exclude_configs" {
  description = "Only return with the active configuration, no drafts."
  type        = bool
  default     = false
}
variable "project_event_notification_complete" {
  description = "The flag to determine if full metadata should be returned."
  type        = bool
  default     = false
}
