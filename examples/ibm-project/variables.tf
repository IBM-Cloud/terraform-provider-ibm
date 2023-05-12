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
  description = "The unique project ID."
  type        = string
  default     = "id"
}
