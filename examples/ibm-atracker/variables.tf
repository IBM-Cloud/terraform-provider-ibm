variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for atracker_target
variable "atracker_target_name" {
  description = "The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`."
  type        = string
  default     = "my-cos-target"
}
variable "atracker_target_target_type" {
  description = "The type of the target."
  type        = string
  default     = "target_type"
}

// Resource arguments for atracker_route
variable "atracker_route_name" {
  description = "The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`."
  type        = string
  default     = "my-route"
}
variable "atracker_route_receive_global_events" {
  description = "Indicates whether or not all global events should be forwarded to this region."
  type        = bool
  default     = false
}
variable "atracker_route_rules" {
  description = "Routing rules that will be evaluated in their order of the array."
  type        = list(object({ example=string }))
  default     = [ { "target_ids" : [ "target_ids" ] } ]
}

// Data source arguments for atracker_targets
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

// Data source arguments for atracker_endpoints
