variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for config_aggregator_settings
variable "config_aggregator_settings_resource_collection_enabled" {
  description = "The field denoting if the resource collection is enabled."
  type        = bool
  default     = true
}
variable "config_aggregator_settings_trusted_profile_id" {
  description = "The trusted profile id that provides Reader access to the App Configuration instance to collect resource metadata."
  type        = string
  default     = "Profile-7d935dbb-7ee5-44e1-9e85-38e65d722398"
}
variable "config_aggregator_settings_regions" {
  description = "The list of regions across which the resource collection is enabled."
  type        = list(string)
  default     = ["all"]
}

// Data source arguments for config_aggregator_configurations
variable "account_id" {
  description = "Account ID for the IBM Cloud instance"
  type        = string
  default     = "18b006922637405389523fb06338e363"
}

variable "config_type" {
  description = "Configuration type for the resource"
  type        = string
  default     = "instance"
}

variable "last_config_refresh_time" {
  description = "Last configuration refresh time"
  type        = string
  default     = "2024-09-16T00:30:22Z"
}

variable "location" {
  description = "Location of the resource"
  type        = string
  default     = "us-south"
}

variable "resource_crn" {
  description = "Resource CRN"
  type        = string
  default     = "crn:v1:staging:public:project:us-south:a/18b006922637405389523fb06338e363:3fdba864-9ab0-4cbd-a705-a78ce1e757ea::"
}

variable "resource_group_id" {
  description = "Resource Group ID"
  type        = string
  default     = "c5450ce8dde54581bdfb8e785d27d292"
}

variable "resource_name" {
  description = "Resource Name"
  type        = string
  default     = "logdna"
}

variable "service_name" {
  description = "Service Name"
  type        = string
  default     = "project"
}

variable "event_notification_enabled" {
  description = "Whether event notification is enabled"
  type        = bool
  default     = true
}