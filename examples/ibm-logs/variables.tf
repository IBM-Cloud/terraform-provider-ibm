variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}
variable "resource_name_prefix" {
  description = "Prefix to the name of the resource"
  type        = string
  default     = "logs"
}
variable "en_instance_guid" {
  description = "GUID of event notifications instance guid to create webhook"
  type        = string
}
variable "en_instance_region" {
  description = "region of event notifications instance guid to create webhook"
  type        = string
}


# // Data source arguments for logs_dashboard
# variable "data_logs_dashboard_dashboard_id" {
#   description = "The ID of the dashboard."
#   type        = string
#   default     = "dashboard_id"
# }

# // Data source arguments for logs_e2m
# variable "data_logs_e2m_logs_e2m_id" {
#   description = "id of e2m to be deleted."
#   type        = string
#   default     = "logs_e2m_id"
# }


# // Data source arguments for logs_view
# variable "data_logs_view_logs_view_id" {
#   description = "View id."
#   type        = number
#   default     = 2
# }


# // Data source arguments for logs_view_folder
# variable "data_logs_view_folder_logs_view_folder_id" {
#   description = "Folder id."
#   type        = string
#   default     = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
# }

