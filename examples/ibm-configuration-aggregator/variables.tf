variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "region"{
  description="Config Aggregator Instance ID"
  type=string
}

variable "instance_id"{
  description="Config Aggregator Instance ID"
  type=string
}

// Resource arguments for config_aggregator_settings
variable "config_aggregator_settings_resource_collection_enabled" {
  description = "The field denoting if the resource collection is enabled."
  type        = bool
}
variable "config_aggregator_settings_trusted_profile_id" {
  description = "The trusted profile id that provides Reader access to the App Configuration instance to collect resource metadata."
  type        = string
}
variable "config_aggregator_settings_regions" {
  description = "The list of regions across which the resource collection is enabled."
  type        = list(string)
  default     = ["all"]
}

// Data source arguments for config_aggregator_configurations
variable "config_aggregator_configurations_config_type" {
  description = "The type of resource configuration that are to be retrieved."
  type        = string
  default     = "placeholder"
}
variable "config_aggregator_configurations_service_name" {
  description = "The name of the IBM Cloud service for which resources are to be retrieved."
  type        = string
  default     = "placeholder"
}
variable "config_aggregator_configurations_resource_group_id" {
  description = "The resource group id of the resources."
  type        = string
  default     = "placeholder"
}
variable "config_aggregator_configurations_location" {
  description = "The location or region in which the resources are created."
  type        = string
  default     = "placeholder"
}
variable "config_aggregator_configurations_resource_crn" {
  description = "The crn of the resource."
  type        = string
  default     = "placeholder"
}


