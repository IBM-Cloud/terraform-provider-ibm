variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
  default     = "<Add IAM Apikey>"
}
variable "region" {
  description = "IBM Cloud region where your App Configuration instance is located."
  type        = string
  default     = "au-syd"
}
variable "app_config_guid" {
  description = "App Configuration instance id or guid."
  type        = string
  default     = "36401ffc-6280-459a-ba98-456aba10d0c7"
}
variable "app_config_environment_id" {
  description = "Id of the environment created in App Configuration instance under the Environments section."
  type        = string
  default     = "dev"
}
variable "app_config_collection_id" {
  description = "Id of the collection created in App Configuration instance under the Collections section."
  type        = string
  default     = "car-rentals"
}
variable "app_config_feature_id" {
  description = "Feature flag id required to be evaluated."
  type        = string
  default     = "weekend-discount"
}
variable "app_config_property_id" {
  description = "Property id required to be evaluated."
  type        = string
  default     = "users-location"
}
variable "app_config_entity_id" {
  description = "Entity Id."
  type        = string
  default     = "user123"
}
variable "app_config_entity_attributes" {
  description = "Entity attributes for evaluation."
  type        = map(any)
  default = {
    city = "Bangalore",
    radius = 60,
  }
}
variable "app_config_feature_flag_ids" {
  description = "List of feature flag id's required to evaluate."
  type        = list(string)
  default     = ["feature-1", "feature-2", "feature-3", "feature-4"]
}
variable "app_config_property_ids" {
  description = "List of property id's required to evaluate."
  type        = list(string)
  default     = ["property-1", "property-2"]
}
