variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for atracker_target
variable "atracker_target_name" {
  description = "The name of the target resource."
  type        = string
  default     = "my-cos-target"
}
variable "atracker_target_target_type" {
  description = "The type of the target."
  type        = string
  default     = "cloud_object_storage"
}
variable "atracker_target_region" {
  description = "Included this optional field if you used it to create a target in a different region other than the one you are connected."
  type        = string
  default     = "us-south"
}

variable "atracker_route_receive_global_events" {
  description = "Indicates whether or not all global events should be forwarded to this region."
  type        = bool
  default     = false
}
variable "atracker_route_rules" {
  description = "Routing rules that will be evaluated in their order of the array."
  type        = list(object({
    target_ids = list(string)
    locations = list(string)
  }))
  default     = [ ]
}
// Resource arguments for atracker_route
variable "atracker_route_name" {
  description = "The name of the route."
  type        = string
  default     = "my-route"
}

// Resource arguments for atracker_settings
variable "atracker_settings_default_targets" {
  description = "The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event. Enterprise-managed targets are not supported."
  type        = list(string)
  default     = [ "c3af557f-fb0e-4476-85c3-0889e7fe7bc4" ]
}
variable "atracker_settings_permitted_target_regions" {
  description = "If present then only these regions may be used to define a target."
  type        = list(string)
  default     = [ "us-south" ]
}
variable "atracker_settings_metadata_region_primary" {
  description = "To store all your meta data in a single region."
  type        = string
  default     = "us-south"
}
variable "atracker_settings_metadata_region_backup" {
  description = "To store all your meta data in a backup region."
  type        = string
  default     = "eu-de"
}
variable "atracker_settings_private_api_endpoint_only" {
  description = "If you set this true then you cannot access api through public network."
  type        = bool
  default     = false
}

// Data source arguments for atracker_targets
variable "atracker_targets_region" {
  description = "Limit the query to the specified region."
  type        = string
  default     = "placeholder"
}
variable "atracker_targets_name" {
  description = "The name of the target resource."
  type        = string
  default     = "a-cos-target-us-south"
}

// Data source arguments for atracker_routes
variable "atracker_routes_name" {
  description = "The name of the route."
  type        = string
  default     = "my-route"
}
